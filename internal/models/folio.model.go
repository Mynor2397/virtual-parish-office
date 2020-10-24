package models

type Folio struct {
	IDFolio     int64 `json:"id_folio,omitempty"`
	NumberFolio int64 `json:"number_folio,omitempty"`
	IDBook      int64 `json:"id_book,omitempty"`
}
