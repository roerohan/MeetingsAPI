package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type RouteHandler struct {
	Meeting *mongo.Collection
}
