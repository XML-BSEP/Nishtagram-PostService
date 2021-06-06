package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"post-service/domain"
	"post-service/dto"
)

func GetAllUserFollowing(ctx context.Context, userId string) ([]domain.Profile, error) {
	client := resty.New()
	userDto := dto.UserIdDTO{UserId: userId}
	resp, _ := client.R().
		SetBody(userDto).
		EnableTrace().
		Post("http://127.0.0.1:8089/usersFollowings")

	var responseDTO FollowingResponseDTO
	err := json.Unmarshal(resp.Body(), &responseDTO)
	if err != nil {
		fmt.Println(err)
	}

	return responseDTO.Data, nil
}
