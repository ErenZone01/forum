package serv

import (
	"net/http"

	"main.go/bd"
	"main.go/structs"
)

func Decon(w http.ResponseWriter, r *http.Request) {
	//Bloquer les methodes Get
	if r.Method == "GET"{
		Errors.Title = "Bad Request"
		Errors.Body = "405"
		errors(w, r, Errors)
		return
	}

	// Déconnectez l'utilisateur et supprimez la session de la base de données
	var users = Myaccount(w, r)
	if users.Error == "Veuillez vous connecté d'abord"{
		Errors.Title = "Bad Request"
		Errors.Body = "400"
		errors(w, r, Errors)
		return
	}
	user = bd.DataUser(BD, users.Email)
	session, _ := r.Cookie("session")
	if session != nil {
		sessionMutex.Lock()
		delete(sessions, user.Id)
		sessionMutex.Unlock()
	}
	

	// Supprimez le cookie de session
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	})

	user = structs.Users{}
	user.Error = "Vous êtes deconnecté avec succée"
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
