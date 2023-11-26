package handlers

import (
	"forum/api"
	"html/template"
	"net/http"
)

func ActivityHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/activity" {
		api.NotFound(w, r)
		return
	}
	if r.Method != "GET" {
		api.BadRequest(w, r)
		return
	}

	sess, err := api.CheckAuthorization(w, r)
	if err != nil {
		api.Error(w, r, "403", "Please login to view your activity page.")
		return
	}

	posts, err := api.GetPostsByUser(sess.User_Id)
	if err != nil {
		api.InternalError(w, r)
		return
	}
	likes, err := api.GetLikedPostsByUser(sess.User_Id)
	if err != nil {
		api.InternalError(w, r)
		return
	}
	stats := api.GetStats()
	data := struct {
		Posts []api.PostInfo
		Likes []api.PostInfo
		Stats api.Stats
	}{
		Posts: posts,
		Likes: likes,
		Stats: stats,
	}

	tmpl := template.Must(template.ParseFiles("./ui/html/activity.html", "./ui/html/header.html", "./ui/html/footer.html", "./ui/html/banner.html", "./ui/html/search.html"))
	err = tmpl.ExecuteTemplate(w, "activity", data)
	if err != nil {
		api.InternalError(w, r)
		return
	}
}
