#!/bin/bash

# 色付きの出力用
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}   データベース完全リセット${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# プロジェクトルートに移動
cd "$(dirname "$0")/.."

# 確認
echo -e "${RED}⚠️  警告: これは全てのデータを削除します！${NC}"
read -p "続行しますか？ (y/n): " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "キャンセルしました"
    exit 1
fi

# 1. データベースのリセット（オプション）
echo -e "${YELLOW}データベースをリセットしますか？ (y/n)${NC}"
read -p "" -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo -e "${YELLOW}データベースをリセット中...${NC}"
    psql -U dbgodotask -d postgres -c "DROP DATABASE IF EXISTS dbgodotask;"
    psql -U dbgodotask -d postgres -c "CREATE DATABASE dbgodotask;"
    echo -e "${GREEN}✓ データベースをリセットしました${NC}"
fi

# 2. マイグレーション実行
echo -e "${YELLOW}マイグレーションを実行中...${NC}"
go run ./cmd/migrate/main.go
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ マイグレーションに失敗しました${NC}"
    exit 1
fi
echo -e "${GREEN}✓ マイグレーション完了${NC}"

# 3. 全シードデータ投入
echo -e "${YELLOW}シードデータを投入中...${NC}"

echo "[1/3] Memory contexts..."
go run ./seed/seed.go
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Memory contextsのシードに失敗しました${NC}"
    exit 1
fi

echo "[2/3] Books, Tasks, Assessments..."
go run ./seed/seedModel.go
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Books/Tasks/Assessmentsのシードに失敗しました${NC}"
    exit 1
fi

echo "[3/3] Heuristics..."
go run ./cmd/seed/main.go -only heuristics
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Heuristicsのシードに失敗しました${NC}"
    exit 1
fi

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}✅ データベースのリセットと初期化が完了しました！${NC}"
echo -e "${GREEN}========================================${NC}"