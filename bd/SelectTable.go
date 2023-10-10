package bd

import (
	"database/sql"
	"fmt"

	"main.go/structs"
)

func DataAllUser(db *sql.DB) ([]structs.Users, error) {
	query := "SELECT id, username, mdp, email, role FROM users"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var usersList []structs.Users
	for rows.Next() {
		var user structs.Users
		if err := rows.Scan(&user.Id, &user.Username, &user.Mdp, &user.Email, &user.Role); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		usersList = append(usersList, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}

	return usersList, nil
}
func DataUser(db *sql.DB, email string) structs.Users {
	query := "SELECT id, username, mdp, email, role FROM users WHERE email = ?"
	var user structs.Users
	err := db.QueryRow(query, email).Scan(&user.Id, &user.Username, &user.Mdp, &user.Email, &user.Role)
	if err != nil {
		fmt.Println("Error1", err)
		return structs.Users{}
	}
	return user
}
func DataPost(db *sql.DB, posts structs.Posts) structs.Posts {
	query := "SELECT Id, Title, Body, N_like, N_dislike, N_com, Users_id, created_at FROM posts WHERE id = ?"
	var post structs.Posts
	err := db.QueryRow(query, 1).Scan(&post.Id, &post.Title, &post.Body, &post.N_like, &post.N_dislike, &post.N_com, &post.Users_id, &posts.CreatedPost)
	if err != nil {
		fmt.Println("Error1", err)
		return structs.Posts{}
	}
	return post
}
func DataCom(db *sql.DB, posts_Id int) ([]structs.Comments, error) {
	query := "SELECT Id, Body, nbr_like, nbr_dislike, users_id, posts_id FROM comments WHERE posts_id = ?"
	rows, err := db.Query(query, posts_Id)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var comList []structs.Comments
	for rows.Next() {
		var coms structs.Comments
		if err := rows.Scan(&coms.Id, &coms.Body, &coms.N_like, &coms.N_dislike, &coms.Users_id, &coms.Posts_id); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		comList = append(comList, coms)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return comList, nil
}
func DataAllCom(db *sql.DB) ([]structs.Comments, error) {
	query := "SELECT Id, Body, nbr_like, nbr_dislike, users_id, posts_id FROM comments"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var comList []structs.Comments
	for rows.Next() {
		var coms structs.Comments
		if err := rows.Scan(&coms.Id, &coms.Body, &coms.N_like, &coms.N_dislike, &coms.Users_id, &coms.Posts_id); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		comList = append(comList, coms)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return comList, nil
}
func DataNbrCom(db *sql.DB, id int) ([]structs.Comments, error) {
	query := "SELECT Id, Body, nbr_like, nbr_dislike, users_id, posts_id FROM comments WHERE posts_id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var comList []structs.Comments
	for rows.Next() {
		var coms structs.Comments
		if err := rows.Scan(&coms.Id, &coms.Body, &coms.N_like, &coms.N_dislike, &coms.Users_id, &coms.Posts_id); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		comList = append(comList, coms)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return comList, nil
}
func DataAllPost(db *sql.DB) ([]structs.Posts, error) {
	query := "SELECT id, title, body, nbr_like, nbr_dislike, nbr_com, users_id, created_at FROM posts ORDER BY created_at DESC"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var postList []structs.Posts
	for rows.Next() {
		var post structs.Posts
		if err := rows.Scan(&post.Id, &post.Title, &post.Body, &post.N_like, &post.N_dislike, &post.N_com, &post.Users_id, &post.CreatedPost); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		postList = append(postList, post)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return postList, nil
}
func DataAllLikePost(db *sql.DB) ([]structs.Appreciation_post, error) {
	query := "SELECT id, like, dislike, users_id, posts_id FROM appreciation_post"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var postList []structs.Appreciation_post
	for rows.Next() {
		var post structs.Appreciation_post
		if err := rows.Scan(&post.Id, &post.Like, &post.Dislike, &post.Users_id, &post.Posts_id); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		postList = append(postList, post)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return postList, nil
}
func SelectLikePost(db *sql.DB, like string, dislike string, postsID int, userID int) {
	_, err := db.Exec("UPDATE appreciation_post SET like = ?, dislike = ? WHERE posts_id = ? AND users_id = ?", like, dislike, postsID, userID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
func SelectNbrCom(db *sql.DB, Nbrcom structs.Posts) {
	_, err := db.Exec("UPDATE posts SET  nbr_com = ? WHERE id = ?", Nbrcom.N_com, Nbrcom.Id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
func SelectNLikePost(db *sql.DB, like int, dislike int, postsID int) {
	_, err := db.Exec("UPDATE posts SET nbr_like = ?, nbr_dislike = ? WHERE id = ?", like, dislike, postsID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
func SelectMostLiked(db *sql.DB) ([]structs.Posts, error){
	query := "SELECT id, title, body, nbr_like, nbr_dislike, nbr_com, users_id FROM posts ORDER BY nbr_like DESC "
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var PostsLikedList []structs.Posts
	for rows.Next() {
		var LikeP structs.Posts
		if err := rows.Scan(&LikeP.Id, &LikeP.Title, &LikeP.Body, &LikeP.N_like, &LikeP.N_dislike, &LikeP.N_com, &LikeP.Users_id); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		PostsLikedList = append(PostsLikedList, LikeP)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return PostsLikedList, nil
}
func SelectLikeComs(db *sql.DB, like string, dislike string, comsID int, userID int) {
	_, err := db.Exec("UPDATE appreciation_com SET like = ?, dislike = ? WHERE coms_id = ? AND users_id = ?", like, dislike, comsID, userID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
func DataAllLikeCom(db *sql.DB) ([]structs.Appreciation_com, error) {
	query := "SELECT id, like, dislike, users_id, coms_id FROM appreciation_com"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var ComList []structs.Appreciation_com
	for rows.Next() {
		var Com structs.Appreciation_com
		if err := rows.Scan(&Com.Id, &Com.Like, &Com.Dislike, &Com.Users_id, &Com.Coms_id); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		ComList = append(ComList, Com)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return ComList, nil
}
func DataLikePost(db *sql.DB, post_id int) ([]structs.Appreciation_post, error){
	query := "SELECT id, like, dislike, users_id, posts_id FROM appreciation_post WHERE posts_id = ? AND like = ?"
	rows, err := db.Query(query, post_id, "true")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var postList []structs.Appreciation_post
	for rows.Next() {
		var post structs.Appreciation_post
		if err := rows.Scan(&post.Id, &post.Like, &post.Dislike, &post.Users_id, &post.Posts_id); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		postList = append(postList, post)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return postList, nil
}
func DataDisLikePost(db *sql.DB, post_id int) ([]structs.Appreciation_post, error){
	query := "SELECT id, like, dislike, users_id, posts_id FROM appreciation_post WHERE posts_id = ? AND dislike = ?"
	rows, err := db.Query(query, post_id, "true")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var postList []structs.Appreciation_post
	for rows.Next() {
		var post structs.Appreciation_post
		if err := rows.Scan(&post.Id, &post.Like, &post.Dislike, &post.Users_id, &post.Posts_id); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		postList = append(postList, post)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return postList, nil
}
func DataLikeCom(db *sql.DB, com_id int) ([]structs.Appreciation_com, error){
	query := "SELECT id, like, dislike, users_id, coms_id FROM appreciation_com WHERE coms_id = ? AND like = ?"
	rows, err := db.Query(query, com_id, "true")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var ComsList []structs.Appreciation_com
	for rows.Next() {
		var com structs.Appreciation_com
		if err := rows.Scan(&com.Id, &com.Like, &com.Dislike, &com.Users_id, &com.Coms_id); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		ComsList = append(ComsList, com)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return ComsList, nil
}
func DataDislikeCom(db *sql.DB, com_id int) ([]structs.Appreciation_com, error){
	query := "SELECT id, like, dislike, users_id, coms_id FROM appreciation_com WHERE coms_id = ? AND dislike = ?"
	rows, err := db.Query(query, com_id, "true")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var ComsList []structs.Appreciation_com
	for rows.Next() {
		var com structs.Appreciation_com
		if err := rows.Scan(&com.Id, &com.Like, &com.Dislike, &com.Users_id, &com.Coms_id); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		ComsList = append(ComsList, com)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return ComsList, nil
}
func DataMyPost(db *sql.DB, user_id int) ([]structs.Posts, error) {
	query := "SELECT id, title, body, nbr_like, nbr_dislike, nbr_com, users_id, created_at FROM posts WHERE users_id = ?"
	rows, err := db.Query(query, user_id)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var postList []structs.Posts
	for rows.Next() {
		var post structs.Posts
		if err := rows.Scan(&post.Id, &post.Title, &post.Body, &post.N_like, &post.N_dislike, &post.N_com, &post.Users_id, &post.CreatedPost); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		postList = append(postList, post)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return postList, nil
}
func DataMyCom(db *sql.DB, user_id int) ([]structs.Comments, error) {
	query := "SELECT id, Body, nbr_like, nbr_dislike, users_id, posts_id FROM comments WHERE users_id = ?"
	rows, err := db.Query(query, user_id)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var comList []structs.Comments
	for rows.Next() {
		var com structs.Comments
		if err := rows.Scan(&com.Id, &com.Body, &com.N_like, &com.N_dislike, &com.Users_id, &com.Posts_id); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		comList = append(comList, com)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return comList, nil
}
func DataPostC(db *sql.DB, id int) structs.Posts {
	query := "SELECT Id, Title, Body, nbr_like, nbr_dislike, nbr_com, Users_id, created_at FROM posts WHERE Id = ?"
	var post structs.Posts
	err := db.QueryRow(query, id).Scan(&post.Id, &post.Title, &post.Body, &post.N_like, &post.N_dislike, &post.N_com, &post.Users_id, &post.CreatedPost)
	if err != nil {
		fmt.Println("Error1", err)
		return structs.Posts{}
	}
	return post
}