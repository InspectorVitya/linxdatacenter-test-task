package reader

import (
	"encoding/json"
	"github.com/inspectorvitya/linxdatacenter-test-task/internal/model"
	"io"
	"os"
)

type JSONReader struct {
	decoder *json.Decoder
}

func NewJSONReader(file *os.File) (Reader, error) {
	decoder := json.NewDecoder(file)
	_, err := decoder.Token()
	if err != nil {
		return nil, err
	}
	return &JSONReader{decoder}, nil
}

func (r *JSONReader) Read() (model.Product, error) {
	product := model.Product{}
	if !r.decoder.More() {
		return product, io.EOF
	}
	err := r.decoder.Decode(&product)
	return product, err
}
