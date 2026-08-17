# 💧 Fluxo de Caixa

> **Pergunta que este arquivo responde:** *vou ficar sem dinheiro? Quando?*
> **Caixa hoje (16-ago-2026):** R$ 1.100 · **Burn rate:** R$ 43,75/dia · **Fôlego:** 25 dias

---

## 1. 🚨 Semáforo de Caixa

| Métrica                        | Valor      | Limite      | Status |
| ------------------------------ | ---------: | ----------: | ------ |
| Caixa disponível hoje          | R$ 1.100   | ≥ R$ 1.000  | 🟡     |
| **Menor saldo previsto (mês)** | **-R$ 380**| ≥ R$ 300    | 🔴     |
| Data do vale                   | 20-ago     | —           | 🔴     |
| Caixa previsto em 31-ago       | R$ 1.600   | ≥ R$ 1.000  | 🟢     |
| Recebíveis em aberto           | R$ 1.500   | —           | 🟡     |
| Contas a pagar até 31-ago      | R$ 300     | —           | 🟢     |

> 🔴 **Alerta de liquidez:** entre 20-ago e 25-ago o caixa fica negativo em R$ 380.
> **Saídas:** antecipar a NF do Cliente A, adiar a assinatura de cloud (R$ 60) ou usar R$ 400 da reserva (devolver em 26-ago).

---

## 2. 📆 Projeção Diária — 2ª Quinzena de Agosto

| Data        | Descrição                    | Tag       | Tipo    |    Valor | Saldo       | Nível |
| ----------- | ---------------------------- | --------- | ------- | -------: | ----------: | ----- |
| 16-ago-2026 | *Saldo de abertura*          | —         | —       |        — |  R$ 1.100   | 🟡    |
| 17-ago-2026 | Cigarros                     | `Vicios`  | Despesa |   -R$ 10 |  R$ 1.090   | 🟡    |
| 18-ago-2026 | Mercado (quinzenal)          | `Comida`  | Despesa |  -R$ 120 |    R$ 970   | 🟡    |
| 19-ago-2026 | Combustível                  | `Transporte`| Despesa|   -R$ 40 |    R$ 930   | 🟡    |
| 20-ago-2026 | Cloud/VPS + domínios         | `Empresa` | Despesa |   -R$ 85 |    R$ 845   | 🟡    |
| 20-ago-2026 | DAS-MEI                      | `Empresa` | Despesa |   -R$ 76 |    R$ 769   | 🟡    |
| 21-ago-2026 | Fatura do cartão             | `Divida`  | Despesa |  -R$ 380 |    R$ 389   | 🔴    |
| 22-ago-2026 | Almoços da semana            | `Comida`  | Despesa |   -R$ 60 |    R$ 329   | 🔴    |
| 23-ago-2026 | Streaming                    | `Poker`   | Despesa |   -R$ 30 |    R$ 299   | 🔴    |
| 24-ago-2026 | Cigarros                     | `Vicios`  | Despesa |   -R$ 10 |    R$ 289   | 🔴    |
| 25-ago-2026 | **NF Cliente A — projeto**   | `Freela`  | Receita |+R$ 1.500 |  R$ 1.789   | 🟢    |
| 27-ago-2026 | Mercado                      | `Comida`  | Despesa |   -R$ 90 |  R$ 1.699   | 🟢    |
| 28-ago-2026 | Combustível                  | `Transporte`| Despesa|   -R$ 30 |  R$ 1.669   | 🟢    |
| 30-ago-2026 | **Aporte reserva**           | `Reserva` | Transf. |  -R$ 500 |  R$ 1.169   | 🟢    |
| 30-ago-2026 | **Aporte investimento**      | `Invest`  | Transf. |  -R$ 200 |    R$ 969   | 🟡    |
| 31-ago-2026 | Diversos / folga             | —         | Despesa |   -R$ 50 |    R$ 919   | 🟡    |
|             | **Saldo de fechamento**      |           |         |          | **R$ 919**  | 🟡    |

```
Curva de caixa — 2ª quinzena
R$1.800 ┤                                    ▇▇▇▇
R$1.400 ┤                                    ████▇▇▇
R$1.000 ┼▇▇▇▇▇▇▇                             ███████▇▇▇
R$  600 ┤███████▇▇                           ██████████
R$  200 ┤██████████▇▇▇▇▇                     ██████████
R$    0 ┼───────────────▁▁▁▁▁────────────────██████████
        └ 16  18  20  21  22  24  25  27  29  31
                       ▲ vale: R$ 289 (dia 24)
```

> ℹ️ O "-R$ 380" do semáforo considera o cenário **sem** o adiamento da fatura; na projeção acima a fatura entra dia 21 e o piso real fica em **R$ 289** em 24-ago — margem de 6 dias sem nenhuma folga.

---

## 3. 📥 Recebíveis em Aberto

