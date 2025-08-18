# Backend Specification Document - betaTasker

## 概要
betaTaskerバックエンドは、タスク管理、学習記録、評価機能を提供するGo言語ベースのREST APIアプリケーションです。

## 技術スタック

### コア技術
- **言語**: Go (v1.21.0+)
- **Webフレームワーク**: Gin (v1.7.4)
- **データベース**: PostgreSQL (v16)
- **ORM**: GORM (v1.30.0)
- **認証**: JWT + bcrypt
- **AI統合**: Ollama (RAG機能)
- **コンテナ化**: Docker + Docker Compose
- **テスト**: Testify

## アーキテクチャ

### レイヤー構造
```
Controller層 (HTTPリクエスト/レスポンス処理)
    ↓
Service層 (ビジネスロジック)
    ↓
Repository層 (データアクセス)
    ↓
Model層 (データベースエンティティ)
```

### ディレクトリ構造
```
godotask/
├── main.go                 # エントリーポイント
├── controller/             # HTTPハンドラー
│   ├── task/              # タスク管理
│   ├── memory/            # メモリー管理
│   ├── assessment/        # 評価管理
│   ├── book/              # 書籍管理
│   ├── user/              # ユーザー認証
│   └── top/               # トップページ
├── service/               # ビジネスロジック
├── repository/            # データアクセス層
├── model/                 # データモデル
├── server/                # ルーター設定
├── rag/                   # AI統合
├── static/                # 静的ファイル
├── view/                  # HTMLテンプレート
└── seed/                  # データベースシード
```

## データベーススキーマ

### Users テーブル
| カラム名 | 型 | 制約 | 説明 |
|---------|-----|------|------|
| id | SERIAL | PRIMARY KEY | ユーザーID |
| username | VARCHAR | UNIQUE, NOT NULL | ユーザー名 |
| email | VARCHAR | UNIQUE, NOT NULL | メールアドレス |
| password_hash | VARCHAR | NOT NULL | パスワードハッシュ |
| role | VARCHAR | DEFAULT 'user' | ユーザー権限 |
| is_active | BOOLEAN | DEFAULT true | アクティブ状態 |
| factor | TEXT | | 因子情報 |
| process | TEXT | | プロセス情報 |
| evaluation_axis | TEXT | | 評価軸 |
| information_amount | TEXT | | 情報量 |
| created_at | TIMESTAMP | | 作成日時 |
| updated_at | TIMESTAMP | | 更新日時 |

### Tasks テーブル
| カラム名 | 型 | 制約 | 説明 |
|---------|-----|------|------|
| id | SERIAL | PRIMARY KEY | タスクID |
| user_id | INTEGER | FOREIGN KEY | ユーザーID |
| memory_id | INTEGER | FOREIGN KEY (NULL可) | メモリーID |
| title | VARCHAR | NOT NULL | タイトル |
| description | TEXT | | 説明 |
| date | TIMESTAMP | NULL可 | 期限日 |
| status | VARCHAR | | ステータス (todo/in_progress/completed) |
| priority | INTEGER | | 優先度 (1-5) |
| created_at | TIMESTAMP | | 作成日時 |
| updated_at | TIMESTAMP | | 更新日時 |

### Memory テーブル
| カラム名 | 型 | 制約 | 説明 |
|---------|-----|------|------|
| id | SERIAL | PRIMARY KEY | メモリーID |
| user_id | INTEGER | FOREIGN KEY | ユーザーID |
| source_type | VARCHAR | DEFAULT 'book' | ソースタイプ |
| title | VARCHAR | NOT NULL | タイトル |
| author | VARCHAR | | 著者 |
| notes | TEXT | | ノート |
| tags | VARCHAR | | タグ (カンマ区切り) |
| read_status | VARCHAR | DEFAULT 'unread' | 読了状態 |
| read_date | TIMESTAMP | NULL可 | 読了日 |
| factor | TEXT | | 因子情報 |
| process | TEXT | | プロセス情報 |
| evaluation_axis | TEXT | | 評価軸 |
| information_amount | TEXT | | 情報量 |
| created_at | TIMESTAMP | | 作成日時 |
| updated_at | TIMESTAMP | | 更新日時 |

### Assessment テーブル
| カラム名 | 型 | 制約 | 説明 |
|---------|-----|------|------|
| id | SERIAL | PRIMARY KEY | 評価ID |
| task_id | INTEGER | FOREIGN KEY | タスクID |
| user_id | INTEGER | FOREIGN KEY | ユーザーID |
| effectiveness_score | INTEGER | | 効果スコア (0-100) |
| effort_score | INTEGER | | 努力スコア (0-100) |
| impact_score | INTEGER | | インパクトスコア (0-100) |
| qualitative_feedback | TEXT | | 定性的フィードバック |
| created_at | TIMESTAMP | | 作成日時 |
| updated_at | TIMESTAMP | | 更新日時 |

