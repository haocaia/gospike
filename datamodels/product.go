package datamodels

type Product struct {
	ID           int64  `json:"id" sql:"ID" show:"ID"`
	ProductName  string `json:"ProductName" sql:"productName" show:"ProductName"`
	ProductNum   int64  `json:"ProductNum" sql:"productNum" show:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" show:"ProductImage"`
	ProductUrl   string `json:"ProductUrl" sql:"productUrl" show:"ProductUrl"`
}
