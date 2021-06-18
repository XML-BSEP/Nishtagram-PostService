package usecase

import (
	"context"
	"post-service/dto"
	"post-service/repository"
)

type PostReportUseCase interface {
	ReportPost(report dto.CreateReportDTO, ctx context.Context) error
	ReviewReport(report dto.ReviewReportDTO, ctx context.Context) error
	GetAllPendingReports(ctx context.Context) ([]dto.ReportDTO, error)
	GetAllApprovedReports(ctx context.Context) ([]dto.ReportDTO, error)
	GetAllRejectedReports(ctx context.Context) ([]dto.ReportDTO, error)
	GetAllReportType(ctx context.Context) ([]string, error)
}

type postReportUseCase struct {
	postReportRepository repository.ReportRepo
}

func (p postReportUseCase) GetAllReportType(ctx context.Context) ([]string, error) {
	return p.postReportRepository.GetAllReportTypes(ctx)
}

func (p postReportUseCase) ReportPost(report dto.CreateReportDTO, ctx context.Context) error {
	return p.postReportRepository.ReportPost(report, context.Background())
}

func (p postReportUseCase) ReviewReport(report dto.ReviewReportDTO, ctx context.Context) error {
	return p.postReportRepository.ReviewReport(report, context.Background())
}

func (p postReportUseCase) GetAllPendingReports(ctx context.Context) ([]dto.ReportDTO, error) {
	return p.postReportRepository.GetAllPendingReports(context.Background())
}

func (p postReportUseCase) GetAllApprovedReports(ctx context.Context) ([]dto.ReportDTO, error) {
	return p.postReportRepository.GetAllApprovedReports(context.Background())
}

func (p postReportUseCase) GetAllRejectedReports(ctx context.Context) ([]dto.ReportDTO, error) {
	return p.postReportRepository.GetAllRejectedReports(context.Background())
}

func NewPostReportUseCase(postReportRepository repository.ReportRepo) PostReportUseCase {
	return &postReportUseCase{postReportRepository: postReportRepository}
}