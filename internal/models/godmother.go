package models

import (
	"github.com/gobuffalo/nulls"
)

type Godmother struct {
	IDGodMother             nulls.String `json:"id_godmother,omitempty"`
	FirstnameGodmother      nulls.String `json:"firstname_godmother,omitempty"`
	SecondnameGodmother     nulls.String `json:"secondname_godmother,omitempty"`
	LastnameGodmother       nulls.String `json:"lastname_godmother,omitempty"`
	SecondlastnameGodmother nulls.String `json:"secondlastname_godmother,omitempty"`
}
