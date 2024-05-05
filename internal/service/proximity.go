package service

import (
	"context"
	"github.com/uber/h3-go/v4"
	"go.mongodb.org/mongo-driver/bson"
	"math"
	"proximity_service_go/cmd/constant"
	"proximity_service_go/internal/model"
	"proximity_service_go/pkg/db"
)

func FindNearbyBusinesses(
	lat, lng, radius float64,
	resolution int,
	businessType *constant.BusinessType,
) (*[]*model.BusinessResponse, error) {
	edgeLength := h3.HexagonEdgeLengthAvgKm(resolution) * 1000
	numRings := int(math.Ceil((radius * 1000) / edgeLength))

	// Convert the center coordinates to an H3 index
	centerIndex := h3.LatLngToCell(h3.LatLng{
		Lat: lat,
		Lng: lng,
	}, resolution)

	// Get the hexagon indexes within the radius
	nearbyIndexes := h3.GridDisk(centerIndex, numRings)

	// Get the MongoDB client instance
	client := db.GetClient()
	collection := client.Database(constant.ProximityServiceDatabase).Collection(constant.BusinessCollection)

	// Prepare the MongoDB query to find businesses with matching H3 indexes
	resolutionType := constant.H3ResolutionStrings[resolution]

	var query bson.M
	query = bson.M{resolutionType: bson.M{"$in": nearbyIndexes}}

	if businessType != nil && *businessType != constant.All {
		query = bson.M{"type": businessType, resolutionType: bson.M{"$in": nearbyIndexes}}
	}

	// Execute the query
	cursor, err := collection.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Decode the results into a slice of Business structs
	var businesses []model.Business
	err = cursor.All(context.Background(), &businesses)
	if err != nil {
		return nil, err
	}

	businessResponses := make([]*model.BusinessResponse, len(businesses))
	for i, business := range businesses {
		businessResponses[i] = business.ToBusinessResponse()
	}

	return &businessResponses, nil
}