### MemoryContext テーブル
| カラム名 | 型 | 制約 | 説明 |
|---------|-----|------|------|
| id | SERIAL | PRIMARY KEY | コンテキストID |
| user_id | INTEGER | FOREIGN KEY | ユーザーID |
| task_id | INTEGER | FOREIGN KEY | タスクID |
| level | INTEGER | | レベル |
| work_target | TEXT | | 作業対象 |
| machine | TEXT | | 機械情報 |
| material_spec | TEXT | | 材料仕様 |
| change_factor | TEXT | | 変更要因 |
| goal | TEXT | | 目標 |
| created_at | TIMESTAMP | | 作成日時 |

### Book テーブル
| カラム名 | 型 | 制約 | 説明 |
|---------|-----|------|------|
| id | SERIAL | PRIMARY KEY | 書籍ID |
| title | VARCHAR | | タイトル |
| name | VARCHAR | | 名前 |
| text | TEXT | | テキスト |
| disc | TEXT | | 説明 |
| img_path | VARCHAR | | 画像パス |
| status | VARCHAR | | ステータス |

## API エンドポイント

### 認証 API

#### POST /api/login
ユーザーログイン
```json
Request:
{
  "email": "user@example.com",
  "password": "password123"
}

Response:
{
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user_id": 1
}
```

#### POST /api/register
新規ユーザー登録
```json
Request:
{
  "username": "newuser",
  "email": "new@example.com",
  "password": "password123"
}

Response:
{
  "message": "User registered successfully",
  "user_id": 2
}
```

#### POST /api/logout
ログアウト（JWTトークンをブラックリストに追加）
```json
Headers: Authorization: Bearer <token>

Response:
{
  "message": "Logged out successfully"
}
```

#### GET /api/user/profile
ユーザープロファイル取得（認証必須）
```json
Headers: Authorization: Bearer <token>

Response:
{
  "id": 1,
  "username": "user",
  "email": "user@example.com",
  "role": "user"
}
```

### タスク管理 API

#### GET /api/task
タスク一覧取得
```json
Headers: Authorization: Bearer <token>

Response:
[
  {
    "id": 1,
    "user_id": 1,
    "title": "タスク1",
    "description": "説明",
    "status": "todo",
    "priority": 3,
    "created_at": "2024-01-01T00:00:00Z"
  }
]
```

#### POST /api/task
タスク作成
```json
Headers: Authorization: Bearer <token>

Request:
{
  "title": "新規タスク",
  "description": "タスクの説明",
  "status": "todo",
  "priority": 3,
  "memory_id": 1
}

Response:
{
  "id": 2,
  "message": "Task created successfully"
}
```

#### PUT /api/task
タスク更新
```json
Headers: Authorization: Bearer <token>

Request:
{
  "id": 1,
  "title": "更新されたタスク",
  "status": "in_progress"
}

Response:
{
  "message": "Task updated successfully"
}
```

#### DELETE /api/task
タスク削除
```json
Headers: Authorization: Bearer <token>

Request:
{
  "id": 1
}

Response:
{
  "message": "Task deleted successfully"
}
```

### メモリー管理 API

#### GET /api/memory
メモリー一覧取得
```json
Headers: Authorization: Bearer <token>

Response:
[
  {
    "id": 1,
    "user_id": 1,
    "title": "学習メモ",
    "notes": "内容",
    "tags": "tag1,tag2",
    "read_status": "unread"
  }
]
```

#### POST /api/memory
メモリー作成
```json
Headers: Authorization: Bearer <token>

Request:
{
  "title": "新規メモリー",
  "notes": "メモ内容",
  "tags": "学習,プログラミング",
  "source_type": "book"
}

Response:
{
  "id": 2,
  "message": "Memory created successfully"
}
```

#### GET /api/memory/:id
特定メモリー取得
```json
Headers: Authorization: Bearer <token>

Response:
{
  "id": 1,
  "title": "メモリータイトル",
  "notes": "詳細内容",
  "tags": "tag1,tag2"
}
```

#### PUT /api/memory
メモリー更新
```json
Headers: Authorization: Bearer <token>

Request:
{
  "id": 1,
  "title": "更新されたメモリー",
  "read_status": "read"
}

Response:
{
  "message": "Memory updated successfully"
}
```

#### DELETE /api/memory
メモリー削除
```json
Headers: Authorization: Bearer <token>

Request:
{
  "id": 1
}

Response:
{
  "message": "Memory deleted successfully"
}
```

### 評価 API

#### GET /api/assessment
評価一覧取得
```json
Headers: Authorization: Bearer <token>

Response:
[
  {
    "id": 1,
    "task_id": 1,
    "user_id": 1,
    "effectiveness_score": 80,
    "effort_score": 70,
    "impact_score": 90
  }
]
```

#### POST /api/assessment
評価作成
```json
Headers: Authorization: Bearer <token>

Request:
{
  "task_id": 1,
  "effectiveness_score": 85,
  "effort_score": 75,
  "impact_score": 90,
  "qualitative_feedback": "良い成果"
}

Response:
{
  "id": 2,
  "message": "Assessment created successfully"
}
```

