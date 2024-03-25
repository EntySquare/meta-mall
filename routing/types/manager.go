package types

type LoginMangerReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginManagerResp struct {
	Token string `json:"token"`
}
type InsertNftReq struct {
	Name            string  `json:"name"`
	NftNumbers      []int64 `json:"nft_number"`       //编号
	TokenId         []int64 `json:"token_id"`         //NFT线上ID
	ContractAddress string  `json:"contract_address"` //合约地址
	OwnerAddress    string  `json:"owner_address"`    //拥有者地址
	Price           float64 `json:"price"`            //价格
	TokenName       string  `json:"token_name"`       //代币
	Power           float64 `json:"power"`            //算力
	TypeNum         int64   `json:"type_num"`         //种类
	ImgUrl          string  `json:"img_url"`          //图片地址
}
type TokenIdFromResp struct {
	TokenId int64 `json:"token_id"`
}
