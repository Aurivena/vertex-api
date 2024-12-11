swag init -g cmd/main.go
cd docs
npx @redocly/cli build-docs swagger.json
cd ..
