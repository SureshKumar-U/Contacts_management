package Routes

import (
	"github.com/SureshKumar-U/Contacts_management/Controllers"
	"net/http"
)

var ContactRoutes = func(router *http.ServeMux) {
	router.HandleFunc("POST /api/v2/contacts", Controllers.CreateNewContact)
	router.HandleFunc("GET /api/v2/contacts", Controllers.GetAllContacts)
	router.HandleFunc("GET /api/v2/contacts/{id}", Controllers.GetContactByID)
	router.HandleFunc("PUT /api/v2/contacts/{id}", Controllers.UpdateContactById)
	router.HandleFunc("DELETE /api/v2/contacts/{id}", Controllers.DeleteContactByID)
	router.HandleFunc("GET /api/v2/contacts/search", Controllers.SearchContacts)  //  search contacts by id, name, email e.t.c
	router.HandleFunc("POST /api/v2/contacts/import", Controllers.ImportContacts) //This endpoint would allow users to import contacts from CSV files .
	router.HandleFunc("GET /api/v2/contacts/export", Controllers.ExportContacts)  // This endpoint would allow users to export contacts to CSV files.

}
