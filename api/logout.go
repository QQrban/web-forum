package api

import (
	"net/http"
)

func LogoutHandle(w http.ResponseWriter, r *http.Request) {
	sid, err := r.Cookie("goforum")
	if err != nil {
		Respond(w,
			http.StatusOK,
			Response{
				Message: "Session not found: user has logged out",
				Ok:      true,
			})
		return
	}
	DeleteSession(sid.Value)
	cookie := getGuestCookie()
	http.SetCookie(w, cookie)
	Respond(w, http.StatusOK, Response{Message: "Logout successful", Role: "Guest", Ok: true})
}
