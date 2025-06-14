package models

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
	Question string
	Answers  string
	Category QuestionCategory
}

type ProductType string

var (
	Saving          ProductType = "储蓄"
	CriticalIllness ProductType = "重疾"
	Healthcare      ProductType = "医疗"
)

type Products struct {
	Type  ProductType
	Name  string
	files []string
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
