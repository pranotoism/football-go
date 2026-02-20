package dto

type CreatePlayerRequest struct {
	Name         string `json:"name" binding:"required"`
	HeightCM     int    `json:"height_cm"`
	WeightKG     int    `json:"weight_kg"`
	Position     string `json:"position" binding:"required,oneof=penyerang gelandang bertahan penjaga_gawang"`
	JerseyNumber int    `json:"jersey_number" binding:"required,min=1,max=99"`
}

type UpdatePlayerRequest struct {
	Name         string `json:"name"`
	HeightCM     int    `json:"height_cm"`
	WeightKG     int    `json:"weight_kg"`
	Position     string `json:"position" binding:"omitempty,oneof=penyerang gelandang bertahan penjaga_gawang"`
	JerseyNumber int    `json:"jersey_number" binding:"omitempty,min=1,max=99"`
}
