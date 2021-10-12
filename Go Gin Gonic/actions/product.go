package actions

import (
	"encoding/json"
	"fmt"
	"goGin/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProductsList(c *gin.Context) {

	var products []entities.Product
	var product entities.Product

	dbProducts := DbConecction()

	result, err := dbProducts.Query("select * from product_go")
	if err != nil {
		//manejar errores
		fmt.Println("Error 1", err.Error())
		return
	}

	for result.Next() {

		err := result.Scan(&product.Id_product, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			//manejar errores
			fmt.Println("Error 2")
			return
		}

		products = append(products, product)
	}

	defer result.Close()
	defer dbProducts.Close()

	c.JSON(http.StatusOK, products)

}

//CRUDE
func ShowProduct(c *gin.Context) {

	var product entities.Product

	productId := c.Param("id")

	dbProducts := DbConecction()

	result, err := dbProducts.Query("select id_product,name,price,stock from product_go where id_product = ?", productId)
	if err != nil {
		//manejar errores
	}

	for result.Next() {

		err := result.Scan(&product.Id_product, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			//manejar errores
		}
	}

	defer dbProducts.Close()
	defer result.Close()

	c.JSON(http.StatusOK, product)
}

//CRUDE
func ProductAdd(c *gin.Context) {

	var productData entities.Product

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&productData)
	if err != nil {
		//manejar errores
		fmt.Println("Error 1", err.Error())
		return
	}

	dbProducts := DbConecction()

	stmt, err := dbProducts.Prepare("insert into product_go (id_product, name, price, stock) values (?,?,?,?)")
	if err != nil {
		//manejar errores
		fmt.Println("Error 2", err.Error())
		return
	}

	_, err = stmt.Exec(nil, productData.Name, productData.Price, productData.Stock)
	if err != nil {
		//manejar errores
		fmt.Println("Error 3", err.Error())
		return
	}

	defer dbProducts.Close()
	defer c.Request.Body.Close()
	defer stmt.Close()

	c.JSON(http.StatusOK, nil)
}

//
////CRUDE
//func ProductUpdate(w http.ResponseWriter, r *http.Request) {
//
//  var productData entities.Product
//
//  params := mux.Vars(r)
//  productId := params["id"]
//
//  dbProducts:= DbConecction()
//
//  decoder :=  json.NewDecoder(r.Body)
//  err := decoder.Decode(&productData)
//  if (err!=nil) {
//     //manejar errores
//      w.WriteHeader(500)
//      return
//  }
//
//  stmt, err := dbProducts.Prepare("update product_go set id_product=?, name=?, stock=? where id_product = ? ")
//  if err != nil {
//      //manejar errores
//      w.WriteHeader(500)
//      return
//  }
//
//   _, err = stmt.Exec(productData.Id_product, productData.Name, productData.Price, productData.Stock, productId)
//  if err != nil {
//      //manejar errores
//      w.WriteHeader(500)
//      return
//  }
//
//  defer stmt.Close()
//  defer r.Body.Close()
//  defer dbProducts.Close()
//
//  w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(200)
//  json.NewEncoder(w).Encode(productData)
//}
//
////CRUDE
//func ProductRemove(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "application/json")
//
//  params := mux.Vars(r)
//  productId := params["id"]
//
//  dbProducts:= DbConecction()
//
//  stmt, err := dbProducts.Prepare("delete from product where id_product = ?")
//  if err != nil {
//        //manejar errores
//      w.WriteHeader(500)
//      return
//  }
//
//  _, err = stmt.Exec(productId)
//  if err != nil {
//      //manejar errores
//      w.WriteHeader(500)
//      return
//  }
//
//  defer stmt.Close()
//  defer dbProducts.Close()
//  defer r.Body.Close()
//
//  w.WriteHeader(200)
//}
