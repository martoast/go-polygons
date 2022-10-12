package models

import (
	"github.com/martoast/go-polygons/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Polygon struct {
	Id        string `gorm:""json:"id"`
	Parcel_id string `json:"parcel_id"`
	Geojson   string `json:"geojson"`
}

func init() {
	config.Connect()
	db = config.GetDB()
}

func GetPolygonById(Id int64) (*Polygon, *gorm.DB) {
	var foundPolygon Polygon
	db := db.Where("ID=?", Id).Find(&foundPolygon)
	return &foundPolygon, db
}

func GetPolygonByLatLng(Lat string, Lng string) (*Polygon, *gorm.DB) {
	var foundPolygon Polygon
	db := db.Select("ST_AsGeoJSON(polygons) as geojson, parcel_id as parcel_id, id as id").Where("ST_Contains(polygons, ST_GeomFromText('POINT(" + Lng + " " + Lat + ")', 4326))").Find(&foundPolygon)
	return &foundPolygon, db
}
