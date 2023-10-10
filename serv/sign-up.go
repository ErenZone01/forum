package serv

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"main.go/bd"
	"main.go/structs"
)

func SignUp(db *sql.DB, w http.ResponseWriter, r *http.Request) structs.Users {
	//recuperer le user
	var Allusers, _ = bd.DataAllUser(db)
	var Username = r.FormValue("Username")
	var Mdp = r.FormValue("Mdp")
	var Email = r.FormValue("Email")
	var CMdp = r.FormValue("confirmMdp")
	var Newuser structs.Users

	var power = IsValidMail(Email)
	if !power {
		Newuser.Error = "Your Email is incorrect"
		return Newuser
	}

	if !IsValid(Mdp) {
		Newuser.Error = "Your password must have 1 character different to space"
		return Newuser
	}

	if Mdp != CMdp {
		Newuser.Error = "The password does not match the password confirmation"
		return Newuser
	}

	//hashage de mdp
	passe := []byte(Mdp)
	cons := 10
	HashPass, err := bcrypt.GenerateFromPassword(passe, cons)
	if err != nil {
		fmt.Println("Erreur lors de la génération du hachage:", err)
		os.Exit(0)
	}
	hash := string(HashPass)

	//initialisation

	Newuser.Username = Username
	Newuser.Mdp = hash
	Newuser.Email = Email
	Newuser.Error = ""

	//verification de si l'email ou le password eest vide
	if !IsValid(Newuser.Email) {
		Newuser.Error = "Your Email is incorrect"
		return Newuser
	} else if !IsValid(Newuser.Mdp) {
		Newuser.Error = "Your password must have 1 character different to space"
		return Newuser
	} else if !IsValid(Newuser.Username) {
		Newuser.Error = "Your username must have 1 character different to space"
		return Newuser
	}

	if Mdp != CMdp { //confirmation de password
		Newuser = structs.Users{}
		Newuser.Error = "Password not correct !"
		return Newuser
	}
	//verification si l'utilisateur n'est pas deja enregistré
	for _, v := range Allusers {
		if v.Username == Newuser.Username || v.Email == Newuser.Email {
			Newuser = structs.Users{}
			Newuser.Error = "Allready use it, please change !"
			return Newuser
		}
	}
	bd.NewUser(db, Newuser) //Inserer le nouvelle utilisateur dans  la base de donnée
	presentUser := bd.DataUser(db, Newuser.Email)
	actif := session(w, r, presentUser, db)
	if !actif {
		Newuser = structs.Users{}
		Newuser.Error = "L'utilisateur est déjà connecté depuis un autre endroit."
		return Newuser
	}
	return presentUser
}

func IsValidMail(t string) bool {
	//si l'email contient que des chiffres lettres et 2 caracteres speciaux
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(t)
}
