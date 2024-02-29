package types

type WithdrawNgtReq struct {
	Num     float64 `json:"num"`
	Address string  `json:"address"`
	Hash    string  `json:"hash"`
	Chain   string  `json:"chain"`
}

type DepositNgtReq struct {
	Num     float64 `json:"num"`
	Address string  `json:"address"`
	Hash    string  `json:"hash"`
	Chain   string  `json:"chain"`
}

type CheckHashApiReq struct {
	Hash string `json:"hash"`
}
