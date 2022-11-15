package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Document struct {
    Id       primitive.ObjectID `json:"id,omitempty"`
	UserId 	  string 			`json:"userId,omitempty" validate:"required"`
	Timestamp string            `json:"timestamp,omitempty" validate:"required"` 
    SignedBy  string			`json:"signedBy,omitempty" validate:"required"`
	TransactionHash string		`json:"transactionHash,omitempty" validate:"required"`
}