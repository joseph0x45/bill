package main

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(
	db *sqlx.DB,
) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) InsertProduct() error {
	return nil
}

func (r *Repository) GetAllProducts() error {
	return nil
}

func (r *Repository) UpdateProduct() error

func (r *Repository) DeleteProduct() error
