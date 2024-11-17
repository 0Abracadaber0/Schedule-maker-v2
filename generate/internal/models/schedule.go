package models

type LessonType string

const (
	Lecture    LessonType = "Lecture"
	Practice   LessonType = "Practice"
	Laboratory LessonType = "Laboratory"
)

type Subject struct {
	Name     string   `json:"name"`
	Teachers []string `json:"teachers"`
}

type Group struct {
	Name       string `json:"name"`
	Curriculum string `json:"curriculum"`
}

type Curriculum struct {
	Name     string `json:"name"`
	Subjects []Plan `json:"subjects"`
}

type Plan struct {
	Name         string   `json:"name"`
	Lectures     []string `json:"lectures"`
	Practices    []string `json:"practices"`
	Laboratories []string `json:"laboratories"`
	Flow         string   `json:"flow"`
}

type Classroom struct {
	Name     string     `json:"name"`
	Type     LessonType `json:"type"`
	Subjects []string   `json:"subjects"`
}
