export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgresql://postgres:qwerty@localhost:5432/vertex?sslmode=disable
swag init -g cmd/main.go
cd docs
npx @redocly/cli build-docs swagger.json
cd ..
