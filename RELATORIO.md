# Relatório de avaliação — wraith v0.5.0

Teste exploratório executado em 18-ago-2026 sobre a vault `/Users/mateusmoutinho/Documents/teste`,
mais 5 vaults descartáveis criadas para isolar casos específicos.

**Cobertura:** 10/10 tasks e 4/4 visualizations exercitadas · ~94 transações · 6 contas ·
2 cartões · 10 categorias (com hierarquia de 3 níveis) · 5 recorrências ·
parcelamentos, transferências, `payment_date`, `tick`/`run`/`render`/`watch`,
`Task.yaml` e `Visualization.yaml` malformados e overflow numérico.
A execução concorrente foi deixada de fora: o sistema não se propõe a ser concorrente hoje.

**Veredito curto:** a camada de validação é a melhor parte do produto — mensagens de erro
excelentes e integridade referencial rigorosa. Os problemas sérios estão **abaixo** dela:
na persistência (sem validação de faixa numérica) e no fato de o cartão de crédito ser
apenas meio implementado.

---

## 1. Bugs críticos

### 1.1 Overflow silencioso de int64 envenena a vault inteira 🔴

Valores monetários são guardados em centavos como int64 (bom), mas **não há validação de
faixa nem na entrada nem no parse**, e o parse aparentemente passa por float64:

```
wraith run AddAccount --account OverflowAcct --opening 99999999999999999999
→ AddAccount done
$ cat data/account-N-values-opening
-9223372036854775808          # int64 mínimo
```

Consequência: o total da vault virou `-R$ 92,232,720,368,516,639.08` e **todo** o dashboard,
forecast e páginas mensais ficaram errados. Nenhum aviso. Aceita também `1e20`,
`9223372036854775808`, `-9223372036854775809`.

Um zero a mais digitado por engano destrói o vault inteiro em silêncio.

**Sugestão:** rejeitar na validação da task qualquer valor fora de uma faixa sã
(ex.: ±10¹²) com a mesma qualidade de mensagem dos outros erros.

### 1.2 O formatador de moeda produz texto quebrado 🔴

Independente do overflow, o formatador falha em magnitudes altas:

```
| OverflowAcct | -R$ -92,233,720,368,547,758.0-8 | ...
```

Repare em três defeitos numa string só: sinal duplicado (`-R$ -`), um `-` **dentro** da parte
decimal (`.0-8`) e a casa decimal deslocada. Reproduzido também com
`--opening 92233720368547758` (valor **positivo**, que renderizou negativo).

### 1.3 Quebra de linha na descrição destrói a tabela markdown 🟠

`|` é corretamente bloqueado (`description may not contain |` — ótimo), mas `\n` não:

```
| 18-aug-2026 | 86 | Wallet | Food | line1
line2 | -R$ 4 | on the date |
```

A linha da tabela racha em duas e o `Statement.md` fica corrompido a partir dali.
Markdown em geral também passa cru — `**bold**`, `[link](http://x)` e `<script>` são
renderizados como markup na página.

**Sugestão:** aplicar ao `\n` (e a `\r`) a mesma guarda já existente para `|`, e escapar
markdown nas células.

---

## 2. Lacunas funcionais

### 2.1 Não existe `RemoveTransaction` — e isso trava as outras remoções 🔴

`wraith tasks` lista 10 tasks; nenhuma apaga uma transação. Mas:

```
error: RemoveAccount: Bank still holds transactions — remove them before removing the account
error: RemoveCategory: Food still classifies transactions — move them to another category first
```

A mensagem manda fazer algo que a ferramenta não oferece. `RemoveAccount` e `RemoveCategory`
tornam-se **permanentemente impossíveis** para qualquer conta ou categoria que já tenha sido
usada uma vez — ou seja, para todas as que importam. Um lançamento digitado errado fica no
ledger para sempre; o máximo que dá para fazer é `ModifyTransaction` para mascará-lo.

