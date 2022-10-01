test_CardDeck:
	go test ./CardDeck -cover -v

test_Watten:
	go test ./Watten -cover -v

test_all: test_CardDeck test_Watten

build_WattenServer_Linux:
	go build -o WattenServer/bin/WattenServer ./WattenServer

build_WattenServer_Windows:
	GOOS=windows GOARCH=amd64 go build -o WattenServer/bin/WattenServer.exe ./WattenServer

build_all: build_WattenServer_Linux build_WattenServer_Windows

zip_program: build_all
	- rm watten.zip
	zip -j watten.zip WattenServer/bin/WattenServer WattenServer/bin/WattenServer.exe WattenServer/*.js WattenServer/*.html