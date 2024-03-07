package types

type NftDetail struct {
	Id              uint   `json:"id"`
	Name            string `json:"name"`
	NftNumber       int64  `json:"nft_number"`
	TokenId         int64  `json:"token_id"`
	ContractAddress string `json:"contract_address"`
	OwnerAddress    string `json:"owner_address"`
	Price           int64  `json:"price"`
	TypeNum         int64  `json:"type_num"`
	ImgUrl          string `json:"img_url"`
}
type NftListResp struct {
	List []NftDetail `json:"nft_list"`
}
