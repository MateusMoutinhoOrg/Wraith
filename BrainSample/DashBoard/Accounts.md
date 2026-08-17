# Accounts

> **Updated:** 16-aug-2026 · 4 accounts · Net available: **R$ 900**

Registry of every account. Each active account has an isolated monthly statement under
[`Month/Accounts/`](Month/Accounts/).

---

## 1. Account registry

| ID     | Account          | Type        | Opening (01-aug) | Current (16-aug) | Projected (31-aug) | Status | Statement |
| ------ | ---------------- | ----------- | ---------------: | ---------------: | -----------------: | :----: | --------- |
| `BANK` | Bank account     | Checking    |         R$ 2,000 |         R$ 1,995 |          R$ 3,895* |   🟢   | [`Month/Accounts/Bank.md`](Month/Accounts/Bank.md) |
| `CASH` | Cash             | Cash        |           R$ 600 |           R$ 255 |             R$ 50* |   🔴   | [`Month/Accounts/Cash.md`](Month/Accounts/Cash.md) |
| `CARD` | Credit card      | Credit card |          -R$ 800 |        -R$ 1,350 |         -R$ 1,650* |   🟡   | [`Month/Accounts/Credit-Card.md`](Month/Accounts/Credit-Card.md) |
| `SAVE` | Emergency savings| Savings     |         R$ 2,000 |         R$ 2,000 |          R$ 3,500* |   🟢   | — (movements in [`Net-Worth.md`](Net-Worth.md)) |

| Totals                          |        Value |
| ------------------------------- | -----------: |
| Available (`BANK` + `CASH`)     | **R$ 2,250** |
| Card outstanding (`CARD`)       |    -R$ 1,350 |
| **Net available**               |   **R$ 900** |
| Reserves (`SAVE`)               |     R$ 2,000 |

```
Where the money sits
Savings  ████████████████████  R$ 2,000   47%
Bank     ███████████████████▉  R$ 1,995   47%
Cash     ██▌                   R$   255    6%
Card     █████████████▌ (owed) -R$ 1,350  due 05-sep
```

---

## 2. Account rules

| Account | Purpose                                            |
| ------- | -------------------------------------------------- |
| `BANK`  | Income, bills, transfers — main operating account  |
| `CASH`  | Discretionary spending only; not refilled mid-month |
| `CARD`  | Subscriptions and planned purchases; paid in full on day 5 |
| `SAVE`  | Emergency fund — receives R$ 1,500 on day 30; withdrawals only for emergencies |

---

## 3. Credit card cycle

| Item                | Value      |
| ------------------- | ---------- |
| Statement closes    | day 28     |
| Bill due            | day 5      |
| Current bill (aug)  | R$ 1,350 owed — projected R$ 1,650*, due 05-sep |
| Payment policy      | Always in full — revolving interest is 14%/month |

> *Fictional document. Balances derive from [`Month/Statement.md`](Month/Statement.md).*
