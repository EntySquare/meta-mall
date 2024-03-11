package types

type ContractDetail struct {
	Id                 uint    `json:"id"`
	AccumulatedBenefit float64 `json:"accumulated_benefit"`
	Power              float64 `json:"power"` //算力
	StartTime          int64   `json:"start_time"`
	Flag               string  `json:"flag"`
}
type GetMyContractListResp struct {
	List []ContractDetail `json:"contract_list"`
}
