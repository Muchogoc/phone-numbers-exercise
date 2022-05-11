package sqlite

// Customer ...
type Customer struct {
	ID    int
	Name  string
	Phone string
}

// TableName ...
func (c Customer) TableName() string {
	return "customer"
}
