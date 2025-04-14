package main

import ()

type ProductService struct {
	FilePath string
	products map[int]product
}

// Добавление товара
func NewProductServise(FilePath) *ProductService {
	s := ProductService{
		FilePath: path,
		products: make(map[int]product),
	}
	s.load()
	return s
}
