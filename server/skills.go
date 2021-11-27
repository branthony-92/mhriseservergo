package server

import "go.mongodb.org/mongo-driver/bson"

type Skill struct {
	SkillName        string
	SkillLevel       int
	SkillLevelMax    int
	SkillDescription string
	SkillLevels      []string
}

func NewSkill() *Skill {
	s := Skill{}
	return &s
}

func (s *Skill) Encode() *bson.D {
	doc := bson.D{
		{"skill_name", s.SkillName},
		{"skill_level", s.SkillLevel},
		{"skill_level_max", s.SkillLevelMax},
		{"skill_desc", s.SkillDescription},
		{"skill_levels", s.SkillLevels},
	}
	return &doc
}
