package api

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		BadRequest(w, r)
		return
	}
	// Check if user is already logged in
	if duplicateSession(w, r) {
		return
	}
	// Get data from form
	user := &User{}
	json.NewDecoder(r.Body).Decode(user)
	passwd := user.Password
	// Check if user exists
	err := user.GetByUsername()
	if err != nil {
		Respond(w, http.StatusNotFound, Response{Message: "Login failed: username Not found", Ok: false})
		return
	}
	// Check if password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwd))
	if err != nil {
		Respond(w, http.StatusUnauthorized, Response{Message: "Login failed: wrong password", Ok: false})
		return
	}
	// Delete all sessions of this user in other devices
	err = DeleteAllSessionsFor(user.Id)
	if err != nil {
		Respond(w, http.StatusUnauthorized, Response{Message: "Login failed: " + err.Error(), Ok: false})
		return
	}
	// Create new session
	expires := time.Now().Local().Add(time.Hour) // expires in one hour
	session := Session{
		UUID:    GenerateUUID(),
		Expires: expires.Format("2006-01-02 15:04:05"),
		User_Id: user.Id,
	}
	err = session.Create()
	if err != nil {
		Respond(w, http.StatusInternalServerError, Response{Message: "Session creation failed: " + err.Error(), Ok: false})
		return
	}
	// Set cookie
	cookie := http.Cookie{
		Name:     "goforum",
		Value:    session.UUID,
		Expires:  expires,
		HttpOnly: true,
		SameSite: 2, // Lax
		//Secure:   true,
		Path: "/",
	}
	http.SetCookie(w, &cookie)
	// Return response
	Respond(w, http.StatusOK, Response{Message: "Login successful", Ok: true})
}
