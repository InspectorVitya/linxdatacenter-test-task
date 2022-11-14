package app

import (
	"github.com/inspectorvitya/linxdatacenter-test-task/internal/model"
	"github.com/inspectorvitya/linxdatacenter-test-task/internal/utils/reader"
	"io"
)

type App struct {
	reader reader.Reader
}

func New(reader reader.Reader) *App {
	return &App{reader: reader}
}

func (app *App) GetMaxPriceAndRating() (model.Product, model.Product, error) {
	maxPrice, maxRating := model.Product{}, model.Product{}
	for {
		product, err := app.reader.Read()
		if err != nil {
			if err == io.EOF {
				return maxPrice, maxRating, nil
			}
			return model.Product{}, model.Product{}, err
		}
		if product.Price > maxPrice.Price {
			maxPrice = product
		}
		if product.Rating > maxRating.Rating {
			maxRating = product
		}

	}
}
