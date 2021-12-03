package server

type Skill struct {
	SkillName        string   `json:"skill_name" bson:"skill_name"`
	SkillLevelMax    int      `json:"skill_level_max" bson:"skill_level_max"`
	SkillDescription string   `json:"skill_desc" bson:"skill_desc"`
	SkillLevels      []string `json:"skill_levels" bson:"skill_levels"`
}

func NewSkill() *Skill {
	s := Skill{}
	return &s
}

func FindSkill(skills []Skill, name string) *Skill {
	for _, s := range skills {
		if s.SkillName == name {
			return &s
		}
	}
	return nil
}
