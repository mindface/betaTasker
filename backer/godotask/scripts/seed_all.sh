#!/bin/bash

# スクリプトのディレクトリを取得
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_DIR="$( cd "$SCRIPT_DIR/.." && pwd )"

# プロジェクトディレクトリに移動
cd "$PROJECT_DIR"

# 色付きの出力用
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}   全Seedデータ実行スクリプト${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

# エラーハンドリング
set -e
trap 'echo -e "${RED}❌ エラーが発生しました${NC}"' ERR

# 1. Memory関連データのシード
echo -e "${YELLOW}[1/3] Memory関連データ（製造工程Level1-3）をシード中...${NC}"
go run ./seed/seed.go
echo -e "${GREEN}✓ Memory関連データのシード完了${NC}"
echo ""

# 2. Book/Task/Assessmentデータのシード
echo -e "${YELLOW}[2/3] Book, Memory, Task, Assessmentデータをシード中...${NC}"
go run ./seed/seedModel.go
echo -e "${GREEN}✓ Book/Task/Assessmentデータのシード完了${NC}"
echo ""

# 3. Heuristicsデータのシード
echo -e "${YELLOW}[3/3] Heuristicsデータをシード中...${NC}"
go run ./cmd/seed/main.go -only heuristics
echo -e "${GREEN}✓ Heuristicsデータのシード完了${NC}"
echo ""

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}✅ 全てのSeedデータの投入が完了しました！${NC}"
echo -e "${GREEN}========================================${NC}"

# データ件数の確認（オプション）
echo ""
echo -e "${BLUE}データ件数を確認しますか？ (y/n)${NC}"
read -r response
if [[ "$response" =~ ^[Yy]$ ]]; then
    echo -e "${YELLOW}データ件数を確認中...${NC}"
    psql -U dbgodotask -d dbgodotask -c "
        SELECT 'memory_contexts' as table_name, COUNT(*) as count FROM memory_contexts
        UNION ALL
        SELECT 'technical_factors', COUNT(*) FROM technical_factors
        UNION ALL
        SELECT 'knowledge_transformations', COUNT(*) FROM knowledge_transformations
        UNION ALL
        SELECT 'books', COUNT(*) FROM books
        UNION ALL
        SELECT 'memories', COUNT(*) FROM memories
        UNION ALL
        SELECT 'tasks', COUNT(*) FROM tasks
        UNION ALL
        SELECT 'assessments', COUNT(*) FROM assessments
        UNION ALL
        SELECT 'heuristics_analyses', COUNT(*) FROM heuristics_analyses
        UNION ALL
        SELECT 'heuristics_trackings', COUNT(*) FROM heuristics_trackings
        UNION ALL
        SELECT 'heuristics_insights', COUNT(*) FROM heuristics_insights
        UNION ALL
        SELECT 'heuristics_patterns', COUNT(*) FROM heuristics_patterns
        UNION ALL
        SELECT 'heuristics_models', COUNT(*) FROM heuristics_models
        ORDER BY table_name;
    "
fi