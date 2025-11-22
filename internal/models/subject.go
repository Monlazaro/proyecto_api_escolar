package models

type Subject struct{
	SubjectID	int			`json:"subject_id" gorm:"primaryKey;autoIncrement"`
	Name		string		`json:"name" gorm:"not null"`
}