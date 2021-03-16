## Тестовое задание Untypical

Задание:
Необходимо создать Http сервис - key-value хранилище.
Сервис должен содержать четыре метода в апи:
- Upsert (вставка либо обновление)
- Delete
- Get
- List

Хранить данные можно просто в оперативной памяти при помощи map.

### Как запустить

```mkdir untyp-test && cd untyp-test && git clone github.com/nikolayoskin/untyp-test .```

```make build && make port=8095 start```

Upsert endpoint принимает валидный json запрос, где и key и value являются строками.

Пример валидного upsert запроса:
```
curl --location --request POST 'localhost:8095/upsert' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "foo",
    "value": "bar"
}'
```

Ответ:
```
{
    "status": "ok",
    "message": "key and value added/updated"
}
```


В get и delete endpoint'ах ключ передается в URL адрес.

GET endpoints примеры:

- http://localhost:8095/list
- http://localhost:8095/get/foo

Удалить по ключу:
```
curl --location --request DELETE 'localhost:8095/delete/foo'
```