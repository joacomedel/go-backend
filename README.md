# Backend Go de ejemplo

Servicio HTTP mínimo con `net/http`, sin dependencias externas. Sirve como
base para backends Go y como proyecto de prueba del ecosistema `devctl`.

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

## Pasar esto a un repo público propio

1. Copiá el contenido de `examples/go-backend` a la raíz de un repo nuevo.
2. Cambiá el módulo:

```bash
go mod edit -module github.com/TU_USUARIO/TU_REPO
go mod tidy
```

3. El workflow ya está en `.github/workflows/go.yml`; GitHub lo detecta al
   pushear a la rama por defecto.
4. Seguí `docs/github.md` para autenticar `gh`, crear el repo y verificar el
   primer run.

## CI

El workflow:

- corre `gofmt`, `go vet` y `go test -race` en `ubuntu-24.04` y
  `ubuntu-24.04-arm`;
- compila binarios estáticos para `linux/amd64` y `linux/arm64`;
- sube los artefactos por 7 días;
- cancela corridas viejas del mismo ref.

En repos públicos los runners estándar de GitHub son gratis.
