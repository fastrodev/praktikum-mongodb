package main

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	uri := "mongodb+srv://admin:admin@cluster0.xtwwu.mongodb.net"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}

// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func main() {
// 	clientOptions := options.Client().
// 		ApplyURI("mongodb+srv://admin:admin@cluster0.u16np.mongodb.net")
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	coll := client.Database("sample_mflix").Collection("movies")
// 	title := "Back to the Future"
// 	var result bson.M
// 	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
// 	if err == mongo.ErrNoDocuments {
// 		fmt.Printf("No document was found with the title %s\n", title)
// 		return
// 	}
// 	if err != nil {
// 		panic(err)
// 	}
// 	jsonData, err := json.MarshalIndent(result, "", "    ")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", jsonData)

// }
