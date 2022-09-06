rm -rf build
env GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o build/quantum-linux-386 &&
env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o build/quantum-linux-amd64 &&
env GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o build/quantum-linux-arm &&
env GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o build/quantum-linux-arm64 &&
env GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o build/quantum-darwin-amd64 &&
env GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o build/quantum-darwin-arm64 &&
env GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o build/quantum-js-wasm &&
env GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o build/quantum-windows-386 &&
env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o build/quantum-windows-amd64
