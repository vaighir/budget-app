package handlers

import (
	"net/http"

	"github.com/vaighir/go-diet/app/pkg/db_helpers"
	"github.com/vaighir/go-diet/app/pkg/models"
	"github.com/vaighir/go-diet/app/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {

	var user = db_helpers.GetUserById(1)

	stringMap := make(map[string]string)
	stringMap["username"] = user.Username

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
