package user

type UpdateUserRequest struct {
	Role string `json:"role"` // optional, normally only admin can update
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
