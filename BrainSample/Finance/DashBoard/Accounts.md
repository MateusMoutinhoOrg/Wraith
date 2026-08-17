# Accounts

> **Updated:** 16-aug-2026 · 4 accounts · Net available: **R$ 720**

Registry of every account. Each active account has an isolated monthly statement under
[`Month/Accounts/`](Month/Accounts/).

---

## 1. Account registry

| ID     | Account          | Type        | Opening (01-aug) | Current (16-aug) | Projected (31-aug) | Status | Statement |
| ------ | ---------------- | ----------- | ---------------: | ---------------: | -----------------: | :----: | --------- |
| `BANK` | Bank account     | Checking    |         R$ 1,000 |           R$ 963 |          R$ 1,803* |   🟢   | [`Month/Accounts/Bank.md`](Month/Accounts/Bank.md) |
| `CASH` | Cash             | Cash        |           R$ 300 |           R$ 137 |             R$ 17* |   🔴   | [`Month/Accounts/Cash.md`](Month/Accounts/Cash.md) |
| `CARD` | Credit card      | Credit card |          -R$ 380 |          -R$ 380 |           -R$ 400* |   🟡   | [`Month/Accounts/Credit-Card.md`](Month/Accounts/Credit-Card.md) |
| `SAVE` | Emergency savings| Savings     |         R$ 3,200 |         R$ 3,200 |          R$ 3,700* |   🟢   | — (movements in [`Net-Worth.md`](Net-Worth.md)) |

| Totals                          |        Value |
| ------------------------------- | -----------: |
| Available (`BANK` + `CASH`)     | **R$ 1,100** |
| Card outstanding (`CARD`)       |      -R$ 380 |
| **Net available**               |   **R$ 720** |
| Reserves (`SAVE`)               |     R$ 3,200 |

```
Where the money sits
Savings  ████████████████████  R$ 3,200   74%
Bank     ██████                R$   963   22%
Cash     █                     R$   137    4%
Card     ██▍ (owed)           -R$   380   due 05-sep
```

---

## 2. Account rules

| Account | Purpose                                            |
| ------- | -------------------------------------------------- |
| `BANK`  | Income, bills, transfers — main operating account  |
| `CASH`  | Discretionary spending only; not refilled mid-month |
| `CARD`  | Subscriptions and planned purchases; paid in full on day 5 |
| `SAVE`  | Emergency fund — receives R$ 500 on day 30; withdrawals only for emergencies |

---

## 3. Credit card cycle

| Item                | Value      |
| ------------------- | ---------- |
| Statement closes    | day 28     |
| Bill due            | day 5      |
| Current bill (aug)  | R$ 380 — due 05-sep |
| Payment policy      | Always in full — revolving interest is 14%/month |

> *Fictional document. Balances derive from [`Month/Statement.md`](Month/Statement.md).*
