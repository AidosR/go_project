package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidWeightFormat = errors.New("invalid weight format")

type Weight float32

func (w Weight) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%.3f kg", w)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (w *Weight) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidWeightFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "kg" {
		return ErrInvalidWeightFormat
	}

	i, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		return ErrInvalidWeightFormat
	}

	*w = Weight(i)
	return nil
}
