package routes

import (
	"github.com/gorilla/mux"
	"github.com/martoast/go-polygons/pkg/controllers"
)

var RegisterPolygonRoutes = func(router *mux.Router) {
	router.HandleFunc("/polygons/{polygonId}", controllers.GetPolygonById).Methods("GET")
	router.HandleFunc("/polygons/{Lat}/{Lng}", controllers.GetPolygonByLatLng).Methods("GET")
}
