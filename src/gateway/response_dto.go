package gateway

import "post-service/domain"

type FollowingResponseDTO struct {
	Data []domain.Profile `json:"data"`
}
