# 📐 KPIs — Dicionário e Série Histórica

> **Referência:** ago/2026 (parcial, dia 16) · Todo indicador tem fórmula, meta e histórico.

---

## 1. 🩺 Score de Saúde Financeira

**48 / 100** — `████▊░░░░░` 🟡 *Frágil, mas recuperável*

| Pilar             | Peso | Nota | Contribuição | Leitura                                  |
| ----------------- | ---: | ---: | -----------: | ---------------------------------------- |
| 💧 Liquidez       |  25% |   30 |         7,5  | Reserva cobre só 3,2 meses               |
| 📦 Orçamento      |  20% |   62 |        12,4  | 5 de 7 envelopes no teto                 |
| 🐖 Poupança       |  25% |   20 |         5,0  | Taxa negativa no parcial                 |
| 🌐 Diversificação |  15% |   25 |         3,8  | 100% da receita em 1 cliente             |
| 🎯 Disciplina     |  15% |   70 |        10,5  | Lançamentos em dia, poucos estouros      |
| **Total**         |100%  |   —  |   **48,2**   | 🟡                                        |

**Faixas:** 0–39 🔴 crítico · 40–59 🟡 frágil · 60–79 🟢 saudável · 80–100 💎 sólido

```
Score  jan  fev  mar  abr  mai  jun  jul  ago
       52   55   58   56   61   66   44   48
       ██▌  ██▊  ██▉  ██▊  ███  ███▎ ██▏  ██▍
```

---

## 2. 📖 Dicionário de Indicadores

### 💧 Liquidez e Segurança

| KPI                   | Fórmula                                | Atual   | Meta    | Status | Tend. |
| --------------------- | -------------------------------------- | ------: | ------: | ------ | ----- |
| **Runway da Reserva** | Reserva ÷ despesa mensal média          | 3,2 m   | ≥ 6 m   | 🟡     | ↗     |
| **Liquidez Imediata** | Caixa ÷ despesa dos próximos 30 dias    | 1,10x   | ≥ 1,5x  | 🟡     | ↘     |
| **Vale de Caixa**     | Menor saldo projetado no mês            | -R$ 380 | ≥ R$ 0  | 🔴     | ↘     |
| **Dias de Fôlego**    | Caixa ÷ burn rate diário                | 25 d    | ≥ 45 d  | 🟡     | →     |

### 🐖 Acumulação

| KPI                     | Fórmula                              | Atual  | Meta   | Status | Tend. |
| ----------------------- | ------------------------------------ | -----: | -----: | ------ | ----- |
| **Taxa de Poupança**    | (Receita − Despesa) ÷ Receita        | -40,0% | ≥ 30%  | 🔴     | ↘     |
| **Índice de Cobertura** | Receita ÷ Despesa                    |  0,71x | ≥ 1,5x | 🔴     | ↘     |
| **Δ Patrimônio (mês)**  | (PL_fim − PL_ini) ÷ PL_ini           |    -2% | ≥ +2%  | 🔴     | ↘     |
| **Aporte Efetivo**      | Valor investido ÷ meta de aporte     |     0% | 100%   | 🔴     | →     |

### 📦 Controle de Gastos

| KPI                       | Fórmula                                | Atual   | Meta    | Status | Tend. |
| ------------------------- | -------------------------------------- | ------: | ------: | ------ | ----- |
| **Aderência ao Orçamento**| Envelopes dentro do teto ÷ total       |   5 / 7 |   7 / 7 | 🟡     | ↘     |
| **Custo Fixo / Receita**  | Fixos ÷ receita projetada              |   29,1% |   ≤ 50% | 🟢     | →     |
| **Discricionário**        | (Vícios + Poker) ÷ despesa total       |   19,3% |   ≤ 15% | 🟡     | ↗     |
| **Ticket Médio**          | Despesa ÷ nº de transações             | R$ 25,9 | ≤ R$ 30 | 🟢     | ↘     |
| **Dias sem Gastar**       | Contagem de dias com R$ 0 de saída     |  4 / 16 |    ≥ 8  | 🟡     | ↘     |
| **Ritmo vs. Calendário**  | % orçamento usado − % do mês decorrido | +18 p.p.| ≤ 0 p.p.| 🔴     | ↗     |

### 💼 Negócio (MEI)

| KPI                      | Fórmula                             | Atual   | Meta    | Status | Tend. |
| ------------------------ | ----------------------------------- | ------: | ------: | ------ | ----- |
| **Concentração Cliente** | Maior cliente ÷ faturamento         |    100% |   ≤ 60% | 🔴     | ↗     |
| **Margem Operacional**   | (Faturamento − custos) ÷ faturamento|   85,0% |   ≥ 70% | 🟢     | →     |
| **Volatilidade Receita** | Desvio-padrão ÷ média (12m)         |     ±38%|   ≤ 20% | 🔴     | ↗     |
| **Uso do Limite MEI**    | Faturamento 12m ÷ R$ 81.000         |   16,7% |  ≤ 100% | 🟢     | ↗     |

