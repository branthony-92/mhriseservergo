package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WeightedParam struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
}

type OptimizationFilters struct {
	Skills      []WeightedParam `json:"skills"`
	Resistences []WeightedParam `json:"resistences"`
	MinRarity   int             `json:"min_rarity"`
	MaxRarity   int             `json:"max_rarity"`
	ExcludeSets []string        `json:"exclude_sets"`

	RemainingPoints map[string]int
	MaxPoints       map[string]int
}

func isInList(item string, list []string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

var URL string = ""

var cachedSkills []Skill = make([]Skill, 0)
var cachedHelms []ArmourPiece = make([]ArmourPiece, 0)
var cachedMails []ArmourPiece = make([]ArmourPiece, 0)
var cachedCoils []ArmourPiece = make([]ArmourPiece, 0)
var cachedVambraces []ArmourPiece = make([]ArmourPiece, 0)
var cachedGreaves []ArmourPiece = make([]ArmourPiece, 0)

func QueryAllSkills() ([]Skill, error) {
	if len(cachedSkills) != 0 {
		return cachedSkills, nil
	}

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
		return skillList, err
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
	cachedSkills = skillList
	return cachedSkills, nil
}

func QueryAllArmour() ([]*ArmourSet, error) {
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
		return armourList, err
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
	return armourList, nil
}

func QueryFilter(q []byte) ([]ArmourPiece, error) {
	pieces := []ArmourPiece{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database("EquipmentInfo").Collection("armour")

	var filter bson.M

	if err := json.Unmarshal(q, &filter); err != nil {
		return pieces, err
	}

	opts := options.Find().SetSort(bson.D{{"rating", 1}})

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return pieces, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		// Try to unmarshall directly into the skill struct
		var a ArmourPiece
		err := cursor.Decode(&a)
		if err != nil {
			log.Fatal(err)
			continue
		}
		pieces = append(pieces, a)
	}
	return pieces, nil
}

func Optimize(filters []byte) (*ArmourSet, error) {
	set := ArmourSet{}

	skills, err := QueryAllSkills()
	if err != nil {
		return nil, err
	}

	// get a list of each individual piece type
	if len(cachedHelms) == 0 {
		cachedHelms, err = QueryFilter([]byte(`{"piece_type":"helm"}`))
		if err != nil {
			return nil, err
		}
	}
	if len(cachedMails) == 0 {
		cachedMails, err = QueryFilter([]byte(`{"piece_type":"mail"}`))
		if err != nil {
			return nil, err
		}
	}
	if len(cachedCoils) == 0 {
		cachedCoils, err = QueryFilter([]byte(`{"piece_type":"coil"}`))
		if err != nil {
			return nil, err
		}
	}
	if len(cachedVambraces) == 0 {
		cachedVambraces, err = QueryFilter([]byte(`{"piece_type":"vambraces"}`))
		if err != nil {
			return nil, err
		}
	}
	if len(cachedGreaves) == 0 {
		cachedGreaves, err = QueryFilter([]byte(`{"piece_type":"greaves"}`))
		if err != nil {
			return nil, err
		}
	}

	f := OptimizationFilters{}

	if ok := json.Valid(filters); !ok {
		return nil, fmt.Errorf("Input filter structure is invalid")
	}
	// the filter will be coming in as a json string so unmarshall it
	if err := json.Unmarshal([]byte(filters), &f); err != nil {
		return nil, err
	}

	f.MaxPoints = make(map[string]int)
	f.RemainingPoints = make(map[string]int)
	for _, s := range f.Skills {
		skillObj := FindSkill(skills, s.Name)
		if skillObj == nil {
			continue
		}
		f.RemainingPoints[s.Name] = skillObj.SkillLevelMax
		f.MaxPoints[s.Name] = skillObj.SkillLevelMax
	}

	// using the filters, and the next list calculate a weight value representing each armour piece to select the next best part
	bestHelm, err := calculateBest(cachedHelms, &f)
	if err != nil {
		return nil, err
	}
	bestMail, err := calculateBest(cachedMails, &f)
	if err != nil {
		return nil, err
	}
	bestCoil, err := calculateBest(cachedCoils, &f)
	if err != nil {
		return nil, err
	}
	bestVambraces, err := calculateBest(cachedVambraces, &f)
	if err != nil {
		return nil, err
	}
	bestGreaves, err := calculateBest(cachedGreaves, &f)
	if err != nil {
		return nil, err
	}

	set.Pieces = append(set.Pieces, *bestHelm)
	set.Pieces = append(set.Pieces, *bestMail)
	set.Pieces = append(set.Pieces, *bestCoil)
	set.Pieces = append(set.Pieces, *bestVambraces)
	set.Pieces = append(set.Pieces, *bestGreaves)

	// return the custom set
	return &set, nil
}

func calculateBest(pieces []ArmourPiece, filters *OptimizationFilters) (*ArmourPiece, error) {
	if len(pieces) == 0 {
		return nil, fmt.Errorf("No Armour in setType")
	}

	p := ArmourPiece{}
	bestsScore := 0.0
	for _, piece := range pieces {

		if piece.Rarity < filters.MinRarity {
			continue
		}
		if piece.Rarity > filters.MaxRarity {
			continue
		}
		if isInList(piece.SetName, filters.ExcludeSets) {
			continue
		}

		currentScore := 0.0
		skillScore := 0.0

		resScore := 0.0
		for _, skill := range filters.Skills {
			if filters.RemainingPoints[skill.Name] <= 0 {
				skillScore = 1.0
			} else {
				skillScore = calculateSkilleWeight(piece, skill.Name, filters)
			}
			currentScore += skillScore * skill.Weight
		}
		for _, res := range filters.Resistences {
			resScore = calculateResistenceWeight(&piece, res.Name)
			currentScore += resScore * res.Weight
		}

		if currentScore > bestsScore {
			p = piece
			bestsScore = currentScore
		}
	}

	// update the remaining skill points with whatever we chose
	for _, skill := range p.Skills {
		if _, ok := filters.RemainingPoints[skill.Name]; ok {
			filters.RemainingPoints[skill.Name] -= skill.Level
		}
	}
	return &p, nil
}

func calculateSkilleWeight(a ArmourPiece, skill string, filters *OptimizationFilters) float64 {
	w := 0.0
	skillWeight := float64(filters.RemainingPoints[skill]) / float64(filters.MaxPoints[skill])
	for _, s := range a.Skills {
		if s.Name == skill {
			currentWeight := float64(s.Level) * skillWeight
			w += currentWeight
			return w
		}
	}
	return w
}

func calculateResistenceWeight(a *ArmourPiece, res string) float64 {
	w := 0.0

	switch res {
	case "fire_res":
		w += float64(a.FireRes)
	case "water_res":
		w += float64(a.WaterRes)
	case "thunder_res":
		w += float64(a.ThunderRes)
	case "ice_res":
		w += float64(a.IceRes)
	case "dragon_res":
		w += float64(a.DragonRes)
	default:
		// do nothing
	}
	return w
}
