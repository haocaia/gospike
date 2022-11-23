package repository

import (
	sq "github.com/Masterminds/squirrel"
	"gospike/datamodels"
	"gospike/utils"
	"log"
	"strconv"
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

func (p *ProductImpl) Delete(productID int64) bool {
	db := utils.GetConn()
	result, err := sq.Delete(TableName).Where(sq.Eq{
		"ID": strconv.FormatInt(productID, 10),
	}).RunWith(db).Exec()
	if err != nil {
		log.Panicf("Delete Product(id: %v) error: %v", productID, err.Error())
		return false
	}
	_, err = result.RowsAffected()
	if err != nil {
		log.Panicf("Delete Product(id: %v) error: %v", productID, err.Error())
		return false
	}

	return true

}

func (p *ProductImpl) Update(product *datamodels.Product) error {
	db := utils.GetConn()

	_, err := sq.Update(TableName).SetMap(map[string]interface{}{
		"productName":  product.ProductName,
		"productNum":   product.ProductNum,
		"productImage": product.ProductImage,
	}).Where(sq.Eq{
		"ID": strconv.FormatInt(product.ID, 10),
	}).RunWith(db).Exec()

	if err != nil {
		return err
	}
	return nil
}

func (p *ProductImpl) SelectByKey(productID int64) (*datamodels.Product, error) {
	db := utils.GetConn()
	raws, err := sq.Select("*").From(TableName).Where(sq.Eq{
		"ID": productID,
	}).RunWith(db).Query()
	if err != nil {
		return nil, err
	}

	var product = datamodels.Product{}
	for raws.Next() {
		err = raws.Scan(&product.ID, &product.ProductName, &product.ProductNum, &product.ProductImage, &product.ProductUrl)
		if err != nil {
			return nil, err
		}
	}
	return &product, nil
}

func (p *ProductImpl) SelectAll() ([]*datamodels.Product, error) {
	db := utils.GetConn()
	raws, err := sq.Select("*").From(TableName).RunWith(db).Query()
	if err != nil {
		log.Printf("product select all sql query error: %v", err.Error())
		return nil, err
	}

	var queryResult = make([]*datamodels.Product, 0)
	for raws.Next() {
		var product = datamodels.Product{}
		err = raws.Scan(&product.ID, &product.ProductName, &product.ProductNum, &product.ProductImage, &product.ProductUrl)
		if err != nil {
			log.Printf("product select all scan error: %v", err.Error())
			return nil, err
		}
		queryResult = append(queryResult, &product)
	}
	return queryResult, nil
}
