package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

func readParameters(request Request) (float64, error) {
	val, ok := request.QueryStringParameters["faturamento-anual"]
	if !ok {
		return 0, errors.New("Parâmetro faturamento-anual deve ser inteiro")
	}
	valInt, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, err
	}
	return valInt, nil
}

func createSuccessResponse(data map[string]interface{}) (Response, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return createErrorResponse(err)
	}
	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
}

func createErrorResponse(err error) (Response, error) {
	body, err := json.Marshal(map[string]interface{}{
		"success": false,
		"message": err.Error(),
	})
	if err != nil {
		return Response{StatusCode: 400}, err
	}
	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      400,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
}
