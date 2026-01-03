APP=stegia

.PHONY: tidy build run-companies run-suppliers test clean

tidy:
	go mod tidy

build:
	go build -o $(APP) .

run-companies: build
	./$(APP) totvs companies list

run-suppliers: build
	./$(APP) totvs suppliers add -f ./examples/coca-cola.toon

test:
	go test ./...

clean:
	rm -f $(APP)
