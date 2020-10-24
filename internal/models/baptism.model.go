package models

// Baptism implementa todas las caracteristicas del un bautismo como
// de de las partes asociadas a el
type Baptism struct {
	Position       int    `json:"position"`
	NumberBaptism  int    `json:"number_baptism"`
	IDBaptized     string `json:"id_baptized,omitempty"`
	Firstname      string `json:"firstname,omitempty"`
	Secondname     string `json:"secondname,omitempty"`
	Lastname       string `json:"lastname,omitempty"`
	Secondlastname string `json:"secondlastname,omitempty"`
	Borndate       string `json:"borndate,omitempty"`
	DPI            string `json:"dpi,omitempty"`
	Sex            string `json:"sex,omitempty"`
	IDBaptism      string `json:"id_baptism,omitempty"`
	Folio          int    `json:"folio,omitempty"`
	Book           int    `json:"book,omitempty"`
	BaptismDate    string `json:"baptism_date,omitempty"`
	Priest         `json:"priest,omitempty"`
	Place          `json:"place,omitempty"`
	IDAddress      string `json:"id_address,omitempty"`
	Address        string `json:"address,omitempty"`
	Father         `json:"father,omitempty"`
	Mother         `json:"mother,omitempty"`
	Godfather      `json:"godfather,omitempty"`
	Godmother      `json:"godmother,omitempty"`
	Manager        `json:"manager,omitempty"`
}
