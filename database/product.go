package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURl      string  `json:"imageUrl"`
}


func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Tangy and rich in Vitamin C — a juicy boost to your day.",
		Price:       100,
		ImgURl:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcRwJf4xPOftZLLgkWjr2eMMumu9XuRdKiGc1eZXFku9OmA4lSymIPBm_vU0bFIM_BjpYOsI_pX7O9c9sRUxfkXxE0N1x_bWSERw7SXHP3A",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Crisp and fresh with a natural sweetness — great for your health",
		Price:       200,
		ImgURl:      "https://static.libertyprim.com/files/varietes/pomme-pink-lady-large.jpg?1569321939",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Mango",
		Description: "Juicy, sweet, and full of tropical flavor — the king of fruits.",
		Price:       120,
		ImgURl:      "https://belescooverseas.com/wp-content/uploads/2024/04/Mango.jpeg",
	}
	prd4 := Product{
		ID:          4,
		Title:       "Watermelon",
		Description: "JRefreshing and hydrating summer fruit, perfect for hot days.",
		Price:       400,
		ImgURl:      "https://images.squarespace-cdn.com/content/v1/5f56423f4aca615934476295/523d9dc5-b24e-404e-b8f4-f2c4a29b6ad2/Ancient-Egyptians-enjoyed-watermelon-as-far-back-as-3500-years-ago.jpg",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Banana",
		Description: "Soft, energy-rich fruit — perfect for a quick snack anytime.",
		Price:       30,
		ImgURl:      "https://www.cuisinelangelique.com/infotheque/wp-content/uploads/2023/03/banane-1a-1200x838.jpg",
	}

	prd6 := Product{
		ID:          6,
		Title:       "Strawberry",
		Description: "Sweet, tangy, and full of antioxidants — a delicious treat for all ages.",
		Price:       500,
		ImgURl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRlksT4qy5XemFBOZJxNsPHqz_KdkZBlRw32g&s",
	}

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
	ProductList = append(ProductList, prd4)
	ProductList = append(ProductList, prd5)
	ProductList = append(ProductList, prd6)

}
