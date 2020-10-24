package models

import "github.com/gobuffalo/nulls"

type Priest struct {
	IDPriest             string `json:"id_priest,omitempty"`
	FirstnamePriest      string `json:"firstname_priest,omitempty"`
	SecondnamePriest     string `json:"secondname_priest,omitempty"`
	LastnamePriest       string `json:"lastname_priest,omitempty"`
	SecondLastnamePriest string `json:"second_lastname_priest,omitempty"`
	Credentials          nulls.String `json:"credentials,omitempty"`
	ParishOrigin         string `json:"parish_origin,omitempty"`
}
