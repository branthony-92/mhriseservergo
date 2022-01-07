package server

import "fmt"

type PieceTypeID = int

const (
	PieceTypeInvalid PieceTypeID = -1 + iota
	PieceTypeHelm
	PieceTypeMail
	PieceTypeVambraces
	PieceTypeCoil
	PieceTypeGreaves
)

type ArmourSkill struct {
	Name  string `json:"name" bson:"name"`
	Level int    `json:"level" bson:"level"`
}

type ArmourPiece struct {
	ID         int           `json:"_id" bson:"_id"`
	PieceName  string        `json:"piece_name" bson:"piece_name"`
	SetName    string        `json:"set_name" bson:"set_name"`
	PieceType  string        `json:"piece_type" bson:"piece_type"`
	Rarity     int           `json:"rarity" bson:"rarity"`
	Defense    int           `json:"defence" bson:"defence"`
	FireRes    int           `json:"fire_res" bson:"fire_res"`
	WaterRes   int           `json:"water_res" bson:"water_res"`
	ThunderRes int           `json:"thunder_res" bson:"thunder_res"`
	IceRes     int           `json:"ice_res" bson:"ice_res"`
	DragonRes  int           `json:"dragon_res" bson:"dragon_res"`
	Skills     []ArmourSkill `json:"skills" bson:"skills"`
}

func NewArmourPiece() *ArmourPiece {
	armour := ArmourPiece{}
	armour.Skills = make([]ArmourSkill, 0)
	return &armour
}

type ArmourSet struct {
	Helm      *ArmourPiece `json:"helm"`
	Mail      *ArmourPiece `json:"mail"`
	Coil      *ArmourPiece `json:"coil"`
	Vambraces *ArmourPiece `json:"vambraces"`
	Greaves   *ArmourPiece `json:"greaves"`
	SetName   string       `json:"set_name"`
	Summary   string       `json:"summary"`
}

func NewArmourSet() *ArmourSet {
	armour := ArmourSet{
		Helm:      nil,
		Mail:      nil,
		Coil:      nil,
		Vambraces: nil,
		Greaves:   nil,
		SetName:   "",
		Summary:   "",
	}
	return &armour
}

func Summarize(set ArmourSet) string {
	str := "Summary\n--------------\n\n"
	skills := make(map[string]int)
	Defense := 0
	ThunderRes := 0
	FireRes := 0
	WaterRes := 0
	IceRes := 0
	DragonRes := 0

	str += "Parts:\n"
	/*
		for _, p := range set.Pieces {
			str += fmt.Sprintf("\t%v (%v)\n", p.PieceName, p.PieceType)
			Defense += p.Defense
			ThunderRes += p.ThunderRes
			FireRes += p.FireRes
			WaterRes += p.WaterRes
			IceRes += p.IceRes
			DragonRes += p.DragonRes

			for _, inSkill := range p.Skills {
				_, ok := skills[inSkill.Name]
				if ok {
					skills[inSkill.Name] += inSkill.Level
					} else {
						skills[inSkill.Name] = inSkill.Level
					}
				}
			}
	*/

	str += "\nStats:\n"
	str += fmt.Sprintf("\tDefence: %v\n", Defense)
	str += fmt.Sprintf("\tFire Res: %v\n", FireRes)
	str += fmt.Sprintf("\tWater Res: %v\n", WaterRes)
	str += fmt.Sprintf("\tThunder Res: %v\n", ThunderRes)
	str += fmt.Sprintf("\tIce Res: %v\n", IceRes)
	str += fmt.Sprintf("\tDragon Res: %v\n", DragonRes)
	str += "\nSkills:\n"

	for skill, lvl := range skills {
		skills, _ := QueryAllSkills()
		s := FindSkill(skills, skill)

		str += fmt.Sprintf("\t%v: %v/%v\n", skill, lvl, s.SkillLevelMax)
	}

	return str
}
