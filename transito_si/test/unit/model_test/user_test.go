package model_test

import (
	"testing"

	"unal.edu.co/rest-example/api/model"
)

func Test_UserTableName(t *testing.T) {
	user := &model.User{}
	expect := "users"
	got := user.TableName()

	if user.TableName() != expect {
		t.Errorf("UserTableName failed, expected %s, got %s", expect, got)
		return
	}

	t.Logf("UserTableName success, expected %s, got %s", expect, got)
}
