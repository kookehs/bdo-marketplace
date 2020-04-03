GOOS=darwin GOARCH=386 go build -o bdo-marketplace-darwin32 &&
GOOS=darwin GOARCH=amd64 go build -o bdo-marketplace-darwin64 &&
GOOS=linux GOARCH=386 go build -o bdo-marketplace-linux32 &&
GOOS=linux GOARCH=amd64 go build -o bdo-marketplace-linux64 &&
GOOS=windows GOARCH=386 go build -o bdo-marketplace-windows32.exe &&
GOOS=windows GOARCH=amd64 go build -o bdo-marketplace-windows64.exe