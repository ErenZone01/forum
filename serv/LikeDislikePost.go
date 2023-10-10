package serv

import (
	"fmt"
	"net/http"
	"strconv"

	"main.go/bd"
	"main.go/structs"
)

func likePost(user structs.Users, w http.ResponseWriter, r *http.Request) {
	var Id_post = r.FormValue("id_PostL")
	var app_Post structs.Appreciation_post
	Id_posts, err := strconv.Atoi(Id_post)
	if err != nil {
		fmt.Println("id incorrect")
	}

	var AlllikePost, _ = bd.DataAllLikePost(BD)
	for i, v := range AlllikePost {
		if v.Posts_id == Id_posts && user.Id == v.Users_id { //si le user a une fois liker ou disliker le post
			if v.Like == "false" && v.Dislike == "false" {
				bd.SelectLikePost(BD, "true", "false", Id_posts, user.Id)
				break
			} else if v.Like == "true" && v.Dislike == "false" {
				bd.SelectLikePost(BD, "false", "false", Id_posts, user.Id)
				break
			} else if v.Like == "false" && v.Dislike == "true" {
				bd.SelectLikePost(BD, "true", "false", Id_posts, user.Id)
				break
			}
		} else if i == len(AlllikePost)-1 { // si l'utilisateur n'a jamais liker le post
			app_Post.Users_id = user.Id
			app_Post.Posts_id = Id_posts
			app_Post.Like = "true"
			app_Post.Dislike = "false"
			bd.NewLikeDislikePost(BD, app_Post)
		}
	}

	if len(AlllikePost) == 0 {
		app_Post.Users_id = user.Id
		app_Post.Posts_id = Id_posts
		app_Post.Like = "true"
		app_Post.Dislike = "false"
		bd.NewLikeDislikePost(BD, app_Post)
	}

	for i, v := range All_Forum.Posts {
		AllLikepost, _ := bd.DataLikePost(BD, v.Id)
		AllDislikePost, _ := bd.DataDisLikePost(BD, v.Id)
		v.N_like = len(AllLikepost)
		v.N_dislike = len(AllDislikePost)
		bd.SelectNLikePost(BD, len(AllLikepost), len(AllDislikePost), i+1)
		All_Forum.Posts[i].N_like = len(AllLikepost)
		All_Forum.Posts[i].N_dislike = len(AllDislikePost)
	}

	http.Redirect(w, r, "/#P "+Id_post, http.StatusSeeOther)

}
func dislikePost(user structs.Users, w http.ResponseWriter, r *http.Request) {
	var Id_post = r.FormValue("id_PostD")
	var app_Post structs.Appreciation_post
	Id_posts, err := strconv.Atoi(Id_post)
	if err != nil {
		fmt.Println("id incorrect")
	}

	var AlllikePost, _ = bd.DataAllLikePost(BD)
	for i, v := range AlllikePost {
		if v.Posts_id == Id_posts && user.Id == v.Users_id { //si le user a une fois liker ou disliker le post
			if v.Like == "false" && v.Dislike == "false" {
				bd.SelectLikePost(BD, "false", "true", Id_posts, user.Id)
				break
			} else if v.Like == "true" && v.Dislike == "false" {
				bd.SelectLikePost(BD, "false", "true", Id_posts, user.Id)
				break
			} else if v.Like == "false" && v.Dislike == "true" {
				bd.SelectLikePost(BD, "false", "false", Id_posts, user.Id)
				break
			}
		} else if i == len(AlllikePost)-1 { // si l'utilisateur n'a jamais liker le post
			app_Post.Users_id = user.Id
			app_Post.Posts_id = Id_posts
			app_Post.Like = "false"
			app_Post.Dislike = "true"
			bd.NewLikeDislikePost(BD, app_Post)
		}
	}
	if len(AlllikePost) == 0 {
		app_Post.Users_id = user.Id
		app_Post.Posts_id = Id_posts
		app_Post.Like = "false"
		app_Post.Dislike = "true"
		bd.NewLikeDislikePost(BD, app_Post)
	}

	for i, v := range All_Forum.Posts {
		AllLikepost, _ := bd.DataLikePost(BD, v.Id)
		AllDislikePost, _ := bd.DataDisLikePost(BD, v.Id)
		v.N_like = len(AllLikepost)
		v.N_dislike = len(AllDislikePost)
		bd.SelectNLikePost(BD, len(AllLikepost), len(AllDislikePost), i+1)
		All_Forum.Posts[i].N_like = len(AllLikepost)
		All_Forum.Posts[i].N_dislike = len(AllDislikePost)
	}

	http.Redirect(w, r, "/#P "+Id_post, http.StatusSeeOther)
}
