package repo

import "product/repository"

type mysqlProduct struct {
	repo repository.ProductRepo
}

func NewProductRepo() repository.ProductRepo {
	return &mysqlProduct{}
}

func (p *mysqlProduct) Create() {

}

func (p *mysqlProduct) Get() {

}
