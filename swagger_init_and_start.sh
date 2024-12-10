export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgresql://postgres:qwerty@localhost:5432/vertex?sslmode=disable
swag init -g cmd/main.go
cd docs
npx @redocly/cli build-docs swagger.json
xdg-open http://localhost:63342/vertex-api/docs/redoc-static.html?_ijt=8oisu34614haqesc5tpncta6bd&_ij_reload=RELOAD_ON_SAVE
cd ..
