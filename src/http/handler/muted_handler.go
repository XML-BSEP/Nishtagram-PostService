package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/dto"
	"post-service/http/middleware"
	"post-service/usecase"
)

type MutedContentHandler interface {
	MuteUser(ctx *gin.Context)
	UnmuteUser(ctx *gin.Context)
	SeeIfMuted(ctx *gin.Context)
}

type mutedContentHandler struct {
	mutedContentUseCase usecase.MutedContentUseCase
	logger *logger.Logger
}

func (m mutedContentHandler) MuteUser(ctx *gin.Context) {
	var muteDTO dto.MuteUserDTO
	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(&muteDTO); err != nil {
		ctx.JSON(400, "invalid request")
		ctx.Abort()
		return
	}
	userRequested, _ := middleware.ExtractUserId(ctx.Request, m.logger)
	muteDTO.BlockedFor = userRequested

	err := m.mutedContentUseCase.MuteUser(muteDTO.BlockedFor, muteDTO.Blocked, ctx)

	if err != nil {
		ctx.JSON(500, "server error")
		ctx.Abort()
		return
	}
	ctx.JSON(200, gin.H{"message" : "Successfully muted"})
}

func (m mutedContentHandler) UnmuteUser(ctx *gin.Context) {
	var muteDTO dto.MuteUserDTO
	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(&muteDTO); err != nil {
		ctx.JSON(400, "invalid request")
		ctx.Abort()
		return
	}
	userRequested, _ := middleware.ExtractUserId(ctx.Request, m.logger)
	muteDTO.BlockedFor = userRequested

	err := m.mutedContentUseCase.UnmuteUser(muteDTO.BlockedFor, muteDTO.Blocked, ctx)

	if err != nil {
		ctx.JSON(500, "server error")
		ctx.Abort()
		return
	}
	ctx.JSON(200, gin.H{"message" : "Successfully unmute"})
}

func (m mutedContentHandler) SeeIfMuted(ctx *gin.Context) {
	var muteDTO dto.MuteUserDTO
	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(&muteDTO); err != nil {
		ctx.JSON(500, "invalid request")
		ctx.Abort()
		return
	}
	userRequested, _ := middleware.ExtractUserId(ctx.Request, m.logger)
	muteDTO.BlockedFor = userRequested

	isMuted := m.mutedContentUseCase.SeeIfMuted(muteDTO.BlockedFor, muteDTO.Blocked, ctx)

	if !isMuted {
		ctx.JSON(400, gin.H{"message" : "user not muted"})
		ctx.Abort()
		return
	}
	ctx.JSON(200, gin.H{"message" : "user muted"})
}

func NewMuteContentHandler(mutedContentUseCase usecase.MutedContentUseCase) MutedContentHandler {
	return &mutedContentHandler{mutedContentUseCase: mutedContentUseCase}
}
