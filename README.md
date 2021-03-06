# E.L.R.A.
[![Go Report Card](https://goreportcard.com/badge/github.com/JediWed/ELRA)](https://goreportcard.com/report/github.com/JediWed/ELRA)
## **E**asy **L**ND **R**EST **A**PI

**ELRA** is an open source REST extension for [LND](https://github.com/lightningnetwork/lnd) with a lot of additional features compared to LND's default REST or gRPC API built on Go.

## Comparison

![Comparison Table](https://raw.githubusercontent.com/JediWed/ELRA/master/docs/comparison.png "Comparison Table")

## Build

### Prerequisites

1. macOS
2. Install Go and Cross Compiler

```bash
brew install go 
brew install mingw-w64
brew install FiloSottile/musl-cross/musl-cross
brew reinstall FiloSottile/musl-cross/musl-cross --with-arm-hf
```
3. Set $GOPATH system environment
4. Get Go Dependencies
```bash
go get github.com/mattn/go-sqlite3
go get github.com/dgrijalva/jwt-go
go get github.com/gorilla/mux
go get github.com/golang/protobuf/proto
go get github.com/joncalhoun/qson
go get github.com/lightningnetwork/lnd/lnrpc
go get github.com/lightningnetwork/lnd/macaroons
go get golang.org/x/crypto/bcrypt
go get golang.org/x/sys/unix
```
5. Clone InventorizerAPI
```bash
cd %GOPATH/src
git clone https://github.com/JediWed/ELRA.git
cd ELRA
```

6. (Optional) Install and Start via PM2
```bash
npm install pm2@latest -g
pm2 start pm2-ELRA.json
```


### Build

|Command|Description|
|:------|:----------|
|make|Builds executable. Binary can be found in dist/|
|make clean|Cleanup previous builds and releases|
|make serve|Build and start API|
|make release|Builds binaries for macOS, Windows, Linux and Raspberry Pi. Release builds can be found in release/|

### Build Releases on other Hosts than macOS

To Cross Compile on other Hosts (e.g. Windows or Linux) edit Makefile and replace CC Binaries with your installed Cross Compile Binaries.

## License

ELRA is released under the terms of the GPL-3.0 license. See [LICENSE](LICENSE) for more information or see https://www.gnu.org/licenses/.
