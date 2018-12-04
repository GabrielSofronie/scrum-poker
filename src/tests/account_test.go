package tests

import (
	"entities"
	"services/account"
	"testing"
)

func Test_UserCreatesAccount(t *testing.T) {
	setup()

	testUser := entities.NewDeveloper("test-user")

	accountMaker := account.NewMaker(userRepository)
	model := account.CreateModel{
		Name: testUser.Name(),
		Role: testUser.Role(),
	}
	response := accountMaker.Execute(model)

	if response.Err != nil {
		t.Errorf(response.Err.Error())
	}

	if response.Name != testUser.Name() {
		t.Errorf("Expected to have name: [%s] , got [%s] ", testUser.Name(), response.Name)
	}
}
