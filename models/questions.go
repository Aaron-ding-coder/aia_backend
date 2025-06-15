package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func SaveQuestion(ctx context.Context, question Question) error {
	question.CreatedAt = time.Now()
	question.UpdatedAt = time.Now()
	_, err := GetCollection("questions").InsertOne(ctx, question)
	return err
}

func ListQuestions(ctx context.Context) ([]Question, error) {
	cur, err := GetCollection("questions").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	q := []Question{}
	if err = cur.All(ctx, &q); err != nil {
		return nil, err
	}

	return q, nil
}
