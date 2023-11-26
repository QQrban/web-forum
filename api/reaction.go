package api

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
)

type NewResponse struct {
	Role     string `json:"Role"`
	Ok       bool   `json:"Ok"`
	Add      string `json:"Add"`
	UpdateTo string `json:"UpdateTo"`
	Delete   string `json:"Delete"`
}

func ReactionHandle(w http.ResponseWriter, r *http.Request) {
	sess, err := CheckAuthorization(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user_id := sess.User_Id
	if r.Method != "GET" {
		BadRequest(w, r)
		return
	}
	//fmt.Println("path", r.URL.Path)
	path := r.URL.Path[len("/api/reaction/"):]
	parts := strings.Split(path, "/")
	if len(parts) != 3 {
		BadRequest(w, r)
		return
	}
	table := parts[0]
	idstr := parts[1]
	react := parts[2]
	rmap := map[string]int{
		"like":    1,
		"dislike": -1,
	}
	if _, ok := rmap[react]; !ok {
		NotFound(w, r)
		return
	}
	mapr := map[int]string{
		1:  "like",
		-1: "dislike",
	}
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		BadRequest(w, r)
		return
	}
	var rid int64
	var reaction int
	//var tbl, col,
	var qry string
	switch table {
	case "post":
		//tbl = "Post_Reaction"
		//col = "Post_Id"
		qry = "SELECT Comment_Reaction.Id, Reaction FROM Comment, Comment_Reaction WHERE Post_Id= ? AND Is_Post=1 AND Comment.Id=Comment_Id AND Comment_Reaction.User_Id = ?"
	case "comment":
		//tbl = "Comment_Reaction"
		//col = "Comment_Id"
		qry = "SELECT Id, Reaction FROM Comment_Reaction WHERE Comment_Id= ? AND User_Id = ?"
	default:
		NotFound(w, r)
		return
	}
	err = DB.QueryRow(qry, id, user_id).Scan(&rid, &reaction)
	if err == nil && reaction == rmap[react] {
		switch table {
		case "post":
			err = deletePostReaction(rid)
		case "comment":
			err = deleteCommentReaction(rid)
		}
		if err != nil {
			InternalError(w, r)
			return
		}
		Respond(w, http.StatusOK, NewResponse{
			Delete: react,
			Role:   sess.Role,
			Ok:     true,
		})
		return
	}
	switch err {
	case nil:
		_, err := DB.Exec("UPDATE Comment_Reaction SET Reaction = ? WHERE Id = ?", rmap[react], rid)
		if err != nil {
			InternalError(w, r)
			return
		}
		Respond(w, http.StatusOK, NewResponse{
			UpdateTo: mapr[-reaction],
			Role:     sess.Role,
			Ok:       true,
		})
		return
	case sql.ErrNoRows:
		if table == "post" {
			err := DB.QueryRow("SELECT Id FROM Comment WHERE Post_Id = ? AND Is_Post = 1", id).Scan(&id)
			if err != nil {
				InternalError(w, r)
				return
			}
		}
		_, err := DB.Exec("INSERT INTO Comment_Reaction (Comment_Id, User_Id, Reaction) VALUES (?, ?, ?)", id, user_id, rmap[react])
		if err != nil {
			InternalError(w, r)
			return
		}
		Respond(w, http.StatusOK, NewResponse{
			Add:  react,
			Role: sess.Role,
			Ok:   true,
		})
		return
	default:
		InternalError(w, r)
		return
	}
}
