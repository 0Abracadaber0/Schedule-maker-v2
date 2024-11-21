package models

type LessonType string

const (
	Lecture    LessonType = "Lecture"
	Practice   LessonType = "Practice"
	Laboratory LessonType = "Laboratory"
)

type Teacher struct {
	Name     string   `json:"name"`
	Load     int      `json:"load"`
	Subjects []string `json:"subjects"`
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
	Name         string `json:"name"`
	Lectures     int    `json:"lectures"`
	Practices    int    `json:"practices"`
	Laboratories int    `json:"laboratories"`
	Flow         string `json:"flow"`
}

type Classroom struct {
	Name     string     `json:"name"`
	Type     LessonType `json:"type"`
	Subjects []string   `json:"subjects"`
}

type Lesson struct {
	Subject string
	Group   string
	Teacher string
	IsLab   bool
}
