package entities

type Order struct{		
	Id_order *int 						`json:"id_order"`
	Id_user *int						`json:"id_user"`
	Id_product *int						`json:"id_product"`
	Order_price *int					`json:"order_price"`
	Quantity_products *int				`json:"quantity_products"`
	State *int							`json:"state"`
}

type Orders [] Order