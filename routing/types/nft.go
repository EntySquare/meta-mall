package types

type PledgeNgtReq struct {
	NftId    string `json:"nft_id"`
	Duration string `json:"duration"`
	Hash     string `json:"hash"`
	Chain    string `json:"chain"`
}
type PledgeNgtResp struct {
	Code int64 `json:"code"`
}
type CancelCovenantReq struct {
	CovenantId uint `json:"covenant_id"`
}
type CancelCovenantResp struct {
	Code int64 `json:"code"`
}
type GetAllNftInfoResp struct {
	List []NftInfo `json:"list"`
}
type NftInfo struct {
	NftInfoId uint   `json:"nft_info_id"`
	Name      string `json:"name"`
	TypeNum   int64  `json:"type_num"`
	DayRate   string `json:"day_rate"`
	ImgUrl    string `json:"img_url"`
}
type UpdateNftInfoReq struct {
	NftInfoId uint   `json:"nft_info_id"`
	DayRate   string `json:"day_rate"`
}
type UpdateNftInfoResp struct {
	Code int64 `json:"code"`
}
