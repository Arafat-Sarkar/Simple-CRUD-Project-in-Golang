package hendlers

import (
	"encoding/json"
	"net/http"
	"ecommerce/database"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)
	encoder.Encode(database.ProductList)
}
