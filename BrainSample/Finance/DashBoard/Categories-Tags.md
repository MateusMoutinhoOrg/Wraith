# 🏷️ Categorias e Tags

> Plano de contas da dashboard. Toda transação em `Month-Balance.md` **precisa** de exatamente 1 tag.

---

## 1. 🌳 Plano de Contas

```
FINANÇAS
├── 📥 RECEITAS
│   ├── Freela ............ serviços PJ/MEI faturados
│   ├── Salario ........... vínculo CLT (inativo)
│   ├── Rendimento ........ juros, dividendos, cashback
│   └── Extra ............. venda de itens, reembolso
│
├── 📤 DESPESAS
│   ├── 🔒 Essenciais
│   │   ├── Casa .......... aluguel, contas, internet
│   │   ├── Comida ........ mercado, refeições
│   │   └── Transporte .... combustível, app, manutenção
│   ├── 💼 Empresa
│   │   └── Empresa ....... cloud, domínios, DAS, ferramentas
│   ├── 🎈 Discricionárias
│   │   ├── Poker ......... lazer, apostas, streaming
│   │   ├── Vicios ........ cigarro, bebida
│   │   └── Estudo ........ cursos, livros
│   └── 💳 Dívidas
│       └── Divida ........ juros e amortização
│
└── 🔄 TRANSFERÊNCIAS (não entram no resultado)
    ├── Reserva ........... aporte na emergência
    └── Invest ............ aporte em carteira
```

---

## 2. 📇 Catálogo de Tags

| Tag          | Emoji | Grupo          | `positive` | `negative` | Teto/mês | Ativa | Criada em   |
| ------------ | ----- | -------------- | ---------- | ---------- | -------: | ----- | ----------- |
| `Freela`     | 💼    | Receita        | ✅ true    | ❌ false   |        — | ✅    | 03-jan-2026 |
| `Rendimento` | 🪙    | Receita        | ✅ true    | ❌ false   |        — | ✅    | 03-jan-2026 |
| `Extra`      | 🎁    | Receita        | ✅ true    | ❌ false   |        — | ✅    | 11-fev-2026 |
| `Casa`       | 🏠    | Essencial      | ❌ false   | ✅ true    |   R$ 300 | ✅    | 03-jan-2026 |
| `Comida`     | 🍽️    | Essencial      | ❌ false   | ✅ true    |   R$ 250 | ✅    | 03-jan-2026 |
| `Transporte` | 🚗    | Essencial      | ❌ false   | ✅ true    |   R$ 120 | ✅    | 03-jan-2026 |
| `Empresa`    | 💻    | Empresa        | ❌ false   | ✅ true    |   R$ 150 | ✅    | 05-jan-2026 |
| `Poker`      | 🃏    | Discricionária | ✅ true    | ✅ true    |    R$ 80 | ✅    | 18-jan-2026 |
| `Vicios`     | 🚬    | Discricionária | ❌ false   | ✅ true    |    R$ 50 | ✅    | 22-jan-2026 |
| `Estudo`     | 📚    | Discricionária | ❌ false   | ✅ true    |    R$ 50 | ✅    | 02-fev-2026 |
| `Divida`     | 💳    | Dívida         | ❌ false   | ✅ true    |   R$ 120 | ✅    | 14-mar-2026 |
| `Reserva`    | 🛡️    | Transferência  | ✅ true    | ✅ true    |        — | ✅    | 03-jan-2026 |
| `Invest`     | 📈    | Transferência  | ✅ true    | ✅ true    |        — | ✅    | 03-jan-2026 |
| `Salario`    | 🧾    | Receita        | ✅ true    | ❌ false   |        — | ⚪ não| 03-jan-2026 |

**Tags bidirecionais** (`positive` **e** `negative` = true): `Poker`, `Reserva`, `Invest`.
São as únicas que aceitam entrada e saída no mesmo agrupamento.

---

## 3. 📊 Desempenho por Tag (ago/2026, dia 16)

