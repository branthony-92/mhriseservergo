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

type ArmourPiece struct {
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
	Skills     []interface{} `json:"skills" bson:"skills"`
}

func NewArmourPiece() *ArmourPiece {
	armour := ArmourPiece{}
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
