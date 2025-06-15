package handler

import (
	"aia_backend/api"
	"aia_backend/models"
	"encoding/json"
	"fmt"
	"github.com/alioygur/gores"
	"github.com/google/uuid"
	"net/http"
)

func PingHandler(rw http.ResponseWriter, _ *http.Request) {
	_ = gores.JSON(rw, http.StatusOK, fmt.Sprintf("Welcome to aia program"))
}

func UploadFrequentQuestionHandler(rw http.ResponseWriter, rq *http.Request) {
	ctx := rq.Context()
	var params api.UploadQuestionRequest
	if err := json.NewDecoder(rq.Body).Decode(&params); err != nil {
		_ = gores.JSON(rw, 400, fmt.Sprintf("decode error %v", err))
		return
	}
	uid, err := uuid.NewUUID()
	if err != nil {
		_ = gores.JSON(rw, http.StatusInternalServerError, fmt.Sprintf("create uuid error %v", err))
		return
	}

	for _, question := range params.Questions {
		q := models.Question{
			QuestionID: uid.String(),
			Question:   question.Question,
			Answer:     question.Answer,
			Category:   question.Category,
		}
		if err = models.SaveQuestion(ctx, q); err != nil {
			_ = gores.JSON(rw, http.StatusInternalServerError, fmt.Sprintf("save question error %v", err))
			return
		}
	}

	_ = gores.JSON(rw, http.StatusOK, fmt.Sprintf("ok"))
}

func ListFrequentQuestionsHandler(rw http.ResponseWriter, rq *http.Request) {
	ctx := rq.Context()
	questions, err := models.ListQuestions(ctx)
	if err != nil {
		_ = gores.JSON(rw, http.StatusInternalServerError, err)
		return
	}

	_ = gores.JSON(rw, http.StatusOK, questions)
}

func DeleteFrequentQuestionHandler(rw http.ResponseWriter, _ *http.Request) {
	_ = gores.JSON(rw, http.StatusInternalServerError, fmt.Sprintf("not implemented"))
}

func PutQuestionHandler(rw http.ResponseWriter, _ *http.Request) {
	_ = gores.JSON(rw, http.StatusInternalServerError, fmt.Sprintf("not implemented"))
}
