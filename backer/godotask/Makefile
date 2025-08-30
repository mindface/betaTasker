.PHONY: help build run test seed migrate clean

# デフォルトターゲット
help:
	@echo "Available commands:"
	@echo "  make build     - Build the application"
	@echo "  make run       - Run the application"
	@echo "  make test      - Run tests"
	@echo "  make seed      - Seed the database"
	@echo "  make seed-clean - Clean and seed the database"
	@echo "  make seed-heuristics - Seed only heuristics data"
	@echo "  make migrate   - Run database migrations"
	@echo "  make clean     - Clean build artifacts"

# ビルド
build:
	go build -o bin/godotask ./cmd/server

# 実行
run:
	go run ./cmd/server

# テスト
test:
	go test -v ./...

# データベースシード
seed:
	go run ./cmd/seed/main.go

# クリーンシード（データベースをクリアしてからシード）
seed-clean:
	go run ./cmd/seed/main.go -clean

# Heuristicsのみシード
seed-heuristics:
	go run ./cmd/seed/main.go -only heuristics

# seed配下の全ファイルを実行
seed-all:
	@echo "=== Seeding all data ==="
	@echo "[1/3] Seeding Memory contexts (Level 1-3)..."
	@go run ./seed/seed.go
	@echo "[2/3] Seeding Books, Tasks, Assessments..."
	@go run ./seed/seedModel.go
	@echo "[3/3] Seeding Heuristics data..."
	@go run ./cmd/seed/main.go -only heuristics
	@echo "✅ All seed data completed!"

# 全データをクリーンしてからシード
seed-all-clean:
	@echo "=== Cleaning and seeding all data ==="
	@echo "⚠️  This will DELETE all existing data!"
	@read -p "Continue? (y/n): " confirm && [ "$$confirm" = "y" ] || exit 1
	@echo "[1/3] Seeding Memory contexts (Level 1-3)..."
	@go run ./seed/seed.go
	@echo "[2/3] Seeding Books, Tasks, Assessments..."
	@go run ./seed/seedModel.go
	@echo "[3/3] Seeding Heuristics data (clean mode)..."
	@go run ./cmd/seed/main.go -clean
	@echo "✅ All data cleaned and seeded!"

# マイグレーション
migrate:
	go run ./cmd/migrate/main.go

# データベースリセット＆初期化
db-reset:
	chmod +x ./scripts/reset_db.sh
	./scripts/reset_db.sh

# マイグレーション後に全シード実行
db-setup: migrate seed-all
	@echo "✅ Database setup completed!"

# クリーンアップ
clean:
	rm -rf bin/
	go clean

# Docker関連
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

# 開発環境セットアップ
dev-setup: docker-up migrate seed
	@echo "Development environment is ready!"

# 開発環境リセット
dev-reset: docker-down docker-up migrate seed-clean
	@echo "Development environment has been reset!"