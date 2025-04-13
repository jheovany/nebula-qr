package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type QRRegistry struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Text           string             `bson:"text" json:"text"`
	ImageData      []byte             `bson:"image_data,omitempty" json:"image_data"`
	Format         string             `bson:"format" json:"format"` // PNG o SVG
	Size           int                `bson:"size" json:"size"`     // Tamaño en píxeles
	CreatedAt      primitive.DateTime `bson:"created_at" json:"created_at"`
	ExpiresAt      primitive.DateTime `bson:"expires_at,omitempty" json:"expires_at"`
	LastDownloaded primitive.DateTime `bson:"last_downloaded,omitempty" json:"last_downloaded"`
	DownloadCount  int                `bson:"download_count" json:"download_count"` // Número de descargas
}
