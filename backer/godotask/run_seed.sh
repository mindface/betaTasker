#!/bin/bash

# betaTasker Seed Script
# このスクリプトはseedデータの投入を実行します

# カラー設定
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo "🌱 betaTasker Seed Script"
echo "========================="

# プロジェクトディレクトリに移動
cd /Users/asdfghjkl/project/betaTasker/backer/godotask

# オプション処理
if [ "$1" = "clean" ]; then
    echo -e "${GREEN}Running clean seed...${NC}"
    go run cmd/seed/main.go -clean
elif [ "$1" = "only" ] && [ -n "$2" ]; then
    echo -e "${GREEN}Running seed only for: $2${NC}"
    go run cmd/seed/main.go -only "$2"
else
    echo -e "${GREEN}Running standard seed...${NC}"
    go run cmd/seed/main.go
fi

# 結果確認
if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Seed completed successfully!${NC}"
else
    echo -e "${RED}❌ Seed failed!${NC}"
    exit 1
fi