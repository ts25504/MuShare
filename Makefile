install:
	go install

run-prod:
	MARTINI_ENV=production go run app.go

run-dev:
	MARTINI_ENV=development go run app.go

run-test:
	MARTINI_ENV=test go run app.go

test:
	go test

