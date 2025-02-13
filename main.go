package main

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Document struct {
	Id    string `json:"id" bson:"_id"` //attribuisce a "bson.ObjectID" la variabile bson "_id" dal DB in fase di decode; guarda "Decode" dopo
	Name  string `json:"name" bson:"name"`
	Value int    `json:"value" bson:"value"`
}

func main() {
	ctx := context.Background()
	uri := "mongodb://localhost:27017/" //var connessione
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	rerr := client.Ping(ctx, nil)
	if rerr != nil {
		panic(rerr)
	}
	coll := client.Database("nuovoDB").Collection("nuovaCartella") //variabile che funge da raccoglitore
	var result Document = Document{}
	filter := bson.M{"name": "Esempio"}

	res := coll.FindOne(ctx, filter)
	if res.Err() != nil {
		panic(res.Err().Error())
	}

	err = res.Decode(&result)
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

}
