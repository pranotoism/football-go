package dto

type CreateTeamRequest struct {
	Name        string `json:"name" binding:"required"`
	LogoURL     string `json:"logo_url"`
	FoundedYear int    `json:"founded_year" binding:"required"`
	HQAddress   string `json:"hq_address"`
	HQCity      string `json:"hq_city"`
}

type UpdateTeamRequest struct {
	Name        string `json:"name"`
	LogoURL     string `json:"logo_url"`
	FoundedYear int    `json:"founded_year"`
	HQAddress   string `json:"hq_address"`
	HQCity      string `json:"hq_city"`
}
