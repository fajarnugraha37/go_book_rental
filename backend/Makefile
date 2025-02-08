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

generate-cert:
	@openssl genrsa -out key.pem 2048
	@openssl req -x509 -new -nodes -key key.pem -out cert.pem -days 365 -subj "/C=US/ST=YourState/L=YourCity/O=YourOrganization/OU=YourUnit/CN=localhost"
	@openssl x509 -in cert.pem -text -noout