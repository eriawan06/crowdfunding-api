package dto

type CreateCampaignRequest struct {
	UserID       uint    `json:"user_id" binding:"required"`
	CategoryID   uint    `json:"category_id" binding:"required"`
	Title        string  `json:"title" binding:"required"`
	Deadline     string  `json:"deadline" binding:"required"`
	TargetAmount uint    `json:"target_amount" binding:"required"`
	Image        *string `json:"image"`
	Description  *string `json:"description"`
}

type UpdateCampaignRequest struct {
	CategoryID   uint    `json:"category_id" binding:"required"`
	Title        string  `json:"title" binding:"required"`
	Deadline     string  `json:"deadline" binding:"required"`
	TargetAmount uint    `json:"target_amount" binding:"required"`
	Image        *string `json:"image"`
	Description  *string `json:"description"`
}

type CampaignResponse struct {
	ID            uint    `json:"id"`
	Title         string  `json:"title"`
	Image         *string `json:"image"`
	Creator       string  `json:"creator"`
	Category      string  `json:"category"`
	DayRemaining  int     `json:"day_remaining"`
	TargetAmount  uint    `json:"target_amount"`
	CurrentAmount uint    `json:"current_amount"`
	IsCompleted   bool    `json:"is_completed"`
}

type CampaignDetailResponse struct {
	ID            uint    `json:"id"`
	Title         string  `json:"title"`
	Image         *string `json:"image"`
	Creator       string  `json:"creator"`
	Category      string  `json:"category"`
	DayRemaining  int     `json:"day_remaining"`
	TargetAmount  uint    `json:"target_amount"`
	CurrentAmount uint    `json:"current_amount"`
	Description   *string `json:"description"`
	IsCompleted   bool    `json:"is_completed"`
	TotalDonation uint    `json:"total_donation"`
}
