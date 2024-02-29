package api

import (
	"gorm.io/gorm"
	"meta-mall/model"
)

type SingleUser struct {
	UID           string
	RecommenderId uint
	Level         int64
	PledgeCount   int64
	Branch        []uint
}

var UserTree map[uint]SingleUser

// AddNewBranch 增加新的用户树枝干 treeId:推荐人 branchId:当前用户
func AddNewBranch(treeId uint, branchId uint) {
	branch := make([]uint, 0)
	UserTree[branchId] = SingleUser{
		RecommenderId: treeId,
		PledgeCount:   0,
		Branch:        branch,
	}
	//newBranch := make([]uint, 0)
	newBranch := append(UserTree[treeId].Branch, branchId)
	UserTree[treeId] = SingleUser{
		RecommenderId: UserTree[treeId].RecommenderId,
		PledgeCount:   UserTree[treeId].PledgeCount,
		Branch:        newBranch,
	}
}

// GetBranchAccumulatedInput 用户累计投入
func GetBranchAccumulatedPledgeCount(userId uint) int64 {
	AccumulatedPledgeCount := int64(0)
	for _, branch := range UserTree[userId].Branch {
		AccumulatedPledgeCount += UserTree[branch].PledgeCount
	}
	return AccumulatedPledgeCount
}

// GetAssociate 用户下级用户
func GetAssociate(userId uint) []uint {
	firstLevelAssociate := make([]uint, 0)

	recommendId := UserTree[userId].RecommenderId
	if recommendId != 0 {
		for _, branch := range UserTree[recommendId].Branch {
			firstLevelAssociate = append(firstLevelAssociate, branch)
		}
	}
	return firstLevelAssociate
}

// UserInput 用户新增投入
func UserPledge(userId uint, count int64) {
	UserTree[userId] = SingleUser{
		UID:           UserTree[userId].UID,
		RecommenderId: UserTree[userId].RecommenderId,
		PledgeCount:   UserTree[userId].PledgeCount + count,
		Level:         UserTree[userId].Level,
		Branch:        UserTree[userId].Branch,
	}
}

// InitUserTree 初始化用户树
func InitUserTree(db *gorm.DB) error {
	UserTree = make(map[uint]SingleUser)
	// 查询出所有用户
	users, err := model.SelectAllUser(db)
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.ID == 0 {
			UserTree[user.ID] = SingleUser{
				UID:           user.UID,
				RecommenderId: 0,
				PledgeCount:   user.PledgeCount,
				Level:         user.Level,
				Branch:        nil,
			}
		} else {
			UserTree[user.ID] = SingleUser{
				UID:           user.UID,
				RecommenderId: user.RecommendId,
				PledgeCount:   user.PledgeCount,
				Level:         user.Level,
				Branch:        nil,
			}
		}
	}
	initUserBranch(users)
	return nil
}

func initUserBranch(users []model.User) {
	for _, u := range users {
		if u.ID == 0 {
			continue
		}
		newBranch := make([]uint, 0)
		if UserTree[u.RecommendId].Branch != nil {
			newBranch = append(UserTree[u.RecommendId].Branch, u.ID)
		} else {
			newBranch = append(newBranch, u.ID)
		}

		UserTree[u.RecommendId] = SingleUser{
			UID:           UserTree[u.RecommendId].UID,
			RecommenderId: UserTree[u.RecommendId].RecommenderId,
			PledgeCount:   UserTree[u.RecommendId].PledgeCount,
			Level:         UserTree[u.RecommendId].Level,
			Branch:        newBranch,
		}
	}
}
