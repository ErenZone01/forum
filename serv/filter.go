package serv

import (
	"database/sql"
	"fmt"

	"main.go/structs"
)

func DataAllPostByCategory(db *sql.DB, categorie string) ([]structs.Categories, error) {
	query := "SELECT posts_id FROM categories WHERE type = ?"
	rows, err := db.Query(query, categorie)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var categorieList []structs.Categories
	for rows.Next() {
		var categorie structs.Categories
		if err := rows.Scan(&categorie.Posts_id); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		categorieList = append(categorieList, categorie)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return categorieList, nil
}
