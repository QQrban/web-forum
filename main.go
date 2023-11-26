package main

import (
	"forum/api"
	"forum/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/query/", api.QueryHandle)                        // api/query.go
	http.HandleFunc("/register", api.RegisterHandle)                   // api/register.go
	http.HandleFunc("/login", api.LoginHandle)                         // api/login.go
	http.HandleFunc("/logout", api.LogoutHandle)                       // api/logout.go
	http.HandleFunc("/check-session", api.CheckSessionHandle)          // api/checksession.go
	http.HandleFunc("/create-new-post", api.CreateNewPostHandle)       // api/posts.go
	http.HandleFunc("/create-new-comment", api.CreateNewCommentHandle) // api/posts.go
	http.HandleFunc("/api/topic/", api.TopicHandle)                    // api/posts.go
	http.HandleFunc("/api/post/", api.PostHandle)                      // api/posts.go
	http.HandleFunc("/api/all-posts", api.AllPosts)                    // api/posts.go
	http.HandleFunc("/api/reaction/", api.ReactionHandle)              // api/reaction.go
	http.HandleFunc("/post/", handlers.PostyTosty)                     // handlers/post.go
	http.HandleFunc("/topic/", handlers.TopicHandle)                   // handlers/topic.go
	http.HandleFunc("/about", handlers.AboutHandle)                    // handlers/about.go
	http.HandleFunc("/activity", handlers.ActivityHandle)              // handlers/activity.go
	http.HandleFunc("/", handlers.HomeHandle)                          // handlers/home.go
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./ui/static"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
