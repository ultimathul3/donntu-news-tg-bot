## Использование

Отправьте следующий GET-запрос на Telegram API:

```shell
https://api.telegram.org/bot<ACCESS_TOKEN>/setWebhook?url=<DOMAIN>&secret_token=<SECRET_TOKEN>
```

Создайте `.env` файл в корне проекта со следующим содержимым:

```shell
DOMAIN=<ДОМЕННОЕ_ИМЯ> (для получения SSL сертификата от Let's Encrypt)
ACCESS_TOKEN=<ТОКЕН_ДОСТУПА> (из BotFather)
SECRET_TOKEN=<СЕКРЕТНЫЙ_ТОКЕН> (который был указан в методе setWebhook)
CHECK_PERIOD=<ПЕРИОД_ПРОВЕРКИ_НОВОСТЕЙ> (в минутах)
```

Запустите проект:
```shell
go run .
```
