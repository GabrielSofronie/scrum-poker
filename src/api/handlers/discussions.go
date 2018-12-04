package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"services/discussions"

	"github.com/gorilla/mux"
)

// DiscussionHandler interface
type DiscussionHandler interface {
	RegisterHandlers()
}

type discussionView struct {
	Developer string
	Story     string
	Question  string
	Errors    string
}

// AnswerView model
type answerView struct {
	Owner    string
	Question string
	Answer   string
	Errors   string
}

type discuss struct {
	*mux.Router
	discussions.Inquirer
	discussions.Answer
}

// NewDiscussionHandler returns a DiscussionHandler
func NewDiscussionHandler(r *mux.Router, inquirer discussions.Inquirer, answer discussions.Answer) DiscussionHandler {
	return &discuss{r, inquirer, answer}
}

func (d *discuss) RegisterHandlers() {
	d.Handle("/v1/stories/{sid}/comment", d.comment()).Methods("POST", "OPTIONS")
	d.Handle("/v1/stories/{sid}/comment", d.reply()).Methods("PUT", "OPTIONS")
}

func (d discuss) comment() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error asking question"
		var model discussions.QuestionModel

		err := json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		model.StoryID = mux.Vars(r)["sid"]

		res := d.Inquirer.Execute(model)
		view := generateQuestionView(res)
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

func (d discuss) reply() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error answering question"
		var model discussions.AnswerModel

		err := json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		model.StoryID = mux.Vars(r)["sid"]

		res := d.Answer.Execute(model)
		view := generateAnswerView(res)
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

func generateQuestionView(response discussions.QuestionResponse) discussionView {
	errors := ""
	if response.Err != nil {
		errors = response.Err.Error()
	}
	return discussionView{
		Developer: response.Developer,
		Story:     response.Story,
		Question:  response.Question,
		Errors:    errors,
	}
}

func generateAnswerView(response discussions.AnswerResponse) answerView {
	errors := ""
	if response.Err != nil {
		errors = response.Err.Error()
	}
	return answerView{
		Owner:    response.Owner,
		Question: response.Question,
		Answer:   response.Answer,
		Errors:   errors,
	}
}
