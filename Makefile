run.test:
	go run pms.test --dsn=./services/test/data/test.db

run.users:
	go run pms.users/cmd/app --dsn=./services/users/data/users.db