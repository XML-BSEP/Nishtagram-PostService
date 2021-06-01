package repository

import (
	"github.com/gocql/gocql"
	"post-service/domain"
)

const (
	CreateReportTable = "CREATE TABLE if not exists post_service.Reports (id, post_id, timestamp, num_of_likes, report_by, type, status, " +
		"PRIMARY KEY (id, status));"
)

type ReportRepo interface {
	ReportPost(report *domain.PostReport) error
	ReviewReport(report *domain.PostReport) error
	GetAllPendingReports() ([]domain.PostReport, error)
	GetAllApprovedReports() ([]domain.PostReport, error)
	GetAllRejectedReports() ([]domain.PostReport, error)
}

type reportRepository struct {
	cassandraSession *gocql.Session
}

func (r reportRepository) ReviewReport(report *domain.PostReport) error {
	panic("implement me")
}

func (r reportRepository) GetAllPendingReports() ([]domain.PostReport, error) {
	panic("implement me")
}

func (r reportRepository) GetAllApprovedReports() ([]domain.PostReport, error) {
	panic("implement me")
}

func (r reportRepository) GetAllRejectedReports() ([]domain.PostReport, error) {
	panic("implement me")
}

func (r reportRepository) ReportPost(report *domain.PostReport) error {
	panic("implement me")
}


func NewReportRepository(cassandraSession *gocql.Session) ReportRepo {
	var r = &reportRepository{
		cassandraSession: cassandraSession,
	}
	r.cassandraSession.Query(CreateReportTable).Exec()
	return r
}