#!/bin/bash

set -e

# 環境変数の読み込み
source .env

# ログ関数
log() {
    echo "[$(date +'%Y-%m-%d %H:%M:%S')] $1"
}

# エラーハンドリング
error_exit() {
    log "ERROR: $1"
    exit 1
}

# ヘルスチェック関数
health_check() {
    local url=$1
    local max_attempts=30
    local attempt=1

    log "ヘルスチェック開始: $url"
    
    while [ $attempt -le $max_attempts ]; do
        if curl -f -s "$url/health" > /dev/null; then
            log "ヘルスチェック成功: $url"
            return 0
        fi
        
        log "ヘルスチェック試行 $attempt/$max_attempts: $url"
        sleep 10
        attempt=$((attempt + 1))
    done
    
    error_exit "ヘルスチェック失敗: $url"
}

# メイン処理
main() {
    log "デプロイ開始"
    
    # 必要なファイルの存在確認
    if [ ! -f "docker-compose.prod.yml" ]; then
        error_exit "docker-compose.prod.yml が見つかりません"
    fi
    
    if [ ! -f ".env" ]; then
        error_exit ".env ファイルが見つかりません"
    fi
    
    # 古いコンテナの停止
    log "古いコンテナを停止中..."
    docker-compose -f docker-compose.prod.yml down || true
    
    # 新しいイメージのプル
    log "新しいイメージをプル中..."
    docker-compose -f docker-compose.prod.yml pull
    
    # コンテナの起動
    log "コンテナを起動中..."
    docker-compose -f docker-compose.prod.yml up -d
    
    # ヘルスチェック
    health_check "http://localhost"
    
    # 古いイメージのクリーンアップ
    log "古いイメージをクリーンアップ中..."
    docker image prune -f
    
    log "デプロイ完了"
}

# スクリプト実行
main "$@" 