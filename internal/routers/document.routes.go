package routers

import (
	"github.com/gorilla/mux"

	"github.com/Mynor2397/virtual-parish-office/internal/controller"
	"github.com/Mynor2397/virtual-parish-office/internal/service"
	"github.com/Mynor2397/virtual-parish-office/internal/storage"
)

var (
	documentStorage    storage.DocumentStorage       = storage.NewDocumentStorage()
	documentService    service.DocumentService       = service.NewDocumentService(documentStorage)
	documentController controller.DocumentController = controller.NewDocumentController(documentService)
)

// SetDocumentRoutes registra la rutas a usar para los controladires de usuario
func SetDocumentRoutes(router *mux.Router) *mux.Router {
	document := router.PathPrefix("/documents").Subrouter()
	document.HandleFunc("/folio/{idbook}", documentController.GetFolio).Methods("GET")
	document.HandleFunc("/baptisms", documentController.Baptism).Methods("POST")
	document.HandleFunc("/baptisms", documentController.GetBaptisms).Methods("GET")
	document.HandleFunc("/books", documentController.GetBook).Methods("GET")
	document.HandleFunc("/books", documentController.CreateBook).Methods("POST")
	document.HandleFunc("/mail/send", controller.SendMail).Methods("POST")
	document.HandleFunc("/books/all", documentController.GetCountBook).Methods("GET")

	document.HandleFunc("/audits", documentController.CreateAudi).Methods("POST")
	document.HandleFunc("/audits", documentController.GetAudit).Methods("GET")

	return router
}
