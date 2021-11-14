package handlers

import (
	"log"
	"net/http"

	"github.com/vaighir/go-diet/app/pkg/db_helpers"
	"github.com/vaighir/go-diet/app/pkg/helpers"
	"github.com/vaighir/go-diet/app/pkg/models"
	"github.com/vaighir/go-diet/app/pkg/render"
	"golang.org/x/crypto/bcrypt"
)

func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)

	render.RenderTemplate(w, "register.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func Register(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	passwordRepeat := r.Form.Get("password-repeat")

	//	TODO prohibit creating a user if username already exists
	if helpers.CheckIfUserExists(username) {
		app.Session.Put(r.Context(), "warning", "User with this username already exists")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	userCheck, userMsg := helpers.CheckUsername(username)

	if !userCheck {
		log.Println("Bad username")

		app.Session.Put(r.Context(), "warning", userMsg)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	pwdCheck, pwdMsg := helpers.CheckPassword(password, passwordRepeat)

	if !pwdCheck {
		app.Session.Put(r.Context(), "warning", pwdMsg)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	user := models.User{
		Username: username,
		Password: password,
	}

	db_helpers.CreateUser(user)

	log.Printf("Created user: %s", username)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	getSessionMsg(r, stringMap)
	render.RenderTemplate(w, "login.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	user := db_helpers.GetUserByUsername(username)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err)
		app.Session.Put(r.Context(), "warning", "Wrong username or password")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	} else {
		app.Session.Put(r.Context(), "user_id", user.Id)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	app.Session.Clear(r.Context())
	app.Session.Put(r.Context(), "info", "You have been logged out")

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
