package models

import (
	"github.com/gobuffalo/nulls"
)

type Manager struct {
	IDManager             nulls.String `json:"id_manager,omitempty"`
	FirstnameManager      nulls.String `json:"firstname_manager,omitempty"`
	SecondnameManager     nulls.String `json:"secondname_manager,omitempty"`
	LastnameManager       nulls.String `json:"lastname_manager,omitempty"`
	SecondlastnameManager nulls.String `json:"secondlastname_manager,omitempty"`
}
