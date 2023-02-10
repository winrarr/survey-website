package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Survey struct {
	Description string             `binding:"required"`
	Created     primitive.DateTime `binding:"required"`
	Questions   []Question         `binding:"required"`
}

type Question struct {
	Text         string   `binding:"required"`
	QuestionType string   `binding:"required"`
	Choices      []string `binding:"required"`
}
