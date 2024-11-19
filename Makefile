docker-build:
	@echo "Building docker..."
	@docker build -t freight-trackr-backend .
docker-run:
	@docker run --name nama-container -d --env-file .env -p 3000:3000 freight-trackr-backend
build:
	@echo "Building..."
	@go build -o main.exe .
run:
	@go run main.go
seed:
	@go run cmd/seed/seed.go