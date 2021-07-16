package domain

import "time"

type PostReport struct {
	Id string
	PostId string
	Timestamp time.Time
	ReportBy Profile
	ReportedPostBy Profile
	ReportType MediaReportType
	ReportStatus ReportStatus
}
