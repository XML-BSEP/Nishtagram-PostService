package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"post-service/dto"
	"post-service/usecase"
)

type ReportPostHandler interface {
	ReportPost(context *gin.Context)
	ReviewReport(context *gin.Context)
	GetAllPendingReports(context *gin.Context)
	GetAllApprovedReports(context *gin.Context)
	GetAllRejectedReports(context *gin.Context)

}

type reportPostHandler struct {
	reportPostUseCase usecase.PostReportUseCase
}

func (r reportPostHandler) ReviewReport(context *gin.Context) {
	var reviewReportDTO dto.ReviewReportDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&reviewReportDTO); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := r.reportPostUseCase.ReviewReport(reviewReportDTO, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, "ok")
}

func (r reportPostHandler) GetAllPendingReports(context *gin.Context) {

	reports, err := r.reportPostUseCase.GetAllPendingReports(context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, reports)
}

func (r reportPostHandler) GetAllApprovedReports(context *gin.Context) {

	reports, err := r.reportPostUseCase.GetAllApprovedReports(context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, reports)
}

func (r reportPostHandler) GetAllRejectedReports(context *gin.Context) {

	reports, err := r.reportPostUseCase.GetAllRejectedReports(context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, reports)
}

func (r reportPostHandler) ReportPost(context *gin.Context) {
	var createReport dto.CreateReportDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&createReport); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := r.reportPostUseCase.ReportPost(createReport, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, "ok")
}

func NewReportPostHandler(reportPostUseCase usecase.PostReportUseCase) ReportPostHandler {
	return &reportPostHandler{reportPostUseCase: reportPostUseCase}
}
