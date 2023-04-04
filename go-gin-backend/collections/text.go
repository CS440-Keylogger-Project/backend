package collections

import (
	"context"
	"log"
	"time"

	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var textCollection *mongo.Collection = config.OpenCollection(config.Client, "texts")

func RetrieveAllTexts(skip int64, slice int64) (text []models.Text, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: 1}}).SetLimit(slice).SetSkip(skip)

	cursor, err := textCollection.Find(ctx, bson.D{}, opts)
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &text); err != nil {
		panic(err)
	}

	return text, err
}

func CreateText(text models.Text) (result interface{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err = textCollection.InsertOne(ctx, text)
	if err != nil {
        log.Println("Insert error =", err.Error())
		return
    }

	return result, err
}