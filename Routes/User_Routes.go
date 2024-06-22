package Routes

import (
	"net/http"

	"github.com/SureshKumar-U/Contacts_management/controllers"
)

var UserRoutes = func(router *http.ServeMux) {
	router.HandleFunc("POST /api/v2/user/register", controllers.RegisterUser)
	router.HandleFunc("POST /api/v2/login", controllers.Loginuser)
	router.HandleFunc("GET /api/v2/users/profile", controllers.GetUserProfile)
	router.HandleFunc("PUT /api/v2/users/profile", controllers.UpdateUserProfile)
	router.HandleFunc("POST /api/v2/users/forgotpassword", controllers.ForgotPassword)
	router.HandleFunc("POST /api/v2/users/resetpassword", controllers.ResetPassword)
	router.HandleFunc("POST /api/v2/logout", controllers.LogoutUser)
}
