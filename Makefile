new-sql:
	@read -p "Enter SQL File Name: " sqlname; \
	go run main.go new-sql $$sqlname

serve:
	@go run main.go

up:
	docker compose up -d

down:
	docker compose down