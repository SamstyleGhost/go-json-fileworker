package models

type geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
	Geo     geo    `json:"geo"`
}

type company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	BS          string `json:"bs"`
}

type Users struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  company `json:"company"`
}
