package handlers

import (
	"forum/api"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func PostyTosty(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		api.NotFound(w, r)
		return
	}

	basePath := path.Base(r.URL.Path)

	cid, err := strconv.ParseInt(basePath, 10, 64)
	if err != nil {
		api.NotFound(w, r)
		return
	}

	post, err := api.GetPostWithAuthor(cid)
	if err != nil {
		api.NotFound(w, r)
		return
	}

	stats := api.GetStats()

	cookie, err := r.Cookie("goforum")
	var current any
	if err != nil {
		current = nil
	} else {
		uuid := cookie.Value
		s, err := api.GetSession(uuid)
		if err != nil {
			current = nil
		}
		current, err = api.GetUserInfo(s.User_Id)
		if err != nil {
			current = nil
		}
	}

	data := struct {
		Post    api.PostWithAuthor
		Stats   api.Stats
		Current any
	}{
		Post:    post,
		Stats:   stats,
		Current: current,
	}

	tmpl := template.Must(template.ParseFiles("./ui/html/post.html", "./ui/html/header.html", "./ui/html/footer.html", "./ui/html/banner.html", "./ui/html/search.html"))

	err = tmpl.ExecuteTemplate(w, "post", data)
	if err != nil {
		api.NotFound(w, r)
		log.Printf("Error fetching post with ID %d: %v", cid, err)
	}
}
