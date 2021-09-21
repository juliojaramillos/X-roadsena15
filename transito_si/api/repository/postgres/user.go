package postgres

import (
	"database/sql"
	"log"
	"unal.edu.co/rest-example/api/dto"

	"unal.edu.co/rest-example/api/model"
)

type Repo struct {
	DB *sql.DB
}

func (ur *Repo) SearchUser(user dto.SearchUser) ([]model.User, error) {
	println("Ingreso al repo")
	println(user.DocumentNumber)
	println(user.DocumentType)
	rows, err := ur.DB.Query("SELECT * FROM citizen WHERE document_number=$1 and document_type= $2",user.DocumentNumber,user.DocumentType)
	if err != nil {
		return nil, err
	}

	var users []model.User
	var userCounter int
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.LastName, &user.DocumentNumber, &user.DocumentType)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
		userCounter++
	}

	log.Printf("Got %d rows\n", userCounter)
	return users, nil
}

func (ur *Repo) CreateUser(user model.User) (model.User, error) {

	sqlStatement := `INSERT INTO citizen (name, lastname, document_number, document_type) VALUES ($1, $2, $3, $4) RETURNING id`
	id := 0

	var err = ur.DB.QueryRow(sqlStatement, user.Name, user.LastName, user.DocumentNumber, user.DocumentType).Scan(&id)

	if err != nil {
		return model.User{}, err
	}

	log.Printf("%d Citizen id \n", id)

	user.Id = uint(id)

	return user, nil
}

func (ur *Repo) CreateTable() () {

	sqlStatement := `CREATE TABLE citizen (id serial PRIMARY KEY,document_number VARCHAR ( 50 ) UNIQUE NOT NULL,document_type VARCHAR ( 50 ) NOT NULL,name VARCHAR ( 255 ) NOT NULL,lastname VARCHAR ( 255 ) NOT NULL);
`

	var err = ur.DB.QueryRow(sqlStatement)

	if err != nil {
		return
	}

	log.Printf("Created citizen table \n")

}
