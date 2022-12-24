package main

import (
	"WeddingUtilities/app/api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPIv1(Container *Container) http.Handler {
	router := api.NewRouter()
	v1 := router.Group("/api/v1")

	AddressRouter(v1)

	return router
}
func UploadFileRouter() {
	// route.Run()

	router := mux.NewRouter()
	router.HandleFunc("/api/upload", api.HandleUpload).Methods("POST")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}

func AddressRouter(parent *api.Router) {
	AddressHandler := api.AddressHandler{
		AddressRepository: container.AddressRepository,
	}
	router := parent.Group("/")
	router.GET("/detail", AddressHandler.GetByID)
	router.GET("", AddressHandler.GetAll)
	router.GET("/detail/name", AddressHandler.GetByName)
	router.GET("/detail/phoneCode", AddressHandler.GetByPhoneCode)
	router.GET("/detail/codeName", AddressHandler.GetByCodeName)
	router.GET("/detail/divisiontype", AddressHandler.GetByDivisiontype)
	router.GET("/detail/level", AddressHandler.GetByLevel)
	router.GET("/detail/parentId", AddressHandler.GetByParentId)
}
