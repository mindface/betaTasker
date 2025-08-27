#!/bin/bash

# =====================================================
# betaTasker - Local Production Test Script
# =====================================================
# このスクリプトはローカル環境で本番ビルドをテストします
# 使用方法: ./scripts/test-production-local.sh

set -e  # エラーが発生したら停止

# カラー出力の設定
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# ヘルパー関数
print_step() {
    echo -e "${BLUE}[STEP]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

# プロジェクトルートに移動
cd "$(dirname "$0")/.."
PROJECT_ROOT=$(pwd)

echo "================================================"
echo "    betaTasker - Production Build Test"
echo "================================================"
echo ""

# 1. 環境確認
print_step "環境確認中..."

# Docker確認
if ! command -v docker &> /dev/null; then
    print_error "Dockerがインストールされていません"
    exit 1
fi

# Docker Compose確認
if ! command -v docker-compose &> /dev/null; then
    print_error "Docker Composeがインストールされていません"
    exit 1
fi

print_success "環境確認完了"
echo ""

# 2. 既存のコンテナを停止
print_step "既存のコンテナを停止中..."
docker-compose down 2>/dev/null || true
docker-compose -f docker-compose.prod.yml down 2>/dev/null || true
print_success "既存のコンテナを停止しました"
echo ""

# 3. 本番用イメージをビルド
print_step "本番用Dockerイメージをビルド中..."
echo "これには数分かかる場合があります..."

# バックエンドのビルド
print_step "バックエンド (Go) をビルド中..."
docker build -f backer/godotask/Dockerfile.prod -t betatasker-backend:prod backer/godotask
if [ $? -eq 0 ]; then
    print_success "バックエンドのビルド完了"
else
    print_error "バックエンドのビルドに失敗しました"
    exit 1
fi

# フロントエンドのビルド
print_step "フロントエンド (Next.js) をビルド中..."
docker build -f fronter/Dockerfile.prod -t betatasker-frontend:prod fronter
if [ $? -eq 0 ]; then
    print_success "フロントエンドのビルド完了"
else
    print_error "フロントエンドのビルドに失敗しました"
    exit 1
fi

# Nginxのビルド
print_step "Nginx をビルド中..."
docker build -f nginx/Dockerfile -t betatasker-nginx:prod nginx
if [ $? -eq 0 ]; then
    print_success "Nginxのビルド完了"
else
    print_error "Nginxのビルドに失敗しました"
    exit 1
fi

echo ""
print_success "全てのイメージのビルドが完了しました"
echo ""

# 4. 本番環境を起動
print_step "本番環境を起動中..."
docker-compose -f docker-compose.prod.yml up -d

# 起動を待つ
print_step "サービスの起動を待機中..."
sleep 10

# 5. ヘルスチェック
print_step "ヘルスチェックを実行中..."
echo ""

# データベースの確認
echo -n "  PostgreSQL... "
if docker exec db-prod pg_isready -U dbgodotask &>/dev/null; then
    echo -e "${GREEN}OK${NC}"
else
    echo -e "${RED}FAILED${NC}"
    HEALTH_ERROR=1
fi

# バックエンドの確認
echo -n "  Backend API... "
if curl -f -s http://localhost:8080/health > /dev/null 2>&1; then
    echo -e "${GREEN}OK${NC}"
else
    # APIのヘルスエンドポイントがない場合は、通常のエンドポイントで確認
    if curl -f -s http://localhost:8080/api/task; then
        echo -e "${GREEN}OK${NC}"
    else
        echo -e "${RED}FAILED${NC}"
        HEALTH_ERROR=1
    fi
fi

# フロントエンドの確認
echo -n "  Frontend... "
if curl -f -s http://localhost:3000/api/test; then
    echo -e "${GREEN}OK${NC}"
else
    echo -e "${RED}FAILED${NC}"
    HEALTH_ERROR=1
fi

# Nginxの確認
echo -n "  Nginx... "
if curl -f -s http://localhost/health > /dev/null 2>&1; then
    echo -e "${GREEN}OK${NC}"
else
    # ヘルスエンドポイントがない場合
    if curl -f -s http://localhost > /dev/null 2>&1; then
        echo -e "${GREEN}OK${NC}"
    else
        echo -e "${RED}FAILED${NC}"
        HEALTH_ERROR=1
    fi
fi

echo ""

# 6. APIテスト
print_step "APIエンドポイントをテスト中..."

# ユーザー登録テスト
echo -n "  ユーザー登録... "
REGISTER_RESPONSE=$(curl -s -X POST http://localhost:8080/api/register \
    -H "Content-Type: application/json" \
    -d '{
        "username": "testuser_'$(date +%s)'",
        "password": "testpass123",
        "email": "test_'$(date +%s)'@example.com",
        "role": "user"
    }' 2>/dev/null || echo "FAILED")

if [[ $REGISTER_RESPONSE == *"token"* ]] || [[ $REGISTER_RESPONSE == *"already exists"* ]]; then
    echo -e "${GREEN}OK${NC}"
else
    echo -e "${RED}FAILED${NC}"
    echo "    Response: $REGISTER_RESPONSE"
fi

echo ""

# 7. ログの確認
print_step "コンテナログの最新5行:"
echo ""
echo "--- Backend ---"
docker logs backend-prod --tail 5 2>&1 | sed 's/^/  /'
echo ""
echo "--- Frontend ---"
docker logs frontend-prod --tail 5 2>&1 | sed 's/^/  /'
echo ""

# 8. 結果サマリー
echo ""
echo "================================================"
echo "              テスト結果サマリー"
echo "================================================"

if [ -z "$HEALTH_ERROR" ]; then
    print_success "全てのサービスが正常に起動しています！"
    echo ""
    echo "アクセス可能なURL:"
    echo "  - Frontend: http://localhost:3000"
    echo "  - Backend API: http://localhost:8080"
    echo "  - Nginx Proxy: http://localhost"
    echo ""
    echo "次のステップ:"
    echo "  1. ブラウザで http://localhost:3000 にアクセス"
    echo "  2. APIテスト: ./test_error_codes.sh"
    echo "  3. ログ確認: docker-compose -f docker-compose.prod.yml logs -f"
    echo "  4. 停止: docker-compose -f docker-compose.prod.yml down"
else
    print_error "一部のサービスで問題が発生しています"
    echo ""
    echo "トラブルシューティング:"
    echo "  - ログ確認: docker-compose -f docker-compose.prod.yml logs"
    echo "  - コンテナ状態: docker-compose -f docker-compose.prod.yml ps"
    echo "  - 再起動: docker-compose -f docker-compose.prod.yml restart"
fi

echo ""
echo "================================================"

# 9. オプション: インタラクティブモード
read -p "コンテナを停止しますか？ (y/N): " -n 1 -r
echo ""
if [[ $REPLY =~ ^[Yy]$ ]]; then
    print_step "コンテナを停止中..."
    docker-compose -f docker-compose.prod.yml down
    print_success "停止完了"
else
    print_warning "コンテナは起動したままです"
    echo "手動で停止する場合: docker-compose -f docker-compose.prod.yml down"
fi