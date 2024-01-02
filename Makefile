composeName:=subscription-weather-bot
composeFile:=docker/compose.yaml

start:
	docker compose -p $(composeName) -f $(composeFile) --env-file .env up

restart:
	docker compose -p $(composeName) -f $(composeFile) stop 
	docker rm bot-server
	docker rmi $(composeName)-server
	docker compose -p $(composeName) -f $(composeFile) --env-file .env up
