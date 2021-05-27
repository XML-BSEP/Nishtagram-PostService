package domain

import "time"

type PostReport struct {
	Id uint
	PostId uint
	Timestamp time.Time
	ReportBy Profile
	ReportType MediaReportType
	ReportStatus ReportStatus
}
