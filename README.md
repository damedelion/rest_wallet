# rest_wallet

Приложение, которое по REST принимает следующие запросы:

- создание кошелька
```
GET api/v1/wallet/create
```
в теле ответа возвращает uuid созданного кошелька

- изменение баланса кошелька
```
POST api/v1/wallet
{
wallet_id: UUID,
operation_type: DEPOSIT or WITHDRAW,
amount: 1000
}
```
выполняет логику по изменению счета в базе данных

- получение баланса кошелька
```
GET api/v1/wallets/{WALLET_UUID}
```
баланс выводится в теле ответа

## Запуск
```
docker compose up -d postgres
docker compose build
docker compose up wallet_app
```
