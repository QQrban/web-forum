package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func TopicHandle(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Path[len("/api/topic/"):]
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		BadRequest(w, r)
		return
	}
	posts, err := GetTopicPosts(id)
	if err != nil || len(posts) == 0 {
		NotFound(w, r)
		return
	}
	//fmt.Println("posts:", posts)
	Respond(w, http.StatusOK, posts)
}

func PostHandle(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Path[len("/api/post/"):]
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		BadRequest(w, r)
		return
	}
	comments, err := GetCommentsWithAuthors(id)
	if err != nil || len(comments) == 0 {
		NotFound(w, r)
		return
	}
	//fmt.Println("post:", post)
	Respond(w, http.StatusOK, comments)
}

func CreateNewPostHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/create-new-post" {
		NotFound(w, r)
		return
	}
	if r.Method != "POST" {
		BadRequest(w, r)
		return
	}
	sess, err := CheckAuthorization(w, r) // CheckAuthorization is in checksession.go
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user_id := sess.User_Id

	type newPost struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		Tags       string `json:"tags"`
		CategoryId string `json:"categoryId"`
		Created    string `json:"created"`
	}
	var NewPost = newPost{}
	json.NewDecoder(r.Body).Decode(&NewPost)
	category_id, err := strconv.ParseInt(NewPost.CategoryId, 10, 64)
	if err != nil {
		InternalError(w, r)
		return
	}
	post := &Post{
		Title:       NewPost.Title,
		User_Id:     user_id,
		Category_Id: category_id,
		Tags:        NewPost.Tags,
	}
	//var id int64
	_, err = post.Create()
	if err != nil {
		InternalError(w, r)
		return
	}
	//fmt.Println("post:", post)
	comm := &Comment{
		User_Id: user_id,
		Post_Id: post.Id,
		Content: NewPost.Content,
		Is_Post: true,
	}
	cid, err := comm.Create()
	//fmt.Println("comm:", cid, comm)
	if err != nil {
		InternalError(w, r)
		return
	}
	post.Content_Id = cid
	post.Update()
	//fmt.Printf("post: %q\n", post)
	if len(post.Tags) > 0 {
		tagSlice := strings.Split(post.Tags, ",")
		for _, tag := range tagSlice {
			t := &Tag{Name: tag}
			t.Create()
		}
	}

	Respond(w, http.StatusCreated, Response{
		Message: "Post Created",
		Role:    sess.Role,
		Ok:      true,
	})
}

func CreateNewCommentHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/create-new-comment" {
		NotFound(w, r)
		return
	}
	if r.Method != "POST" {
		BadRequest(w, r)
		return
	}
	sess, err := CheckAuthorization(w, r) // CheckAuthorization is in checksession.go
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user_id := sess.User_Id
	type newComment struct {
		Content string `json:"text"`
		PostId  string `json:"Id"`
		//Created string `json:"created"`
	}
	var NewComment = newComment{}
	json.NewDecoder(r.Body).Decode(&NewComment)
	//fmt.Println("NewComment:", NewComment)
	post_id, err := strconv.ParseInt(NewComment.PostId, 10, 64)
	if err != nil {
		InternalError(w, r)
		return
	}
	comm := &Comment{
		User_Id: user_id,
		Post_Id: post_id,
		Content: NewComment.Content,
	}
	_, err = comm.Create()
	if err != nil {
		InternalError(w, r)
		return
	}
	//fmt.Println("comm:", comm)
	Respond(w, http.StatusCreated, Response{
		Message: "Comment Created",
		Role:    sess.Role,
		Ok:      true,
	})
}

func AllPosts(w http.ResponseWriter, r *http.Request) {
	posts := GetAllPosts()
	Respond(w, http.StatusOK, posts)
}
