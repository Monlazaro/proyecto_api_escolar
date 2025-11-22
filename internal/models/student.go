package models

type Student struct {
	StudentID int		`json:"student_id" gorm:"primaryKey;autoIncrement"`
	Name	  string 	`json:"name" gorm:"not null"`
	Group	  string 	`json:"group"`
	Email	  string 	`json:"email" gorm:"not null;uniqueIndex"`
}