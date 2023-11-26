package api

import (
	"errors"
	"net/http"
	"time"
)

const (
	GUEST          = "36626261343036362d376434612d343430322d393335612d313562643939386439386337"
	CHECK_DURATION = time.Minute * 1 //5
)

var (
	lastCheck = time.Now()
)

func CheckSessionHandle(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("check: %v\nlastCheck: %v\nNow: %v\nSince: %v\nCHK_DUR: %v\nlastCheck+dur: %v\n", time.Since(lastCheck) >= CHECK_DURATION, lastCheck, time.Now().Local(), time.Since(lastCheck), CHECK_DURATION, lastCheck.Add(CHECK_DURATION))
	if time.Since(lastCheck) >= CHECK_DURATION {
		err := flushExpiredSessions()
		if err != nil {
			InternalError(w, r)
			return
		}
		lastCheck = time.Now()
	}
	// check session
	sid, err := r.Cookie("goforum")
	//fmt.Println("sid:", sid)
	if err != nil || sid.Value == GUEST { // if cookie does not exist or is guest
		cookie := getGuestCookie()
		http.SetCookie(w, cookie)
		Respond(w, http.StatusOK, Response{Message: "Guest session", Role: "Guest", Ok: false})
		return
	}
	s, err := getSessionWithRole(sid.Value)
	// if session is valid, renew cookie and return true
	if err == nil {
		if s.Expires < time.Now().Local().Format("2006-01-02T15:04:05Z") {
			Respond(w, http.StatusUnauthorized, Response{Message: "Session has expired", Role: s.Role, Ok: false})
			return
		}
		expires := time.Now().Local().Add(time.Hour * 24)
		cookie := http.Cookie{
			Name:     "goforum",
			Value:    sid.Value,
			Expires:  expires,
			SameSite: 2, // Lax
			//Secure:   true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, &cookie)
		Respond(w, http.StatusOK, Response{Message: "Session is active", Role: s.Role, Ok: true})
		return
	}
	// if session is invalid, delete cookie and return false
	cookie := getGuestCookie()
	http.SetCookie(w, cookie)
	Respond(w, http.StatusNotFound, Response{Message: "Session not found: user is logged out", Role: "Guest", Ok: false})
}

func duplicateSession(w http.ResponseWriter, r *http.Request) bool {
	sid, err := r.Cookie("goforum")
	if err == nil && sid.Value != GUEST { // if cookie exists and is not guest
		s := Session{UUID: sid.Value}
		err = s.Get()
		if err != nil { // if session does not exist in DB
			cookie := getGuestCookie()
			http.SetCookie(w, cookie)
			Respond(w, http.StatusOK, Response{Message: "Session not found, user logged out", Ok: true})
			return false
		}
		if s.Expires < time.Now().Local().Format("2006-01-02T15:04:05Z") { // if session has expired
			expires := time.Now().Local().Add(time.Hour * 24)
			cookie := http.Cookie{
				Name:     "goforum",
				Value:    sid.Value,
				Expires:  expires,
				SameSite: 2, // Lax
				//Secure:   true,
				HttpOnly: true,
				Path:     "/",
			}
			http.SetCookie(w, &cookie)
			s.Expires = expires.Format("2006-01-02T15:04:05Z")
			s.Update()
			return false
		}
		Respond(w, http.StatusUnauthorized, Response{Message: "Already logged in", Ok: false})
		return true
	}
	return false
}

func CheckAuthorization(w http.ResponseWriter, r *http.Request) (SessionWithRole, error) {
	sid, err := r.Cookie("goforum")
	if err != nil || sid.Value == GUEST {
		return SessionWithRole{}, errors.New("Unauthorized")
	}
	s, err := getSessionWithRole(sid.Value)
	if err != nil || s.Role == "Guest" {
		return SessionWithRole{}, errors.New("Unauthorized")
	}
	if s.Expires < time.Now().Local().Format("2006-01-02 15:04:05") {
		return SessionWithRole{}, errors.New("Session has expired")
	}
	return s, nil
}
