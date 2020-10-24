package models

// Person implements all propertyes for person enntity
type Person struct {
	ID             string `json:"id,omitempty"`
	Firstname      string `json:"firstname,omitempty"`
	Secondname     string `json:"secondname,omitempty"`
	Thirdname      string `json:"thirdname,omitempty"`
	Lastname       string `json:"lastname,omitempty"`
	Secondlastname string `json:"secondlastname,omitempty"`
	DPI            string `json:"dpi,omitempty"`
	Sexo           string `json:"sexo,omitempty"`
	IDAddress      string `json:"id_address,omitempty"`
	Address        string `json:"address,omitempty"`
	Component      string `json:"component,omitempty"`
	Count          int    `json:"count,omitempty"`
}
