package entities

import (
	"regexp"
	customError "rodrigoorlandini/urlshortener/shortener/internal/domain/custom-error"
)

type URL struct {
	OriginalURL string
	ShortURL    string
}

type URLOption func(*URL)

func NewURL(originalURL string, options ...URLOption) (*URL, error) {
	if originalURL == "" || !isValid(originalURL) {
		return nil, &customError.InvalidEntityCreationError{
			EntityName: "URL",
			Field:      "originalURL",
			Value:      originalURL,
		}
	}

	url := &URL{
		OriginalURL: originalURL,
		ShortURL:    "",
	}

	for _, option := range options {
		option(url)
	}

	return url, nil
}

func WithShortURL(shortURL string) URLOption {
	return func(url *URL) {
		url.ShortURL = shortURL
	}
}

func isValid(url string) bool {
	re := regexp.MustCompile(`^(https?:\/\/)?([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(:\d{1,5})?(\/[^\s?#]*)?(\?[^\s#]*)?(#[^\s]*)?$`)
	return re.MatchString(url)
}
