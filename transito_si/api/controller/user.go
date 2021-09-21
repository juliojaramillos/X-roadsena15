package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"unal.edu.co/rest-example/api/dto"
	"unal.edu.co/rest-example/api/model"
	"unal.edu.co/rest-example/api/repository"
)

var repo repository.UserRepository

func InitRepository(r repository.UserRepository) {
	repo = r
	PopulateDatabase(repo)
}

func GETUser(c *gin.Context) {
	var document = c.Query("document")
	var docType = c.Query("type")
	var user = dto.SearchUser{DocumentNumber: document, DocumentType: docType}

	users, err := repo.SearchUser(user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var responseUser = dto.ResponseUser{HasRegistry: false}

	if users != nil {
		responseUser = dto.ResponseUser{HasRegistry: true,RegistryID: users[0].Id}
	}



	c.JSON(http.StatusOK, responseUser)
}

func POSTUser(c *gin.Context) {
	var user model.User
	println(c.Request.Body)
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	user, err = repo.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user on database:\n%v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	createdUser := dto.CreatedUser{RuntID: user.Id}


	c.JSON(http.StatusCreated, createdUser)

}

func PopulateDatabase(r repository.UserRepository) {
	var user1 = model.User{Name: "Salvador",LastName: "Pastor Espinoza",DocumentNumber: "11036516300",DocumentType: "CC"}
	var user2 = model.User{Name: "Tomás",LastName: "Pérez Gil",DocumentNumber: "15129143371",DocumentType: "CC"}
	var user3 = model.User{Name: "Iván",LastName: "Rodríguez Tovar",DocumentNumber: "16337688",DocumentType: "TI"}
	var user4 = model.User{Name: "Nancy",LastName: "Tapia Árias",DocumentNumber: "10900945",DocumentType: "CE"}
	var user5 = model.User{Name: "Juan Pedro",LastName: "Sosa Vaquero",DocumentNumber: "11773819393",DocumentType: "NUIP"}

	r.CreateTable()
	r.CreateUser(user1)
	r.CreateUser(user2)
	r.CreateUser(user3)
	r.CreateUser(user4)
	r.CreateUser(user5)
}
