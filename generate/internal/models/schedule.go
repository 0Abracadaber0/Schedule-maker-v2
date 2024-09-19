package models

type LessonType string

const (
	Lecture    LessonType = "Lecture"
	Practice   LessonType = "Practice"
	Laboratory LessonType = "Laboratory"
)

type Lesson struct {
	Subject string
	Teacher string
	Group   int
	Type    LessonType
}

type Subject struct {
	Name            string `json:"name"`
	LectureCount    int    `json:"lecture_count"`
	PracticeCount   int    `json:"practice_count"`
	LaboratoryCount int    `json:"laboratory_count"`
}

type Curriculum struct {
	Group    string    `json:"group"`
	Shift    int       `json:"shift"`
	Subjects []Subject `json:"subjects"`
}

type Teacher struct {
	Name            string   `json:"name"`
	Subjects        []string `json:"subjects"`
	CanGiveLectures bool     `json:"can_give_lectures"`
}
