package handlers

import (
	"encoding/json"
	"forum/api"
	"net/http"
)

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/api/home" {
		api.NotFound(w, r)
		return
	}
	if r.Method != "GET" {
		api.BadRequest(w, r)
		return
	}

	categories := api.GetCategoryStructure()
	stats := api.GetStats()
	data := struct {
		Categories  []api.CategoryWithChildren `json:"categories"`
		Stats       api.Stats                  `json:"stats"`
		TopCoders   []api.UserInfo             `json:"topCoders"`
		LatestPosts []api.PostWithAuthor       `json:"latestPosts"`
		HotTopics   []api.PostWithAuthor       `json:"hotTopics"`
	}{
		Categories:  categories,
		Stats:       stats,
		TopCoders:   api.GetTopCoders(),
		LatestPosts: api.GetLatestPosts(),
		HotTopics:   api.GetHotTopics(),
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		api.InternalError(w, r)
	}
}
