package domain

type Customer struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	Country Country `json:"country"`
	IsValid bool    `json:"valid"`
}

type PaginationInput struct {
	Limit  int
	Offset int
}

type FilterInput struct {
	Country *Country
	State   *bool
}
