package ui

import (
	"html/template"
	"net/http"

	"dfeprado.dev/rpg-master/rpgmaster"
	"dfeprado.dev/rpg-master/rpgmaster/master/ui/components"
)

type Layout struct {
	Title         string
	PlayerAddress string
	Navigator     []components.MenuButton
}

var layout Layout = Layout{
	Title: "Master view - RPG-Master",
	Navigator: []components.MenuButton{
		{Id: "items", Name: "Items", Icon: "category"},
		{Name: "Players", Icon: "group"},
	},
}

func Render(w http.ResponseWriter, r *http.Request) {
	if layout.PlayerAddress == "" {
		layout.PlayerAddress = "http://" + rpgmaster.GetApplication().JoinHostAndPort()
	}

	// TODO move to global declaration after dev
	var templates *template.Template = template.Must(
		template.ParseFiles(
			"rpgmaster/master/ui/index.html",
			"rpgmaster/master/ui/components/menu-button.html",
		),
	)
	templates.Execute(w, layout)
}
