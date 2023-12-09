all: dev

dev: 
	air -c .air.conf

test:
	go test -v ./... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

logt:
	tail -f log/app.log

doc:
	swag init

benchmark:
	ab -k -n 1000 -c 10 127.0.0.1:6969/api/v1/employees
