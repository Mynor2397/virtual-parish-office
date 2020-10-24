package controller

import (
	"encoding/json"
	"github.com/Mynor2397/virtual-parish-office/internal/helper"
	"github.com/Mynor2397/virtual-parish-office/internal/models"
	"net/http"
)

func SendMail(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	person := models.Person{}

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	err := helper.SendEmail(person)
	if err != nil {
		respondError(w, err)
		return
	}

	if err == nil {
		respond(w, response{
			Ok:      true,
			Message: "Enviado correctamente",
		}, http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
