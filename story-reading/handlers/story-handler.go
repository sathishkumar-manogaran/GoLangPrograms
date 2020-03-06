package handlers

import (
	"github.com/sathishkumar-manogaran/GoLangPrograms/story-reading/models"
	"html/template"
	"net/http"
)

func NewHandler(s models.Story) handler {
	return handler{s}
}

type handler struct {
	s models.Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parsedTemplate := template.Must(template.New("").Parse(models.DefaultHandlerTemplate))
	err := parsedTemplate.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}

}
