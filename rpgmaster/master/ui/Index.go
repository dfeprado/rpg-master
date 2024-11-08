package ui

import (
	"html/template"
	"net/http"

	"dfeprado.dev/rpg-master/rpgmaster"
)

type Layout struct {
	Title         string
	PlayerAddress string
}

var templates *template.Template = template.Must(
	template.ParseFiles(
		"rpgmaster/master/ui/index.html",
	),
)

var layout Layout = Layout{
	Title: "Master view - RPG-Master",
}

func Render(w http.ResponseWriter, r *http.Request) {
	if layout.PlayerAddress == "" {
		layout.PlayerAddress = "http://" + rpgmaster.GetApplication().JoinHostAndPort()
	}
	templates.Execute(w, layout)
}
