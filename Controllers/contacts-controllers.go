package Controllers

import "net/http"

func CreateNewContact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("createNewContact"))
}

func GetAllContacts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getAllContacts"))
}

func GetContactByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getContactByID"))
}

func UpdateContactById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updateContactById"))
}

func DeleteContactByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleteContactByID"))
}

func SearchContacts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("searchContacts"))
}

func ImportContacts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("importContacts"))
}

func ExportContacts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("exportContacts"))
}
