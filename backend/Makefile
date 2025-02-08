migrate:
	@go run .\cmd\migrator\

seed:
	@go run .\cmd\migrator\
# make run service=[api|migrator|seeder]

run:
	@go run .\cmd\api\
	
# make run service=[api|migrator|seeder]
build:
	@go build -o .\$(service).exe -ldflags="-s -w" .\cmd\$(service)\