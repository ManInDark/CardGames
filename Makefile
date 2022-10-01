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
	- rm WattenServer/watten.zip
	cd WattenServer; zip -r watten.zip *.js *.html svg; zip -jr watten.zip bin