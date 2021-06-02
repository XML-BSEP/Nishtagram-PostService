package dto

type ReviewReportDTO struct {
	ReportId string `json:"report_id" validate:"required"`
	Status string `json:"status" validate:"required"`
	DeletePost bool `json:"delete_post" validate:"required"`
}
