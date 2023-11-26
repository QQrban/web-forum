package handlers

import (
	"forum/api"
	"html/template"
	"net/http"
	"strconv"
)

func TopicHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idstr := r.URL.Path[len("/topic/"):]
	tid, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		api.NotFound(w, r)
		return
	}
	category, err := api.GetCategoryWithParent(tid)
	if err != nil {
		api.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("./ui/html/topic.html", "./ui/html/header.html", "./ui/html/footer.html", "./ui/html/banner.html", "./ui/html/search.html", "./ui/html/post.html"))
	stats := api.GetStats()
	data := struct {
		Topic api.CategoryWithParent
		Stats api.Stats
	}{
		Topic: category,
		Stats: stats,
	}
	err = tmpl.ExecuteTemplate(w, "topic", data)
	if err != nil {
		api.NotFound(w, r)
		return
	}
}
