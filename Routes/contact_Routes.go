package Routes

import (
	"net/http"

	"github.com/SureshKumar-U/Contacts_management/controllers"
)

var ContactRoutes = func(router *http.ServeMux) {

	router.HandleFunc("POST /api/v2/contacts", controllers.CreateNewContact)
	router.HandleFunc("GET /api/v2/contacts", controllers.GetAllContacts)
	router.HandleFunc("GET /api/v2/contacts/{id}", controllers.GetContactByID)
	router.HandleFunc("PUT /api/v2/contacts/{id}", controllers.UpdateContactById)
	router.HandleFunc("DELETE /api/v2/contacts/{id}", controllers.DeleteContactByID)
	router.HandleFunc("GET /api/v2/contacts/search", controllers.SearchContacts)  //  search contacts by id, name, email e.t.c
	router.HandleFunc("POST /api/v2/contacts/import", controllers.ImportContacts) //This endpoint would allow users to import contacts from CSV files .
	router.HandleFunc("GET /api/v2/contacts/export", controllers.ExportContacts)  // This endpoint would allow users to export contacts to CSV files.

}