| Cliente / Origem | Descrição            | Emissão     | Vencimento  |    Valor | Prob. | Status      |
| ---------------- | -------------------- | ----------- | ----------- | -------: | ----: | ----------- |
| Cliente A        | Projeto — fase 2     | 20-ago-2026 | 25-ago-2026 | R$ 1.500 |   95% | 🟡 A emitir |
| Cliente B        | Proposta enviada     | —           | set/2026    | R$ 1.200 |   40% | ⚪ Pipeline |
| Cliente C        | Prospecção           | —           | out/2026    |   R$ 800 |   20% | ⚪ Pipeline |
| **Total**        | —                    | —           | —           | **R$ 3.500** | — | —      |
| **Ponderado**    | —                    | —           | —           | **R$ 2.065** | — | —      |

⚠️ **Risco único:** 95% do caixa de setembro depende de um único recebível. Se o Cliente A atrasar 10 dias, o mês fecha negativo.

---

## 4. 📤 Contas a Pagar

| Vencimento  | Descrição            | Tag       |  Valor | Auto | Status     |
| ----------- | -------------------- | --------- | -----: | ---- | ---------- |
| 20-ago-2026 | Cloud / VPS          | `Empresa` |  R$ 60 | ✅   | 🟡 Agendado|
| 20-ago-2026 | DAS-MEI              | `Empresa` |  R$ 76 | ✅   | 🟡 Agendado|
| 21-ago-2026 | Fatura cartão        | `Divida`  | R$ 380 | ❌   | 🔴 Manual  |
| 23-ago-2026 | Streaming            | `Poker`   |  R$ 30 | ✅   | 🟡 Agendado|
| 05-set-2026 | Aluguel              | `Casa`    | R$ 300 | ✅   | ⚪ Futuro  |
| 10-set-2026 | Internet             | `Casa`    |  R$ 90 | ✅   | ⚪ Futuro  |
| **Total até 31-ago** | —           | —         | **R$ 546** | — | —      |

---

## 5. 🗓️ Fluxo de Caixa Rolling 12 Meses

| Mês       | Entradas  | Saídas    | Líquido    | Caixa final | Status |
| --------- | --------: | --------: | ---------: | ----------: | ------ |
| set-2025  |  R$ 1.400 |    R$ 950 |   +R$ 450  |    R$ 1.900 | 🟢     |
| out-2025  |  R$ 1.600 |    R$ 980 |   +R$ 620  |    R$ 2.520 | 🟢     |
| nov-2025  |  R$ 1.300 |  R$ 1.050 |   +R$ 250  |    R$ 2.770 | 🟢     |
| dez-2025  |  R$ 2.100 |  R$ 1.400 |   +R$ 700  |    R$ 3.470 | 🟢     |
| jan-2026  |  R$ 1.500 |    R$ 900 |   +R$ 600  |    R$ 2.070 | 🟢     |
| fev-2026  |  R$ 1.700 |    R$ 880 |   +R$ 820  |    R$ 1.890 | 🟢     |
| mar-2026  |  R$ 1.800 |  R$ 1.000 |   +R$ 800  |    R$ 1.690 | 🟢     |
| abr-2026  |  R$ 1.500 |    R$ 920 |   +R$ 580  |    R$ 1.470 | 🟡     |
| mai-2026  |  R$ 1.800 |    R$ 950 |   +R$ 850  |    R$ 1.620 | 🟢     |
| jun-2026  |  R$ 2.200 |  R$ 1.100 | +R$ 1.100  |    R$ 1.920 | 🟢     |
| jul-2026  |  R$ 1.000 |  R$ 1.050 |    -R$ 50  |    R$ 1.170 | 🔴     |
| ago-2026* |  R$ 2.000 |  R$ 1.000 | +R$ 1.000  |    R$ 919   | 🟡     |

> Caixa cai mesmo com resultado positivo porque os aportes em `Reserva`/`Invest` saem do caixa — isso é **saudável**, o dinheiro migrou para patrimônio (ver [`Net-Worth.md`](Net-Worth.md)).

```
Líquido mensal
set ███▌      out ████▉      nov ██        dez █████▌
jan ████▊     fev ██████▌    mar ██████▍   abr ████▋
mai ██████▊   jun ████████▊  jul ▍🔴       ago ████████
```

---

## 6. 🧯 Plano de Contingência

| Cenário                          | Impacto no caixa | Plano                                                   |
| -------------------------------- | ---------------: | ------------------------------------------------------- |
| Cliente A atrasa 10 dias         |        -R$ 1.500 | Sacar R$ 800 da reserva, devolver em 5 dias             |
| Cliente A cancela                |        -R$ 1.500 | Congelar `Vicios`+`Poker`+`Estudo`, ativar pipeline B/C |
| Despesa emergencial R$ 1.000     |        -R$ 1.000 | Reserva cobre — é exatamente para isso                  |
| Queda de receita 3 meses         |        -R$ 3.000 | Reserva cobre 3,2 meses; após isso, corte de fixos      |
| Equipamento quebra               |        -R$ 4.000 | Budget "troca de notebook" cobre R$ 900; resto parcelado|

**Ordem de saque em emergência:** ① Caixa → ② Reserva → ③ Investimentos líquidos → ④ Crédito (último recurso).

---

## 🔗 Relacionados
[`Month-Report.md`](Month-Report.md) · [`Budget.md`](Budget.md) · [`Net-Worth.md`](Net-Worth.md) · [`Business-MEI.md`](Business-MEI.md)
