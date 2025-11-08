package enums

type TodoStatus string

const (
	PENDING   TodoStatus = "PENDING"
	COMPLETED TodoStatus = "COMPLETED"
	CANCELLED TodoStatus = "CANCELLED"
)
