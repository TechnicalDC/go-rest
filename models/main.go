package models

type User struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	City		string	`json:"city"`
	Dob		string	`json:"dateofbirth"`
	Email 	string	`json:"email"`
	Password string	`json:"-"`
	Active	bool  	`json:"active"`
}

type Users struct {
	User		[]User	`json:"users"`
}

type Config struct {
	Ip			string
	Port		string
	DbIp		string
	DbPort	string
	DbUser	string
	DbPass	string
}
