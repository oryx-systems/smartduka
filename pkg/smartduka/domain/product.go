package domain

// Product is used to display product info
type Product struct {
	ID           string `json:"id"`
	Active       bool   `json:"active"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	Quantity     string `json:"quantity"`
	Unit         string `json:"unit"`
	Price        string `json:"price"`
	Description  string `json:"description"`
	Manufacturer string `json:"manufacturer"`
	InStock      bool   `json:"inStock"`
}

// Sale is used to show sales data
type Sale struct {
	ID          string `json:"id"`
	ProductName string `json:"productName"`
	Quantity    string `json:"quantity"`
	Unit        string `json:"unit"`
	Price       string `json:"price"`
	SoldBy      string `json:"soldBy"`
}
