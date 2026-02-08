# BetaTasker

タスク管理とメモリ支援機能を提供するWebアプリケーション

## アーキテクチャ

- **フロントエンド**: Next.js + TypeScript + Redux
- **バックエンド**: Go (Gin) + PostgreSQL
- **インフラ**: Docker + Docker Compose + Nginx

## CI/CDパイプライン

### 概要

このプロジェクトはGitHub Actionsを使用したCI/CDパイプラインを採用しています。

### パイプラインの流れ

1. **コードプッシュ/PR作成** → パイプライン開始
2. **バックエンドテスト・ビルド**
   - Goテスト実行
   - コード品質チェック（go vet）
3. **フロントエンドテスト・ビルド**
   - 依存関係インストール
   - リンティング
   - 型チェック
   - Next.jsビルド
4. **統合テスト** (対応予定)
   - PostgreSQLサービスを使用した統合テスト
5. **セキュリティスキャン** (対応予定)
   - Trivyによる脆弱性スキャン
6. **デプロイ**（mainブランチのみ）[費用が確保されてから]
   - 本番環境へのデプロイ

### 環境

- **開発環境**: `develop`ブランチ
- **本番環境**: `main`ブランチ

## ローカル開発

### 前提条件

- Docker & Docker Compose
- Go 1.21+
- Node.js 18+
- Yarn

### セットアップ

1. リポジトリのクローン
```bash
git clone <repository-url>
cd betaTasker
```

2. 環境変数の設定
```bash
cp env.example .env
# .envファイルを編集して必要な値を設定
```

3. 開発環境の起動
```bash
# バックエンド
cd backer
docker-compose up -d

# フロントエンド
cd fronter
yarn install
yarn dev
```

## デプロイ

### 本番環境へのデプロイ

1. コードを`main`ブランチにプッシュ
2. GitHub Actionsが自動的にデプロイを実行
3. デプロイスクリプトの実行（手動の場合）
```bash
chmod +x scripts/deploy.sh
./scripts/deploy.sh
```

### 環境変数

本番環境では以下の環境変数を設定してください：

- `DATABASE_URL`: PostgreSQL接続文字列
- `JWT_SECRET`: JWT署名用のシークレット
- `DB_PASSWORD`: データベースパスワード
- `REGISTRY`: コンテナレジストリ（デフォルト: ghcr.io）
- `IMAGE_NAME`: イメージ名（デフォルト: betaTasker）
- `TAG`: イメージタグ（デフォルト: latest）

## セキュリティ

- セキュリティヘッダーの設定
- レート制限の実装
- 非rootユーザーでのコンテナ実行
- 定期的な脆弱性スキャン

## 監視・ログ

- ヘルスチェックエンドポイント: `/health`
- アプリケーションログは標準出力に出力
- Nginxアクセスログ

## トラブルシューティング

### よくある問題

1. **ポート競合**
   - 既存のサービスが使用しているポートを確認
   - `docker-compose down`でコンテナを停止

2. **データベース接続エラー**
   - 環境変数の設定を確認
   - データベースコンテナの起動状態を確認

3. **ビルドエラー**
   - 依存関係のバージョンを確認
   - キャッシュのクリア: `docker system prune`

## 貢献

1. フィーチャーブランチを作成
2. 変更をコミット
3. プルリクエストを作成
4. CI/CDパイプラインの通過を確認
5. レビュー後にマージ 