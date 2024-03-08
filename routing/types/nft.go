package types

type NftDetail struct {
	Id              uint   `json:"id"`
	Name            string `json:"name"`
	NftNumber       int64  `json:"nft_number"`       //编号
	TokenId         int64  `json:"token_id"`         //NFT线上ID
	ContractAddress string `json:"contract_address"` //合约地址
	OwnerAddress    string `json:"owner_address"`    //拥有者地址
	Price           int64  `json:"price"`            //价格
	TypeNum         int64  `json:"type_num"`         //种类
	ImgUrl          string `json:"img_url"`          //图片地址
}
type NftListResp struct {
	List []NftDetail `json:"nft_list"`
}
type PurchaseNftReq struct {
	NftId uint `json:"id"`
}
type PurchaseNftResp struct {
}
