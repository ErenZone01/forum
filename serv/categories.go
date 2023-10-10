package serv

import (
	"database/sql"
	"fmt"
)

func addCategorie(db *sql.DB, categorie string, post int) {
	_, err := db.Exec(`
    INSERT INTO categories (type, posts_id) VALUES (?,?)`, categorie, post)
	if err != nil {
		fmt.Println(err)
		return
	}
}
