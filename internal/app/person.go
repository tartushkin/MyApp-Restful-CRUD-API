package app

// Модель
type Person struct {
	ID        int    `json:"id"        db:"id"`
	Email     string `json:"email"     db:"email"`
	Phone     string `json:"phone"     db:"phone"`
	FirstName string `json:"firstName" db:"firstName"`
	LastName  string `json:"lastName"  db:"lastName"`
}
