# make run service=[api|migrator|seeder]
run:
	@go run cmd/$(service)/main.go
	
# make run service=[api|migrator|seeder]
build:
	@go build -o ./$(service).exe -ldflags="-s -w" cmd/$(service)/main.go