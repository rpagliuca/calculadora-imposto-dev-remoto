package main

import (
	"math"
)

func gerarEsquemaFiscal(faturamentoAnual float64) (map[string]interface{}, error) {
	impostoSimples := 0.0
	restante := faturamentoAnual
	limiteAnterior := 0.0

	// Cálculo progressivo
	for _, faixa := range faixasAnexo3 {
		larguraFaixa := faixa.LimiteMaximo - limiteAnterior
		impostoSimples += float64(math.Min(restante, larguraFaixa)) * faixa.Aliquota / 100 * PERCENTUAL_CPP_CSLL_IRPJ / 100
		restante = restante - larguraFaixa
		limiteAnterior = faixa.LimiteMaximo
		if restante <= 0 {
			break
		}
	}

	proLabore := faturamentoAnual * PERCENTUAL_FATOR_R / 100

	// Cálculo INSS
	inss := math.Min(TETO_BASE_INSS_ANUAL, proLabore) * PERCENTUAL_INSS / 100

	// Cálculo IR do pró-labore para fator R
	restante = proLabore - inss // INSS deve ser abatido da base cálculo do IR
	impostoRendaPF := 0.0
	limiteAnterior = 0.0
	for _, faixa := range faixasImpostoRendaPF {
		larguraFaixa := faixa.LimiteMaximo - limiteAnterior
		impostoRendaPF += float64(math.Min(restante, larguraFaixa)) * faixa.Aliquota / 100.0
		restante = restante - larguraFaixa
		limiteAnterior = faixa.LimiteMaximo
		if restante <= 0 {
			break
		}
	}

	descontoTotal := impostoSimples + impostoRendaPF + inss

	// Dados de saída
	data := map[string]interface{}{
		"input": map[string]interface{}{
			"faturamento-anual-em-reais": faturamentoAnual,
		},
		"output": map[string]interface{}{
			"impostos-detalhados": map[string]interface{}{
				"imposto-simples-nacional":    impostoSimples,
				"imposto-renda-pessoa-fisica": impostoRendaPF,
				"inss":                        inss,
			},
			"imposto-total-em-reais":   descontoTotal,
			"imposto-total-percentual": descontoTotal / faturamentoAnual * 100,
		},
		"sobre": map[string]interface{}{
			"repositorio-github":                 "https://github.com/rpagliuca/calculadora-imposto-dev-remoto",
			"constantes-utilizadas-na-simulacao": "https://github.com/rpagliuca/calculadora-imposto-dev-remoto/blob/master/function1/fiscal.go",
		},
	}
	return data, nil
}

// Fator de ajuste considerando que todo o faturamento anual
// é proveniente de clientes no exterior
const PERCENTUAL_CPP_CSLL_IRPJ = 50.9

const PERCENTUAL_FATOR_R = 28.0

const PERCENTUAL_INSS = 11.0

const TETO_BASE_INSS_ANUAL = 6101.06 * 12

type Faixa struct {
	LimiteMaximo float64 `json:"limite-maximo"`
	Aliquota     float64 `json:"aliquota"`
}

var faixasAnexo3 = []Faixa{
	Faixa{180000.0, 6},
	Faixa{360000.0, 11.2},
	Faixa{720000.0, 13.5},
	Faixa{1800000.0, 16},
}

var faixasImpostoRendaPF = []Faixa{
	Faixa{22847.76, 0},
	Faixa{33919.80, 7.5},
	Faixa{45012.60, 15},
	Faixa{55976.16, 22.5},
	Faixa{999999999999999999.99, 27.5},
}
