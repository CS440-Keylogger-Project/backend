package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBinstance()

func DBinstance() (client *mongo.Client) {

	user := os.Getenv("MONGO_USERNAME")
	pass := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	conn := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, pass, host, port)

	replicaSetQueryString := "/?replicaSet=replica-set"
	tlsQueryString := ""
	secondaryQueryString := ""
	//var tlsConfig *tls.Config

	if os.Getenv("GIN_MODE") == "release" {
		replicaSetQueryString = "/?replicaSet=rs0"
		//tlsQueryString = "&tls=true"
		secondaryQueryString = "&readPreference=secondaryPreferred&retryWrites=false"

		//// configure tls
		//var filename = "rds-combined-ca-bundle.pem"
		//tlsConfig := new(tls.Config)
		//certs, err := ioutil.ReadFile(filename)

		//if err != nil {
		//fmt.Println("Failed to read CA file")
		//return
		//}

		//tlsConfig.RootCAs = x509.NewCertPool()
		//ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

		//if !ok {
		//fmt.Println("Failed to append CA file")
		//return
		//}

		//if tlsConfig != nil {
		//fmt.Println("Successfully set TLS config")
		//clientOptions.SetTLSConfig(tlsConfig)
		//}

	}
	conn = fmt.Sprintf("%s%s%s%s", conn, replicaSetQueryString, tlsQueryString, secondaryQueryString)

	fmt.Printf("Attempting connection with: %s\n", conn)
	//serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	//clientOptions := options.Client().ApplyURI(conn).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect to mongodb
	//client, err := mongo.Connect(ctx, clientOptions)
	//if err != nil {
	//log.Fatal(err)
	//}
	client, err := mongo.NewClient(options.Client().ApplyURI(conn))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Println("Connecting to MongoDB ...")
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	fmt.Println("Success!")
	// Force a connection to verify our connection string

	fmt.Println("Pinging server ...")
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}
	fmt.Println("Success!")

	fmt.Println("Initialising indexes ...")
	// initialise indexes
	InitIndexes(client)
	fmt.Println("Success!")
	return client
}

func InitIndexes(client *mongo.Client) {

	// texts_texts_-1 index
	textCollection := OpenCollection(client, "text")

	textIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "windows_name", Value: -1}},
		Options: options.Index().SetUnique(false),
	}
	textIndexCreated, err := textCollection.Indexes().CreateOne(context.Background(), textIndexModel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created Text Index %s\n", textIndexCreated)
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("keylogger").Collection(collectionName)

	return collection
}
