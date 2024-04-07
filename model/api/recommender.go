package api

import (
	"gorm.io/gorm"
	"meta-mall/model"
)

type SingleUser struct {
	Address       string
	RecommenderId uint
	Level         int64
	Power         float64
	Branch        []uint
}

var UserTree map[uint]SingleUser

// AddNewBranch 增加新的用户树枝干 treeId:推荐人 branchId:当前用户
func AddNewBranch(treeId uint, address string, branchId uint) {
	branch := make([]uint, 0)
	UserTree[branchId] = SingleUser{
		Address:       address,
		RecommenderId: treeId,
		Power:         0,
		Branch:        branch,
	}
	//newBranch := make([]uint, 0)
	newBranch := append(UserTree[treeId].Branch, branchId)
	UserTree[treeId] = SingleUser{
		Address:       address,
		RecommenderId: UserTree[treeId].RecommenderId,
		Power:         UserTree[treeId].Power,
		Branch:        newBranch,
	}
}

// GetBranchAccumulatedInput 累计算力
func GetBranchAccumulatedPower(userId uint) float64 {
	AccumulatedPower := float64(0)
	for _, branch := range UserTree[userId].Branch {
		if len(UserTree[branch].Branch) > 0 {
			AccumulatedPower += GetBranchAccumulatedPower(branch)
		}
		AccumulatedPower += UserTree[branch].Power
	}
	return AccumulatedPower
}

// GetUserPromotionPower  推广算力
func GetUserPromotionPower(userId uint) float64 {
	AccumulatedPower := float64(0)
	for _, branch := range UserTree[userId].Branch {
		AccumulatedPower += UserTree[branch].Power * 0.6
		AccumulatedPower += getBranchPromotionPower(branch, 0)
		if len(UserTree[branch].Branch) > 0 {
			AccumulatedPower += getBranchPromotionPower(branch, 0)
			if UserTree[branch].Level < UserTree[userId].Level {
				differential := UserTree[userId].Level - UserTree[branch].Level
				AccumulatedPower += getCommunityPromotionPower(branch, differential)
			}
		}
	}
	return AccumulatedPower
}
func getBranchPromotionPower(userId uint, generation int64) float64 {
	AccumulatedPower := UserTree[userId].Power * 0.05
	generation++
	if generation == 5 {
		return AccumulatedPower
	}
	for _, branch := range UserTree[userId].Branch {
		if len(UserTree[branch].Branch) > 0 {
			AccumulatedPower += getBranchPromotionPower(branch, generation)
		}
	}
	return AccumulatedPower
}

// 534
func getCommunityPromotionPower(userId uint, differential int64) float64 {
	AccumulatedPower := UserTree[userId].Power * 0.1 * float64(differential)
	if len(UserTree[userId].Branch) > 0 {
		for _, branch := range UserTree[userId].Branch {
			if UserTree[branch].Level < UserTree[userId].Level {
				AccumulatedPower += getCommunityPromotionPower(branch, differential)
			} else if UserTree[branch].Level >= UserTree[userId].Level && UserTree[branch].Level < UserTree[userId].Level+differential {
				AccumulatedPower += getCommunityPromotionPower(branch, UserTree[userId].Level+differential-UserTree[branch].Level)
			} else {
				return AccumulatedPower
			}
		}
	}

	return AccumulatedPower

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
func UserPurchase(userId uint, count float64) {
	UserTree[userId] = SingleUser{
		Address:       UserTree[userId].Address,
		RecommenderId: UserTree[userId].RecommenderId,
		Power:         UserTree[userId].Power + count,
		Level:         UserTree[userId].Level,
		Branch:        UserTree[userId].Branch,
	}
}

// UserLevelCheck 用户等级核算
func UserLevelCheck(userId uint) int64 {
	if UserTree[userId].Level < 5 {
		switch UserTree[userId].Level {
		case 0:
			maxBranchPower := float64(0)
			allBranchPower := float64(0)
			if len(UserTree[userId].Branch) > 0 {
				for _, branch := range UserTree[userId].Branch {
					branchPower := GetBranchAccumulatedPower(branch)
					if branchPower > maxBranchPower {
						maxBranchPower = branchPower
					}
					allBranchPower += branchPower
				}
			}
			if allBranchPower > 25 {
				return 1
			}
		default:
			count := 0
			if len(UserTree[userId].Branch) > 0 {
				for _, branch := range UserTree[userId].Branch {
					if UserTree[branch].Level > UserTree[userId].Level {
						count++
					}
				}
			}
			if count > 2 {
				return UserTree[userId].Level + 1
			}

		}

	}
	return UserTree[userId].Level
}

// PowerLose 算力合约到期
func ContractEnd(userId uint, count float64) {
	UserTree[userId] = SingleUser{
		Address:       UserTree[userId].Address,
		RecommenderId: UserTree[userId].RecommenderId,
		Power:         UserTree[userId].Power - count,
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
				Address:       user.WalletAddress,
				RecommenderId: 0,
				Power:         user.Power,
				Level:         user.Level,
				Branch:        nil,
			}
		} else {
			UserTree[user.ID] = SingleUser{
				Address:       user.WalletAddress,
				RecommenderId: user.RecommendId,
				Power:         user.Power,
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
			Address:       UserTree[u.RecommendId].Address,
			RecommenderId: UserTree[u.RecommendId].RecommenderId,
			Power:         UserTree[u.RecommendId].Power,

			Level:  UserTree[u.RecommendId].Level,
			Branch: newBranch,
		}
	}
}