Esse é o buraco mais visível no dia a dia. A categoria ao menos tem saída ("move them to
another category"); a conta não tem nenhuma.

### 2.2 `closing_day` e `due_day` são coletados e nunca usados 🟠

`AddCreditCard` **exige** os dois campos, valida a faixa 1–31, mostra-os em
`Credit-Cards.md`... e é só. Testei um ciclo completo:

| Compra | Data | Fatura real (fecha dia 28) |
| --- | --- | --- |
| Jul purchase | 20-jul | fatura de julho, vence 05-ago |
| After close | 29-jul | fatura de agosto, vence 05-set |
| Aug purchase | 10-ago | fatura de agosto, vence 05-set |

Nenhuma página do dashboard separa essas compras por fatura. Não existe "fatura atual",
"fatura fechada", "valor a vencer em 05-set", nem agrupamento no extrato do cartão. O usuário
tem que somar na mão — que é exatamente o trabalho pelo qual instalou a ferramenta.

Relacionado: no `Forecast.md`, a coluna **Card bills** joga **toda** a dívida acumulada dos
cartões no primeiro mês projetado, ignorando os ciclos:

```
| sep-2026 | ... | Card bills R$ 3,316.55 |   ← Nubank (fecha 28) e Inter (fecha 10) somados
```

Também assume implicitamente que a fatura é sempre paga integralmente, embora o ledger nunca
registre esse pagamento — premissa que não está documentada em lugar nenhum.

### 2.3 `Available` do cartão ignora parcelas futuras 🟡

Compra de R$ 3.600 em 6x no Nubank (limite R$ 8.000): a página mostra
`Outstanding R$ 1,060.65 · Available R$ 6,939.35`. Um cartão real já teria reservado os
R$ 3.600 inteiros. O limite disponível exibido é otimista em R$ 3.000.

### 2.4 Parcelamento não é uma entidade 🟠

As parcelas viram N transações independentes; o "(2/6)" é texto colado na descrição.
Consequências verificadas:

- `ModifyTransaction --id 8 --amount -1000` alterou **uma** parcela para R$ 1.000 enquanto as
  outras cinco continuaram R$ 600, e o rótulo seguiu dizendo "(2/6)". Nenhum aviso.
- `ModifyTransaction --description` numa parcela apagaria o marcador "(n/6)".
- Não há como cancelar ou renegociar a série inteira — só editar 6 linhas na mão (e, sem
  `RemoveTransaction`, nem apagar).

### 2.5 Hierarquia de categorias não agrega 🟡

`Groceries` e `Restaurants` foram criadas com `parent: Food`. No resumo mensal:

```
| Food        | -R$ 25     |   ← só o que foi lançado direto em Food
| Groceries   | -R$ 340.75 |
| Restaurants | -R$ 89.90  |
```

`Food` não soma os filhos. Aninhamento de 3 níveis é aceito (`GC → C → P`), mas a coluna
`Parent` mostra só o pai imediato e não existe nenhuma visão em árvore. Ou seja: o campo
`parent` hoje é decorativo — custa uma decisão ao usuário e não entrega nada.

---

## 3. Corretude de números

### 3.1 Percentuais somam 101%

```
| Bank    | R$ 5,200.50  | 30% |
| Savings | R$ 12,000    | 69% |
| Wallet  | R$ 300       |  2% |     total: 101%
```

Arredondamento independente por linha. Sugestão: método de maior resto.

### 3.2 Barras e percentuais quebram com saldo negativo

Conta com saldo -R$ 900 exibiu `-2%` e uma barra vazia. "Share of the money you hold" não tem
significado para um valor negativo — e ele ainda entra no denominador do total.

### 3.3 "Pending settlement" mede outra coisa

O rótulo diz "movements with a payment date still ahead". Sem nenhum `payment_date`
cadastrado o indicador já marcava **-R$ 3.400**, que são as parcelas futuras. Depois de
cadastrar um `payment_date` real (-R$ 450) foi para -R$ 3.850, ou seja: mistura
"parcelas a vencer" com "movimentos com liquidação futura" sob um rótulo que só descreve o
segundo. Não há nenhuma página que liste *quais* são esses movimentos.

### 3.4 Arredondamento de sub-centavos é silencioso e inconsistente

| Entrada | Guardado | Direção |
| --- | ---: | --- |
| `1.005` | 100 | ↓ |
| `2.675` | 268 | ↑ |
| `0.145` | 14 | ↓ |

Artefato de float64. Nenhum aviso de que a precisão foi descartada.
Além disso `--amount -0.001` retorna `amount may not be zero`, mensagem que culpa o zero
quando o problema é a precisão abaixo do centavo.

### 3.5 `payment_date` anterior a `date` é aceito

```
wraith run AddTransaction ... --date 2026-08-18 --payment_date 2026-07-01
→ AddTransaction done
```

Dinheiro liquidado antes de a transação existir. Provavelmente deveria ser recusado, ou ao
menos sinalizado.

---

## 4. UX da CLI e dos arquivos

### 4.1 O `tick` destrói comentários do `Task.yaml` 🟠

Este é o atrito mais chato de todos, porque atinge o arquivo que o produto coloca no centro
("a second brain you drive with two files"). Entrada:

```yaml
# my precious comment
name: AddAccount
account: TickTest      # inline comment
opening: 42
apply: true
```

Depois de um `tick`:

```yaml
name: AddAccount
account: TickTest
opening: 42
apply: false
```

Comentários apagados e campos reordenados (num teste anterior `amount` pulou para cima de
`category`). O próprio `Task.yaml` que o `wraith start` gera vem com um bloco de comentários
de ajuda — que some no primeiro tick, exatamente quando o usuário ainda precisa dele.

**Sugestão:** reescrever apenas a linha `apply:` in-place, preservando o resto do arquivo byte
a byte.

### 4.2 `tick` diz "tick done" para quatro situações diferentes 🟠

| Situação | Saída |
| --- | --- |
| Task executada com sucesso | `tick done` |
| `apply: false` (nada a fazer) | `tick done` |
| `Task.yaml` inexistente | `tick done` |
| `Visualization.yaml` vazio (nada renderizado) | `tick done` |

Com `wraith watch --time 1s` rodando, é impossível saber pela saída se a task rodou. Note o
contraste com `wraith run`, que imprime `AddAccount done` — informativo. O `tick` deveria
dizer o quê rodou, ou `nothing to apply`.

### 4.3 `render` mascara erros do `Visualization.yaml` 🟠

Todos os erros de configuração abaixo são reportados corretamente pelo `tick`:

```
error: Visualization.yaml: DashBoard writes outside the vault: ../ESCAPE
error: Visualization.yaml: DashBoard carries no dest
error: Visualization.yaml: unknown visualization: Nonexistent
error: DashBoard: unknown field: bogus-arg — accepted fields are prev-months, future-months
```

(Excelente, inclusive o bloqueio de `dest` para fora da vault e para caminho absoluto.)

Mas com o **mesmo** arquivo, o `render` diz:

```
DashBoard is not declared in Visualization.yaml — give it a --dest
```

...seguido das 30 linhas da tela de ajuda. A entrada **está** declarada; o problema é outro. A
mensagem manda o usuário para o caminho errado e esconde a causa real.

### 4.4 Incoerência no fallback de `dest`

- `Visualization.yaml` **ausente** + `wraith render DashBoard` sem `--dest` → renderiza em
  `DashBoard/`, sem reclamar.
- `Visualization.yaml` **presente mas sem a entrada** + o mesmo comando → recusa pedindo
  `--dest`.

A documentação (`Help/Visualization.md` §4) descreve só o segundo comportamento.

### 4.5 Argumentos inválidos ora erram, ora são ignorados

```
wraith render DashBoard --future-months abc  → error: must be a number   ✅
wraith render DashBoard --future-months 0    → silenciosamente vira 8
wraith render DashBoard --future-months -5   → silenciosamente vira 8
wraith render DashBoard --future-months 500  → gera 500 linhas, sem teto
```

### 4.6 Tela de ajuda inteira em todo erro de nome

`unknown task: NotATask`, `unknown visualization: Nope` e `watch` sem `--time` despejam as 30
linhas de usage. Uma linha de erro + "run `wraith tasks`" seria mais legível, sobretudo em
script.

### 4.7 `enabled: false` é ignorado pelo `render`

`Help/Visualization.md` diz que `false` "silencia a entrada". O `tick` respeita; o
`render DashBoard` renderiza assim mesmo. Pode ser intencional (pedido explícito vence), mas
não está documentado.

### 4.8 `Error.md` fica para trás após o sucesso

É explicitamente documentado no próprio arquivo, mas na prática um `Error.md` de dias atrás
fica na raiz da vault dando a impressão de que algo está quebrado agora. Sugestão: apagá-lo no
tick bem-sucedido, ou renomeá-lo para `Error.last.md`.

### 4.9 Duplicidade de `Task-List` e `Help/Task.md`

`Tasks/README.md` e `Help/Task.md` mantêm a **mesma** tabela de 10 tasks, geradas por
visualizations diferentes. Duas fontes para o mesmo conteúdo.

### 4.10 `ModifyTransaction --id N` sem nenhum campo retorna "done"

Um no-op que consome o tick e re-renderiza tudo. Deveria avisar que nada foi alterado.

### 4.11 Mensagem confusa no conflito conta × cartão

```
wraith run AddCreditCard --account Bank ...
→ error: AddCreditCard: credit card Bank already exists
```

`Bank` é uma **conta**, não um cartão. O namespace é compartilhado (correto), mas a mensagem
afirma algo falso. Compare com as mensagens do caminho inverso, que são exemplares:

```
error: RemoveAccount: Nubank is a credit card — remove it with RemoveCreditCard
error: RemoveCreditCard: Bank is an account, not a credit card — remove it with RemoveAccount
```

---

## 5. Armazenamento

Um por campo, um arquivo por valor:

```
data/transaction-4-values-amount  → -34075
data/transaction-4-values-detail  → 000004|Nubank|Groceries|Supermarket run
data/account-4-values-limit       → 800000
```

**Amplificação de disco de ~650×.** Medido com ~90 transações:

| Métrica | Valor |
| --- | ---: |
| Conteúdo real | 5.655 bytes |
| Arquivos | 918 |
| Ocupação em disco | 3,6 MB |

São ~9 arquivos por transação. Cinco anos de uso doméstico (≈6.000 lançamentos) dão
~54.000 arquivos e ~220 MB para guardar menos de 1 MB de dados. Backup, `rsync`, Dropbox,
Time Machine e Spotlight sofrem bem antes disso. A performance de escrita, ao menos, está boa
(~50 ms por task com 90 transações).

**Riscos de modelagem:**

- **Conta e categoria são gravadas por nome de exibição**, dentro de um campo `detail`
  delimitado por `|`, e não por id. Não existe task de rename — e se existisse, ela teria de
  reescrever todas as transações. O `|` está protegido na entrada (bom), mas o esquema é
  frágil por construção.
- O parcelamento não tem chave de agrupamento (ver 2.4): "(1/6)" está dentro da string de
  descrição.

**Nota positiva:** valores são inteiros em centavos no disco, não floats. A escolha certa —
só o *parse* da entrada é que passa por float (ver 1.1 e 3.4).

---

## 6. O que está muito bom

Vale registrar, porque é a parte que sustenta o produto:

- **Mensagens de erro.** Entre as melhores que já vi numa CLI. Dizem o que houve, com qual
  valor e o que fazer:
  ```
  a negative amount needs a category with expenses: true — Salary does not accept expenses
  installments and payment_date cannot be combined — each part settles on its own date
  to_account only belongs on a transfer category — one with revenues: false and expenses: false
  P is still the parent of C — remove the child first
  RecAcct is still named by the recurrence RefTest — remove it first
  ```
- **Integridade referencial completa.** Todas as remoções verificam transações, filhos de
  categoria *e* recorrências. Não consegui deixar um ponteiro órfão em lugar nenhum.
- **Nomes case-insensitive e com trim** (`"Bank "` → conflito com `Bank`; `"  Padded  "` →
  `Padded`) — comportamento certo, mas não documentado.
- **Unicode e emoji** funcionam sem problema em nomes de conta.
- **`dest` fora da vault é bloqueado**, tanto relativo (`../`) quanto absoluto.
- **`|` bloqueado na entrada** — defesa consciente do formato de armazenamento.
- **Validação de datas de verdade**: `2026-02-30` é recusado, não só o formato.
- **Documentação bem escrita**, com voz consistente e navegação cruzada entre as páginas.
- A separação task/visualization é conceitualmente limpa e o `--database` de fato permite
  múltiplas vaults com um binário só.

---

## 7. Prioridades sugeridas

| # | Item | Seção | Por quê |
| --- | --- | --- | --- |
| 1 | Faixa numérica válida + parse decimal exato | 1.1, 3.4 | Um zero a mais corrompe a vault inteira sem aviso |
| 2 | `RemoveTransaction` | 2.1 | Destrava `RemoveAccount`/`RemoveCategory`, hoje inalcançáveis |
| 3 | Corrigir o formatador de moeda | 1.2 | Produz texto inválido |
| 4 | Usar `closing_day`/`due_day`: página de fatura | 2.2 | Campos obrigatórios que não entregam nada |
| 5 | Preservar comentários do `Task.yaml` | 4.1 | Atrito diário no arquivo central do produto |
| 6 | Bloquear `\n` na descrição | 1.3 | Corrompe a renderização |
| 7 | `tick` distinguir aplicado / nada a fazer | 4.2 | Torna o `watch` legível |
| 8 | `render` propagar os erros de config | 4.3 | Diagnóstico enganoso |
| 9 | Agregar categorias filhas no pai | 2.5 | `parent` hoje é decorativo |

---

## Apêndice — estado da vault de teste

A vault em `/Users/mateusmoutinho/Documents/teste` ficou com os dados da simulação
(≈94 transações, incluindo lixo proposital: descrições com quebra de linha, datas em 1900 e
2099, valores extremos). Removi a conta `OverflowAcct`, que envenenava todos os totais, e as
demais contas de teste que não tinham lançamentos. As contas `casetest` e `Bank` ainda
carregam transações de teste e **não podem ser removidas** — pelo motivo descrito na seção 2.1.
Para um vault limpo, apagar o diretório `data/` e rodar `wraith start` de novo.
