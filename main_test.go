package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"forum/api"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestCheckSessionHandle(t *testing.T) {
	// create a test session
	sess := createTestSession()

	// create a test request
	req, err := http.NewRequest("GET", "/api/checksession", nil)
	if err != nil {
		t.Fatal(err)
	}

	// set the session cookie
	cookie := http.Cookie{Name: "goforum", Value: sess.UUID}
	req.AddCookie(&cookie)

	// create a test response recorder
	rr := httptest.NewRecorder()

	// call the CheckSessionHandle function
	handler := http.HandlerFunc(api.CheckSessionHandle)
	handler.ServeHTTP(rr, req)

	// check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// check the response body
	expected := `{"Message":"Session is active","Role":"Admin","Ok":true}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \ngot  %q \nwant %q", rr.Body.String(), expected)
	}

	api.DB.Exec("DELETE FROM Session WHERE UUID = 'testsession'")
}
func TestCheckAuthorization(t *testing.T) {
	// create a test session
	sess := createTestSession()

	// create a test request
	req, err := http.NewRequest("GET", "/api/checksession", nil)
	if err != nil {
		t.Fatal(err)
	}

	// set the session cookie
	cookie := http.Cookie{Name: "goforum", Value: sess.UUID}
	req.AddCookie(&cookie)

	// create a test response recorder
	rr := httptest.NewRecorder()

	// call the CheckAuthorization function
	s, err := api.CheckAuthorization(rr, req)

	// check the response status code
	if err != nil {
		t.Errorf("handler returned unexpected error: %v", err)
	}

	// check the session role
	if s.Role != "Admin" {
		t.Errorf("handler returned unexpected role: got %v want %v", s.Role, "Admin")
	}

	// check the session UUID
	if s.UUID != sess.UUID {
		t.Errorf("handler returned unexpected UUID: got %v want %v", s.UUID, sess.UUID)
	}

	api.DB.Exec("DELETE FROM Session WHERE UUID = 'testsession'")
}

func createTestSession() *api.Session {
	api.DB.Exec("INSERT INTO Session (UUID, User_Id, Expires) VALUES ('testsession', 1, '2023-12-01T00:00:00Z')")
	sess := api.Session{
		UUID:    "testsession",
		User_Id: 1,
		Expires: "2023-12-01T00:00:00Z",
	}
	return &sess
}

func TestFormatDate(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2022-01-01T12:00:00Z", "01 Jan 2022, 12:00"},
		{"2023-01-02T12:00:00Z", "02 Jan, 12:00"},
		{time.Now().Format(time.RFC3339), "Today, " + time.Now().Format("15:04")},
		{time.Now().AddDate(0, 0, -1).Format(time.RFC3339), "Yesterday, " + time.Now().AddDate(0, 0, -1).Format("15:04")},
		{"invalid date", "invalid date"},
	}

	for _, test := range tests {
		result := api.FormatDate(test.input)
		if result != test.expected {
			t.Errorf("Expected '%s', but got '%s' for input '%s'", test.expected, result, test.input)
		}
	}
}

func TestLogoutHandle(t *testing.T) {
	// create a test session
	sess := createTestSession()

	// create a test request
	req, err := http.NewRequest("GET", "/api/logout", nil)
	if err != nil {
		t.Fatal(err)
	}

	// set the session cookie
	cookie := http.Cookie{Name: "goforum", Value: sess.UUID}
	req.AddCookie(&cookie)

	// create a test response recorder
	rr := httptest.NewRecorder()

	// call the LogoutHandle function
	handler := http.HandlerFunc(api.LogoutHandle)
	handler.ServeHTTP(rr, req)

	// check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// check the response body
	expected := `{"Message":"Logout successful","Role":"Guest","Ok":true}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \ngot  %q \nwant %q", rr.Body.String(), expected)
	}

	// check that the session was deleted
	_, err = api.GetSession(sess.UUID)
	if err == nil {
		t.Errorf("session was not deleted")
	}
}

