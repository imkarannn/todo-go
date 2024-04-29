package model

import
(
	"time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// Define a struct to represent your data model
type UserMaster struct {
    ID         primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
    FirstName  string                `json:"firstName,omitempty" bson:"firstName,omitempty"`
    LastName   string                `json:"lastName,omitempty" bson:"lastName,omitempty"`
    UserEmail  string                `json:"userEmail,omitempty" bson:"userEmail,omitempty"`
    Password   string                `json:"password,omitempty" bson:"password,omitempty"`
    IsActive   bool                  `json:"isActive,omitempty" bson:"isActive,omitempty"`
    CreatedAt  time.Time             `json:"created_at,omitempty" bson:"created_at,omitempty"`
    UserRole   string                `json:"userRole,omitempty" bson:"_userRole,omitempty"`
}
// Define a struct to represent your data model
type UserRoleMaster struct {
    UserRoleID primitive.ObjectID    `json:"userRoleId,omitempty" bson:"_userRoleId,omitempty"`
    UserRole   string                `json:"userRole,omitempty" bson:"userRole,omitempty"`
    IsActive   bool                  `json:"isActive,omitempty" bson:"isActive,omitempty"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}