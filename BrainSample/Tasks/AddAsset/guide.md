# AddAsset

Adds an asset to the net-worth sheet. Accounts (`BANK`, `CASH`, `SAVE`) are already counted
automatically — only register assets that are not accounts (car, equipment, investments).

## Fields

| Field         | Required | Type   | Description                                  |
| ------------- | :------: | ------ | -------------------------------------------- |
| `name`        |    ✅    | string | Must be `AddAsset`                           |
| `asset`       |    ✅    | string | Asset name — unique, e.g. `Car`              |
| `liquidity`   |    ✅    | string | `Immediate`, `Days` or `Low`                 |
| `value`       |    ✅    | number | Current value in R$                          |
| `appreciates` |    ✅    | bool   | `false` for depreciating goods               |
| `apply`       |    ✅    | bool   | Set `true` to execute on the next tick       |

## Renders

`Net-Worth.md`, `README.md`

## Errors

- Asset already exists → `Error.md` (re-run `AddAsset` after `RemoveAsset` to revalue, or remove first)
