.PHONY: test bench

test:
	# Tags bounds,noasm,safe are added for safety checks in Gonum.
	go test -count 1 -tags bounds,noasm,safe -race -v ./...

bench:
	go test -v -bench=. ./...

compare:
	go test -v -bench='FullFill_PP|FFull_PP|FullFill_II|FFull_II|FullFill_IP|FFull_IP' ./...
