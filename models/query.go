package models

// 用户查询请求对象
type UserQueryReqDto struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

// 商品查询请求对象
type GoodsQueryReqDto struct {
	GoodsName string `json:"goodsName"`
	PageNo    int    `json:"pageNo"`
	PageSize  int    `json:"pageSize"`
}
