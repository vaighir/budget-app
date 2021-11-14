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
	render.RenderTemplate(w, "register.page.tmpl", &models.TemplateData{})
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

	userCheck, userMsg := helpers.CheckUsername(username)

	if !userCheck {
		w.Write([]byte(userMsg))
		return
	}

	pwdCheck, pwdMsg := helpers.CheckPassword(password, passwordRepeat)

	if !pwdCheck {
		w.Write([]byte(pwdMsg))
		return
	}

	user := models.User{
		Username: username,
		Password: password,
	}

	db_helpers.CreateUser(user)

	w.Write([]byte("<h1>User created</h1>"))
	// TODO add redirection to login page
}

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.page.tmpl", &models.TemplateData{})
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
		w.Write([]byte("<h1>Wrong username or password</hq>"))
		// TODO redirect to login page
	} else {
		w.Write([]byte("<h1>You're logged in</hq>"))
		// TODO save user to session
		// TODO redirect to home page
	}

}
