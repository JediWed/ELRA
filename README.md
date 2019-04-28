# E.L.R.A.
## **E**asy **L**ND **R**EST **A**PI

**ELRA** is an open source REST extension for [LND](https://github.com/lightningnetwork/lnd) with a lot of additional features compared to LND's default REST or gRPC API built on Go.

## Comparison

|                                   |ELRA|LND REST|LND gRPC|
|-----------------------------------|:--:|:------:|:------:|
|Easy to use                        |<span style="color:green">**yes**</span>|<span style="color:red">**no**  |<span style="color:red">**no**        |
|User Management                    |<span style="color:green">**yes**</span>|<span style="color:red">**no**  |<span style="color:red">**no**        |
|Role Management                    |<span style="color:green">**yes**</span>|<span style="color:red">**no**  |<span style="color:red">**no**        |
|Reveals Macaroons to Frontend      |<span style="color:green">**no**</span> |<span style="color:green">**no**|<span style="color:red">**yes**       |
|Reveals RPC Credentials to Frontend|<span style="color:green">**no**</span> |<span style="color:red">**yes** |<span style="color:green">**no**</span>|


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
go get golang.org/x/crypto/bcrypt
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

## License

ELRA is released under the terms of the GPL-3.0 license. See [LICENSE](LICENSE) for more information or see https://www.gnu.org/licenses/.