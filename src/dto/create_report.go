package dto

type CreateReportDTO struct {
	PostId string `json:"postId" validate:"required"`
	ReportedBy string `json:"reportedBy" validate:"required"`
	ReportType string `json:"reportType" validate:"required"`
	ReportedPostBy string `json:"reportedPostBy" validate:"required"`
}
