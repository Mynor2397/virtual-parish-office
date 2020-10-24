package models

type Audit struct {
	IdBaptismHistory string `json:"id_baptism_history"`
	DateEmited       string `json:"date_emited"`
	IdUser           string `json:"id_user"`
	IdBaptism        string `json:"id_baptism"`
	UserName         string `json:"user_name"`
}
