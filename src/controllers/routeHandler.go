package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// RouteHandler stores mongo client
type RouteHandler struct {
	Meeting *mongo.Collection
}

// NewRouteHandler returns a RouteHandler object
func NewRouteHandler(Meeting *mongo.Collection) *RouteHandler {
	return &RouteHandler{
		Meeting: Meeting,
	}
}
