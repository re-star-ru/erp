# Backend doc

```mermaid
graph TD
  go[golang backend] --> s3[minio]
  go --> search[meilisearch]
  go --> db[(postgres ? mongodb)]
  
  frontend[flutter app]-- grpc -->go
  web[web shop]-- rest -->go
 
  go --> telemetry[jaeger] --> graphana
  go --> logs[stdout logs] --> loki --> graphana
```
можно без метрик обойтись в принципе, логи и телеметрия

## GitHub actions

add to commit message ```[skip ci]``` to skip commit deploy

## Caddy reverse proxy

casbin plugin for caddy :

```http
https://github.com/casbin/caddy-authz/blob/master/authz.go#L100
```

add oauth or jwt auth to plugin

## caddy setup for reverse proxy

```text
Get backends dynamically from SRV record DNS queries:

reverse_proxy {
dynamic srv _api._tcp.example.com
}
```

## Oprox

is proxy for once services and some buisness logic

## Search entry

```json5
{
    "charGUID": "guid string", // primary key
    "GUID": "guid of item",
    "typeGUID": "guid of type",

    // searchable
    "sku": "analog group key",
    "name": "name",
    "type": "type of item, starter or generator or else",
    "amount": "current amount",
    "firstIncome": "first time when item entered the warehouse",
    "lastExpance": "last time item sold"
}
```

```js
filterPurple = `filter = "lastExpance < ${time.Now()}"`
```
