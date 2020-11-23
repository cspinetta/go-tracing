Playground for tracing in Golang
=================================

## Relevant dependencies

- `go 1.14`

**API and DB:**

- `github.com/gin-gonic/gin v1.6.3`
- `github.com/jmoiron/sqlx v1.2.0`

**Tracing implementation:**

- `go.opentelemetry.io/otel v0.14.0`

**Instrumentation for API and DB:**

- `go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.14.0`
- `github.com/cspinetta/otsql v0.3.0-alpha-2`

**Exporter:**

- `go.opentelemetry.io/otel/exporters/trace/jaeger v0.14.0`

## Run docker compose

```
docker-compose up --build --force-recreate --no-deps -d
```

## Example requests:

```
curl --request GET \
  --url http://localhost:8080/ping
```

```
curl --request POST \
  --url http://localhost:8080/user \
  --header 'content-type: application/json' \
  --data '{
	"name": "Charles Darwin",
	"birthday": "1809-02-12"
}'
```

```
curl --request GET \
  --url http://localhost:8080/user/1
```

```
curl --request GET \
  --url 'http://localhost:8080/user?limit=100&offset=0'
```

## Inspect traces on Jaeger UI

Open http://localhost:16686/
