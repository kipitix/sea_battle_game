
server_docker_debug_up:
	@cd ./sbg_server/deployments/debug ; docker-compose up -d

server_docker_debug_down:
	@cd ./sbg_server/deployments/debug ; docker-compose down

