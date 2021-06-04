package dto

type CreateReportDTO struct {
	PostId string `json:"post_id" validate:"required"`
	ReportedBy string `json:"reported_by" validate:"required"`
	ReportType string `json:"report_type" validate:"required"`
	ReportedPostBy string `json:"reported_post_by" validate:"required"`
}
