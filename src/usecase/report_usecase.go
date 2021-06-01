package usecase

import (
	"context"
	"post-service/domain"
	"post-service/dto"
	"post-service/repository"
)

type PostReportUseCase interface {
	ReportPost(report *dto.ReportDTO, ctx context.Context) error
	ReviewReport(report *dto.ReportDTO, ctx context.Context) error
	GetAllPendingReports(ctx context.Context) ([]domain.PostReport, error)
	GetAllApprovedReports(ctx context.Context) ([]domain.PostReport, error)
	GetAllRejectedReports(ctx context.Context) ([]domain.PostReport, error)
}

type postReportUseCase struct {
	postReportRepository repository.ReportRepo
}

func (p postReportUseCase) ReportPost(report *dto.ReportDTO, ctx context.Context) error {
	panic("implement me")
}

func (p postReportUseCase) ReviewReport(report *dto.ReportDTO, ctx context.Context) error {
	panic("implement me")
}

func (p postReportUseCase) GetAllPendingReports(ctx context.Context) ([]domain.PostReport, error) {
	panic("implement me")
}

func (p postReportUseCase) GetAllApprovedReports(ctx context.Context) ([]domain.PostReport, error) {
	panic("implement me")
}

func (p postReportUseCase) GetAllRejectedReports(ctx context.Context) ([]domain.PostReport, error) {
	panic("implement me")
}

func NewPostReportUseCase(postReportRepository repository.ReportRepo) PostReportUseCase {
	return &postReportUseCase{postReportRepository: postReportRepository}
}