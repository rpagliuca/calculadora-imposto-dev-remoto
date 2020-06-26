package main

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"testing"
)

type testCase struct {
	inputFaturamentoAnual      string
	expectedDescontoPercentual float64
	expectedImpostoRendaPF     float64
}

func Test(t *testing.T) {
	testCases := []testCase{
		// Teste baseado em https://medium.com/@hnordt/como-pagar-6-de-impostos-dentro-da-lei-a91c23868ec6
		testCase{"120000", 6.575, 0},
		// IR calculado usando simulador http://www.receita.fazenda.gov.br/Aplicacoes/ATRJO/Simulador/simulador.asp?tipoSimulador=A
		testCase{"240000", 9.302, 6014.88},
		testCase{"280000", 10.059, 8912.99},
	}
	for _, tc := range testCases {
		request := Request{
			QueryStringParameters: map[string]string{
				"faturamento-anual": tc.inputFaturamentoAnual,
			},
		}
		out, err := Handler(context.TODO(), request)
		if err != nil {
			t.Fatal(err)
		}
		var data map[string]interface{}
		err = json.Unmarshal([]byte(out.Body), &data)
		log.Printf(out.Body + "\n\n")
		if err != nil {
			t.Fatal(err)
		}
		desc := data["output"].(map[string]interface{})["imposto-total-percentual"].(float64)
		irpf := data["output"].(map[string]interface{})["impostos-detalhados"].(map[string]interface{})["imposto-renda-pessoa-fisica"].(float64)
		if !compare(desc, tc.expectedDescontoPercentual, 0.001) {
			t.Fatalf(
				"inputFaturamentoAnual: %s. Invalid expectedDescontoPercentual. Expected: %f. Got: %f\n",
				tc.inputFaturamentoAnual,
				tc.expectedDescontoPercentual,
				desc,
			)
		}
		if tc.expectedImpostoRendaPF > 0 && !compare(tc.expectedImpostoRendaPF, irpf, 0.01) {
			t.Fatalf(
				"inputFaturamentoAnual: %s. Invalid expectedImpostoRendaPF. Expected: %f. Got: %f\n",
				tc.inputFaturamentoAnual,
				tc.expectedImpostoRendaPF,
				irpf,
			)
		}
	}
}

func compare(a, b, tolerance float64) bool {
	if math.Abs(a-b) < tolerance {
		return true
	}
	return false
}
