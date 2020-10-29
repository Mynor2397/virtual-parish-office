package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Mynor2397/virtual-parish-office/internal/lib"
	"github.com/Mynor2397/virtual-parish-office/internal/middleware"
	"github.com/Mynor2397/virtual-parish-office/internal/models"
)

func (*personController) GetPlaces(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	data, err := personService.GetPlaces(r.Context())

	if err == lib.ErrNotFound {
		respond(w, response{
			Ok:      false,
			Message: lib.ErrNotFound.Error(),
		}, http.StatusOK)
		return
	}

	if err == nil {
		respond(w, response{
			Ok:   true,
			Data: data,
		}, http.StatusOK)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (*personController) CreatePlace(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()
	var place models.Place
	if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
		respond(w, response{
			Ok:      false,
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	res, err := personService.CreatePlace(r.Context(), place)

	if err == nil {
		respond(w, response{
			Ok:   true,
			Data: res,
		}, http.StatusOK)
		return
	}
	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (*personController) GetCountPlace(w http.ResponseWriter, r *http.Request) {
	id, ok := middleware.IsAuthenticated(r.Context())
	if !ok {
		respond(w, response{Message: lib.ErrUnauthenticated.Error()}, http.StatusUnauthorized)
		return
	}
	log.Println(id)
	data, err := personService.GetCountPlace(r.Context())

	if err == lib.ErrNotFound {
		respond(w, response{
			Ok:      false,
			Message: lib.ErrNotFound.Error(),
		}, http.StatusNotFound)
		return
	}

	if err == nil {
		respond(w, response{
			Ok:   true,
			Data: data,
		}, http.StatusOK)
		return
	}

	if err != nil {
		respondError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