| Tag          | Mov. | Entradas | Saídas | Líquido  | % da despesa | vs. jul  |
| ------------ | ---: | -------: | -----: | -------: | -----------: | -------- |
| `Freela`     |    1 |   R$ 500 |   R$ 0 | +R$ 500  |          — | ↘ -50%   |
| `Casa`       |    1 |     R$ 0 | R$ 300 | -R$ 300  |        42,9% | → 0%     |
| `Comida`     |    7 |     R$ 0 | R$ 180 | -R$ 180  |        25,7% | ↗ +12%   |
| `Vicios`     |    9 |     R$ 0 |  R$ 90 |  -R$ 90  |        12,9% | ↗ +80% 🔴|
| `Poker`      |    3 |    R$ 30 |  R$ 75 |  -R$ 45  |        10,7% | ↘ -10%   |
| `Transporte` |    4 |     R$ 0 |  R$ 45 |  -R$ 45  |         6,4% | ↘ -20%   |
| `Empresa`    |    2 |     R$ 0 |  R$ 40 |  -R$ 40  |         5,7% | → 0%     |
| `Estudo`     |    0 |     R$ 0 |   R$ 0 |    R$ 0  |         0,0% | ⚪ ocioso|
| `Divida`     |    0 |     R$ 0 |   R$ 0 |    R$ 0  |         0,0% | ⚪       |

```
Distribuição da despesa
Casa        ██████████████████████  42,9%
Comida      █████████████░░░░░░░░░  25,7%
Vicios      ██████▌░░░░░░░░░░░░░░░  12,9%
Poker       █████▌░░░░░░░░░░░░░░░░  10,7%
Transporte  ███▎░░░░░░░░░░░░░░░░░░   6,4%
Empresa     ██▉░░░░░░░░░░░░░░░░░░░   5,7%
```

---

## 4. 🔥 Ranking de Frequência (nº de transações)

| Pos | Tag        | Mov. | Ticket médio | Leitura                                    |
| --: | ---------- | ---: | -----------: | ------------------------------------------ |
|  1º | `Vicios`   |    9 |      R$ 10,0 | 🔴 Micro-gastos diários — maior vazamento  |
|  2º | `Comida`   |    7 |      R$ 25,7 | 🟡 Compras fracionadas, sem lista          |
|  3º | `Transporte`|   4 |      R$ 11,3 | 🟢 Saudável                                |
|  4º | `Poker`    |    3 |      R$ 25,0 | 🟢 Dentro do envelope                      |
|  5º | `Empresa`  |    2 |      R$ 20,0 | 🟢 Recorrentes                             |
|  6º | `Casa`     |    1 |     R$ 300,0 | 🟢 Fixo único                              |

> 💡 **Insight:** `Vicios` custa R$ 10/dia. São **R$ 3.650/ano** no ritmo atual — mais que o aluguel anual.

---

## 5. 📐 Regras de Categorização

1. **Uma tag por transação.** Se couber em duas, vence a mais específica (`Empresa` > `Casa`).
2. **Tag duplicada é erro.** Criar tag existente retorna `Error: Tag <X> already exists`.
3. **Transferências não afetam o resultado** — `Reserva` e `Invest` movem dinheiro, não o consomem.
4. **Toda tag nova** exige teto definido antes do primeiro lançamento.
5. **Tag sem movimento por 3 meses** é desativada (`Ativa: ⚪ não`), não deletada — preserva histórico.
6. **Renomear tag é proibido.** Desative a antiga e crie a nova, para não corromper meses fechados.
7. **Reembolso** entra como valor positivo na **mesma tag** da despesa original, nunca como receita.

---

## 6. 🧾 Log de Alterações

| Data        | Operação   | Tag        | Resultado                          |
| ----------- | ---------- | ---------- | ---------------------------------- |
| 14-mar-2026 | AddTag     | `Divida`   | ✅ OK                              |
| 02-fev-2026 | AddTag     | `Estudo`   | ✅ OK                              |
| 11-fev-2026 | AddTag     | `Extra`    | ✅ OK                              |
| 09-ago-2026 | AddTag     | `Poker`    | ❌ `Error: Tag Poker already exists` |
| 12-ago-2026 | RemoveTag  | `Salario`  | ⚪ Desativada (sem uso desde jan)  |

---

## 🔗 Relacionados
[`Budget.md`](Budget.md) · [`Month-Balance.md`](Month-Balance.md) · [`Alerts-Rules.md`](Alerts-Rules.md)
