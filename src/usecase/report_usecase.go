package usecase

import (
	"post-service/domain"
	"post-service/dto"
	"post-service/repository"
)

type PostReportUseCase interface {
	ReportPost(report *dto.ReportDTO) error
	ReviewReport(report *dto.ReportDTO) error
	GetAllPendingReports() ([]domain.PostReport, error)
	GetAllApprovedReports() ([]domain.PostReport, error)
	GetAllRejectedReports() ([]domain.PostReport, error)
}

type postReportUseCase struct {
	postReportRepository repository.ReportRepo
}

func (p postReportUseCase) ReportPost(report *dto.ReportDTO) error {
	panic("implement me")
}

func (p postReportUseCase) ReviewReport(report *dto.ReportDTO) error {
	panic("implement me")
}

func (p postReportUseCase) GetAllPendingReports() ([]domain.PostReport, error) {
	panic("implement me")
}

func (p postReportUseCase) GetAllApprovedReports() ([]domain.PostReport, error) {
	panic("implement me")
}

func (p postReportUseCase) GetAllRejectedReports() ([]domain.PostReport, error) {
	panic("implement me")
}

func NewPostReportUseCase(postReportRepository repository.ReportRepo) PostReportUseCase {
	return &postReportUseCase{postReportRepository: postReportRepository}
}