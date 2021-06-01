package repository

import (
	"context"
	"github.com/gocql/gocql"
	"post-service/domain"
)

const (
	CreateReportTable = "CREATE TABLE if not exists post_keyspace.Reports (id text, post_id text, timestamp timestamp, report_by text, type text, status text, " +
		"PRIMARY KEY (id, status));"
	InsertReportStatement = "INSERT INTO post_keyspace.Reports (id, post_id, timestamp, report_by, type, status) VALUES (?, ?, ?, ?, ?, ?) IF NOT EXISTS;"
)

type ReportRepo interface {
	ReportPost(report *domain.PostReport, ctx context.Context) error
	ReviewReport(report *domain.PostReport, ctx context.Context) error
	GetAllPendingReports(ctx context.Context) ([]domain.PostReport, error)
	GetAllApprovedReports(ctx context.Context) ([]domain.PostReport, error)
	GetAllRejectedReports(ctx context.Context) ([]domain.PostReport, error)
}

type reportRepository struct {
	cassandraSession *gocql.Session
}

func (r reportRepository) ReviewReport(report *domain.PostReport, ctx context.Context) error {
	panic("implement me")
}

func (r reportRepository) GetAllPendingReports(ctx context.Context) ([]domain.PostReport, error) {
	panic("implement me")
}

func (r reportRepository) GetAllApprovedReports(ctx context.Context) ([]domain.PostReport, error) {
	panic("implement me")
}

func (r reportRepository) GetAllRejectedReports(ctx context.Context) ([]domain.PostReport, error) {
	panic("implement me")
}

func (r reportRepository) ReportPost(report *domain.PostReport, ctx context.Context) error {
	panic("implement me")
}


func NewReportRepository(cassandraSession *gocql.Session) ReportRepo {
	var r = &reportRepository{
		cassandraSession: cassandraSession,
	}
	err := r.cassandraSession.Query(CreateReportTable).Exec()
	if err != nil {
		return nil
	}
	return r
}