package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"services/stories"

	"github.com/gorilla/mux"
)

// StoryHandler interface
type StoryHandler interface {
	RegisterHandlers()
}

type story struct {
	*mux.Router
	stories.Creator
	stories.Lister
}

type createStoryView struct {
	Title  string
	Errors string
}

type listStoriesView struct {
	Titles []string
	Errors string
}

// NewStoryHandler returns a QuestionHandler
func NewStoryHandler(r *mux.Router, creator stories.Creator, lister stories.Lister) StoryHandler {
	return &story{r, creator, lister}
}

func (s *story) RegisterHandlers() {
	s.Handle("/v1/stories", s.create()).Methods("POST", "OPTIONS")
	s.Handle("/v1/stories", s.list()).Methods("GET", "OPTIONS")
}

func (s story) create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error creating story"
		var model stories.StoryModel

		err := json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		res := s.Creator.Execute(model)
		view := generateCreateStoryView(res)
		if view.Errors != "" {
			log.Printf("%s", view.Errors)
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

func (s story) list() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error retrieving stories"
		res := s.Lister.Execute()
		view := generateListStoriesView(res)
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

func generateCreateStoryView(response stories.CreateResponse) createStoryView {
	errors := ""
	if response.Err != nil {
		errors = response.Err.Error()
	}
	return createStoryView{
		Title:  response.Title,
		Errors: errors,
	}
}

func generateListStoriesView(response stories.ListResponse) listStoriesView {
	errors := ""
	if response.Err != nil {
		errors = response.Err.Error()
	}
	return listStoriesView{
		Titles: response.Titles,
		Errors: errors,
	}
}
