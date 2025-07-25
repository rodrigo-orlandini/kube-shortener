package repositories

import (
	"rodrigoorlandini/urlshortener/analytics/internal/application/repositories"
	customError "rodrigoorlandini/urlshortener/analytics/internal/domain/custom-error"
	"rodrigoorlandini/urlshortener/analytics/internal/domain/entities"
	"rodrigoorlandini/urlshortener/analytics/internal/infrastructure/database"
	"rodrigoorlandini/urlshortener/analytics/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

type PostgresURLsRepository struct {
	db *gorm.DB
}

func NewPostgresURLsRepository() (repositories.URLsRepository, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	return &PostgresURLsRepository{
		db: db.Connection,
	}, nil
}

func (r *PostgresURLsRepository) Create(url entities.URL) error {
	return r.db.Create(&models.URL{
		ShortURL: url.ShortURL,
		Visits:   url.Visits,
	}).Error
}

func (r *PostgresURLsRepository) GetTopRanked(limit int) ([]repositories.TopRankedURL, error) {
	var results []models.URL

	err := r.db.Order("visits DESC").Limit(limit).Find(&results).Error
	if err != nil {
		return nil, err
	}

	topRankedURLs := make([]repositories.TopRankedURL, len(results))
	for i, result := range results {
		topRankedURLs[i] = repositories.TopRankedURL{
			ShortURL:   result.ShortURL,
			VisitCount: result.Visits,
		}
	}

	return topRankedURLs, nil
}

func (r *PostgresURLsRepository) FindByShortURL(shortURL string) (entities.URL, error) {
	var result models.URL

	err := r.db.Where("short_url = ?", shortURL).First(&result).Error
	if err != nil {
		if err.Error() == "record not found" {
			return entities.URL{}, &customError.NotFoundError{
				Entity: "URL",
				Field:  "short_url",
				Value:  shortURL,
			}
		}

		return entities.URL{}, err
	}

	url, err := entities.NewURL(result.ShortURL, entities.WithVisits(result.Visits))
	if err != nil {
		return entities.URL{}, err
	}

	return *url, nil
}

func (r *PostgresURLsRepository) IncrementVisits(shortURL string) error {
	return r.db.Model(&models.URL{}).
		Where("short_url = ?", shortURL).
		Update("visits", gorm.Expr("visits + 1")).Error
}
