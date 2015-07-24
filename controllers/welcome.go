package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func Welcome() *welcomeController {
	return &welcomeController{
		templates: template.Must(template.ParseFiles("./templates/welcome/index.html.tpl")),
	}
}

type welcomeController struct {
	templates *template.Template
}

func (c *welcomeController) Index(w http.ResponseWriter, r *http.Request) {
	err := c.templates.ExecuteTemplate(w, "welcome/index", nil)

	if err != nil {
		log.Print(err)
	}
}
