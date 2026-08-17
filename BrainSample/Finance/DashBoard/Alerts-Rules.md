# 🚨 Alertas e Regras

> **Última varredura:** 16-ago-2026 · **Abertos:** 3 🔴 · 3 🟡 · 1 ⚪ · 1 ℹ️

---

## 1. 🔔 Alertas Abertos

| Sev. | ID     | Alerta                                                  | Disparado em | Regra | Ação                                          | Prazo  |
| ---- | ------ | ------------------------------------------------------- | ------------ | ----- | --------------------------------------------- | ------ |
| 🔴   | `A-01` | `Vicios` estourou o teto em 80% (R$ 90 / R$ 50)         | 14-ago-2026  | R-03  | Congelar tag até 01-set                       | 18-ago |
| 🔴   | `A-02` | 100% da receita concentrada em 1 cliente                | 01-ago-2026  | R-11  | Enviar 3 propostas                            | 31-ago |
| 🔴   | `A-03` | Vale de caixa previsto: R$ 289 em 24-ago                | 16-ago-2026  | R-06  | Antecipar NF ou adiar assinatura              | 20-ago |
| 🟡   | `A-04` | Ritmo de gasto 18 p.p. à frente do calendário           | 12-ago-2026  | R-02  | Realocar envelopes (ver `Budget.md` §3)       | 18-ago |
| 🟡   | `A-05` | Runway travado em 3,2 meses há 3 meses                  | 01-ago-2026  | R-05  | Executar aporte de R$ 500                     | 30-ago |
| 🟡   | `A-06` | Dívida cara em aberto: R$ 380 a 14% a.m.                | 01-ago-2026  | R-08  | Quitar fatura integral                        | 21-ago |
| ⚪   | `A-07` | Tag `Estudo` sem movimento há 2 meses                   | 01-ago-2026  | R-04  | Realocar R$ 50 ou assumir compromisso         | 31-ago |
| ℹ️   | `A-08` | `AddTag Poker` retornou erro — tag já existe            | 09-ago-2026  | R-13  | Usar a tag existente; nenhuma ação necessária | —      |

```
Alertas por severidade
🔴 Crítico  ███████████████  3
🟡 Atenção  ███████████████  3
⚪ Info     █████            1
ℹ️ Sistema  █████            1
```

---

## 2. 📜 Catálogo de Regras

### 📦 Orçamento

| ID   | Regra                                                        | Gatilho                          | Sev. | Ação automática              |
| ---- | ------------------------------------------------------------ | -------------------------------- | ---- | ---------------------------- |
| R-01 | Envelope acima de 80% do teto                                | uso ≥ 80%                        | 🟡   | Avisar                       |
| R-02 | Ritmo de gasto à frente do calendário                        | %uso − %mês ≥ 10 p.p.            | 🟡   | Sugerir realocação           |
| R-03 | **Envelope estourado**                                       | uso > 100%                       | 🔴   | **Congelar a tag**           |
| R-04 | Envelope ocioso                                              | 0 movimentos por 2 meses         | ⚪   | Sugerir realocação           |
| R-14 | Teto global do mês ultrapassado                              | despesa > R$ 1.000               | 🔴   | Bloquear gastos livres       |

### 💧 Liquidez

| ID   | Regra                                                        | Gatilho                          | Sev. | Ação automática              |
| ---- | ------------------------------------------------------------ | -------------------------------- | ---- | ---------------------------- |
| R-05 | Runway abaixo da meta                                        | reserva < 6 meses de despesa     | 🟡   | Avisar mensalmente           |
| R-06 | **Vale de caixa**                                            | saldo projetado < R$ 300         | 🔴   | Plano de contingência        |
| R-07 | Caixa negativo projetado                                     | saldo projetado < R$ 0           | 🔴   | Sacar da reserva             |
| R-15 | Recebível vencido                                            | NF sem pagamento há 5 dias       | 🔴   | Cobrar cliente               |

### 💳 Dívida

| ID   | Regra                                                        | Gatilho                          | Sev. | Ação automática              |
| ---- | ------------------------------------------------------------ | -------------------------------- | ---- | ---------------------------- |
| R-08 | **Dívida cara em aberto**                                    | juros > 3% a.m. e saldo > R$ 0   | 🟡   | Prioridade sobre aportes     |
| R-09 | Comprometimento de renda alto                                | parcelas > 30% da receita        | 🔴   | Proibir nova dívida          |
| R-10 | Parcelamento com reserva incompleta                          | reserva < 6 m e nova parcela     | 🔴   | **Bloquear compra**          |

### 💼 Negócio

