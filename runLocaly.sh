export POSTGRES_CONNECTION_STRING="postgresql://admin:password@1234@localhost:55432/postgres"
export PASSWORD_PEPPER="MyH0rse1sAm@zing"
export PRIVATE_KEY="MyPr!v@tKey"
export DOMAIN="localhost:8081"
export ACCESS_TOKEN_ROUTE="/api/auth"

go run main.go :8081
