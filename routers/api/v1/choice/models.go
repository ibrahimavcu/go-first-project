package choice

import "time"

type choice struct {
	ID        uint
	Limit     int
	Interval  int
	OrderDate *time.Time
	IsSubject bool
	Days      string
	LessonId  uint
	UserId    uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type choices []choices
