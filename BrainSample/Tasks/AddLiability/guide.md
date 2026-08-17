# AddLiability

Adds a debt to the net-worth sheet. The credit card outstanding is already counted
automatically from the `CARD` account — do not register it here.

## Fields

| Field       | Required | Type   | Description                              |
| ----------- | :------: | ------ | ---------------------------------------- |
| `name`      |    ✅    | string | Must be `AddLiability`                   |
| `debt`      |    ✅    | string | Debt name — unique, e.g. `Car loan`      |
| `balance`   |    ✅    | number | Outstanding balance in R$                |
| `interest`  |    ✅    | string | e.g. `1.2%/month`                        |
| `payment`   |    ✅    | number | Monthly payment in R$                    |
| `remaining` |    ✅    | number | Installments left                        |
| `apply`     |    ✅    | bool   | Set `true` to execute on the next tick   |

## Renders

`Net-Worth.md`, `README.md`

## Errors

- Debt already exists → `Error.md`
