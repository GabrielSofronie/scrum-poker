package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"services/estimates"

	"github.com/gorilla/mux"
)

// EstimateHandler interface
type EstimateHandler interface {
	RegisterHandlers()
}

type estimator struct {
	*mux.Router
	estimates.Estimator
	estimates.Changer
	estimates.Lister
}

type estimateView struct {
	Developer string
	Story     string
	Estimate  int
	Errors    string
}

type estimateChangeView struct {
	Developer string
	Story     string
	Estimate  int
	Errors    string
}

type estimateListView struct {
	Estimates map[string]int
	Errors    string
}

// NewEstimateHandler returns an EstimateHandler
func NewEstimateHandler(r *mux.Router, e estimates.Estimator, c estimates.Changer, l estimates.Lister) EstimateHandler {
	return &estimator{r, e, c, l}
}

func (e *estimator) RegisterHandlers() {
	e.Handle("/v1/stories/{sid}/estimates", e.estimate()).Methods("POST", "OPTIONS")
	e.Handle("/v1/stories/{sid}/estimates", e.change()).Methods("PUT", "OPTIONS")
	e.Handle("/v1/stories/{sid}/estimates", e.list()).Methods("GET", "OPTIONS")
}

func (e estimator) estimate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error estimating story"
		var model estimates.EstimateModel

		err := json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		model.StoryID = mux.Vars(r)["sid"]

		res := e.Estimator.Execute(model)
		view := generateEstimateView(res)
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

func (e estimator) change() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error changing story estimation"
		var model estimates.ChangeModel

		err := json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		model.StoryID = mux.Vars(r)["sid"]

		res := e.Changer.Execute(model)
		view := generateEstimateChangeView(res)
		if view.Errors != "" {
			log.Printf("%s", view.Errors)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(view); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func (e estimator) list() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error retrieving story estimates"
		model := estimates.ListModel{
			StoryID: mux.Vars(r)["sid"],
		}

		res := e.Lister.Execute(model)
		view := generateEstimateListView(res)
		if view.Errors != "" {
			log.Printf("%s", view.Errors)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(view); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func generateEstimateView(response estimates.EstimateResponse) estimateView {
	errors := ""
	if response.Err != nil {
		errors = response.Err.Error()
	}
	return estimateView{
		Developer: response.Developer,
		Story:     response.Story,
		Estimate:  response.Estimate,
		Errors:    errors,
	}
}

func generateEstimateChangeView(response estimates.ChangeResponse) estimateChangeView {
	errors := ""
	if response.Err != nil {
		errors = response.Err.Error()
	}
	return estimateChangeView{
		Developer: response.Developer,
		Story:     response.Story,
		Estimate:  response.Estimate,
		Errors:    errors,
	}
}

func generateEstimateListView(response estimates.ListResponse) estimateListView {
	errors := ""
	if response.Err != nil {
		errors = response.Err.Error()
	}
	return estimateListView{
		Estimates: response.Estimates,
		Errors:    errors,
	}
}
