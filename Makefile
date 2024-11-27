new-sql:
	@read -p "Enter SQL File Name: " sqlname; \
	go run main.go new-sql $$sqlname