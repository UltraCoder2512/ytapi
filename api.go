package api

import (
	"context"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func GetAPISearchResults(query string, max int64) (map[string]string, int, error) {
	count := 0
	videoMap := map[string]string{}
	baseURL := "https://www.youtube.com/watch?v="
	service, err := youtube.NewService(context.Background(), option.WithAPIKey(os.Getenv("YT_API_KEY")))
	if err != nil {
		return nil, count, err
	}
	call := service.Search.List([]string{"snippet"}).Q(query).MaxResults(max)
	response, err := call.Do()
	if err != nil {
		return nil, count, err
	}
	for _, v := range response.Items {
		if v.Id.VideoId == "" {
			continue
		}
		count++
		videoURL := baseURL + v.Id.VideoId
		title := v.Snippet.Title
		videoMap[title] = videoURL
	}
	return videoMap, count, nil
}
