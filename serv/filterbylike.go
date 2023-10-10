package serv

import (
	"database/sql"
	"fmt"

	"main.go/structs"
)


func DataAllLikedPosts(db *sql.DB, userID int) ([]structs.Posts, error) {
    query := "SELECT id, title, body, nbr_like, nbr_dislike, nbr_com, users_id  FROM posts WHERE users_id = ?"
    rows, err := db.Query(query, userID)
    if err != nil {
        fmt.Println("Error:", err)
        return nil, err
    }
    defer rows.Close()

    var postList []structs.Posts
    for rows.Next() {
        var post structs.Posts
        if err := rows.Scan(&post.Id, &post.Title, &post.Body , &post.N_like, &post.N_dislike, &post.N_com, &post.Users_id); err != nil {
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

func DataLikedPosts(db *sql.DB, userID int) ([]structs.Appreciation_post, error) {
    query := "SELECT id, like, dislike, users_id, posts_id  FROM appreciation_post WHERE users_id = ? AND like = ?"
    rows, err := db.Query(query, userID, "true")
    if err != nil {
        fmt.Println("Error:", err)
        return nil, err
    }
    defer rows.Close()

    var APList []structs.Appreciation_post
    for rows.Next() {
        var likepost structs.Appreciation_post
        if err := rows.Scan(&likepost.Id, &likepost.Like, &likepost.Dislike , &likepost.Users_id, &likepost.Posts_id); err != nil {
            fmt.Println("Error scanning row:", err)
            continue
        }
        APList = append(APList, likepost)
    }

    if err := rows.Err(); err != nil {
        fmt.Println("Error iterating rows:", err)
        return nil, err
    }

    return APList, nil
}