### 💳 Endividamento

| KPI                       | Fórmula                          | Atual  | Meta   | Status | Tend. |
| ------------------------- | -------------------------------- | -----: | -----: | ------ | ----- |
| **Dívida / Patrimônio**   | Passivos ÷ ativos                | 21,6%  | ≤ 30%  | 🟢     | ↘     |
| **Comprometimento Renda** | Parcelas ÷ receita mensal        |  9,0%  | ≤ 30%  | 🟢     | ↘     |
| **Dívida Cara**           | Saldo com juros > 3% a.m.        | R$ 380 |  R$ 0  | 🟡     | ↘     |

---

## 3. 📈 Série Histórica 2026

| KPI                    | jan   | fev   | mar   | abr   | mai   | jun   | jul   | ago*  | Meta   |
| ---------------------- | ----: | ----: | ----: | ----: | ----: | ----: | ----: | ----: | -----: |
| Receita (R$)           | 1.500 | 1.700 | 1.800 | 1.500 | 1.800 | 2.200 | 1.000 | 2.000 | 2.000  |
| Despesa (R$)           |   900 |   880 | 1.000 |   920 |   950 | 1.100 | 1.050 | 1.000 | 1.000  |
| Resultado (R$)         |  +600 |  +820 |  +800 |  +580 |  +850 |+1.100 |   -50 |+1.000 | +1.000 |
| Taxa de poupança       |  40%  |  48%  |  44%  |  39%  |  47%  |  50%  |  -5%  |  50%  |  ≥30%  |
| Cobertura (x)          | 1,67  | 1,93  | 1,80  | 1,63  | 1,89  | 2,00  | 0,95  | 2,00  |  ≥1,5  |
| Runway (meses)         |  1,2  |  1,8  |  2,3  |  2,6  |  2,9  |  3,3  |  3,1  |  3,2  |   ≥6   |
| Aderência (envelopes)  |  6/7  |  7/7  |  6/7  |  7/7  |  6/7  |  5/7  |  4/7  |  5/7  |   7/7  |
| Score de saúde         |   52  |   55  |   58  |   56  |   61  |   66  |   44  |   48  |  ≥60   |

`*` projetado

```
Taxa de Poupança 2026 (meta 30% = ─────)
50% ┤                    ██              ██
40% ┤ ██  ██  ██  ██  ██  ██              ██
30% ┼─██──██──██──██──██──██──────────────██──
20% ┤ ██  ██  ██  ██  ██  ██              ██
10% ┤ ██  ██  ██  ██  ██  ██              ██
 0% ┼─██──██──██──██──██──██───▁▁─────────██──
-10%┤                          ██
    └ jan fev mar abr mai jun jul ago
```

---

## 4. 🔍 Diagnóstico do Mês

| Tipo         | Observação                                                                    |
| ------------ | ----------------------------------------------------------------------------- |
| 🔴 Problema  | **Julho quebrou a série.** Receita caiu 55% e a despesa não acompanhou.       |
| 🔴 Problema  | Concentração de cliente subiu de 60% → 100%: risco existencial da receita.    |
| 🟡 Risco     | Runway travado em ~3,2 meses há 3 meses — os aportes não estão acontecendo.   |
| 🟢 Positivo  | Custo fixo controlado em 29% da receita, bem abaixo do teto de 50%.           |
| 🟢 Positivo  | Ticket médio caindo — compras grandes por impulso desapareceram.              |
| 💡 Alavanca  | Cortar `Vicios` pela metade sobe a taxa de poupança em ~2,3 p.p./mês.         |
| 💡 Alavanca  | 2º cliente de R$ 1.000/mês derruba volatilidade de ±38% para ~±18%.           |

---

## 5. 🎚️ Metas de KPI para dez/2026

| KPI                | Hoje  | Meta dez | Gap      | Como fechar                            |
| ------------------ | ----: | -------: | -------: | -------------------------------------- |
| Score de saúde     |    48 |       65 | +17 pts  | Runway + diversificação                |
| Runway             | 3,2 m |    6,0 m | +2,8 m   | R$ 700/mês de aporte por 4 meses       |
| Taxa de poupança   |  -40% |     +35% | +75 p.p. | Fechar ago no azul e manter            |
| Concentração       |  100% |      60% | -40 p.p. | Fechar 2 clientes até out              |
| Aderência          |   5/7 |      7/7 | +2       | Congelar `Vicios`, realocar `Estudo`   |
| Dívida cara        | R$ 380|     R$ 0 | -R$ 380  | Quitar fatura em set                   |

---

## 🔗 Relacionados
[`Month-Report.md`](Month-Report.md) · [`Year-Report.md`](Year-Report.md) · [`Net-Worth.md`](Net-Worth.md) · [`Business-MEI.md`](Business-MEI.md)
