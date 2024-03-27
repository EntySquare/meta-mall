package types

type NftDetail struct {
	Id              uint    `json:"id"`
	Name            string  `json:"name"`
	NftNumber       int64   `json:"nft_number"`       //编号
	TokenId         int64   `json:"token_id"`         //NFT线上ID
	ContractAddress string  `json:"contract_address"` //合约地址
	OwnerAddress    string  `json:"owner_address"`    //拥有者地址
	Price           float64 `json:"price"`            //价格
	TokenName       string  `json:"token_name"`       //代币
	Power           float64 `json:"power"`            //算力
	TypeNum         int64   `json:"type_num"`         //种类
	ImgUrl          string  `json:"img_url"`          //图片地址
	Flag            string  `json:"flag"`             //标记 1-未售出 2- 确认中 0-已售出
}
type NftListResp struct {
	List []NftDetail `json:"nft_list"`
}
type PurchaseCheackReq struct {
	NftId uint `json:"id"`
}
type CancelCheackReq struct {
	NftId uint `json:"id"`
}
type PurchaseNftReq struct {
	NftId     uint   `json:"id"`
	Hash      string `json:"hash"`
	TokenName string `json:"token_name"`
}

type PurchaseNftResp struct {
	Price float64 `json:"price"`
}
type GetMyNftListReq struct {
}
type NftOrderDetail struct {
	Id              uint    `json:"id"`
	Name            string  `json:"name"`
	NftNumber       int64   `json:"nft_number"`       //编号
	TokenId         int64   `json:"token_id"`         //NFT线上ID
	ContractAddress string  `json:"contract_address"` //合约地址
	OwnerAddress    string  `json:"owner_address"`    //拥有者地址
	Price           float64 `json:"price"`            //价格
	TokenName       string  `json:"token_name"`       //代币
	Power           float64 `json:"power"`            //算力
	TypeNum         int64   `json:"type_num"`         //种类
	ImgUrl          string  `json:"img_url"`          //图片地址
	OrderId         uint    `json:"id"`               //订单id
	OrderFlag       string  `json:"order_flag"`       //标记 (1-处理中 2-已完成 0-取消)
}
type GetMyNftListResp struct {
	List []NftOrderDetail `json:"order_list"`
}
