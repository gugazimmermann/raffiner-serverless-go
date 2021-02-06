package types

var (
	ErrorInvalidID      = "invalid ID"
	ErrorInvalidData    = "invalid data"
	ErrorInvalidEmail   = "invalid email"
	ErrorCouldNotDelete = "could not delete"
)

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

type Address struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zipcode"`
}

type Client struct {
	ID      int     `json:"id,omitempty"`
	Email   string  `json:"email"`
	Name    string  `json:"Name"`
	Address Address `json:"address"`
	Phone   string  `json:"phone"`
}

type Product struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Supplier    string   `json:"supplier"`
	Tags        []string `json:"tags"`
	Category    string   `json:"category"`
	Subcategory string   `json:"subcategory,omitempty"`
	CostPrice   float32  `json:"costprice"`
	RentalPrice float32  `json:"rentalprice"`
	Quantity    int      `json:"quantity"`
}

type Supplier struct {
	ID      int     `json:"id,omitempty"`
	Email   string  `json:"email"`
	Name    string  `json:"Name"`
	Address Address `json:"address"`
	Phone   string  `json:"phone"`
}
