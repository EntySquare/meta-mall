package types

type ContractDetail struct {
	Id                 uint    `json:"id"`                  //合约ID
	AccumulatedBenefit float64 `json:"accumulated_benefit"` //累计收益
	Power              float64 `json:"power"`               //算力
	TokenName          string  `json:"token_name"`          //收益token名字
	StartTime          int64   `json:"start_time"`          //开始时间
	Flag               string  `json:"flag"`                //状态（1-待确认 2- 进行中 ）
}
type GetMyContractListResp struct {
	List []ContractDetail `json:"contract_list"`
}
type GetMiningIncomeResultResp struct {
	AllPowers                 float64 `json:"all_powers"`
	AllAccumulatedUSDTBenefit float64 `json:"all_accumulated_usdt"` //usdt全网总收益
	MyAccumulatedUSDTBenefit  float64 `json:"my_accumulated_usdt"`  //usdt我的总收益
	MyAvailableUSDTBenefit    float64 `json:"my_available_usdt"`    //usdt可领取收益
	AllAccumulatedUNCBenefit  float64 `json:"all_accumulated_unc"`  //unc全网总收益
	MyAccumulatedUNCBenefit   float64 `json:"my_accumulated_unc"`   //unc我的总收益
	MyAvailableUNCBenefit     float64 `json:"my_available_unc"`     //unc可领取收益
}
