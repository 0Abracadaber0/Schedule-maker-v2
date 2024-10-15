package models

type LessonType string

const (
	Lecture    LessonType = "Lecture"
	Practice   LessonType = "Practice"
	Laboratory LessonType = "Laboratory"
)

type Lesson struct {
	Subject   string
	Teacher   string
	Group     int
	Type      LessonType
	TimeSlot  int
	Classroom string
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
	Load            int      `json:"load"`
}

type Classroom struct {
	Building string     `json:"building"`
	Number   string     `json:"number"`
	Capacity int        `json:"capacity"`
	Type     LessonType `json:"type"`
}

type TimeSlot struct {
	Slot int `json:"slot"`
}
