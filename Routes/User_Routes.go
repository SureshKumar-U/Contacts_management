package Routes

import (
	"github.com/SureshKumar-U/Contacts_management/Controllers"
	"net/http"
)

var UserRoutes = func(router *http.ServeMux) {
	router.HandleFunc("POST /api/v2/user/register", Controllers.RegisterUser)
	router.HandleFunc("POST /api/v2/login", Controllers.Loginuser)
	router.HandleFunc("GET /api/v2/users/profile", Controllers.GetUserProfile)
	router.HandleFunc("PUT /api/v2/users/profile", Controllers.UpdateUserProfile)
	router.HandleFunc("POST /api/v2/users/forgotpassword", Controllers.ForgotPassword)
	router.HandleFunc("POST /api/v2/users/resetpassword", Controllers.ResetPassword)
	router.HandleFunc("POST /api/v2/logout", Controllers.LogoutUser)
}
