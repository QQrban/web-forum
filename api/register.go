package api

import (
	"fmt"
	"net/http"
	"time"
)

func RegisterHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Check if user is already logged in
	if duplicateSession(w, r) {
		return
	}
	// Get data from form
	u := getUserdataFrom("signUp__form__", w, r)
	user := User{
		Username: u["username"],
		Password: u["password"],
		Email:    u["email"],
		Avatar:   u["chosenAvatar"],
		Role_Id:  1, // 1 = user, 2 = moderator, 3 = admin
	}
	//fmt.Println("user:", user)
	ok, err := user.duplicateRegister() // Check if username or email already exist (again)
	if ok {
		// Check email
		if !checkEmail(user.Email) {
			Respond(w, http.StatusInternalServerError, Response{Message: "Invalid email", Ok: false})
			return
		}
		// Encrypt password
		pwd, err := EncryptPwd(user.Password)
		if err != nil {
			Respond(w, http.StatusInternalServerError, Response{Message: "Password encryption failed: " + err.Error(), Ok: false})
			return
		}
		user.Password = pwd
		// Create new user
		id, err := user.Create()
		if err != nil {
			Respond(w, http.StatusInternalServerError, Response{Message: "User creation failed: " + err.Error(), Ok: false})
			return
		}

		// Create new session
		expires := time.Now().Local().Add(time.Hour) // expires in one hour
		session := Session{
			UUID:    GenerateUUID(),
			Expires: expires.Format("2006-01-02 15:04:05"),
			User_Id: id,
		}
		err = session.Create()
		if err != nil {
			Respond(w,
				http.StatusInternalServerError,
				Response{Message: "Session creation failed: " + err.Error(), Role: "User", Ok: false},
			)
			return
		}

		// Set cookie
		cookie := http.Cookie{
			Name:    "goforum",
			Value:   session.UUID,
			Expires: expires,
			//MaxAge:   0,
			SameSite: 2, // Lax
			//Secure:   true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, &cookie)

		// Return response
		//w.Header().Set("Location", fmt.Sprintf("/%s/%d", "user", id))
		var response = Response{
			Message: fmt.Sprintf("User created (%d) and logged in", id),
			Role:    "user",
			Ok:      true,
		}
		Respond(w, http.StatusCreated, response)
		return
	}
	if err == nil {
		Respond(w, http.StatusInternalServerError, Response{Message: "Username or email already exist", Ok: false})
		return
	}
	Respond(w, http.StatusInternalServerError, Response{Message: "User creation failed: " + err.Error(), Ok: false})
}
