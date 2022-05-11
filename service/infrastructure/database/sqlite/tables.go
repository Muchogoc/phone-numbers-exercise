package sqlite

type Customer struct {
	ID    int
	Name  string
	Phone string
}

func (c Customer) TableName() string {
	return "customer"
}
