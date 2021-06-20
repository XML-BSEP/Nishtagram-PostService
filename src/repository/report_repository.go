package repository

import (
	"context"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"post-service/dto"
	"strings"
	"time"
)

const (
	CreateReportTable = "CREATE TABLE if not exists post_keyspace.Reports (id text, post_id text, timestamp timestamp, report_by text, reported_post_by text, type text, status text, " +
		"PRIMARY KEY (status, id));"
	InsertReportStatement = "INSERT INTO post_keyspace.Reports (id, post_id, timestamp, report_by, reported_post_by, type, status) VALUES (?, ?, ?, ?, ?, ?, ?) IF NOT EXISTS;"
	GetAllRequestsByStatus = "SELECT  id, post_id, timestamp, report_by, reported_post_by, type, status FROM post_keyspace.Reports WHERE status = ?;"
	DeleteReport = "DELETE FROM post_keyspace.Reports where status = ? and id = ?;"
	GetPendingReportById = "SELECT id, post_id, reported_post_by, reported_by, type, timestamp FROM post_keyspace.Reports " +
		"WHERE status = 'CREATED' AND id = ?;"
	SelectAllTypes = "SELECT * FROM post_keyspace.ReportType LIMIT 300000000;"
)

type ReportRepo interface {
	ReportPost(report dto.CreateReportDTO, ctx context.Context) error
	ReviewReport(report dto.ReviewReportDTO, ctx context.Context) error
	GetAllPendingReports(ctx context.Context) ([]dto.ReportDTO, error)
	GetAllApprovedReports(ctx context.Context) ([]dto.ReportDTO, error)
	GetAllRejectedReports(ctx context.Context) ([]dto.ReportDTO, error)
	GetAllReportTypes(ctx context.Context) ([]string, error)
}

type reportRepository struct {
	cassandraSession *gocql.Session
}

func (r reportRepository) GetAllReportTypes(ctx context.Context) ([]string, error) {
	var retVal []string

	iter := r.cassandraSession.Query(SelectAllTypes).Iter().Scanner()

	var reportType string
	for iter.Next() {

		err := iter.Scan(&reportType)
		if err != nil {
			return nil, err
		}

		retVal = append(retVal, reportType)
	}

	return retVal, nil
}

func (r reportRepository) ReviewReport(report dto.ReviewReportDTO, ctx context.Context) error {
	var reportId, postId, reportedPostBy, reportedBy, reportType string
	var timestamp time.Time

	iter := r.cassandraSession.Query(GetPendingReportById, report.ReportId).Iter()

	if iter == nil {
		return fmt.Errorf("no such element")
	}

	for iter.Scan(&reportId, &postId, &reportedPostBy, &reportedBy, &reportType, &timestamp) {

		if report.DeletePost {
			r.cassandraSession.Query(DeletePost, reportedPostBy, postId).Exec()
		}

		updatedStatus := strings.ToUpper(report.Status)

		r.cassandraSession.Query(DeleteReport, reportId).Exec()
		var newUUID, err = uuid.NewUUID()

		if err != nil {
			return err
		}
		err = r.cassandraSession.Query(InsertReportStatement, newUUID, postId, time.Now(), reportedBy, reportedPostBy, reportType, updatedStatus).Exec()

		if err != nil {
			return err
		}

		return nil
	}

	return nil

}

func (r reportRepository) GetAllPendingReports(ctx context.Context) ([]dto.ReportDTO, error) {
	iter := r.cassandraSession.Query(GetAllRequestsByStatus, "CREATED").Iter().Scanner()

	if iter == nil {
		return nil, fmt.Errorf("no pending reports")
	}
	var reports []dto.ReportDTO
	var reportId, postId, reportedBy, reportedPostBy, reportType, status string
	var timestamp time.Time

	for iter.Next() {
		iter.Scan(&reportId, &postId, &timestamp, &reportedBy, &reportedPostBy, &reportType, &status)
		reports = append(reports, dto.NewReportDTO(reportId, postId, timestamp, reportedBy, reportedPostBy, reportType, status))
	}
	return reports, nil
}

func (r reportRepository) GetAllApprovedReports(ctx context.Context) ([]dto.ReportDTO, error) {
	iter := r.cassandraSession.Query(GetAllRequestsByStatus, "APPROVED").Iter().Scanner()

	if iter == nil {
		return nil, fmt.Errorf("no pending reports")
	}
	var reports []dto.ReportDTO
	var reportId, postId, reportedBy, reportedPostBy, reportType, status string
	var timestamp time.Time

	for iter.Next() {
		iter.Scan(&reportId, &postId, &timestamp, &reportedBy, &reportedPostBy, &reportType, &status)
		reports = append(reports, dto.NewReportDTO(reportId, postId, timestamp, reportedBy, reportedPostBy, reportType, status))
	}
	return reports, nil
}

func (r reportRepository) GetAllRejectedReports(ctx context.Context) ([]dto.ReportDTO, error) {
	iter := r.cassandraSession.Query(GetAllRequestsByStatus, "REJECTED").Iter().Scanner()

	if iter == nil {
		return nil, fmt.Errorf("no pending reports")
	}
	var reports []dto.ReportDTO
	var reportId, postId, reportedBy, reportedPostBy, reportType, status string
	var timestamp time.Time

	for iter.Next() {
		iter.Scan(&reportId, &postId, &timestamp, &reportedBy, &reportedPostBy, &reportType, &status)
		reports = append(reports, dto.NewReportDTO(reportId, postId, timestamp, reportedBy, reportedPostBy, reportType, status))
	}
	return reports, nil
}

func (r reportRepository) ReportPost(report dto.CreateReportDTO, ctx context.Context) error {
	id := uuid.NewString()

	err := r.cassandraSession.Query(InsertReportStatement, id, report.PostId, time.Now(), report.ReportedBy, report.ReportedPostBy, strings.ToUpper(report.ReportType), "CREATED").Exec()

	if err != nil {
		return err
	}

	return nil
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