#### POST /api/assessmentsForTaskUser
特定タスク・ユーザーの評価取得
```json
Headers: Authorization: Bearer <token>

Request:
{
  "userId": 1,
  "taskId": 1
}

Response:
[
  {
    "id": 1,
    "effectiveness_score": 80,
    "effort_score": 70,
    "impact_score": 90
  }
]
```

### 書籍管理 API

#### GET /api/book
書籍一覧取得
```json
Response:
[
  {
    "id": 1,
    "title": "書籍タイトル",
    "name": "著者名",
    "disc": "説明"
  }
]
```

#### POST /api/book
書籍作成
```json
Request:
{
  "title": "新規書籍",
  "name": "著者",
  "disc": "説明文"
}

Response:
{
  "id": 2,
  "message": "Book created successfully"
}
```

#### POST /api/file
ファイルアップロード
```
Content-Type: multipart/form-data
Field: file (画像ファイル)

Response:
{
  "path": "/images/uploaded_file.jpg"
}
```

## 認証とセキュリティ

### JWT認証
- **秘密鍵**: 環境変数で管理（現在はハードコード）
- **トークン有効期限**: 24時間
- **ブラックリスト**: メモリ内管理（ログアウト時）

### パスワード管理
- **ハッシュ化**: bcrypt (デフォルトコスト)
- **最小文字数**: 実装による

### CORS設定
```go
AllowOrigins:     []string{"http://localhost:3000"}
AllowMethods:     []string{"POST", "HEAD", "PATCH", "OPTIONS", "GET", "PUT", "DELETE"}
AllowHeaders:     []string{"Content-Type", "Authorization"}
AllowCredentials: true
```

## データベース設定

### PostgreSQL接続情報
- **ホスト**: dbgodotask (Dockerコンテナ)
- **ポート**: 5432
- **データベース名**: dbgodotask
- **ユーザー**: dbgodotask
- **パスワード**: dbgodotask

### 自動マイグレーション
起動時に以下のモデルが自動マイグレーションされます：
- User
- Task
- Memory
- Assessment
- MemoryContext
- TechnicalFactor
- KnowledgeTransformation
- Book

## AI統合 (RAG機能)

### Ollama設定
- **モデル**: llama3.1:8b
- **エンドポイント**: http://host.docker.internal:11434/api/generate
- **用途**: RAG（Retrieval Augmented Generation）

## Docker構成

### docker-compose.yml
```yaml
services:
  godotask:
    build:
      context: .
      dockerfile: godotask/DockerFile.compose
    ports:
      - "8080:8080"
    volumes:
      - ./godotask:/usr/local/go/godotask
    environment:
      - DATABASE_DSN=...
    depends_on:
      - dbgodotask

  dbgodotask:
    image: postgres:16
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_DB=dbgodotask
      - POSTGRES_USER=dbgodotask
      - POSTGRES_PASSWORD=dbgodotask
    volumes:
      - db-data:/var/lib/postgresql/data
```

## テスト戦略

### テストフレームワーク
- **Testify**: アサーションとモック
- **カバレッジ**: Controller, Service, Repository層

### テスト実行
```bash
go test ./...
go test -v ./controller/task
go test -cover ./service/...
```

## 環境変数

### 必須環境変数
```bash
DATABASE_DSN=host=dbgodotask user=dbgodotask password=dbgodotask dbname=dbgodotask port=5432 sslmode=disable
JWT_SECRET=your_secret_key  # 本番環境では変更必須
```

### オプション環境変数
```bash
GIN_MODE=release  # 本番環境用
PORT=8080         # APIポート
```

## デプロイメント

### 開発環境
```bash
docker-compose up -d
```

### 本番環境への考慮事項
1. JWT秘密鍵を環境変数化
2. データベース認証情報の外部化
3. HTTPSの実装
4. レート制限の追加
5. ロギングシステムの強化
6. モニタリングの実装

## パフォーマンス最適化

### データベース
- インデックスの追加（user_id, task_id等）
- N+1問題の解決（Eager Loading）
- コネクションプーリング

### API
- ページネーション実装
- キャッシング戦略
- 非同期処理の活用

## セキュリティチェックリスト

- [ ] JWT秘密鍵の環境変数化
- [ ] SQLインジェクション対策（GORM使用）
- [ ] XSS対策
- [ ] CSRF対策
- [ ] レート制限
- [ ] 入力値検証
- [ ] エラーメッセージの適切な処理
- [ ] ログの適切な管理

## 今後の改善点

1. **認証強化**
   - リフレッシュトークンの実装
   - 2要素認証の追加

2. **API改善**
   - GraphQL対応
   - WebSocket対応
   - APIバージョニング

3. **監視・ログ**
   - 構造化ログの実装
   - メトリクス収集
   - 分散トレーシング

4. **テスト強化**
   - E2Eテストの追加
   - 負荷テストの実装
   - カバレッジ向上

---

最終更新日: 2024年
バージョン: 1.0.0