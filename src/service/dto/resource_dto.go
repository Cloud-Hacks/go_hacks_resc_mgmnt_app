package dto

//RequestAddResourceDTO - DTO for add Resource Request
type RequestAddResourceDTO struct {
	ID         int      `json:"id" validate:"min=1"`
	Title      string   `json:"title" validate:"required"`
	Category   string   `json:"category" validate:"required"`
	Status     string   `json:"status" validate:"required,oneof= Draft Published"`
	Type       string   `json:"type" validate:"required,oneof= Text PDF"`
	ImageLinks []string `json:"image_links"`
	FileLink   string   `json:"file_link"`
	Content    string   `json:"content"`
	CreatedBy  int      `json:"created_by" validate:"required"`
}

//RequestGetResourceDTO - DTO for Get Resource Request
type RequestGetResourceDTO struct {
	ID int `json:"id" validate:"min=0"`
}

//ResponseGetListOfResourcesDTO - DTO for Get Resource Response
type ResponseGetResourceDTO struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Category   string `json:"category"`
	Status     string `json:"status"`
	Types      string `json:"types"`
	ImageLinks string `json:"image_links"`
	FileLink   string `json:"file_link"`
	Content    string `json:"content"`
	UpdatedBy  int    `json:"updated_by"`
	UpdatedAt  int64  `json:"updated_at"`
}
