package handlers

import "net/http"

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

}

func SignUp(w *http.ResponseWriter, r *http.Request) {

}

func Login(w *http.ResponseWriter, r *http.Request) {

}

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	return true, ""
}
