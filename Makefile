.PHONY: setup watch build docker check load load10

setup:
	cd ~ && go get -u golang.org/x/tools/cmd/present
	cd ~ && go get -u github.com/rakyll/hey

watch:
	./scripts/dev.sh

build:
	go build -o app

docker:
	GOOS=linux go build -o app_linux
	docker build -t app .
	docker run --rm -p 4000:4000 --memory 100m --memory-swap 100m --name app app

docker-stop:
	docker kill app

check:
	@if [[ -z "${TASK}" ]]; then echo "Missing TASK target"; exit 1; fi

load: check
	hey -cpus=1 -z 1h -q 60 http://localhost:4000/work/${TASK}

multiload: check
	hey -cpus=1 -c 50 -z 1h -q 60 http://localhost:4000/work/${TASK}

load10: check
	hey -cpus=1 -c 1 -z 10s http://localhost:4000/work/${TASK}

multiload10: check
	hey -cpus=1 -c 50 -z 10s http://localhost:4000/work/${TASK}
