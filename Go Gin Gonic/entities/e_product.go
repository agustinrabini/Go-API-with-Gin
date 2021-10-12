package entities 

type Product struct{		//Para que los campos aparezcan en minúscula
	Id_product int 			`json:"id_product"`
	Name string				`json:"name"`
	Price int				`json:"price"`
	Stock int				`json:"stock"`
}

type Products [] Product