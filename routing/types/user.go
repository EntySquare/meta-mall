package types

type SengMsg struct {
	Msg   string `json:"msg"`
	Phone string `json:"phone"`
	Area  string `json:"area"`
}

// 登录注册请求参数
type LoginAndRegisterReq struct {
	WalletAddress string `json:"wallet_address"`
	RecommendId   uint   `json:"recommend_id"`
	Code          string `json:"code"`
}

// 修改支付密码请求参数
type UpdatePwdReq struct {
	WalletAddress string `json:"wallet_address"`
	NewPwd        string `json:"new_pwd"`
	OldPwd        string `json:"old_pwd"`
}

type MyInvestmentResp struct {
	UID                    string               `json:"uid"`
	Level                  int64                `json:"level"`
	AccumulatedPledgeCount int64                `json:"accumulated_pledge_count"`
	InvestmentCount        int64                `json:"investment_count"`
	InvestmentAddress      string               `json:"investment_address"`
	InvestmentUsers        []InvestmentUserInfo `json:"investment_users"`
}
type InvestmentUserInfo struct {
	UID         string `json:"uid"`
	Level       int64  `json:"level"`
	PledgeCount int64  `json:"pledge_count"`
}
type MyNgtResp struct {
	BenefitInfo  BenefitInfo       `json:"benefit_info"`
	Transactions []TransactionInfo `json:"transactions"`
}
type TransactionInfo struct {
	Num             float64 `json:"num"`
	Chain           string  `json:"chain"`
	Status          string  `json:"status"`
	Address         string  `json:"address"`
	Hash            string  `json:"hash"`
	AskForTime      int64   `json:"ask_for_time"`
	AchieveTime     int64   `json:"achieve_time"`
	TransactionType string  `json:"transaction_type"`
}
type BenefitInfo struct {
	Balance            float64 `json:"balance"`
	LastDayBenefit     float64 `json:"last_day_benefit"`
	AccumulatedBenefit float64 `json:"accumulated_benefit"`
}
type MyContractFlowResp struct {
	BenefitInfo BenefitInfo    `json:"benefit_info"`
	Contracts   []ContractInfo `json:"Contract_flows"`
}
type ContractInfo struct {
	ContractId         uint    `json:"Contract_id"`
	NFTName            string  `json:"nft_name"`
	PledgeId           string  `json:"pledge_id"`
	ChainName          string  `json:"chain_name"`
	Duration           string  `json:"duration"`
	Hash               string  `json:"hash"`
	InterestRate       float64 `json:"interest_rate"`
	AccumulatedBenefit float64 `json:"accumulated_benefit"`
	PledgeFee          float64 `json:"pledge_fee"`
	ReleaseFee         float64 `json:"release_fee"`
	StartTime          int64   `json:"start_time"`
	ExpireTime         int64   `json:"expire_time"`
	NFTReleaseTime     int64   `json:"nft_release_time"`
	Flag               string  `json:"flag"`
}
type InviteeInfoReq struct {
	Uid string `json:"uid"`
}
type InviteeInfoResp struct {
	Uid         string         `json:"uid"`
	Level       int64          `json:"level"`
	PledgeCount int64          `json:"pledge_count"`
	CreateTime  int64          `json:"create_time"`
	Contracts   []ContractInfo `json:"Contract_flows"`
}
type ContractDetailReq struct {
	Hash string `json:"hash"`
}
