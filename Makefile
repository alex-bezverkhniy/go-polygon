all: install

install:
	go install -v

test:
	go test ./... -v

fmt:
	go fmt ./...

image:
	docker build -t alex-bezverkhniy/go-polygon .

release:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)
	goreleaser --rm-dist

.PHONY: install test fmt release

hello:
	go run main.go
test-hello:
	go test -v -run TestSayHello

test-text-tempale:
	go test -timeout 30s -v github.com/alex-bezverkhniy/go-polygon/template