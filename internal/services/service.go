package services

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"nebula-qr/internal/models"
)

type QRService struct {
	collection *mongo.Collection
}

func NewQRService(db *mongo.Database) *QRService {
	return &QRService{
		collection: db.Collection("qr_registry"),
	}
}

func (s *QRService) GenerateQR(text string, expiresIn time.Duration) (*models.QRRegistry, error) {
	if expiresIn < time.Minute {
		return nil, fmt.Errorf("la duración mínima de expiración debe ser 1 minuto")
	}
	
	qrc, err := qrcode.New(text)
	if err != nil {
		return nil, err
	}

	shape := newShape(0.85)

	buf := &writeCloserBuffer{Buffer: bytes.NewBuffer(nil)}
	// w := standard.NewWithWriter(buf, standard.WithQRWidth(uint8(size/10)))
	w := standard.NewWithWriter(
		buf,
		standard.WithCustomShape(shape),
	)

	err = qrc.Save(w)
	if err != nil {
		return nil, err
	}

	expiresAt := time.Now().Add(expiresIn)
	qr := &models.QRRegistry{
		ID:            primitive.NewObjectID(),
		Text:          text,
		ImageData:     buf.Bytes(),
		Format:        "PNG",
		Size:          20,
		CreatedAt:     primitive.NewDateTimeFromTime(time.Now()),
		ExpiresAt:     primitive.NewDateTimeFromTime(expiresAt),
		DownloadCount: 0,
	}

	_, err = s.collection.InsertOne(context.Background(), qr)
	if err != nil {
		return nil, err
	}

	return qr, nil
}

func (s *QRService) GetQR(id string) (*models.QRRegistry, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var qr models.QRRegistry
	err = s.collection.FindOne(context.Background(), primitive.M{"_id": objID}).Decode(&qr)
	if err != nil {
		return nil, err
	}

	// Actualizar contador de descargas
	update := primitive.M{
		"$set": primitive.M{
			"last_downloaded": primitive.NewDateTimeFromTime(time.Now()),
			"download_count":  qr.DownloadCount + 1,
		},
	}
	_, err = s.collection.UpdateOne(context.Background(), primitive.M{"_id": objID}, update)
	if err != nil {
		return nil, err
	}

	return &qr, nil
}
