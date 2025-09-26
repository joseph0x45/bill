package main

type Product struct {
	ID    string `json:"id" db:"id"`
	Label string `json:"label" db:"label"`
	Price int    `json:"price" db:"price"`
}
