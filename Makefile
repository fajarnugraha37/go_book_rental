up:
	@docker compose -f ./infrastructure/docker-compose.yml up --force-recreate -d

down:
	@docker compose -f ./infrastructure/docker-compose.yml down
	-docker image rm infrastructure-backend:latest 
	-docker image rm infrastructure-frontend:latest

rebuild:
	@docker compose -f ./infrastructure/docker-compose.yml up --force-recreate -d --build $(service)