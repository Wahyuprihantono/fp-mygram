package params

type CreateComments struct {
	Message string `json:"message" validate:"required"`
	PhotoID uint   `json:"photo_id,omitempty"`
	UserID  uint   `json:"user_id,omitempty"`
}
