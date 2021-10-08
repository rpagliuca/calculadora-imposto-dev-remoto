package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

func Handler(ctx context.Context, request Request) (Response, error) {
	// Read parameters from query string
	faturamentoAnual, err := readParameters(request)
	if err != nil {
		return createErrorResponse(err)
	}

	if faturamentoAnual > 1800000{
		return createErrorResponse(fmt.Errorf("faturamento Anual maior que o teto da faixa 4 (1.800.000)"))
	}
	// Generate fiscal scheme
	data, err := gerarEsquemaFiscal(faturamentoAnual)
	if err != nil {
		return createErrorResponse(err)
	}

	// Return JSON response
	return createSuccessResponse(data)
}

// Allow mocking
var lambdaStart = lambda.Start

func main() {
	lambdaStart(Handler)
}
