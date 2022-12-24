package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Childrenlocal struct {
	ID            bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Order         string        `json:"order"bson:"order"`
	Name          string        `json:"name" bson:"name"`
	CodeName      string        `json:"codename"bson:"codename"`
	DivisonType   string        `json:"division_type"bson:"division_type"`
	PhoneCode     int           `json:"phoneCode"bson:"phoneCode"`
	ShortCodeName string        `json:"short_codename"bson:"short_codename"`
	Level         int           `json:"level"bson:"level"`
	ParentId      string        `json:"parent_id"bson:"parent_id"`
}
type LocationDetail struct {
	ID            bson.ObjectId   `json:"id,omitempty" bson:"_id,omitempty"`
	Order         string          `json:"order"bson:"order"`
	Name          string          `json:"name" bson:"name"`
	Lat           string          `bson:"lat" json:"lat" `
	Lng           string          `json:"lng"bson:"lng"`
	CodeName      string          `json:"codename"bson:"codename"`
	DivisonType   string          `json:"division_type"bson:"division_type"`
	PhoneCode     int             `json:"phoneCode"bson:"phoneCode"`
	ShortCodeName string          `json:"short_codename"bson:"short_codename"`
	Level         int             `json:"level"bson:"level"`
	ParentId      string          `json:"parent_id"bson:"parent_id"`
	CreatedAt     time.Time       `json:"created_at"bson:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at" bson:"update_at"`
	Children      []Childrenlocal `bson:"children"json:"children"`
}

type Address struct {
	ID    bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Order string        `json:"order"bson:"order"`

	Name          string           `json:"name" bson:"name"`
	Lat           string           `bson:"lat" json:"lat" `
	Lng           string           `json:"lng"bson:"lng"`
	CodeName      string           `json:"codename"bson:"codename"`
	DivisonType   string           `json:"division_type"bson:"division_type"`
	PhoneCode     int              `json:"phoneCode"bson:"phoneCode"`
	ShortCodeName string           `json:"short_codename"bson:"short_codename"`
	Level         int              `json:"level"bson:"level"`
	ParentId      string           `json:"parent_id"bson:"parent_id"`
	CreatedAt     time.Time        `json:"created_at"bson:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at" bson:"update_at"`
	Children      []LocationDetail `bson:"children"json:"children"`
}
