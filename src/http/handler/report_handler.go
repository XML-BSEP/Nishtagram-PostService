package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/dto"
	"post-service/http/middleware"
	"post-service/usecase"
)

type ReportPostHandler interface {
	ReportPost(context *gin.Context)
	ReviewReport(context *gin.Context)
	GetAllPendingReports(context *gin.Context)
	GetAllApprovedReports(context *gin.Context)
	GetAllRejectedReports(context *gin.Context)
	GetAllReportTypes(context *gin.Context)

}

type reportPostHandler struct {
	reportPostUseCase usecase.PostReportUseCase
	logger *logger.Logger
}

func (r reportPostHandler) GetAllReportTypes(context *gin.Context) {
	types, err := r.reportPostUseCase.GetAllReportType(context)

	if err != nil {
		context.JSON(500, gin.H{"message" : "server error"})
		context.Abort()
		return
	}

	context.JSON(200, types)
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
		context.Abort()
		return
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
		context.JSON(400, gin.H{"message" : "Bad request"})
		context.Abort()
		return
	}

	createReport.ReportedBy, _ = middleware.ExtractUserId(context.Request, r.logger)

	err := r.reportPostUseCase.ReportPost(createReport, context)

	if err != nil {
		context.JSON(500, gin.H{"message" : "Server error"})
	}

	context.JSON(200, gin.H{"message" : "Thanks for your report!"})
}

func NewReportPostHandler(reportPostUseCase usecase.PostReportUseCase, logger *logger.Logger) ReportPostHandler {
	return &reportPostHandler{reportPostUseCase: reportPostUseCase, logger: logger}
}
