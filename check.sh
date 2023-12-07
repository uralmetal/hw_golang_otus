cd $(git branch --show-current)
golangci-lint run .
export CGO_ENABLED=1 && go test -v -count=1 -race -timeout=1m .
[ -f "test.sh" ] && ./test.sh || echo "skip test.sh"
