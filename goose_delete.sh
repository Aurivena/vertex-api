export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgresql://postgres:qwerty@localhost:5432/vertex?sslmode=disable

cd migrations/updates
goose down
cd ..
cd init
goose down
goose down
cd ../../
