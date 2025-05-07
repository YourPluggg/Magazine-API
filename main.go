package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/productdb")

	if err != nil {
		log.Println(err)
	}
	database = db

	service := NewProductService("products.json")

	// Обработчик функции для нашего пути
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Получить все товары
			products := service.GetAll()
			json.NewEncoder(w).Encode(products)

		case http.MethodPost:
			// Добавить товар
			var p Product
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newProduct := service.Add(p)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newProduct)
		}
	})

	//Метод для работы по пути + ID
	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/products/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			// Поиск товара
			p, err := service.Search(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(p)

		case http.MethodPut:
			// Обновляем товар
			var p Product
			if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			p.ID = id
			updated, err := service.Edit(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(updated)

		case http.MethodDelete:
			// Удаляем товар
			deleted, err := service.Remove(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(deleted)
		}
	})

	// Запуск сервера
	http.ListenAndServe(":8080", nil)
}
