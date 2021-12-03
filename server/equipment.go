package server

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
	PieceName  string         `json:"piece_name" bson:"piece_name"`
	SetName    string         `json:"set_name" bson:"set_name"`
	PieceType  string         `json:"piece_type" bson:"piece_type"`
	Rarity     int            `json:"rarity" bson:"rarity"`
	Defense    int            `json:"defence" bson:"defence"`
	FireRes    int            `json:"fire_res" bson:"fire_res"`
	WaterRes   int            `json:"water_res" bson:"water_res"`
	ThunderRes int            `json:"thunder_res" bson:"thunder_res"`
	IceRes     int            `json:"ice_res" bson:"ice_res"`
	DragonRes  int            `json:"dragon_res" bson:"dragon_res"`
	Skills     []*ArmourSkill `json:"skills" bson:"skills"`
}

func NewArmourPiece() *ArmourPiece {
	armour := ArmourPiece{}
	armour.Skills = make([]*ArmourSkill, 0)
	return &armour
}

type ArmourSet struct {
	Pieces  []ArmourPiece `json:"pieces"`
	SetName string        `json:"set_name"`
}

func NewArmourSet() *ArmourSet {
	armour := ArmourSet{}
	return &armour
}

func Summarize(set *ArmourSet) *ArmourPiece {
	summary := NewArmourPiece()
	for _, p := range set.Pieces {
		summary.Defense += p.Defense
		summary.ThunderRes += p.ThunderRes
		summary.FireRes += p.FireRes
		summary.WaterRes += p.WaterRes
		summary.IceRes += p.IceRes
		summary.DragonRes += p.DragonRes

		for _, inSkill := range p.Skills {
			found := false
			for _, outSkill := range summary.Skills {
				if outSkill.Name == inSkill.Name {
					outSkill.Level += inSkill.Level
					found = true
					break
				}
			}
			if !found {
				summary.Skills = append(summary.Skills, inSkill)
			}
		}
	}
	summary.PieceName = "Summary"
	return summary
}
