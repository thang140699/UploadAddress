package model

import (
	"time"
)

type File struct {
	// FileName    string    `bson :"fileName" json:"fileName"`
	// ImageName   string    `bson :"imageName"json:"imageName"`
	// Size        int       `bson:"size"json:"size"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
	UpdatedTime time.Time `bson:"updatedTime" json:"updatedTime"`
	URL         string    `bson:"url"json:"url"`
}
