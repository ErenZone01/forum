package serv

import (
	"fmt"
	"net/http"
	"strconv"

	"main.go/bd"
	"main.go/structs"
)

func likeCom(user structs.Users, w http.ResponseWriter, r *http.Request) {
	var Id_Com = r.FormValue("id_ComL")
	var app_Com structs.Appreciation_com
	Id_Coms, err := strconv.Atoi(Id_Com)
	if err != nil {
		fmt.Println("id incorrect")
	}
	var AlllikeCom, _ = bd.DataAllLikeCom(BD)
	for i, v := range AlllikeCom {
		if v.Coms_id == Id_Coms && user.Id == v.Users_id { //si le user a une fois liker ou disliker le Com
			if v.Like == "false" && v.Dislike == "false" {
				bd.SelectLikeComs(BD, "true", "false", Id_Coms, user.Id)
				break
			} else if v.Like == "true" && v.Dislike == "false" {
				bd.SelectLikeComs(BD, "false", "false", Id_Coms, user.Id)
				break
			} else if v.Like == "false" && v.Dislike == "true" {
				bd.SelectLikeComs(BD, "true", "false", Id_Coms, user.Id)
				break
			}
		} else if i == len(AlllikeCom)-1 { // si l'utilisateur n'a jamais liker le Com
			app_Com.Users_id = user.Id
			app_Com.Coms_id = Id_Coms
			app_Com.Like = "true"
			app_Com.Dislike = "false"
			bd.NewLikeDislikeCom(BD, app_Com)
		}
	}
	if len(AlllikeCom) == 0 {
		app_Com.Users_id = user.Id
		app_Com.Coms_id = Id_Coms
		app_Com.Like = "true"
		app_Com.Dislike = "false"
		bd.NewLikeDislikeCom(BD, app_Com)
	}

	for i, v := range All_Forum.Coms {
		AllLikeCom, _ := bd.DataLikeCom(BD, v.Id)
		AllDislikeCom, _ := bd.DataDislikeCom(BD, v.Id)
		v.N_like = len(AllLikeCom)
		v.N_dislike = len(AllDislikeCom)
		All_Forum.Coms[i].N_like = len(AllLikeCom)
		All_Forum.Coms[i].N_dislike = len(AllDislikeCom)
	}

	http.Redirect(w, r, "/#C "+Id_Com, http.StatusSeeOther)
}
func dislikeCom(user structs.Users, w http.ResponseWriter, r *http.Request) {
	var Id_Com = r.FormValue("id_ComD")
	var app_Com structs.Appreciation_com
	Id_Coms, err := strconv.Atoi(Id_Com)
	if err != nil {
		fmt.Println("id incorrect")
	}

	var AlllikeCom, _ = bd.DataAllLikeCom(BD)
	for i, v := range AlllikeCom {
		if v.Coms_id == Id_Coms && user.Id == v.Users_id { //si le user a une fois liker ou disliker le Com
			if v.Like == "false" && v.Dislike == "false" {
				bd.SelectLikeComs(BD, "false", "true", Id_Coms, user.Id)
				break
			} else if v.Like == "true" && v.Dislike == "false" {
				bd.SelectLikeComs(BD, "false", "true", Id_Coms, user.Id)
				break
			} else if v.Like == "false" && v.Dislike == "true" {
				bd.SelectLikeComs(BD, "false", "false", Id_Coms, user.Id)
				break
			}
		} else if i == len(AlllikeCom)-1 { // si l'utilisateur n'a jamais liker le Com
			app_Com.Users_id = user.Id
			app_Com.Coms_id = Id_Coms
			app_Com.Like = "false"
			app_Com.Dislike = "true"
			bd.NewLikeDislikeCom(BD, app_Com)
		}
	}

	if len(AlllikeCom) == 0 {
		app_Com.Users_id = user.Id
		app_Com.Coms_id = Id_Coms
		app_Com.Like = "true"
		app_Com.Dislike = "false"
		bd.NewLikeDislikeCom(BD, app_Com)
	}

	for i, v := range All_Forum.Coms {
		AllLikeCom, _ := bd.DataLikeCom(BD, v.Id)
		AllDislikeCom, _ := bd.DataDislikeCom(BD, v.Id)
		v.N_like = len(AllLikeCom)
		v.N_dislike = len(AllDislikeCom)
		All_Forum.Coms[i].N_like = len(AllLikeCom)
		All_Forum.Coms[i].N_dislike = len(AllDislikeCom)
	}
	http.Redirect(w, r, "/#C "+Id_Com, http.StatusSeeOther)
}
