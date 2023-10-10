package serv

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"main.go/bd"
	"main.go/structs"
)

var user = structs.Users{}
var BD *sql.DB
var All_Forum structs.All_Forum
var Errors structs.Error

func Serve(db *sql.DB, errors structs.Error) {
	Errors = errors
	BD = db
	port := ":8080"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	http.HandleFunc("/", Handler)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/decon", Decon)
	//demarrer le serveur
	fmt.Println("server is start at : http://localhost:8080")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		Errors.Title = "Page not Found"
		Errors.Body = "404"
		errors(w, r, Errors)
		return
	}
	if Errors.Body == "Error 500" {
		Errors.Title = "Internal Server Error"
		errors(w, r, Errors)
		return
	}
	t, err := template.ParseFiles("templates/login.html")
	if err != nil {
		Errors.Title = "Internal Server Error"
		Errors.Body = "500"
		errors(w, r, Errors)
		return
	}
	Allusers := []structs.Users{}
	Allusers = append(Allusers, user)
	user = structs.Users{}
	t.Execute(w, Allusers)
}
func Register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		Errors.Title = "Page not Found"
		Errors.Body = "404"
		errors(w, r, Errors)
	}
	if Errors.Body == "Error 500" {
		Errors.Title = "Internal Server Error"
		errors(w, r, Errors)
	}
	t, err := template.ParseFiles("templates/register.html")
	if err != nil {
		Errors.Title = "Internal Server Error"
		Errors.Body = "500"
		errors(w, r, Errors)
		return
	}
	Allusers := []structs.Users{}
	Allusers = append(Allusers, user)
	user = structs.Users{}
	t.Execute(w, Allusers)
}
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if Errors.Body == "500" {
			Errors.Title = "Internal Server Error"
			errors(w, r, Errors)
			return
		}
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			Errors.Title = "Internal Server Error"
			Errors.Body = "500"
			errors(w, r, Errors)
			return

		}
		//Recuperer tout les Utilisateurs et les mettre dans la structure
		var Allusers, err0 = bd.DataAllUser(BD)
		//Recuperer tout les Posts et les mettre dans la structure
		var Allpost, err1 = bd.DataAllPost(BD)
		//Recuperer tout les Coms et les mettre dans la structure
		var Allcom, err2 = bd.DataAllCom(BD)

		//initialisation des structures
		All_Forum.Users = Allusers
		All_Forum.Posts = Allpost
		All_Forum.Coms = Allcom

		//mise a jour des Like et dislike
		for i, v := range All_Forum.Posts {
			AllLikepost, _ := bd.DataLikePost(BD, v.Id)
			AllDislikePost, _ := bd.DataDisLikePost(BD, v.Id)
			v.N_like = len(AllLikepost)
			v.N_dislike = len(AllDislikePost)
			All_Forum.Posts[i].N_like = len(AllLikepost)
			All_Forum.Posts[i].N_dislike = len(AllDislikePost)
		}
		for i, v := range All_Forum.Coms {
			AllLikeCom, _ := bd.DataLikeCom(BD, v.Id)
			AllDislikeCom, _ := bd.DataDislikeCom(BD, v.Id)
			v.N_like = len(AllLikeCom)
			v.N_dislike = len(AllDislikeCom)
			All_Forum.Coms[i].N_like = len(AllLikeCom)
			All_Forum.Coms[i].N_dislike = len(AllDislikeCom)
		}

		presentUser := Myaccount(w, r)

		//mise à jour nmbr de post et de coms
		All_Forum.Utilisateur = []structs.Users{}
		var MyPost, err4 = bd.DataMyPost(BD, presentUser.Id)
		All_Forum.MyPost = MyPost
		var MyCom, err5 = bd.DataMyCom(BD, presentUser.Id)
		All_Forum.MyCom = MyCom

		//Nombre de commentaire et reglage de la date
		for i, v := range All_Forum.Posts {
			//enlever certains caractere de la date
			var timer = v.CreatedPost
			timer1 := strings.Split(timer, "T")
			timer = timer1[0] + " " + timer1[1]
			timer2 := strings.Split(timer, "Z")
			timer = timer2[0] + " " + timer2[1]
			v.CreatedPost = timer
			Allpost[i].CreatedPost = timer

			ComPost, _ := bd.DataNbrCom(BD, v.Id)
			v.N_com = len(ComPost)
			Allpost[i].N_com = len(ComPost)
			bd.SelectNbrCom(BD, v)
		}

		//calcule du nombre de like
		AllLiked, err6 := DataAllLikedPosts(BD, presentUser.Id)
		if err6 != nil {
			fmt.Println("error 1:", err)
			return
		}
		var Nlike int
		for _, v := range AllLiked {
			Nlike += v.N_like
		}
		var tab structs.Mylike
		tab.N_like = Nlike
		All_Forum.Mylike = []structs.Mylike{}
		All_Forum.Mylike = append(All_Forum.Mylike, tab)

		if presentUser.Error != "Veuillez vous connecté d'abord" {
			All_Forum.Utilisateur = append(All_Forum.Utilisateur, presentUser)
		}

		//verification de la methode utilisée
		if r.Method == "POST" {
			switch r.FormValue("submit") {
			case "Sign Up": //ajout d'un utilisateur
				NewUser := SignUp(BD, w, r)
				All_Forum.Utilisateur = []structs.Users{}
				All_Forum.Utilisateur = append(All_Forum.Utilisateur, NewUser)
				Allusers, _ = bd.DataAllUser(BD)
				if NewUser.Error == "Allready use it, please change !" || NewUser.Error == "L'utilisateur est déjà connecté depuis un autre endroit." || NewUser.Error == "Password not correct !" || NewUser.Error == "Your Email is incorrect" || NewUser.Error == "Your password is incorrect" || NewUser.Error == "Your username must have 1 character different to space" || NewUser.Error == "The password does not match the password confirmation" || NewUser.Error == "Your password must have 1 character different to space" {
					user = structs.Users{}
					user.Error = NewUser.Error
					NewUser = structs.Users{}
					http.Redirect(w, r, "/register", http.StatusSeeOther)
					return
				}
				fmt.Println(NewUser.Username, " ,votre compte a été crée avec succée")
				All_Forum.Users = Allusers
				var Allcom, _ = bd.DataAllCom(BD)
				All_Forum.Coms = Allcom
				AllLiked, _ := DataAllLikedPosts(BD, NewUser.Id)
				var Nlike int
				for _, v := range AllLiked {
					Nlike += v.N_like
				}
				var tab structs.Mylike
				//tab = append(tab, Nlike)
				tab.N_like = Nlike
				All_Forum.Mylike = []structs.Mylike{}
				All_Forum.Mylike = append(All_Forum.Mylike, tab)
				//filtre most liked
				var Mostliked, err8 = bd.SelectMostLiked(BD)
				if err8 != nil {
					fmt.Println("error 8: ", err8)
					return
				}
				All_Forum.FilterLiked = Mostliked
				w.WriteHeader(http.StatusOK)
				t.Execute(w, All_Forum)
			case "Sign In": // Authentification
				NewUser := SignIn(BD, w, r)
				All_Forum.Utilisateur = []structs.Users{}
				All_Forum.Utilisateur = append(All_Forum.Utilisateur, NewUser)
				Allusers, _ = bd.DataAllUser(BD)
				if NewUser.Error == "User not found" || NewUser.Error == "L'utilisateur est déjà connecté depuis un autre endroit." || NewUser.Error == "Your Email is incorrect" || NewUser.Error == "Your password is incorrect" || NewUser.Error == "Your password must have 1 character different to space" {
					user = structs.Users{}
					user.Error = NewUser.Error
					NewUser = structs.Users{}
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}
				fmt.Println(NewUser.Username, " ,vous êtes connecté crée avec succée")
				All_Forum.Users = Allusers
				var MyPost, _ = bd.DataMyPost(BD, NewUser.Id)
				All_Forum.MyPost = MyPost
				var MyCom, _ = bd.DataMyCom(BD, NewUser.Id)
				All_Forum.MyCom = MyCom
				AllLiked, _ := DataAllLikedPosts(BD, NewUser.Id)
				var Nlike int
				for _, v := range AllLiked {
					Nlike += v.N_like
				}
				var tab structs.Mylike
				tab.N_like = Nlike
				All_Forum.Mylike = []structs.Mylike{}
				All_Forum.Mylike = append(All_Forum.Mylike, tab)
				//filtre most liked
				var Mostliked, err8 = bd.SelectMostLiked(BD)
				if err8 != nil {
					fmt.Println("error 8: ", err8)
					return
				}
				All_Forum.FilterLiked = Mostliked
				w.WriteHeader(http.StatusOK)
				t.Execute(w, All_Forum)
			case "Post":
				user := Myaccount(w, r)
				if user.Error == "Veuillez vous connecté d'abord" {
					user.Error = ""
					http.Redirect(w, r, "/register", http.StatusSeeOther)
					return
				}
				var NewUser = bd.DataUser(BD, user.Email)
				var statusPost = CreatePost(w, r, NewUser)
				if !statusPost {
					Errors.Title = "Bad Request"
					Errors.Body = "400"
					errors(w, r, Errors)
					return
				}
				Allpost, _ := bd.DataAllPost(BD)
				All_Forum.Posts = Allpost
				for i, v := range All_Forum.Posts {
					//enlever certains caractere de la date
					var timer = v.CreatedPost
					timer1 := strings.Split(timer, "T")
					timer = timer1[0] + " " + timer1[1]
					timer2 := strings.Split(timer, "Z")
					timer = timer2[0] + " " + timer2[1]
					v.CreatedPost = timer
					Allpost[i].CreatedPost = timer
				}
				All_Forum.Posts = Allpost

				var MyPost, _ = bd.DataMyPost(BD, NewUser.Id)
				All_Forum.MyPost = MyPost
				var MyCom, _ = bd.DataMyCom(BD, NewUser.Id)
				All_Forum.MyCom = MyCom
				AllLiked, _ := DataAllLikedPosts(BD, NewUser.Id)
				var Nlike int
				for _, v := range AllLiked {
					Nlike += v.N_like
				}
				var tab structs.Mylike
				//tab = append(tab, Nlike)
				tab.N_like = Nlike
				All_Forum.Mylike = []structs.Mylike{}
				All_Forum.Mylike = append(All_Forum.Mylike, tab)
				//filtre most liked
				var Mostliked, err8 = bd.SelectMostLiked(BD)
				if err8 != nil {
					fmt.Println("error 8: ", err8)
					return
				}
				All_Forum.FilterLiked = Mostliked
				w.WriteHeader(http.StatusOK)
				t.Execute(w, All_Forum)
			case "Add Comment": // Ajout d'un Commentaire
				user := Myaccount(w, r)
				if user.Error == "Veuillez vous connecté d'abord" {
					user.Error = ""
					http.Redirect(w, r, "/register", http.StatusSeeOther)
					return
				}
				users := bd.DataUser(BD, user.Email)

				var statusPost = CreatCom(w, r, users.Id)
				if !statusPost {
					Errors.Title = "Bad Request"
					Errors.Body = "400"
					errors(w, r, Errors)
					return
				}
				var MyPost, _ = bd.DataMyPost(BD, user.Id)
				for i, v := range MyPost {
					MyPost[i].Filter = 0
					v.Filter = 0
				}
				All_Forum.MyPost = MyPost
				var MyCom, _ = bd.DataMyCom(BD, user.Id)
				All_Forum.MyCom = MyCom
				AllLiked, _ := DataAllLikedPosts(BD, user.Id)
				var Nlike int
				for _, v := range AllLiked {
					Nlike += v.N_like
				}
				var tab structs.Mylike
				//tab = append(tab, Nlike)
				tab.N_like = Nlike
				var Allcom, _ = bd.DataAllCom(BD)
				All_Forum.Coms = Allcom
				//Nombre de commentaire
				for i, v := range All_Forum.Posts {
					ComPost, _ := bd.DataNbrCom(BD, v.Id)
					v.N_com = len(ComPost)
					Allpost[i].N_com = len(ComPost)
					bd.SelectNbrCom(BD, v)
				}
				MyPost, _ = bd.DataMyPost(BD, user.Id)
				All_Forum.MyPost = MyPost
				All_Forum.Mylike = []structs.Mylike{}
				All_Forum.Mylike = append(All_Forum.Mylike, tab)
			case "likePost": //like
				user := Myaccount(w, r)
				if user.Error == "Veuillez vous connecté d'abord" {
					user.Error = ""
					http.Redirect(w, r, "/register", http.StatusSeeOther)
					return
				}
				users := bd.DataUser(BD, user.Email)
				likePost(users, w, r)
				AllLiked, _ := DataAllLikedPosts(BD, user.Id)
				var Nlike int
				for _, v := range AllLiked {
					Nlike += v.N_like
				}
				var tab structs.Mylike
				tab.N_like = Nlike
				All_Forum.Mylike = []structs.Mylike{}
				All_Forum.Mylike = append(All_Forum.Mylike, tab)
				//filtre most liked
				var Mostliked, err8 = bd.SelectMostLiked(BD)
				if err8 != nil {
					fmt.Println("error 8: ", err8)
					return
				}
				All_Forum.FilterLiked = Mostliked
			case "dislikePost": //dislike
				user := Myaccount(w, r)
				if user.Error == "Veuillez vous connecté d'abord" {
					user.Error = ""
					http.Redirect(w, r, "/register", http.StatusSeeOther)
					return
				}
				users := bd.DataUser(BD, user.Email)
				dislikePost(users, w, r)
				AllLiked, _ := DataAllLikedPosts(BD, user.Id)
				var Nlike int
				for _, v := range AllLiked {
					Nlike += v.N_like
				}
				var tab structs.Mylike
				//tab = append(tab, Nlike)
				tab.N_like = Nlike
				All_Forum.Mylike = []structs.Mylike{}
				All_Forum.Mylike = append(All_Forum.Mylike, tab)
				//filtre most liked
				var Mostliked, err8 = bd.SelectMostLiked(BD)
				if err8 != nil {
					fmt.Println("error 8: ", err8)
					return
				}
				All_Forum.FilterLiked = Mostliked
			case "likeCom":
				user := Myaccount(w, r)
				if user.Error == "Veuillez vous connecté d'abord" {
					user.Error = ""
					http.Redirect(w, r, "/register", http.StatusSeeOther)
					return
				}
				users := bd.DataUser(BD, user.Email)
				likeCom(users, w, r)
			case "dislikeCom":
				user := Myaccount(w, r)
				if user.Error == "Veuillez vous connecté d'abord" {
					user.Error = ""
					http.Redirect(w, r, "/register", http.StatusSeeOther)
					return
				}
				users := bd.DataUser(BD, user.Email)
				dislikeCom(users, w, r)
			case "Filter by category":
				var categorie = r.FormValue("categorie")
				if categorie != "" {
					var c, _ = DataAllPostByCategory(BD, categorie)
					var final []structs.Posts
					for _, v := range c {
						var f = bd.DataPostC(BD, v.Posts_id)
						//enlever certains caractere de la date
						var timer = f.CreatedPost
						timer1 := strings.Split(timer, "T")
						timer = timer1[0] + " " + timer1[1]
						timer2 := strings.Split(timer, "Z")
						timer = timer2[0] + " " + timer2[1]
						f.CreatedPost = timer
						//filtre
						f.Filter = 1
						final = append(final, f)
					}
					All_Forum.Posts = final
				}
				w.WriteHeader(http.StatusOK)
				t.Execute(w, All_Forum)
			case "Filter by Like":
				//filtre most liked

				user := Myaccount(w, r)

				if user.Error == "Veuillez vous connecté d'abord" {
					user.Error = ""
					http.Redirect(w, r, "/register", http.StatusSeeOther)
					return
				}
				var actu, _ = DataLikedPosts(BD, user.Id)
				var MyPostsLiked = []structs.Posts{}
				for _, v := range actu {
					var Prepost = bd.DataPostC(BD, v.Posts_id)
					//enlever certains caractere de la date
					var timer = Prepost.CreatedPost
					timer1 := strings.Split(timer, "T")
					timer = timer1[0] + " " + timer1[1]
					timer2 := strings.Split(timer, "Z")
					timer = timer2[0] + " " + timer2[1]
					Prepost.CreatedPost = timer
					//Filtre
					Prepost.Filter = 1
					MyPostsLiked = append(MyPostsLiked, Prepost)
				}
				All_Forum.Posts = MyPostsLiked
				t.Execute(w, All_Forum)
			case "Filter by Created Post":
				user := Myaccount(w, r)
				if user.Error == "Veuillez vous connecté d'abord" {
					user.Error = ""
					http.Redirect(w, r, "/register", http.StatusSeeOther)
					return
				}
				var myPost, _ = bd.DataMyPost(BD, user.Id)
				for i, v := range myPost {
					//enlever certains caractere de la date
					var timer = v.CreatedPost
					timer1 := strings.Split(timer, "T")
					timer = timer1[0] + " " + timer1[1]
					timer2 := strings.Split(timer, "Z")
					timer = timer2[0] + " " + timer2[1]
					v.CreatedPost = timer
					myPost[i].CreatedPost = timer
					//filtre
					myPost[i].Filter = 1
					v.Filter = 1
				}
				All_Forum.Posts = myPost
				t.Execute(w, All_Forum)
			}
		} else {
			if err0 != nil || err1 != nil || err2 != nil || err4 != nil || err5 != nil {
				fmt.Println("error: ", err0)
				return
			}
			w.WriteHeader(http.StatusOK)
			t.Execute(w, All_Forum)
		}
	} else {
		Errors.Body = "404"
		Errors.Title = "Page not Found"
		errors(w, r, Errors)
		return
	}
}

// pour les mdp username email post commentaire
func IsValid(s string) bool {
	eff := strings.TrimSpace(s)
	return len(eff) != 0
}
