package server

import "go.mongodb.org/mongo-driver/bson"

type PieceType = int

const (
	PieceTypeInvalid PieceType = -1
	PieceTypeHelm    PieceType = iota
	PieceTypeMail
	PieceTypeVambraces
	PieceTypeCoil
	PieceTypeGreaves
)

type ArmourPiece struct {
	Skills  []Skill
	Name    string
	SetName string
}

func NewArmourPiece() *ArmourPiece {
	armour := ArmourPiece{}
	return &armour
}

func (a *ArmourPiece) Encode() *bson.D {
	doc := bson.D{}
	return &doc
}

type ArmourSet struct {
	Pieces  []ArmourPiece
	SetName string
}

func NewArmourSet() *ArmourSet {
	armour := ArmourSet{}
	return &armour
}

func (a *ArmourSet) Encode() *bson.D {
	doc := bson.D{}
	return &doc
}
