.PHONY: setup watch build docker check load load10

setup:
	cd ~ && go get -u golang.org/x/tools/cmd/present
	cd ~ && go get -u github.com/rakyll/hey

docker:
	GOOS=linux go build -o app_linux
	docker build -t app .
	docker run --rm -p 4000:4000 --memory 100m --memory-swap 100m --name app app

cpu-load:
	hey -cpus=1 -c 1 -z 45s http://localhost:4000/work/cpu

cpu-profile:
	go tool pprof -nodefraction=0 http://localhost:4000/debug/pprof/profile\?seconds\=10 

cpu-bench:
	go test -bench=. -benchtime=5s

mem-load:
	hey -cpus=1 -c 1 -z 45s http://localhost:4000/work/mem

mem-profile:
	go tool pprof -nodefraction=0 http://localhost:4000/debug/pprof/heap

lock-load:
	hey -cpus=1 -c 1 -z 45s http://localhost:4000/work/lock

lock-load-many:
	hey -cpus=1 -c 10 -z 45s http://localhost:4000/work/lock

lock-profile:
	go tool pprof -nodefraction=0 http://localhost:4000/debug/pprof/block

mutex-profile:
	go tool pprof -nodefraction=0 http://localhost:4000/debug/pprof/mutex

panic:
	curl http://localhost:4000/work/invalid