| ID   | Regra                                                        | Gatilho                          | Sev. | Ação automática              |
| ---- | ------------------------------------------------------------ | -------------------------------- | ---- | ---------------------------- |
| R-11 | **Concentração de cliente**                                  | maior cliente > 60% da receita   | 🔴   | Meta de prospecção           |
| R-12 | Pipeline insuficiente                                        | ponderado < 300% da meta do tri  | 🟡   | Bloco de prospecção na agenda|
| R-16 | Limite MEI em risco                                          | faturamento 12m > 80% de R$ 81k  | 🟡   | Planejar migração de regime  |
| R-17 | Mês fechado no vermelho                                      | resultado < R$ 0                 | 🔴   | Revisão de teto no mês seguinte |

### 🧾 Integridade dos Dados

| ID   | Regra                                                        | Gatilho                          | Sev. | Ação automática              |
| ---- | ------------------------------------------------------------ | -------------------------------- | ---- | ---------------------------- |
| R-13 | Tag duplicada                                                | `AddTag` de tag existente        | ℹ️   | `Error: Tag X already exists`|
| R-18 | Transação sem tag                                            | linha sem tag no extrato         | 🟡   | Bloquear fechamento do mês   |
| R-19 | Extrato desatualizado                                        | sem lançamento há 3 dias         | 🟡   | Lembrete                     |
| R-20 | Saldo do extrato ≠ saldo do banco                            | divergência > R$ 1               | 🔴   | Conciliar antes de fechar    |

---

## 3. 🧊 Congelamentos Ativos

| Tag       | Congelada em | Libera em   | Motivo                          | Regra |
| --------- | ------------ | ----------- | ------------------------------- | ----- |
| `Vicios`  | 18-ago-2026  | 01-set-2026 | Estouro de 80% sobre o teto     | R-03  |

> Tag congelada **não aceita novos lançamentos** até a data de liberação. Exceção só via realocação formal registrada em [`Budget.md`](Budget.md) §3.

---

## 4. 🧯 Matriz de Escalonamento

| Situação                                    | Nível | Quem decide  | Ação                                              |
| ------------------------------------------- | ----- | ------------ | ------------------------------------------------- |
| 1 envelope estourado                        | 🟡    | Automático   | Congelar tag, realocar de envelope ocioso         |
| 2+ envelopes estourados                     | 🔴    | Revisão      | Congelar todos os discricionários até o dia 1     |
| Caixa projetado < R$ 300                    | 🔴    | Revisão      | Adiar despesas não essenciais, antecipar recebível|
| Caixa projetado < R$ 0                      | 🔴    | Emergência   | Sacar da reserva (devolver em ≤ 30 dias)          |
| Mês fechado no vermelho                     | 🔴    | Revisão      | Cortar 15% do teto no mês seguinte                |
| 2 meses seguidos no vermelho                | 🔴    | Emergência   | Congelar todo gasto discricionário por 60 dias    |
| Perda do cliente principal                  | 🔴    | Emergência   | Modo sobrevivência: só fixos + comida + transporte|

**Modo sobrevivência** = teto mensal de R$ 700 (`Casa` R$ 300 + `Comida` R$ 250 + `Transporte` R$ 100 + `Empresa` R$ 50), tudo mais congelado. Estende o runway de 3,2 → 4,6 meses.

---

## 5. 📋 Checklist de Fechamento do Mês

Rodar todo dia 1º, na ordem:

- [ ] Conciliar `Month-Balance.md` com o extrato do banco (regra R-20)
- [ ] Verificar que toda transação tem tag (regra R-18)
- [ ] Consolidar totais em `Month-Results.md`
- [ ] Atualizar `Month-Report.md` com os números fechados
- [ ] Lançar o mês em `Year-Report.md` e `KPIs.md`
- [ ] Atualizar saldos em `Net-Worth.md`
- [ ] Revisar tetos e realocações em `Budget.md`
- [ ] Liberar tags congeladas que venceram
- [ ] Rodar todas as regras deste arquivo e registrar novos alertas
- [ ] Marcar progresso em `Goals.md`
- [ ] Arquivar o extrato e zerar `Month-Balance.md`

---

## 6. 🗂️ Histórico de Alertas Resolvidos

| ID     | Alerta                                    | Aberto      | Fechado     | Como resolveu                    |
| ------ | ----------------------------------------- | ----------- | ----------- | -------------------------------- |
| `A-00` | Cheque especial em uso                    | 05-jan-2026 | 18-fev-2026 | Quitado com resultado de fevereiro|
| `J-04` | Mês de julho fechou no vermelho           | 01-ago-2026 | —           | 🔴 Ainda em revisão              |
| `J-02` | `Comida` estourou o teto em julho         | 28-jul-2026 | 01-ago-2026 | Teto renovado no novo mês        |
| `M-07` | Transação sem tag em maio                 | 31-mai-2026 | 01-jun-2026 | Classificada como `Transporte`   |

---

## 🔗 Relacionados
[`Budget.md`](Budget.md) · [`Cash-Flow.md`](Cash-Flow.md) · [`Categories-Tags.md`](Categories-Tags.md) · [`Month-Report.md`](Month-Report.md)
