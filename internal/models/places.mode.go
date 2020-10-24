package models

import (
	"database/sql"
	"encoding/json"

	"github.com/gobuffalo/nulls"
)

type Place struct {
	ID          string       `json:"id,omitempty"`
	Name        nulls.String `json:"name,omitempty"`
	Description nulls.String `json:"description,omitempty"`
}

type NullString struct {
	sql.NullString
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}
