package models

import (
	"github.com/gobuffalo/nulls"
)

type Mother struct {
	IDMother             nulls.String `json:"id_mother,omitempty"`
	FirstnameMother      nulls.String `json:"firstname_mother,omitempty"`
	SecondnameMother     nulls.String `json:"secondname_mother,omitempty"`
	LastnameMother       nulls.String `json:"lastname_mother,omitempty"`
	SecondlastnameMother nulls.String `json:"secondlastname_mother,omitempty"`
}
