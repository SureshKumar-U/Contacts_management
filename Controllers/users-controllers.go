package Controllers

import (
	"github.com/SureshKumar-U/Contacts_management/models"
	"github.com/SureshKumar-U/Contacts_management/utils"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	utils.ParseBody(r, user)
}
func Loginuser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("loginuser"))
}
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getUserProfile"))
}
func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updateUserProfile"))
}
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("forgotPassword"))
}
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("resetPassword"))
}
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logoutUser"))
}
