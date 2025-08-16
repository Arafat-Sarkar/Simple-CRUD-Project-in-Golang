package hendlers

import (
	"ecommerce/database"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	pid, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Please give me valid id", 400)
		return
	}

	for _, product := range database.ProductList {

		if product.ID == pid {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}

	}
	http.Error(w, "Product not found", http.StatusNotFound)
}
