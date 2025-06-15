package models

import (
	"aia_backend/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveProducts(ctx context.Context, p ProductFile) error {
	if p.FIleID == "" {
		p.FIleID = utils.RandomStringWithLength(5)
	}
	filter := bson.M{
		"file_id": p.FIleID,
	}
	GetCollection("products").FindOneAndReplace(ctx, filter, p, options.FindOneAndReplace().SetUpsert(true))
	return nil
}

func GetProductByID(ctx context.Context, fileID string) (*ProductFile, error) {
	file := ProductFile{}
	if err := GetCollection("products").FindOne(ctx, bson.M{"file_id": fileID}).Decode(&file); err != nil {
		return nil, err
	}

	return &file, nil
}

func ListProducts(ctx context.Context) ([]ProductFile, error) {
	cur, err := GetCollection("products").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)
	p := []ProductFile{}
	if err = cur.All(ctx, &p); err != nil {
		return nil, err
	}

	return p, nil
}
