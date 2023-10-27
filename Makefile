# for Window
SHELL=cmd

build:
	@echo Building...
	@go build -o dist/api.exe ./${SHELL}
	@echo Backend built!

start: build
	@echo Starting the Backend ...
	.\dist\api.exe
	@echo Backend running!

stop:
	@echo Stopping the Backend...
	@taskkill /IM api.exe /F
	@echo Stopped Backend