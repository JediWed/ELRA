.DEFAULT_GOAL := build
version := 1.0

clean:
	$(RM) dist/elra dist/elra.exe dist/elra.db
	rm -rf release

build: clean
	go build -o dist/elra elra.go

serve: build
	cd dist && ./elra

release: clean
	mkdir release

	# macOS x64
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o dist/elra elra.go
	zip -rj release/ELRA_$(version)_macOS.zip dist/*
	$(RM) dist/elra dist/elra.exe dist/elra.db
	
	# Windows x64
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o dist/elra.exe elra.go
	zip -rj release/ELRA_$(version)_Windows.zip dist/*
	$(RM) dist/elra dist/elra.exe dist/elra.db
	
	# Linux x64
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-linux-musl-gcc go build -ldflags '-w -extldflags "-static"' -o dist/elra elra.go
	zip -rj release/ELRA_$(version)_Linux.zip dist/*
	$(RM) dist/elra dist/elra.exe dist/elra.db
	
	# ARMHF x86 (Raspberry Pi) 
	GOOS=linux GOARCH=arm GOARM=7 CC=arm-linux-musleabihf-gcc go build -o dist/elra elra.go
	zip -rj release/ELRA_$(version)_RaspberryPi.zip dist/*
	$(RM) dist/elra dist/elra.exe dist/elra.db