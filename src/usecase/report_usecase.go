package usecase

import (
	"context"
	"post-service/domain"
	"post-service/dto"
	"post-service/gateway"
	"post-service/repository"
)

type PostReportUseCase interface {
	ReportPost(report dto.CreateReportDTO, ctx context.Context) error
	ReviewReport(report dto.ReviewReportDTO, ctx context.Context) error
	GetAllPendingReports(ctx context.Context) (*[]dto.ReportDTO, error)
	GetAllApprovedReports(ctx context.Context) (*[]dto.ReportDTO, error)
	GetAllRejectedReports(ctx context.Context) (*[]dto.ReportDTO, error)
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

func (p postReportUseCase) GetAllPendingReports(ctx context.Context) (*[]dto.ReportDTO, error) {
	reports, err := p.postReportRepository.GetAllPendingReports(context.Background())
	if err != nil {
		return nil, err
	}

	for i, r := range *reports {
		reportedBy, err := gateway.GetUser(ctx, r.ReportBy.Id)
		if err == nil {
			r.ReportBy = domain.Profile{Id: r.ReportBy.Id, ProfilePhoto: reportedBy.ProfilePhoto, Username: reportedBy.Username}
		}


		reportedUser, err := gateway.GetUser(ctx, r.ReportedPostBy.Id)
		if err == nil {
			r.ReportedPostBy = domain.Profile{Id: r.ReportedPostBy.Id, ProfilePhoto: reportedUser.ProfilePhoto, Username: reportedUser.Username}
		}

		(*reports)[i] = r

	}

	return reports, nil
}

func (p postReportUseCase) GetAllApprovedReports(ctx context.Context) (*[]dto.ReportDTO, error) {
	reports, err := p.postReportRepository.GetAllApprovedReports(context.Background())

	if err != nil {
		return nil, err
	}

	for i, r := range *reports {
		reportedBy, err := gateway.GetUser(ctx, r.ReportBy.Id)
		if err == nil {
			r.ReportBy = domain.Profile{Id: r.ReportBy.Id, ProfilePhoto: reportedBy.ProfilePhoto, Username: reportedBy.Username}
		}


		reportedUser, err := gateway.GetUser(ctx, r.ReportedPostBy.Id)
		if err == nil {
			r.ReportedPostBy = domain.Profile{Id: r.ReportedPostBy.Id, ProfilePhoto: reportedUser.ProfilePhoto, Username: reportedUser.Username}
		}

		(*reports)[i] = r
	}

	return reports, nil
}

func (p postReportUseCase) GetAllRejectedReports(ctx context.Context) (*[]dto.ReportDTO, error) {
	reports, err := p.postReportRepository.GetAllRejectedReports(context.Background())
	if err != nil {
		return nil, err
	}

	for i, r := range *reports {
		reportedBy, err := gateway.GetUser(ctx, r.ReportBy.Id)
		if err == nil {
			r.ReportBy = domain.Profile{Id: r.ReportBy.Id, ProfilePhoto: reportedBy.ProfilePhoto, Username: reportedBy.Username}
		}


		reportedUser, err := gateway.GetUser(ctx, r.ReportedPostBy.Id)
		if err == nil {
			r.ReportedPostBy = domain.Profile{Id: r.ReportedPostBy.Id, ProfilePhoto: reportedUser.ProfilePhoto, Username: reportedUser.Username}
		}

		(*reports)[i] = r
	}

	return reports, nil

}

func NewPostReportUseCase(postReportRepository repository.ReportRepo) PostReportUseCase {
	return &postReportUseCase{postReportRepository: postReportRepository}
}