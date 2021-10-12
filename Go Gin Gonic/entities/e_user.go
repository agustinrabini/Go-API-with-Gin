package entities 

type User struct{		
	Id_user       *int          	`json:"id_user"`
    Name		*string       		`json:"name"`
    Email         *string        	`json:"email"`
    Phone         *int     			`json:"phone"`	
}

type Users [] User