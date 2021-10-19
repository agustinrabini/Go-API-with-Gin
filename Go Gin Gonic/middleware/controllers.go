package middleware

import (
	"fmt"
	"goGin/actions"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogOut(c *gin.Context) {

	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, "Succes LogOut")
}

func LogIn(c *gin.Context) {

	var userId int

	name := c.DefaultQuery("name", "user")
	phone := c.Query("phone")

	dbProducts := actions.DbConecction()

	result, err := dbProducts.Query("select id_user from user_go where name = ? and phone=?", name, phone)
	if err != nil {
		log.Fatal(err.Error())
	}

	for result.Next() {
		err := result.Scan(&userId)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	defer dbProducts.Close()
	defer result.Close()

	if userId != 0 {

		session := sessions.Default(c)
		session.Set("id_user", userId)

		session.Save()

		c.JSON(http.StatusOK, "Succes Loggin")

	} else {
		c.JSON(http.StatusNotFound, "Unauthorized")
		c.Abort()
	}

	fmt.Println(userId)
}
