package service

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"proximity_service_go/cmd/constant"
	"proximity_service_go/internal/model"
	"proximity_service_go/pkg/db"
	"time"
)

func AddBusiness(businessRequest *model.BusinessRequest) error {
	business, conversionError := businessRequest.ToBusiness()
	if conversionError != nil {
		return fmt.Errorf("failed to create business: %s", conversionError)
	}

	client := db.GetClient()
	collection := client.Database(constant.ProximityServiceDatabase).Collection(constant.BusinessCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, business)
	if err != nil {
		return fmt.Errorf("failed to create business: %s", err)
	}

	return nil
}

func GetBusinessByID(businessId string) (*model.BusinessResponse, error) {
	client := db.GetClient()
	collection := client.Database(constant.ProximityServiceDatabase).Collection(constant.BusinessCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(businessId)
	if err != nil {
		return nil, fmt.Errorf("invalid business ID: %s", err)
	}

	var business model.Business
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&business)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("business not found")
		}
		return nil, fmt.Errorf("failed to get business: %s", err)
	}
	return business.ToBusinessResponse(), nil
}

func GetAllBusinesses() (*[]*model.BusinessResponse, error) {
	client := db.GetClient()
	collection := client.Database(constant.ProximityServiceDatabase).Collection(constant.BusinessCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get businesses: %s", err)
	}
	defer cursor.Close(ctx)

	var businesses []model.Business
	if err = cursor.All(ctx, &businesses); err != nil {
		panic(err)
	}

	businessResponses := make([]*model.BusinessResponse, len(businesses))
	for i, business := range businesses {
		businessResponses[i] = business.ToBusinessResponse()
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate businesses: %s", err)
	}
	return &businessResponses, nil
}
