all: dev

list:
	@LC_ALL=C $(MAKE) -pRrq -f $(firstword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/(^|\n)# Files(\n|$$)/,/(^|\n)# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | grep -E -v -e '^[^[:alnum:]]' -e '^$@$$'

dev: 
	air -c .air.conf

test:
	go test -v ./... | sed ''/PASS/s//$$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

log_tail:
	tail -f app.log

update_doc:
	swag init

run_benchmark:
	ab -k -n 1000 -c 10 127.0.0.1:6969/api/v1/employees

docker_build:
	docker build -t fiber_go_backend .

docker_run:
	docker run -p 6969:6969 fiber_go_backend
