package controller_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"unal.edu.co/rest-example/api/controller"
	"unal.edu.co/rest-example/api/model"
)

type mockRepo struct {
	shouldPass bool
}

var testUsers = []model.User{
	{
		Id:      1,
		Name:    "test name",
		Email:   "test@email.com",
		Phone:   "123123123",
		Address: "street 1",
	},
	{
		Id:      2,
		Name:    "test name 2",
		Email:   "test2@email.com",
		Phone:   "321321321",
		Address: "street 2",
	},
}

func (m mockRepo) ListUsers() ([]model.User, error) {
	if m.shouldPass {
		return testUsers, nil
	} else {
		return nil, errors.New("error on database")
	}
}

func (m mockRepo) CreateUser(user model.User) (model.User, error) {
	panic("implement me")
}

func Test_GETUsers(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	controller.InitRepository(mockRepo{
		shouldPass: true,
	})

	var got []model.User
	controller.GETUsers(c)
	if w.Code != http.StatusOK {
		log.Printf("GETUsers failed, expected %v got %v\n", http.StatusOK, w.Code)
		t.Error()
		return
	}

	body, _ := ioutil.ReadAll(w.Body)
	err := json.Unmarshal(body, &got)
	if err != nil {
		log.Printf("Test failed due bad decoding\n")
		t.Error()
		return
	}

	log.Printf("GETUsers succeed, expected %v got %v\n", []interface{}{testUsers, http.StatusOK}, []interface{}{got, w.Code})
}
