package model

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	LastName   string `json:"lastname"`
	DocumentNumber   string `json:"documentNumber"`
	DocumentType string `json:"documentType"`
}

func (u *User) TableName() string {
	return "citizen"
}
