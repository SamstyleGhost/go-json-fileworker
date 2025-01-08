package models

type geo struct {
	Lat string `json:"lat,omitempty"`
	Lng string `json:"lng,omitempty"`
}

type address struct {
	Street  string `json:"street,omitempty"`
	Suite   string `json:"suite,omitempty"`
	City    string `json:"city,omitempty"`
	ZipCode string `json:"zipcode,omitempty"`
	Geo     geo    `json:"geo,omitempty"`
}

type company struct {
	Name        string `json:"name,omitempty"`
	CatchPhrase string `json:"catchPhrase,omitempty"`
	BS          string `json:"bs,omitempty"`
}

type Users struct {
	ID       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Username string  `json:"username,omitempty"`
	Email    string  `json:"email,omitempty"`
	Address  address `json:"address,omitempty"`
	Phone    string  `json:"phone,omitempty"`
	Website  string  `json:"website,omitempty"`
	Company  company `json:"company,omitempty"`
}
