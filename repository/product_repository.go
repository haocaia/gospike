package repository

import (
	sq "github.com/Masterminds/squirrel"
	"gospike/datamodels"
	"gospike/utils"
)

const (
	TableName = "product"
)

type IProduct interface {
	Insert(*datamodels.Product) (int64, error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64) (*datamodels.Product, error)
	SelectAll() ([]*datamodels.Product, error)
}

type ProductImpl struct {
}

func (p *ProductImpl) Insert(product *datamodels.Product) (int64, error) {
	db := utils.GetConn()
	// "INSERT product SET productName=?,productNum=?,productImage=?,productUrl=?"
	// result: lastInsertId, affect rows
	result, err := sq.Insert(TableName).Columns("productName", "productNum", "productImage", "productUrl").
		Values(product.ProductName, product.ProductNum, product.ProductImage, product.ProductUrl).RunWith(db).Exec()
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()

}

func (p *ProductImpl) Delete(i int64) bool {
	//TODO implement me
	panic("implement me")
}

func (p *ProductImpl) Update(product *datamodels.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProductImpl) SelectByKey(i int64) (*datamodels.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductImpl) SelectAll() ([]*datamodels.Product, error) {
	db := utils.GetConn()
	raws, err := sq.Select("*").From(TableName).RunWith(db).Query()
	if err != nil {
		return nil, err
	}

	var queryResult []*datamodels.Product
	for raws.Next() {
		var product *datamodels.Product
		err = raws.Scan(product)
		if err != nil {
			return nil, err
		}
		queryResult = append(queryResult, product)
	}
	return queryResult, nil
}
