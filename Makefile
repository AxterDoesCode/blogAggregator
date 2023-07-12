build:
	go build -o bin/ ./cmd/blogAggregator
exec:
	./bin/blogAggregator
migrateu:
	goose -dir ./sql/schema/ postgres postgresql://axtr:123@localhost:5432/blogator up
migrated:
	goose -dir ./sql/schema/ postgres postgresql://axtr:123@localhost:5432/blogator down


