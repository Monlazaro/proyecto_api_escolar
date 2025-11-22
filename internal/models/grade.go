package models

type Grade struct{
	GradeID 	int  		`json:"grade_id" gorm:"primaryKey;autoIncrement"`
	StudentID	int			`json:"student_id" gorm:"not null"`
	SubjectID 	int			`json:"subject_id" gorm:"not null"`
	Grade		float64 	`json:"grade" gorm:"not null"`

	//relaciones
	Student Student `json:"student" gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE"`
	Subject Subject `json:"subject" gorm:"foreignKey:SubjectID;constraint:OnDelete:CASCADE"`
}