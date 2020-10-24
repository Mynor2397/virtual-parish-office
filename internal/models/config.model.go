package models

// Config es la configuracion del servidor
type Config struct {
	PORT       string
	HOSTDB     string
	PORTDB     int
	USERDB     string
	PASSWORDDB string
	DATABASE   string
	ADDRESSMAIL string
	NAMEMAIL string
}
