migrateup:
	~/go/bin/migrate -path providers/db/migration -database "postgresql://postgres:local@localhost:5432/finance?sslmode=disable" -verbose up

migratedown:
	~/go/bin/migrate -path providers/db/migration -database "postgresql://postgres:local@localhost:5432/finance?sslmode=disable" -verbose down

sqlgen:
	~/go/bin/sqlc generate