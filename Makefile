reset_docker:
	docker compose down --rmi local --volumes --remove-orphans
	docker compose up

frontend_container:
	docker exec -it frontend_container /bin/sh
