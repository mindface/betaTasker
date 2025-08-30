#!/bin/bash

# =====================================================
# betaTasker - Health Check Script
# =====================================================
# 全サービスの健全性を確認
# 使用方法: ./scripts/health-check.sh [--watch]

# カラー設定
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

# アイコン
CHECK="✓"
CROSS="✗"
LOADING="⟳"

# 引数チェック（--watchオプション）
WATCH_MODE=false
if [ "$1" == "--watch" ]; then
    WATCH_MODE=true
fi

check_service() {
    local name=$1
    local url=$2
    local container=$3
    
    printf "  %-20s" "$name:"
    
    # コンテナが起動しているか確認
    if [ ! -z "$container" ]; then
        if ! docker ps --format "table {{.Names}}" | grep -q "^$container$"; then
            echo -e "${RED}$CROSS コンテナ停止中${NC}"
            return 1
        fi
    fi
    
    # URLへのアクセス確認
    if curl -f -s "$url" > /dev/null 2>&1; then
        echo -e "${GREEN}$CHECK 正常${NC}"
        return 0
    else
        echo -e "${RED}$CROSS 応答なし${NC}"
        return 1
    fi
}

check_container_health() {
    local container=$1
    local name=$2
    
    printf "  %-20s" "$name:"
    
    if docker ps --format "table {{.Names}}" | grep -q "^$container$"; then
        # CPUとメモリ使用率を取得
        STATS=$(docker stats --no-stream --format "table {{.CPUPerc}}\t{{.MemUsage}}" $container | tail -1)
        echo -e "${GREEN}$CHECK 稼働中${NC} [$STATS]"
    else
        echo -e "${RED}$CROSS 停止中${NC}"
    fi
}

perform_check() {
    clear
    echo "================================================"
    echo "         betaTasker Health Check"
    echo "================================================"
    echo "Time: $(date '+%Y-%m-%d %H:%M:%S')"
    echo ""
    
    # 1. Docker環境の確認
    echo "${CYAN}🐳 Docker環境${NC}"
    printf "  %-20s" "Docker:"
    if docker version > /dev/null 2>&1; then
        echo -e "${GREEN}$CHECK 利用可能${NC}"
    else
        echo -e "${RED}$CROSS 利用不可${NC}"
        exit 1
    fi
    
    printf "  %-20s" "Docker Compose:"
    if docker-compose version > /dev/null 2>&1; then
        echo -e "${GREEN}$CHECK 利用可能${NC}"
    else
        echo -e "${RED}$CROSS 利用不可${NC}"
        exit 1
    fi
    echo ""
    
    # 2. コンテナの状態
    echo "${CYAN}📦 コンテナ状態${NC}"
    check_container_health "db" "Database"
    check_container_health "backender" "Backend"
    check_container_health "fronter" "Frontend"
    check_container_health "nginx" "Nginx"
    echo ""

    # 3. サービスの応答確認
    echo "${CYAN}🌐 サービス応答${NC}"
    check_service "PostgreSQL" "localhost" "db"
    check_service "Backend API" "http://localhost:8080" "backender"
    check_service "Frontend" "http://localhost:3000" "fronter"
    check_service "Nginx Proxy" "http://localhost" "nginx"
    echo ""
    
    # 4. API エンドポイントの確認
    echo "${CYAN}🔌 APIエンドポイント${NC}"
    
    # ログインエンドポイント
    printf "  %-20s" "/api/login:"
    LOGIN_STATUS=$(curl -s -o /dev/null -w "%{http_code}" -X POST http://localhost:8080/api/login \
        -H "Content-Type: application/json" \
        -d '{"username":"test","password":"test","email":"test@test.com"}' 2>/dev/null)
    
    if [ "$LOGIN_STATUS" == "200" ] || [ "$LOGIN_STATUS" == "401" ] || [ "$LOGIN_STATUS" == "400" ]; then
        echo -e "${GREEN}$CHECK 応答あり (HTTP $LOGIN_STATUS)${NC}"
    else
        echo -e "${RED}$CROSS エラー (HTTP $LOGIN_STATUS)${NC}"
    fi
    
    # タスクエンドポイント
    printf "  %-20s" "/api/task:"
    TASK_STATUS=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/api/task 2>/dev/null)
    
    if [ "$TASK_STATUS" == "200" ] || [ "$TASK_STATUS" == "401" ]; then
        echo -e "${GREEN}$CHECK 応答あり (HTTP $TASK_STATUS)${NC}"
    else
        echo -e "${RED}$CROSS エラー (HTTP $TASK_STATUS)${NC}"
    fi
    echo ""
    
    # 5. ディスク使用量
    echo "${CYAN}💾 ディスク使用量${NC}"
    echo "  Docker イメージ:"
    docker images --format "table {{.Repository}}\t{{.Tag}}\t{{.Size}}" | grep -E "(betatasker|postgres)" | head -5 | sed 's/^/    /'
    echo ""
    
    echo "  Docker ボリューム:"
    docker volume ls --format "table {{.Name}}\t{{.Driver}}" | grep betatasker | sed 's/^/    /'
    echo ""
    
    # 6. 最新のエラーログ
    echo "${CYAN}⚠️  最新のエラー（もしあれば）${NC}"
    
    # バックエンドのエラー
    BACKEND_ERRORS=$(docker logs backender 2>&1 | grep -i error | tail -3)
    if [ ! -z "$BACKEND_ERRORS" ]; then
        echo "  Backend:"
        echo "$BACKEND_ERRORS" | sed 's/^/    /'
    fi
    
    # フロントエンドのエラー
    FRONTEND_ERRORS=$(docker logs fronter 2>&1 | grep -i error | tail -3)
    if [ ! -z "$FRONTEND_ERRORS" ]; then
        echo "  Frontend:"
        echo "$FRONTEND_ERRORS" | sed 's/^/    /'
    fi
    
    if [ -z "$BACKEND_ERRORS" ] && [ -z "$FRONTEND_ERRORS" ]; then
        echo -e "  ${GREEN}エラーなし${NC}"
    fi
    echo ""
    
    # サマリー
    echo "================================================"
    echo -e "${GREEN}Health check completed${NC}"
    
    if [ "$WATCH_MODE" == "true" ]; then
        echo "Refreshing in 5 seconds... (Ctrl+C to exit)"
    fi
}

# メイン処理
if [ "$WATCH_MODE" == "true" ]; then
    # ウォッチモード
    while true; do
        perform_check
        sleep 5
    done
else
    # 単発実行
    perform_check
    echo ""
    echo "継続的な監視: $0 --watch"
fi