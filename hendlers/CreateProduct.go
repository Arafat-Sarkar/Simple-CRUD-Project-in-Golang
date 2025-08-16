package hendlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ecommerce/database"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	/*
		     1. take body information( description, imageurl,price, title ) from r.body
			 2. create the instance using product struct with the body information
			 3. append the instance into productlist
	*/

	var newProduct  database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please Give me a Valid json", 400)
		return
	}

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)
	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)
}
