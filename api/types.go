package api

import "aia_backend/models"

type UploadQuestionRequest struct {
	Questions []models.Question `json:"questions"`
}

type UploadQuestionResponse struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Category string `json:"category"`
}

type ListQuestionsRequest struct {
}

type ListQuestionsResponse struct {
	Questions []models.Question `json:"questions"`
	Count     int               `json:"count"`
}
