package entities

type Product struct {
	Id_product int    `json:"id_product"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
}

type Products []Product
