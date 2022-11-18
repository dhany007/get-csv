package models

type ShopingCart struct {
	Vegetables string `json:"vegetables"`
	Fruits     string `json:"fruits"`
}

type Object struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}
