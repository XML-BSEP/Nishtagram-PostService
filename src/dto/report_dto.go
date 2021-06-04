package dto

import (
	"post-service/domain"
	"time"
)

type ReportDTO struct {
	Id string `json:"id" validate:"required"`
	PostId string `json:"post_id" validate:"required"`
	Timestamp time.Time `json:"timestamp" validate:"required"`
	ReportBy domain.Profile `json:"reported_by" validate:"required"`
	ReportType domain.MediaReportType `json:"report_type" validate:"required"`
	ReportedPostBy domain.Profile `json:"reported_post_by" validate:"required"`
	ReportStatus domain.ReportStatus `json:"report_status" validate:"required"`
}

func NewReportDTO(id string, postId string, timestamp time.Time, reportedBy string,
	 reportedPostBy string, reportType string, reportStatus string) ReportDTO {
	return ReportDTO{
		Id: id,
		ReportBy: domain.Profile{Id: reportedBy},
		Timestamp: timestamp,
		PostId: postId,
		ReportedPostBy: domain.Profile{Id: reportedPostBy},
		ReportType: domain.MediaReportType{ReportType: reportType},
		ReportStatus: domain.ReportStatus{Status: reportStatus},
	}
}
