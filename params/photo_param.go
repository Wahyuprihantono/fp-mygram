package params

type CreatePhotos struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
	PhotoURL string `json:"photo_url" validate:"required"`
	UserID   uint   `json:"user_id,omitempty"`
}
