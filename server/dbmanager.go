package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var URL string = ""

func QueryAllSkills() []Skill {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	skillList := []Skill{}

	collection := client.Database("EquipmentInfo").Collection("skills")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		// Try to unmarshall directly into the skill struct
		var s Skill
		err := cursor.Decode(&s)
		if err != nil {
			log.Fatal(err)
			continue
		}
		fmt.Printf("Unmarshalled struct \n%+v\n", s)
		skillList = append(skillList, s)
	}
	return skillList
}

func QueryAllArmour() []*ArmourSet {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	armourList := []*ArmourSet{}
	collection := client.Database("EquipmentInfo").Collection("armour")

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"rating", 1}})

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	lastSet := "None"
	var currentSet *ArmourSet
	for cursor.Next(ctx) {
		// Try to unmarshall directly into the skill struct
		var a ArmourPiece
		err := cursor.Decode(&a)
		if err != nil {
			log.Fatal(err)
			continue
		}

		// when the set name changes we need to start a new set
		if a.SetName != lastSet {
			currentSet = NewArmourSet()
			currentSet.SetName = a.SetName
			armourList = append(armourList, currentSet)
		}
		currentSet.Pieces = append(currentSet.Pieces, a)
	}

	// db query logic...

	return armourList
}
