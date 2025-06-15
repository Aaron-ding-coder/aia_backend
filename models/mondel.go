package models

import "time"

type People struct {
	ChineseName string
	EnglishName string
	title       string
	Awards      []string
}

type QuestionCategory string

var (
	SavingWithoutClaims QuestionCategory = "储蓄险(不涉及理赔)"
	Critical            QuestionCategory = "重疾/医疗等(涉及健康理赔)"
	Others              QuestionCategory = "其他"
)

type Question struct {
	QuestionID string           `json:"question_id" bson:"question_id"`
	Question   string           `json:"question" bson:"question"`
	Answer     string           `json:"answer" bson:"answer"`
	Category   QuestionCategory `json:"category" bson:"category"`
	CreatedAt  time.Time        `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at" bson:"updated_at"`
}

type ProductType string

var (
	Saving          ProductType = "储蓄"
	CriticalIllness ProductType = "重疾"
	Healthcare      ProductType = "医疗"
)

type ProductFile struct {
	FIleID     string      `json:"file_id" bson:"file_id"`
	Category   ProductType `json:"category" bson:"category"`
	Name       string      `json:"name" bson:"name"`
	ObjectKeys string      `json:"object_keys" bson:"object_keys"`
}

type Courses struct {
	CourseID    string
	Name        string
	Author      string
	Description string
}

type User struct {
	Name        string
	PhoneNumber string
	Courses     []string
}
