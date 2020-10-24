package middleware

import (
	"context"
	"github.com/Mynor2397/virtual-parish-office/internal/lib"
	"github.com/Mynor2397/virtual-parish-office/internal/models"
	"log"
	"net/http"
	"strings"

	"github.com/Mynor2397/virtual-parish-office/internal/helper"
)

type claim string

var (
	//Claims contendrá el rol del usuario en el contexto
	Claims claim = "claims"
)

// Auth middleware de autenticación
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authorizationHeader, "Bearer") {
			next.ServeHTTP(w, r)
			return
		}

		var (
			idUser  string
			rolUser string
		)

		idUser, rolUser, err := helper.AuthUserID(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// log.Println("Middleware passed auth")
		ctx := r.Context()
		ctx = context.WithValue(ctx, Claims, models.User{ID: idUser, Rol: rolUser})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AuthForAmdmin es el middleware que protege las rutas unicamente disponibles para
// modo administrador
func AuthForAmdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware passed admin")
		User, ok := r.Context().Value(Claims).(models.User)
		if !ok {
			http.Error(w, lib.ErrUnauthenticated.Error(), http.StatusUnauthorized)
			return
		}

		if User.Rol != "admin" {
			http.Error(w, lib.ErrUnauthenticated.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// IsAuthenticated evalua el contexto de la peticion
func IsAuthenticated(ctx context.Context) (string, bool) {
	uid, ok := ctx.Value(Claims).(models.User)
	if !ok {
		return uid.ID, false
	}

	return uid.ID, true
}
