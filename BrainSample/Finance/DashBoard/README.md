# 🧭 Dashboard Financeiro — Wraith

> **Última atualização:** 16-ago-2026 · **Moeda:** BRL (R$)
> **Escopo:** Pessoa Física + MEI (Wraith Software)

---

## ⚡ Visão de 1 Tela

| Bloco                   | Número           | Status | Onde ver                            |
| ----------------------- | ---------------- | ------ | ----------------------------------- |
| Resultado do mês (real) | **-R$ 200**      | 🔴     | [Month-Report](Month-Report.md)     |
| Resultado do mês (proj) | **+R$ 1.000**    | 🟢     | [Month-Report](Month-Report.md)     |
| Orçamento usado         | 70% (R$ 700/1.000) | 🟡   | [Budget](Budget.md)                 |
| Caixa hoje              | R$ 1.100         | 🟡     | [Cash-Flow](Cash-Flow.md)           |
| Menor saldo previsto    | -R$ 380 (20-ago) | 🔴     | [Cash-Flow](Cash-Flow.md)           |
| Patrimônio líquido      | **R$ 18.820**    | 🟡     | [Net-Worth](Net-Worth.md)           |
| Dívida total            | R$ 5.180         | 🟡     | [Net-Worth](Net-Worth.md)           |
| Ano 2026 acumulado      | +R$ 5.700        | 🟢     | [Year-Report](Year-Report.md)       |
| Score de Saúde          | **48 / 100**     | 🟡     | [KPIs](KPIs.md)                     |
| Progresso FIRE          | 2,5%             | ⚪     | [Long-Term](Long-Term.md)           |
| Alertas abertos         | 3 🔴 · 3 🟡      | 🔴     | [Alerts-Rules](Alerts-Rules.md)     |

```
Ano 2026   Receita  █████████████▌░░░░░░  56%  (R$ 13.500 / R$ 24.000)
           Despesa  ██████████████▍░░░░░  64%  (R$  7.800 / R$ 12.100)
           Tempo    ████████████▍░░░░░░░  62%  (16-ago / 31-dez)
```

---

## 📁 Mapa da Dashboard

### 🔵 Curto prazo — o mês corrente
| Arquivo                                  | O que responde                                        | Frequência |
| ---------------------------------------- | ----------------------------------------------------- | ---------- |
| [`Month-Results.md`](Month-Results.md)   | Números crus do mês (realizado vs. projetado)         | Diária     |
| [`Month-Balance.md`](Month-Balance.md)   | Extrato — toda transação do mês, linha a linha        | Diária     |
| [`Month-Report.md`](Month-Report.md)     | **Relatório completo do mês** — KPIs, budget, ações   | Semanal    |
| [`Cash-Flow.md`](Cash-Flow.md)           | Vou ficar sem dinheiro? Quando?                       | Semanal    |

### 🟢 Planejamento — quanto posso gastar
| Arquivo                                    | O que responde                                      | Frequência |
| ------------------------------------------ | --------------------------------------------------- | ---------- |
| [`Budget.md`](Budget.md)                   | **Orçamento de envelopes** — tetos por categoria    | Mensal     |
| [`Categories-Tags.md`](Categories-Tags.md) | Plano de contas + catálogo de tags e regras         | Trimestral |
| [`Alerts-Rules.md`](Alerts-Rules.md)       | Regras que disparam alerta automático               | Trimestral |

### 🟡 Médio prazo — trimestre e ano
| Arquivo                                      | O que responde                                    | Frequência |
| -------------------------------------------- | ------------------------------------------------- | ---------- |
| [`Quarter-Report.md`](Quarter-Report.md)     | Fechamento trimestral e tendências                | Trimestral |
| [`Year-Report.md`](Year-Report.md)           | Consolidado 2026 mês a mês + projeção de fechamento| Mensal     |
| [`Business-MEI.md`](Business-MEI.md)         | A empresa dá lucro? Limite do MEI, DAS, clientes  | Mensal     |
| [`Goals.md`](Goals.md)                       | Metas do ano e progresso                          | Mensal     |

### 🔴 Longo prazo — patrimônio
| Arquivo                            | O que responde                                       | Frequência |
| ---------------------------------- | ---------------------------------------------------- | ---------- |
| [`Net-Worth.md`](Net-Worth.md)     | **Ativos, passivos, dívidas e carteira**             | Mensal     |
| [`Long-Term.md`](Long-Term.md)     | Plano 2026–2031, independência financeira, cenários  | Semestral  |
| [`KPIs.md`](KPIs.md)               | Dicionário de indicadores + série histórica          | Mensal     |

---

## 🔄 Rotina de Manutenção

| Quando         | O que fazer                                                                 |
| -------------- | --------------------------------------------------------------------------- |
| **Todo dia**   | Lançar transações em `Month-Balance.md`                                     |
| **Toda 2ª-f.** | Rodar `Cash-Flow.md`, checar `Alerts-Rules.md`, atualizar `Month-Report.md` |
| **Dia 1**      | Fechar o mês: `Month-Results.md` → `Year-Report.md` → `KPIs.md` → zerar extrato |
| **Dia 1 (mês)**| Revisar tetos em `Budget.md`; realocar envelopes ociosos                    |
| **Fim de tri** | `Quarter-Report.md` + revisar `Goals.md`                                    |
| **Fim de ano** | `Year-Report.md` final + recalibrar `Long-Term.md`                          |

---

## 🧾 Convenções (valem para todos os arquivos)

| Item        | Padrão                                            |
| ----------- | ------------------------------------------------- |
| Data        | `DD-mmm-AAAA` (ex.: `16-ago-2026`)                |
| Valor       | `R$ 1.234,56` · saída sempre com sinal negativo   |
| Tag         | `crase-minúscula` conforme `Categories-Tags.md`   |
| Projeção    | sufixo `*` no número                              |
| Status      | 🟢 ok · 🟡 atenção · 🔴 ação imediata · ⚪ inativo |
| Tendência   | ↗ subindo · ↘ caindo · → estável                  |

> 📌 *Dashboard fictícia, usada como modelo de gestão financeira pessoal + pequena empresa.*
