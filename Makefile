.PHONY: test lint check format bench

test:
	# Tags bounds,noasm,safe are added for safety checks in Gonum.
	go test -count 1 -tags bounds,noasm,safe -race -v ./...

lint:
	gofmt -l .; test -z "$$(gofmt -l .)"
	find . \( -name '*.c' -or -name '*.h' \) -exec clang-format-10 --style=file --dry-run --Werror {} +
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all ./...
	go run github.com/mgechev/revive@latest -set_exit_status ./...
	go run github.com/securego/gosec/v2/cmd/gosec@latest scan -checks=all ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

check: lint test

format:
	gofmt -w -s .
	find . \( -name '*.c' -or -name '*.h' \) -exec clang-format-10 --style=file -i {} +

bench:
	go test -v -bench=. ./...
