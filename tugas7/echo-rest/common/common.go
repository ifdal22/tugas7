package common

type Users struct {
	ID           int    `json:"id,omitempty"`
	NamaDepan    string `json:"nama_depan,omitempty"`
	NamaBelakang string `json:"nama_belakang,omitempty"`
	Email        string `json:"email, omitempty"`
	Username     string `json:"username, omitempty"`
	Password     string `json:"password, omitempty"`
}

type Suppliers struct {
	CompanyName  string `json:"CompanyName, omitempty"`
	ContactName  string `json:"ContactName, omitempty"`
	ContactTitle string `json:"ContactTitle, omitempty"`
	Address      string `json:"Address, omitempty"`
	PostalCode   string `json:"PostalCode, omitempty"`
}
