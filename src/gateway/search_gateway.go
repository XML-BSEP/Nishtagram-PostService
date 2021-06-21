package gateway

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"post-service/dto"
)

func SaveNewPostLocation(postLocation dto.PostLocationProfileDTO, ctx context.Context) error {
	var domain string
	if os.Getenv("DOCKER_ENV") == "" {
		domain = "http://127.0.0.1:8087"
	} else {
		domain = "http://searchms:8087"
	}
	client := resty.New()

	resp, err := client.R().
		SetBody(postLocation).
		SetContext(ctx).
		EnableTrace().
		Post(domain + "/saveNewPostLocation")

	if err != nil {
		return nil
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("oh no i hope i don't fall")
	}

	return nil


}



func SaveNewPostTage(postTag dto.PostTagProfileDTO, ctx context.Context) error {
	var domain string
	if os.Getenv("DOCKER_ENV") == "" {
		domain = "http://127.0.0.1:8087"
	} else {
		domain = "http://searchms:8087"
	}
	client := resty.New()

	resp, err := client.R().
		SetBody(postTag).
		SetContext(ctx).
		EnableTrace().
		Post(domain + "/saveNewPostTag")

	if err != nil {
		return nil
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("oh no i hope i don't fall")
	}

	return nil


}
