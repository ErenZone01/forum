package bd

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"main.go/structs"
)

func CreateBd() *sql.DB {
	//creer une base de donnée
	db, err := sql.Open("sqlite3", "Forum.sqlite")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users(
			id INTEGER PRIMARY KEY,
			username TEXT NOT NULL,
			mdp		 TEXT NOT NULL,
			email	 TEXT NOT NULL,
			role     TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)	
	`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS posts(
		id INTEGER PRIMARY KEY,
		title TEXT IF NOT NULL,
		body  TEXT IF NOT NULL,
		nbr_like INTEGER,
		nbr_dislike INTEGER,
		nbr_com INTEGER,
		users_id INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (users_id) REFERENCES users (id)
	)
	`)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS comments(
		id INTEGER PRIMARY KEY,
		nbr_like INTEGER,
		Body TEXT,
		nbr_dislike INTEGER,
		users_id INTEGER NOT NULL,
		posts_id INTEGER,
		FOREIGN KEY (users_id) REFERENCES users (id),
		FOREIGN KEY (posts_id) REFERENCES posts (id)
	)
	`)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS appreciation_com(
		id INTEGER PRIMARY KEY,
		like TEXT,
		dislike TEXT,
		users_id INTEGER,
		coms_id INTEGER,
		FOREIGN KEY (users_id) REFERENCES users (id)
		FOREIGN KEY (coms_id) REFERENCES comments (id)
	)
	`)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS appreciation_post(
		id INTEGER PRIMARY KEY,
		like TEXT,
		dislike TEXT,
		users_id INTEGER,
		posts_id INTEGER,
		FOREIGN KEY (users_id) REFERENCES users (id)
		FOREIGN KEY (posts_id) REFERENCES posts (id)
	)
	`)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS categories(
		id INTEGER PRIMARY KEY,
		type TEXT NOT NULL,
		posts_id INTEGER,
		FOREIGN KEY (posts_id) REFERENCES post (id)
	)
	`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}
func NewUser(bd *sql.DB, users structs.Users) {
	//Inserer des utilisateurs dans notre table Users
	_, err := bd.Exec(`
    INSERT INTO users (username, mdp, email, role) VALUES (?,?,?,?)
`, users.Username, users.Mdp, users.Email, "utilisateur")
	if err != nil {
		fmt.Println(err)
		return
	}
}
func NewPost(bd *sql.DB, post structs.Posts) {
	//Inserer des utilisateurs dans notre table Users
	_, err := bd.Exec(`
    INSERT INTO posts (title, body, users_id, nbr_like, nbr_dislike, nbr_com) VALUES (?,?,?,?,?,?)
`, post.Title, post.Body, post.Users_id, post.N_like, post.N_dislike, post.N_com)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func NewCom(bd *sql.DB, com structs.Comments) {
	//Inserer des commentaires dans notre table Com
	_, err := bd.Exec(`
	INSERT INTO comments (Body, nbr_like, nbr_dislike, users_id, posts_id) VALUES (?,?,?,?,?)
	`, com.Body, com.N_like, com.N_dislike, com.Users_id, com.Posts_id)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func NewCategorie(bd *sql.DB, categorie string, post_id int) {
	//Inserer des utilisateurs dans notre table Users
	_, err := bd.Exec(`
    INSERT INTO categorie (type, posts_id) VALUES (?,?)`, categorie, post_id)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func NewLikeDislikePost(bd *sql.DB, likePost structs.Appreciation_post){
	//Inserer des likes dans notre table Appreciation_post
	_, err := bd.Exec(`
    INSERT INTO appreciation_post (like, dislike, users_id, posts_id) VALUES (?,?,?,?)`, likePost.Like, likePost.Dislike, likePost.Users_id, likePost.Posts_id)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func NewLikeDislikeCom(bd *sql.DB, likeCom structs.Appreciation_com){
	//Inserer des likes dans notre table Appreciation_post
	_, err := bd.Exec(`
    INSERT INTO appreciation_com (like, dislike, users_id, coms_id) VALUES (?,?,?,?)`, likeCom.Like, likeCom.Dislike, likeCom.Users_id, likeCom.Coms_id)
	if err != nil {
		fmt.Println(err)
		return
	}
}