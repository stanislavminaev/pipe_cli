package jira_wick

import (
	"github.com/go-resty/resty/v2"
	"log/slog"
)

// JiraWickClient интерфейс для restyClient
type JiraWickClient interface {
	CloseRelease(req *CloseReleaseRequest) (*CloseReleaseResponse, error)
}

// restyClient структура для работы с сервисом JiraWick
type restyClient struct {
	client  *resty.Client
	baseURL string
}

// CloseRelease закрытие релиза в жире
func (r *restyClient) CloseRelease(req *CloseReleaseRequest) (*CloseReleaseResponse, error) {
	var response CloseReleaseResponse
	_, err := r.client.R().
		SetBody(req).
		SetResult(&response).
		Post(r.baseURL)

	if err != nil {
		slog.Error("Не удалось подключиться к сервису JiraWick", "error", err)
		return nil, err
	}

	if !response.Status {
		slog.Error("Ошибка JiraWick", "comment", response.Comment)
		return nil, err
	}
	slog.Info("Ответ JiraWick", "comment", response.Comment)
	return &response, nil
}

// NewJiraWickClient конструктор клиента для сервиса JiraWick
func NewJiraWickClient(baseURL string) JiraWickClient {
	r := resty.New()
	r.SetHeader("Content-Type", "application/json")
	return &restyClient{
		client:  r,
		baseURL: baseURL,
	}
}
