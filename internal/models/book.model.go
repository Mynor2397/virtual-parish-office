package models

import "github.com/gobuffalo/nulls"

type Book struct {
	IDBook     int          `json:"id_book,omitempty"`
	NumberBook int          `json:"number_book,omitempty"`
	StartDate  nulls.String `json:"start_date,omitempty"`
	EndDate    nulls.String `json:"end_date,omitempty"`
	Commentary nulls.String `json:"commentary,omitempty"`
	Folios     int          `json:"folios"`
}
