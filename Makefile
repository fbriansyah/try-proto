GO_MODULE := hello-proto

vendor:
	@go mod vendor

tidy:
	@go mod tidy

clean-a :
	ifeq ($(OS), Windows_NT)
		if exist "protogen" rd /s /q protogen
	else
		rm -fR ./protogen
	endif

clean :
	@rm -rf protogen

protoc: 
	@protoc --go_opt=module=${GO_MODULE} --go_out=. \
	./proto/basic/*.proto \
	./proto/dummy/*.proto \
	./proto/jobsearch/*.proto

build: clean protoc tidy

run:
	@go run .\main.go

.PHONY: protoc tidy run vendor build