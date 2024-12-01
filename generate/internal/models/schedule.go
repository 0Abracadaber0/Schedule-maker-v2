package models

type Classroom struct {
	Name     string
	Type     string
	Subjects []string
}

type Course struct {
	Lectures  int
	Practices int
	Labs      int
	Stream    string
}

type ScheduleGenerator struct {
	Subjects       map[string][]string
	Groups         map[string]string
	Plans          map[string]map[string]*Course
	Classrooms     []Classroom
	LessonsPerWeek int
}
