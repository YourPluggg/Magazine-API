package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	//"errors"
	"os"   //Чтение запись
	"sort" //Сортировка слайсов
	// для мьютексов, пока не используется "sync"
)

var database *sql.DB

type ProductService struct {
	filePath string
	products map[int]Product
}

// Добавление товара (конструктор)
func NewProductService(path string) *ProductService {
	//Создаём новый экземпляр
	s := &ProductService{
		filePath: path,
		products: make(map[int]Product),
	}
	s.load()
	return s
}

func (s *ProductService) load() {
	file, err := os.ReadFile(s.filePath)
	if err != nil {
		return
	}

	var list []Product
	if err := json.Unmarshal(file, &list); err != nil {
		return
	}

	for _, p := range list {
		s.products[p.ID] = p
	}
}

func (s *ProductService) save() {
	var list []Product
	for _, p := range s.products {
		list = append(list, p)
	}

	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })
	data, err := json.MarshalIndent(list, "", " ")

	if err != nil {
		println("Ошибка при сериализации ", err.Error())
		return
	}

	err = os.WriteFile(s.filePath, data, 0644)

	if err != nil {
		println("Ошибка при записи", err.Error())
	}
}

func (s *ProductService) Add(p Product) Product {
	maxID := 0
	for id := range s.products {
		if id > maxID {
			maxID = id
		}
	}
	p.ID = maxID + 1
	s.products[p.ID] = p
	s.save()
	return p
}

func (s *ProductService) Remove(id int) (Product, error) {
	p, ok := s.products[id]
	if !ok {
		return Product{}, errors.New("Not found")
	}
	delete(s.products, id)
	s.save()
	return p, nil
}

func (s *ProductService) Edit(p Product) (Product, error) {
	if _, ok := s.products[p.ID]; !ok {
		return Product{}, errors.New("Not found")
	}
	s.products[p.ID] = p
	s.save()
	return p, nil
}

func (s *ProductService) Search(id int) (Product, error) {
	p, ok := s.products[id]
	if !ok {
		return Product{}, errors.New("Not found")
	}
	return p, nil
}

func (s *ProductService) GetAll() []Product {
	var list []Product
	for _, p := range s.products {
		list = append(list, p)
	}
	return list
}
