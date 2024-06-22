package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Contacts struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Company string `json:"company"`
}

func (u *User) CreateUser() *User {
	return &User{}
}

// func (u *User) ExistedUser() bool {
// 	query := "select email from userdetails where email= ?"
// 	row := config.Db.QueryRow(query, u.Email)
// 	if err := row.Scan(&email); err != sql.ErrNoRows { // if now rows found with email
// 		return false
// 	}
// 	return true
// }
