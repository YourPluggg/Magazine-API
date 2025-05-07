package main

import "database/sql"

type Product struct {
	ID         int
	Definition string
	Name       string
	Price      float64
}
