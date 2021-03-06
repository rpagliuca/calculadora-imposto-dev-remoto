package main

import (
	"math"
)

func gerarEsquemaFiscal(faturamentoAnual float64) (map[string]interface{}, error) {

	// Cálculo de cada termo do imposto
	impostoSimplesNacional := calcularImpostoSimplesNacional(faturamentoAnual)
	proLabore := calcularProLabore(faturamentoAnual)
	inss := calcularInss(proLabore)
	impostoRendaPF := calcularImpostoRendaPF(proLabore, inss)

	// Soma de todos os impostos
	impostoTotal := impostoSimplesNacional + impostoRendaPF + inss

	// Formatação da estrutura de dados de saída
	dados := formatarSaida(
		faturamentoAnual,
		impostoSimplesNacional,
		impostoRendaPF,
		inss,
		impostoTotal,
	)

	return dados, nil
}

// Cálculo progressivo do imposto do Simples Nacional, já descontando ISS, PIS e COFINS
func calcularImpostoSimplesNacional(faturamentoAnual float64) float64 {
	impostoSimples := 0.0
	restante := faturamentoAnual
	limiteAnterior := 0.0
	for _, faixa := range FAIXAS_ANEXO_3_SIMPLES_NACIONAL {
		larguraFaixa := faixa.LimiteMaximo - limiteAnterior
		impostoSimples += float64(math.Min(restante, larguraFaixa)) * faixa.Aliquota / 100 * PERCENTUAL_CPP_CSLL_IRPJ / 100
		restante = restante - larguraFaixa
		limiteAnterior = faixa.LimiteMaximo
		if restante <= 0 {
			break
		}
	}
	return impostoSimples
}

// Cálculo do pró-labore necessário para se enquadrar no Anexo 3, com base no FATOR R
func calcularProLabore(faturamentoAnual float64) float64 {
	return faturamentoAnual * PERCENTUAL_FATOR_R / 100
}

// Cálculo do INSS sobre pró-labore
func calcularInss(proLabore float64) float64 {
	return math.Min(TETO_BASE_INSS_ANUAL, proLabore) * PERCENTUAL_INSS / 100
}

// Cálculo do Imposto de Renda Pessoa Física sobre pró-labore para fator R
func calcularImpostoRendaPF(proLabore, inss float64) float64 {
	restante := proLabore - inss // INSS deve ser abatido da base cálculo do IR
	impostoRendaPF := 0.0
	limiteAnterior := 0.0
	for _, faixa := range FAIXAS_IMPOSTO_RENDA_PF {
		larguraFaixa := faixa.LimiteMaximo - limiteAnterior
		impostoRendaPF += float64(math.Min(restante, larguraFaixa)) * faixa.Aliquota / 100.0
		restante = restante - larguraFaixa
		limiteAnterior = faixa.LimiteMaximo
		if restante <= 0 {
			break
		}
	}
	return impostoRendaPF
}

func formatarSaida(
	faturamentoAnual,
	impostoSimplesNacional,
	impostoRendaPF,
	inss,
	impostoTotal float64,
) map[string]interface{} {
	// Dados de saída
	return map[string]interface{}{
		"input": map[string]interface{}{
			"faturamento-anual-em-reais": faturamentoAnual,
		},
		"output": map[string]interface{}{
			"impostos-detalhados": map[string]interface{}{
				"imposto-simples-nacional":    impostoSimplesNacional,
				"imposto-renda-pessoa-fisica": impostoRendaPF,
				"inss":                        inss,
			},
			"imposto-total-em-reais":     impostoTotal,
			"imposto-total-percentual":   impostoTotal / faturamentoAnual * 100,
			"faturamento-menos-impostos": faturamentoAnual - impostoTotal,
		},
		"sobre": map[string]interface{}{
			"repositorio-github":                 "https://github.com/rpagliuca/calculadora-imposto-dev-remoto",
			"constantes-utilizadas-na-simulacao": "https://github.com/rpagliuca/calculadora-imposto-dev-remoto/blob/master/main/fiscal.go",
		},
	}
}

// Fator de ajuste considerando que todo o faturamento anual é proveniente de clientes no exterior
// É o mesmo para as 4 primeiras faixas do Anexo 3, obtido somando CPP + CSLL + IRPJ
// Fonte: https://blog.contabilizei.com.br/contabilidade-online/anexo-3-simples-nacional/
const PERCENTUAL_CPP_CSLL_IRPJ = 50.9

// Fonte: http://normas.receita.fazenda.gov.br/sijut2consulta/link.action?idAto=92278
const PERCENTUAL_FATOR_R = 28.0

// Fonte: https://www.contabilizei.com.br/contabilidade-online/o-que-e-o-pro-labore/
const PERCENTUAL_INSS = 11.0

// Fonte: https://www.inss.gov.br/servicos-do-inss/calculo-da-guia-da-previdencia-social-gps/tabela-de-contribuicao-mensal/
const TETO_BASE_INSS_ANUAL = 6101.06 * 12

type Faixa struct {
	LimiteMaximo float64 `json:"limite-maximo"`
	Aliquota     float64 `json:"aliquota"`
}

// Fonte: https://contabilizei.com.br/contabilidade-online/anexo-3-simples-nacional
var FAIXAS_ANEXO_3_SIMPLES_NACIONAL = []Faixa{
	Faixa{180000.0, 6},
	Faixa{360000.0, 11.2},
	Faixa{720000.0, 13.5},
	Faixa{1800000.0, 16},
}

// Fonte: http://receita.economia.gov.br/acesso-rapido/tributos/irpf-imposto-de-renda-pessoa-fisica#c-lculo-anual-do-irpf
var FAIXAS_IMPOSTO_RENDA_PF = []Faixa{
	Faixa{22847.76, 0},
	Faixa{33919.80, 7.5},
	Faixa{45012.60, 15},
	Faixa{55976.16, 22.5},
	Faixa{999999999999999999.99, 27.5},
}
