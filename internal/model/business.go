package model

import (
	"github.com/uber/h3-go/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"proximity_service_go/cmd/constant"
)

type Business struct {
	ID                  primitive.ObjectID    `json:"_id"`
	Name                string                `json:"name"`
	City                string                `json:"city"`
	Country             string                `json:"country"`
	Type                constant.BusinessType `json:"type"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	H3IndexResolution9  uint64                `json:"h3IndexResolution9"`
	H3IndexResolution12 uint64                `json:"h3IndexResolution12"`
}

type BusinessRequest struct {
	Name      string  `json:"name"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Type      string  `json:"type"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (businessRequest BusinessRequest) ToBusiness() (*Business, error) {
	// Convert the business coordinates to an H3 index at the default resolution
	businessRequestLatLng := h3.LatLng{
		Lat: businessRequest.Latitude,
		Lng: businessRequest.Longitude,
	}
	h3IndexResolution9 := h3.LatLngToCell(businessRequestLatLng, constant.Resolution9)
	h3IndexResolution12 := h3.LatLngToCell(businessRequestLatLng, constant.Resolution12)

	businessType, err := constant.ParseBusinessType(businessRequest.Type)
	if err != nil {
		return nil, err
	}
	return &Business{
		ID:                  primitive.NewObjectID(),
		Name:                businessRequest.Name,
		City:                businessRequest.City,
		Country:             businessRequest.Country,
		Type:                businessType,
		Latitude:            businessRequest.Latitude,
		Longitude:           businessRequest.Longitude,
		H3IndexResolution9:  uint64(h3IndexResolution9),
		H3IndexResolution12: uint64(h3IndexResolution12),
	}, nil
}

func (business Business) ToBusinessResponse() *BusinessResponse {
	return &BusinessResponse{
		Name:                business.Name,
		City:                business.City,
		Country:             business.Country,
		Type:                business.Type,
		Latitude:            business.Latitude,
		Longitude:           business.Longitude,
		H3IndexResolution9:  business.H3IndexResolution9,
		H3IndexResolution12: business.H3IndexResolution12,
	}
}

type BusinessResponse struct {
	Name                string                `json:"name"`
	City                string                `json:"city"`
	Country             string                `json:"country"`
	Type                constant.BusinessType `json:"type"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	H3IndexResolution9  uint64                `json:"h3IndexResolution9"`
	H3IndexResolution12 uint64                `json:"h3IndexResolution12"`
}
