package models

// 商品查询请求对象
type GoodsQueryReqDto struct {
	GoodsName string `json:"goodsName"`
	PageNo    int    `json:"pageNo"`
	PageSize  int    `json:"pageSize"`
}
