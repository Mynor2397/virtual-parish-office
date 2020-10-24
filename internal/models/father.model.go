package models

import (
	"github.com/gobuffalo/nulls"
)

type Father struct {
	IDFather             nulls.String `json:"id_father,omitempty"`
	FirstnameFather      nulls.String `json:"firstname_father,omitempty"`
	SecondnameFather     nulls.String `json:"secondname_father,omitempty"`
	LastnameFather       nulls.String `json:"lastname_father,omitempty"`
	SecondlastnameFather nulls.String `json:"secondlastname_father,omitempty"`
}
