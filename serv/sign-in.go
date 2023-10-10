package serv

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"main.go/bd"
	"main.go/structs"
)

func SignIn(db *sql.DB, w http.ResponseWriter, r *http.Request) structs.Users {
	//initialisation des variables
	var Allusers, _ = bd.DataAllUser(BD)
	var Mdp = r.FormValue("Mdp")
	var Email = r.FormValue("Email")
	var Newuser structs.Users
	Newuser.Mdp = Mdp
	Newuser.Email = Email
	Newuser.Role = "utilisateur"
	Newuser.Error = ""

	//verification de si l'email ou le password eest vide
	if !IsValid(Newuser.Email) {
		Newuser.Error = "Your Email is incorrect"
		return Newuser
	} 
	if !IsValid(Newuser.Mdp) {
		Newuser.Error = "Your password must have 1 character different to space"
		return Newuser
	}

	//verification si l'utilisateur n'est pas deja enregistré
	for _, v := range Allusers {
		auth := bcrypt.CompareHashAndPassword([]byte(v.Mdp), []byte(Newuser.Mdp))
		if v.Email == Newuser.Email && auth == nil {
			presentUser := bd.DataUser(db, Newuser.Email)
			actif := session(w, r, presentUser, db)
			if !actif {
				Newuser = structs.Users{}
				Newuser.Error = "L'utilisateur est déjà connecté depuis un autre endroit."
				return Newuser
			}
			return presentUser
		}
	}
	//sinon retourner une structure vide avec un message d'erreur
	Newuser = structs.Users{}
	Newuser.Error = "User not found"
	return Newuser
}
