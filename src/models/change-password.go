package models

type ChangePassword struct {
	NewPassword string `json:"newPassword" validate:"required,min=6"`
	OldPassword string `json:"oldPassword" validate:"required,min=6"`
}
