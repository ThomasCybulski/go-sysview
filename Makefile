install:
	@go get -u github.com/swaggo/swag/cmd/swag
	@go get -u github.com/asticode/go-astilectron-bundler/...
	@go install github.com/asticode/go-astilectron-bundler/astilectron-bundler
	@echo "[OK] Installation of golang dependencies completed!"

doc:
	@swag init main.go
	@echo "[OK] Swagger files created!"

ng-build:
	@cd ./ui  && npm install &&  npm run build
	@rm -rf resources/app/ && mkdir resources/app
	@cp -R ./ui/dist/* resources/app/
	@echo "[OK] Build Angular frontend complete!"

build: install doc ng-build
	@mv pkg/gui/bind.go pkg/gui/bind_exclude
	@if astilectron-bundler; then \
		mv pkg/gui/bind_exclude pkg/gui/bind.go;  \
		echo "[OK] Bundle application done!"; \
	else \
		mv pkg/gui/bind_exclude pkg/gui/bind.go;  \
		echo "[Error] Bundle application failed!"; \
	fi