export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgresql://postgres:qwerty@localhost:5432/vertex?sslmode=disable

cd migrations/init
goose up
cd ..
cd updates
goose up
cd ../../
