#!/bin/bash

# =====================================================
# betaTasker - Quick Test Script
# =====================================================
# 開発環境を素早くテストするスクリプト
# 使用方法: ./scripts/quick-test.sh

set -e

# カラー出力
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo "=========================================="
echo "    betaTasker - Quick Test"
echo "=========================================="
echo ""

# 1. 開発環境の起動
echo "🚀 開発環境を起動中..."
docker-compose up -d

# 待機
echo "⏳ サービスの起動を待機中 (10秒)..."
sleep 10

# 2. サービスの確認
echo ""
echo "📋 サービス状態:"
docker-compose ps

# 3. バックエンドのテスト
echo ""
echo "🔧 バックエンドAPIテスト:"

# ヘルスチェック
echo -n "  Health Check: "
if curl -f -s http://localhost:8080 > /dev/null 2>&1; then
    echo -e "${GREEN}✓${NC}"
else
    echo -e "${RED}✗${NC}"
fi

# 4. フロントエンドのテスト
echo ""
echo "🎨 フロントエンドテスト:"

echo -n "  Next.js Server: "
if curl -f -s http://localhost:3000 > /dev/null 2>&1; then
    echo -e "${GREEN}✓${NC}"
else
    echo -e "${YELLOW}起動中...${NC}"
fi

# 5. データベースのテスト
echo ""
echo "💾 データベーステスト:"

echo -n "  PostgreSQL: "
if docker exec db pg_isready -U dbgodotask &>/dev/null; then
    echo -e "${GREEN}✓${NC}"
else
    echo -e "${RED}✗${NC}"
fi

# 6. APIエンドポイントテスト
echo ""
echo "🔌 APIエンドポイントテスト:"
echo ""

# テスト用のタイムスタンプ
TIMESTAMP=$(date +%s)

# ユーザー登録
echo "1. ユーザー登録テスト:"
curl -X POST http://localhost:8080/api/register \
    -H "Content-Type: application/json" \
    -d '{
        "username": "test'$TIMESTAMP'",
        "password": "pass123",
        "email": "test'$TIMESTAMP'@example.com",
        "role": "user"
    }' 2>/dev/null | jq '.' || echo "  登録済みまたはエラー"

echo ""

# ログイン
echo "2. ログインテスト:"
TOKEN=$(curl -s -X POST http://localhost:8080/api/login \
    -H "Content-Type: application/json" \
    -d '{
        "username": "test'$TIMESTAMP'",
        "password": "pass123",
        "email": "test'$TIMESTAMP'@example.com"
    }' | jq -r '.token')

if [ ! -z "$TOKEN" ] && [ "$TOKEN" != "null" ]; then
    echo -e "  ${GREEN}✓ トークン取得成功${NC}"
    echo "  Token: ${TOKEN:0:20}..."
else
    echo -e "  ${YELLOW}⚠ トークン取得失敗（既存ユーザーでログインします）${NC}"
    # 既存のテストユーザーでログイン
    TOKEN=$(curl -s -X POST http://localhost:8080/api/login \
        -H "Content-Type: application/json" \
        -d '{
            "username": "testuser",
            "password": "testpass123",
            "email": "test@example.com"
        }' | jq -r '.token')
fi

echo ""

# タスク作成
if [ ! -z "$TOKEN" ] && [ "$TOKEN" != "null" ]; then
    echo "3. タスク作成テスト:"
    curl -X POST http://localhost:8080/api/task \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer $TOKEN" \
        -d '{
            "title": "テストタスク '$TIMESTAMP'",
            "description": "クイックテストで作成",
            "status": "pending"
        }' 2>/dev/null | jq '.' || echo "  エラー"
fi

# 7. ログ表示オプション
echo ""
echo "=========================================="
echo -e "${GREEN}✓ テスト完了！${NC}"
echo ""
echo "📌 次のアクション:"
echo "  • ブラウザで確認: http://localhost:3000"
echo "  • APIテスト実行: ./test_error_codes.sh"
echo "  • ログ確認: docker-compose logs -f"
echo "  • 停止: docker-compose down"
echo ""

# インタラクティブオプション
read -p "ログを表示しますか？ (y/N): " -n 1 -r
echo ""
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "最新のログ (Ctrl+C で終了):"
    docker-compose logs --tail=20 -f
fi