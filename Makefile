rlpl:
	go build -o dist/rlpl cmd/rlpl/main.go
	./dist/rlpl

rppl:
	go build -o dist/rppl cmd/rppl/main.go
	./dist/rppl

repl:
	go build -o dist/repl cmd/repl/main.go
	./dist/repl

vm-repl:
	go build -o dist/vm-repl cmd/vm-repl/main.go
	./dist/vm-repl
