build:
	docker build -t not-five/auth-gateway .
run:
	docker run -d -p 127.0.0.1:8000:8000 --name=auth-gateway not-five/auth-gateway
up:
	docker build -t not-five/auth-gateway .
	docker run -d -p 127.0.0.1:8000:8000 --name=auth-gateway not-five/auth-gateway
down:
	docker stop auth-gateway
	docker remove auth-gateway
local-run:
	CONFIG_PATH='./configs/dev.yml' go run cmd/main.go