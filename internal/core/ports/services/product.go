package services

// todo: must be domain
type ProductDTO struct {
	Id          string  `json:"id"`
	Description string  `json:"description"`
	Ingredients string  `json:"ingredients"`
	Price       float64 `json:"price"`
	IDCategory  string  `json:"category_id"`
	Category    string  `json:"category"`
}

type Product interface {
	// todo: must be domain
	GetProductById(id string) (*ProductDTO, error)
	// todo: must be domain
	GetProducts(description string) ([]*ProductDTO, error)
	// todo: must be domain
	CreateProduct(productDTO *ProductDTO) (*ProductDTO, error)
	GetCategoryID(categoryName string) (string, error)
}
