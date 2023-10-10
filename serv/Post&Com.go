package serv

import (
	"net/http"
	"strconv"

	"main.go/bd"
	"main.go/structs"
)

func CreatePost(w http.ResponseWriter, r *http.Request, Newuser structs.Users) bool {
	var title = r.FormValue("title")
	var body = r.FormValue("body")
	var categorie1 = r.FormValue("categorie1")
	var categorie2 = r.FormValue("categorie2")
	var categorie3 = r.FormValue("categorie3")
	var categorie4 = r.FormValue("categorie4")
	var categorie5 = r.FormValue("categorie5")
	var Allcategorie = []string{categorie1, categorie2, categorie3, categorie4, categorie5}

	if !IsValid(title) || !IsValid(body) || (!IsValid(categorie1) &&  !IsValid(categorie2) && !IsValid(categorie3) && !IsValid(categorie4) && !IsValid(categorie5)){
		return false
	}
	var post = structs.Posts{}
	post.Body = body
	post.Title = title
	post.Users_id = Newuser.Id
	bd.NewPost(BD, post)
	var post_bd, _ = bd.DataAllPost(BD)
	var newPost = post_bd[len(post_bd)-1]
	for _, v := range Allcategorie{
		if IsValid(v){
			addCategorie(BD, v, newPost.Id)
		}
	}
	//likePost(Newuser, w, r)
	return true
}
func CreatCom(w http.ResponseWriter, f *http.Request, user_id int) bool {
	var comment = f.FormValue("Comment")
	var id = f.FormValue("id")
	var com = structs.Comments{}
	com.Body = comment
	if !IsValid(comment) {
		return false
	}
	Id, err := strconv.Atoi(id)
	if err != nil{
		return false
	}
	com.Posts_id = Id
	com.Users_id = user_id
	bd.NewCom(BD, com)
	http.Redirect(w, f, "/#P "+id, http.StatusSeeOther)
	return true
}
