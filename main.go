package main

func main() {
	service := NewProductService("products.json")

	newProduct := Product{
		definition: "Новый телефон",
		name:       "Iphone 16 ",
		price:      15.000,
	}

	service.Add(newProduct)
}
