package handler

import (
	"github.com/gin-gonic/gin"
	"post-service/usecase"
)

type ReportPostHandler interface {
	ReportPost(context *gin.Context)
}

type reportPostHandler struct {
	reportPostUseCase usecase.PostReportUseCase
}

func (r reportPostHandler) ReportPost(context *gin.Context) {
	panic("implement me")
}

func NewReportPostHandler(reportPostUseCase usecase.PostReportUseCase) ReportPostHandler {
	return &reportPostHandler{reportPostUseCase: reportPostUseCase}
}
