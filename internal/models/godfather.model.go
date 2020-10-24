package models

import (
	"github.com/gobuffalo/nulls"
)

type Godfather struct {
	IDGodfather             nulls.String `json:"id_godfather,omitempty"`
	FirstnameGodfather      nulls.String `json:"firstname_godfather,omitempty"`
	SecondnameGodfather     nulls.String `json:"secondname_godfather,omitempty"`
	LastnameGodfather       nulls.String `json:"lastname_godfather,omitempty"`
	SecondlastnameGodfather nulls.String `json:"secondlastname_godfather,omitempty"`
}
