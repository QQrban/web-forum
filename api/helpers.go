package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func getUserdataFrom(form string, w http.ResponseWriter, r *http.Request) map[string]string {
	err := r.ParseForm()
	if err != nil {
		InternalError(w, r)
	}
	var u = map[string]string{}
	for k, v := range r.Form {
		k = k[len(form):]
		u[k] = v[0]
	}
	return u
}

func GenerateUUID() string {
	return fmt.Sprintf("%x", uuid.NewV4())
}

func EncryptPwd(pwd string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(pwd), 4) // 4 = cost (between 4 and 31?)
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

func Respond(w http.ResponseWriter, status int, data any) { //data Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func getGuestCookie() *http.Cookie {
	cookie := http.Cookie{
		Name:     "goforum",
		Value:    GUEST,
		Expires:  time.Now().Local(),
		SameSite: 2, // Lax
		//Secure:   true,
		HttpOnly: true,
		Path:     "/",
	}
	return &cookie
}

func FormatDate(date string) string {
	tm, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return date
	}
	if tm.YearDay() == time.Now().Local().YearDay() {
		date = tm.Format("Today, 15:04")
	} else if tm.YearDay() == time.Now().Local().AddDate(0, 0, -1).YearDay() {
		date = tm.Format("Yesterday, 15:04")
	} else if tm.Year() == time.Now().Local().Year() {
		date = tm.Format("02 Jan, 15:04")
	} else {
		date = tm.Format("02 Jan 2006, 15:04")
	}
	return date
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("ui/html/404.html"))
	tmpl.ExecuteTemplate(w, "404", r.URL.Path)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("ui/html/400.html"))
	tmpl.ExecuteTemplate(w, "400", r.URL.Path)
}

func InternalError(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("ui/html/500.html"))
	tmpl.ExecuteTemplate(w, "500", r.URL.Path)
}

func Error(w http.ResponseWriter, r *http.Request, code, message string) {
	data := struct {
		Code    string
		Message string
	}{
		Code:    code,
		Message: message,
	}
	tmpl := template.Must(template.ParseFiles("ui/html/error.html"))
	tmpl.ExecuteTemplate(w, "error", data)
}

func checkEmail(email string) bool {
	rex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return rex.MatchString(email)
}
