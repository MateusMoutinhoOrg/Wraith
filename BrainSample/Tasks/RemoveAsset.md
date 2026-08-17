# RemoveAsset

Removes an asset from the net-worth sheet (sold, written off, or being revalued).

## Fields

| Field   | Required | Type   | Description                            |
| ------- | :------: | ------ | -------------------------------------- |
| `name`  |    ✅    | string | Must be `RemoveAsset`                  |
| `asset` |    ✅    | string | Exact asset name                       |
| `apply` |    ✅    | bool   | Set `true` to execute on the next tick |

## Renders

`Net-Worth.md`, `README.md`

## Errors

- Asset does not exist → `Error.md`

> If the asset was sold, also record the money received with `AddTransaction`.

## Sample

```yaml
name: RemoveAsset
asset: Car
apply: true
```
