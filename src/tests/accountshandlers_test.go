package tests

import (
	"api/handlers"
	"entities"
	"fmt"
	"net/http"
	"net/http/httptest"
	"services/account"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateAccountHandler(t *testing.T) {
	setup()

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Add test developer and story
	testDeveloper := entities.NewDeveloper("test-dev")
	userRepository.Create(testDeveloper)

	accountMaker := account.NewMaker(userRepository)
	accountHdlr := handlers.NewAccountHandler(router, accountMaker)
	accountHdlr.RegisterHandlers()

	payload := fmt.Sprintf(`{
		"Name": "%s",
		"Role": %d
	}`, testDeveloper.Name(), testDeveloper.Role())

	req, err := http.NewRequest("POST", "/v1/accounts", strings.NewReader(payload))
	if err != nil {
		t.Error(err)
	}

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected to have HTTP Code [%d] , got [%d] ", http.StatusCreated, recorder.Code)
	}
}
