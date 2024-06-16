package dto

type Paginate struct {
	DataPerPage int `json:"data_per_page"`
	TotalData   int `json:"total_data"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}
