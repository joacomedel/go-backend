# Backend Go de ejemplo

[![CI](https://github.com/joacomedel/go-backend/actions/workflows/go.yml/badge.svg)](https://github.com/joacomedel/go-backend/actions)

Servicio HTTP mínimo con `net/http`, sin dependencias externas. Sirve como
base para backends Go y como proyecto de prueba del ecosistema `devctl`.

- Repo: https://github.com/joacomedel/go-backend
- CI: GitHub Actions en `ubuntu-24.04` y `ubuntu-24.04-arm`.

## Endpoints

- `GET /healthz` → `ok`
- `GET /readyz` → `ready`
- `GET /api/v1/hello?name=opencode` → `{"message":"hello, opencode"}`

## Desarrollo local

Go 1.27.3 está instalado en `/opt/go`:

```bash
export PATH=/opt/go/bin:$PATH
cd examples/go-backend
go test ./...
go run ./cmd/api
```

En otra terminal:

```bash
curl -s localhost:8080/healthz
curl -s 'localhost:8080/api/v1/hello?name=opencode'
```

## Con devctl

```bash
cd /root/pipelineCompilacion
./bin/devctl run test --project examples/go-backend
./bin/devctl run vet --project examples/go-backend
./bin/devctl run build --project examples/go-backend
./bin/devctl route --task test-race --project examples/go-backend --explain
```

`test-race` no corre en el Moto porque falta `gcc`; el router lo deriva a PC o
CI. En GitHub Actions sí corre con `-race` en x86 y ARM64.

## Repo y CI

Este proyecto ya vive en
[`joacomedel/go-backend`](https://github.com/joacomedel/go-backend). Cada push
a `main` ejecuta el workflow `.github/workflows/go.yml`:

- `gofmt`, `go vet` y `go test -race` en x86 y ARM64;
- builds estáticos `linux/amd64` y `linux/arm64`;
- artefactos descargables por `gh run download`.

Para crear otro backend, copiá esta estructura a un repo nuevo y cambiá el
módulo:

```bash
go mod edit -module github.com/TU_USUARIO/TU_REPO
go mod tidy
```

Después seguí `docs/github.md` para autenticar `gh`, crear el repo público y
verificar el primer run.

## CI

El workflow:

- corre `gofmt`, `go vet` y `go test -race` en `ubuntu-24.04` y
  `ubuntu-24.04-arm`;
- compila binarios estáticos para `linux/amd64` y `linux/arm64`;
- sube los artefactos por 7 días;
- cancela corridas viejas del mismo ref.

En repos públicos los runners estándar de GitHub son gratis. El último run
verificado terminó en 46 s con los 4 jobs en verde.
