package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/martoast/go-polygons/pkg/models"
)

var NewPolygons models.Polygon

func GetPolygonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	polygonId := vars["polygonId"]
	ID, err := strconv.ParseInt(polygonId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	polygonDetails, _ := models.GetPolygonById(ID)
	res, _ := json.Marshal(polygonDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPolygonByLatLng(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Lat := vars["Lat"]
	Lng := vars["Lng"]
	polygonDetails, _ := models.GetPolygonByLatLng(Lat, Lng)
	res, _ := json.Marshal(polygonDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
