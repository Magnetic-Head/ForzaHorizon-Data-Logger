BACKEND_DIR = backend
FRONTEND_DIR = frontend

.PHONY: dev
dev:
	$(MAKE) -j2 dev-frontend dev-backend

.PHONY: dev-frontend
dev-frontend:
	cd $(FRONTEND_DIR) && npm run dev

.PHONY: dev-backend
dev-backend:
	cd $(BACKEND_DIR) && go run cmd/server/main.go

.PHONY: build-backend
build-backend:
	cd $(BACKEND_DIR) && go build -o app .