func TestLoginHandle(t *testing.T) {
	// create a test user
	password := "test1234"
	pwd, err := api.EncryptPwd("test1234")
	if err != nil {
		t.Fatal(err)
	}
	user := api.User{
		Username: "testuser",
		Password: pwd,
		Email:    "testuser@example.com",
		Avatar:   "testavatar",
		Role_Id:  1,
	}
	_, err = user.Create()
	if err != nil {
		t.Fatal(err)
	}
	defer user.Delete()

	// create a test request
	data := map[string]string{
		"username": user.Username,
		"password": password,
	}
	body, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// create a test response recorder
	rr := httptest.NewRecorder()

	// call the LoginHandle function
	handler := http.HandlerFunc(api.LoginHandle)
	handler.ServeHTTP(rr, req)

	// check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// check the response body
	expected := `{"Message":"Login successful","Role":"","Ok":true}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \ngot  %q \nwant %q", rr.Body.String(), expected)
	}

	// check the session cookie
	cookie := rr.Header().Get("Set-Cookie")
	if cookie == "" {
		t.Errorf("handler did not set session cookie")
	}
	// check that the session was created
	sid := cookie[len("goforum="):strings.Index(cookie, ";")]
	sess, err := api.GetSession(sid)
	if err != nil {
		t.Errorf("handler did not create session: %v", err)
	}
	if sess.User_Id != user.Id {
		t.Errorf("handler created session with wrong user ID: got %v want %v", sess.User_Id, user.Id)
	}
	defer api.DB.Exec("DELETE FROM Session WHERE UUID = ?", sid)

	// test invalid username
	data = map[string]string{
		"username": "invaliduser",
		"password": user.Password,
	}
	body, err = json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	req, err = http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	// test invalid password
	data = map[string]string{
		"username": user.Username,
		"password": "invalidpassword",
	}
	body, err = json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	req, err = http.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}

func TestRegisterHandle(t *testing.T) {
	// create a test request
	form := url.Values{}
	form.Add("signUp__form__username", "testuser")
	form.Add("signUp__form__password", "test1234")
	form.Add("signUp__form__email", "testuser@example.com")
	form.Add("signUp__form__chosenAvatar", "testavatar")
	req, err := http.NewRequest("POST", "/api/register", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// create a test response recorder
	rr := httptest.NewRecorder()

	// call the RegisterHandle function
	handler := http.HandlerFunc(api.RegisterHandle)
	handler.ServeHTTP(rr, req)

	// check the response status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
	var uid int64
	err = api.DB.QueryRow("SELECT Id FROM User WHERE Username = 'testuser'").Scan(&uid)
	if err != nil {
		t.Fatal(err)
	}
	defer api.DB.Exec("DELETE FROM User WHERE Id = ?", uid)
	// check the response body
	expected := fmt.Sprintf("{\"Message\":\"User created (%v) and logged in\",\"Role\":\"user\",\"Ok\":true}\n", uid)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \ngot  %q \nwant %q", rr.Body.String(), expected)
	}

	// check the session cookie
	cookie := rr.Header().Get("Set-Cookie")
	if cookie == "" {
		t.Errorf("handler did not set session cookie")
	}
	// check that the session was created
	sid := cookie[len("goforum="):strings.Index(cookie, ";")]
	sess, err := api.GetSession(sid)
	if err != nil {
		t.Errorf("handler did not create session: %v", err)
	}
	if sess.User_Id != uid {
		t.Errorf("handler created session with wrong user ID: got %v want %v", sess.User_Id, 1)
	}
	defer api.DB.Exec("DELETE FROM Session WHERE UUID = ?", sid)

	// test duplicate registration
	req, err = http.NewRequest("POST", "/api/register", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
	expected = `{"Message":"Username or email already exist","Role":"","Ok":false}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \ngot  %q \nwant %q", rr.Body.String(), expected)
	}

	// test invalid form data
	data := map[string]string{
		"username": "testuser",
		"password": "test1234",
		"email":    "invalidemail",
	}
	body, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	req, err = http.NewRequest("POST", "/api/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
	expected = `{"Message":"Invalid email","Role":"","Ok":false}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: \ngot  %q \nwant %q", rr.Body.String(), expected)
	}
}
