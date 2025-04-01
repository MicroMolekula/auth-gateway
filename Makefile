build:
	docker build -t not-five/auth-gateway .
run:
	docker run -d -p 127.0.0.1:8000:8000 --name=auth-gateway not-five/auth-gateway