package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Chi tiết quyền
type PermissionDetail struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`    // ID của quyền
	Name string             `json:"name,omitempty" bson:"name,omitempty"` // Tên của quyền
}

// Vai trò
type Role struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`                    // ID của vai trò
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`                 // Tên của vai trò
	Describe     string             `json:"describe,omitempty" bson:"describe,omitempty"`         // Mô tả vai trò
	Organization primitive.ObjectID `json:"organization,omitempty" bson:"organization,omitempty"` // ID tổ chức
	CreatedAt    int64              `json:"createdAt,omitempty" bson:"createdAt,omitempty"`       // Thời gian tạo
	UpdatedAt    int64              `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`       // Thời gian cập nhật
}

// ==========================================================================================

// Dữ liệu đầu vào tạo vai trò
type RoleCreateInput struct {
	Name         string `json:"name,omitempty" bson:"name,omitempty" validate:"required"`                 // Tên của vai trò
	Describe     string `json:"describe,omitempty" bson:"describe,omitempty" validate:"required"`         // Mô tả vai trò
	Ogranization string `json:"organization,omitempty" bson:"organization,omitempty" validate:"required"` // ID tổ chức
}

// Dữ liệu đầu vào cập nhật vai trò
type RoleUpdateInput struct {
	Name         string `json:"name,omitempty" bson:"name,omitempty"`                 // Tên của vai trò
	Describe     string `json:"describe,omitempty" bson:"describe,omitempty"`         // Mô tả vai trò
	Ogranization string `json:"organization,omitempty" bson:"organization,omitempty"` // ID tổ chức
}
