package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"services/account"

	"github.com/gorilla/mux"
)

// AccountHandler interface
type AccountHandler interface {
	RegisterHandlers()
}

type accountMaker struct {
	*mux.Router
	account.Maker
}

type createAccountView struct {
	ID     string
	Name   string
	Errors string
}

// NewAccountHandler returns an AccountHandler
func NewAccountHandler(r *mux.Router, service account.Maker) AccountHandler {
	return &accountMaker{r, service}
}

func (a *accountMaker) RegisterHandlers() {
	a.Handle("/v1/accounts", a.createAccount()).Methods("POST", "OPTIONS")
}

func (a accountMaker) createAccount() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error creating account"
		var model account.CreateModel

		err := json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		res := a.Maker.Execute(model)
		view := generateCreateAccountView(res)
		if view.Errors != "" {
			log.Printf("%s - %s", err.Error(), view.Errors)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(view); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func generateCreateAccountView(response account.CreateResponse) createAccountView {
	errors := ""
	if response.Err != nil {
		errors = response.Err.Error()
	}
	return createAccountView{
		ID:     response.ID,
		Name:   response.Name,
		Errors: errors,
	}
}
