# betaTasker Scripts

このディレクトリには、開発とデプロイメントを支援するスクリプトが含まれています。

## 🚀 Quick Start

```bash
# スクリプトに実行権限を付与
chmod +x scripts/*.sh

# 開発環境を素早くテスト
./scripts/quick-test.sh

# ヘルスチェックを実行
./scripts/health-check.sh

# 本番環境をローカルでテスト
./scripts/test-production-local.sh
```

## 📋 スクリプト一覧

### quick-test.sh
開発環境を素早く起動してテストします。

**機能:**
- Docker環境の起動
- サービスの確認
- APIエンドポイントのテスト
- 自動的にユーザー登録・ログイン・タスク作成

**使用方法:**
```bash
./scripts/quick-test.sh
```

### health-check.sh
全サービスの健全性を確認します。

**機能:**
- Docker環境の確認
- コンテナ状態の監視
- サービス応答の確認
- エラーログの表示
- リソース使用状況の表示

**使用方法:**
```bash
# 単発実行
./scripts/health-check.sh

# 継続的な監視（5秒ごとに更新）
./scripts/health-check.sh --watch
```

### test-production-local.sh
ローカル環境で本番ビルドをテストします。

**機能:**
- 本番用Dockerイメージのビルド
- 本番環境の起動
- 全サービスのヘルスチェック
- APIエンドポイントのテスト
- 詳細なログ表示

**使用方法:**
```bash
./scripts/test-production-local.sh
```

### deploy.sh
本番環境へのデプロイを実行します（サーバー設定が必要）。

**使用方法:**
```bash
./scripts/deploy.sh
```

## 🔧 トラブルシューティング

### スクリプトが実行できない場合
```bash
chmod +x scripts/*.sh
```

### Dockerが起動していない場合
```bash
# Dockerの起動を確認
docker version

# Docker Desktopを起動（macOS）
open -a Docker
```

### ポートが使用中の場合
```bash
# 使用中のポートを確認
lsof -i :3000  # Frontend
lsof -i :8080  # Backend
lsof -i :5432  # PostgreSQL

# プロセスを停止
kill -9 <PID>
```

### コンテナをクリーンアップする場合
```bash
# 全てのコンテナを停止・削除
docker-compose down -v
docker system prune -f
```

## 📊 環境変数

スクリプトで使用される主な環境変数：

| 変数名 | デフォルト値 | 説明 |
|--------|------------|------|
| API_BASE_URL | http://localhost:8080 | バックエンドAPIのURL |
| FRONTEND_URL | http://localhost:3000 | フロントエンドのURL |
| DB_HOST | localhost | データベースホスト |
| DB_PORT | 5432 | データベースポート |

## 🎯 使用例

### 開発環境の完全なテスト
```bash
# 1. 環境を起動してテスト
./scripts/quick-test.sh

# 2. ヘルスチェックで確認
./scripts/health-check.sh

# 3. APIテストを実行
./test_error_codes.sh

# 4. ログを監視
docker-compose logs -f
```

### 本番ビルドのテスト
```bash
# 1. 本番環境をビルド・起動
./scripts/test-production-local.sh

# 2. ヘルスチェックで監視
./scripts/health-check.sh --watch

# 3. ブラウザで確認
open http://localhost:3000
```

## 📝 メンテナンス

### ログの確認
```bash
# 全サービスのログ
docker-compose logs -f

# 特定サービスのログ
docker-compose logs -f backender
docker-compose logs -f fronter
```

### データベースへのアクセス
```bash
# PostgreSQLに接続
docker exec -it db psql -U dbgodotask -d dbgodotask

# データベースのバックアップ
docker exec db pg_dump -U dbgodotask dbgodotask > backup.sql
```

### イメージの更新
```bash
# 最新のイメージをビルド
docker-compose build --no-cache

# 本番用イメージをビルド
docker-compose -f docker-compose.prod.yml build
```