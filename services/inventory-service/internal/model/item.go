package model

type Item struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Stock int `json:"stock"`
	Price string `json:"price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}