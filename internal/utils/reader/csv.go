package reader

import (
	"encoding/csv"
	"github.com/inspectorvitya/linxdatacenter-test-task/internal/model"
	"os"
	"strconv"
)

type CSVReader struct {
	reader *csv.Reader
}

func NewCsvReader(file *os.File) (Reader, error) {
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comma = ';'
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}
	return &CSVReader{reader}, nil
}

func (r *CSVReader) Read() (model.Product, error) {
	cols, err := r.reader.Read()
	if err != nil {
		return model.Product{}, err
	}

	price, err := strconv.Atoi(cols[1])
	if err != nil {
		return model.Product{}, err
	}

	rating, err := strconv.Atoi(cols[2])
	if err != nil {
		return model.Product{}, err
	}

	return model.Product{
		Name:   cols[0],
		Price:  price,
		Rating: rating,
	}, nil
}
