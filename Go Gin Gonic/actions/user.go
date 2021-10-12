package actions

import (
  "github.com/gin-gonic/gin"
  "goGin/entities"
  "net/http"
)

func GetUserByNameAndPhone(c *gin.Context) {

  var userData entities.User

  name := c.DefaultQuery("name", "user")
  phone := c.Query("phone")

  dbProducts:= DbConecction()
  
	result, err := dbProducts.Query("select id_user,name,email,phone from user_go where name = ? and phone=?", name, phone)
  if err != nil {
   //manejar errores
  }
 
  for result.Next() {
  
    err := result.Scan(&userData.Id_user, &userData.Name, &userData.Email, &userData.Phone)
    if err != nil {
        //manejar errores
    }
  }
  
  defer dbProducts.Close()
  defer result.Close() 
  
  c.JSON(http.StatusOK,userData)
}