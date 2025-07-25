package handlers

import (
	"encoding/json"
	"rodrigoorlandini/urlshortener/analytics/internal/application/factories"
	useCases "rodrigoorlandini/urlshortener/analytics/internal/application/use-cases"
)

type URLAccessedHandler struct{}

func NewURLAccessedHandler() *URLAccessedHandler {
	return &URLAccessedHandler{}
}

func (h *URLAccessedHandler) Handle(data []byte) error {
	var request useCases.HandleURLAccessedUseCaseRequest

	err := json.Unmarshal(data, &request)
	if err != nil {
		return err
	}

	useCase, err := factories.NewHandleURLAccessedUseCase()
	if err != nil {
		return err
	}

	_, err = useCase.Execute(request)
	return err
}
