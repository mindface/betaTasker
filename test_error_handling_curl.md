# エラーハンドリング確認用curlコマンド集（認証対応版）

## 0. 認証トークンの取得

### ユーザー登録
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "testpass123",
    "email": "test@example.com"
  }' \
  -v

# 期待されるレスポンス:
# {
#   "success": true,
#   "message": "ユーザー登録が完了しました",
#   "data": {
#     "user": {
#       "id": 1,
#       "username": "testuser",
#       "email": "test@example.com"
#     }
#   }
# }
# HTTPステータス: 201
```

### ログインしてトークン取得
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "testpass123"
  }' \
  -v

# 期待されるレスポンス:
# {
#   "success": true,
#   "message": "ログインに成功しました",
#   "data": {
#     "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
#     "user": {
#       "id": 1,
#       "username": "testuser"
#     }
#   }
# }
# HTTPステータス: 200

# トークンを環境変数に保存
export TOKEN="取得したトークンをここに設定"
```

## 1. バリデーションエラーの確認（認証付き）

### VAL_002: 必須フィールド不足
```bash
# タイトルなしでタスク作成（エラーになるはず）
curl -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"description": "タイトルなしタスク"}' \
  -v

# 期待されるレスポンス:
# {
#   "code": "VAL_002",
#   "message": "必須項目が入力されていません",
#   "detail": "タイトルは必須項目です"
# }
# HTTPステータス: 400
```

### VAL_001: 不正な入力形式
```bash
# 不正なJSONを送信
curl -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d 'invalid json format' \
  -v

# 期待されるレスポンス:
# {
#   "code": "VAL_001",
#   "message": "入力値が無効です",
#   "detail": "invalid character 'i' looking for beginning of value"
# }
# HTTPステータス: 400
```

### VAL_004: 重複エントリ（実装依存）
```bash
# 同じタイトルのタスクを2回作成
curl -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"title": "ユニークなタイトル", "description": "テスト"}' \
  -v

# 2回目（重複エラーになる可能性）
curl -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"title": "ユニークなタイトル", "description": "テスト2"}' \
  -v
```

## 2. 認証エラーの確認

### AUTH_001: 認証情報なし
```bash
# Authorizationヘッダーなしでアクセス
curl -X GET http://localhost:8080/api/task \
  -v

# 期待されるレスポンス:
# {
#   "code": "AUTH_001",
#   "message": "認証が必要です",
#   "detail": "このリソースにアクセスするには認証が必要です"
# }
# HTTPステータス: 401
```

### AUTH_002: 無効なトークン
```bash
# 不正なトークンでアクセス
curl -X GET http://localhost:8080/api/task \
  -H "Authorization: Bearer invalid_token_12345" \
  -v

# 期待されるレスポンス:
# {
#   "code": "AUTH_002",
#   "message": "無効な認証情報です",
#   "detail": "トークンが無効または期限切れです"
# }
# HTTPステータス: 401
```

### AUTH_003: 権限不足
```bash
# 他のユーザーのリソースにアクセス（実装依存）
curl -X DELETE http://localhost:8080/api/task/1 \
  -H "Authorization: Bearer $TOKEN" \
  -v

# 期待されるレスポンス（権限がない場合）:
# {
#   "code": "AUTH_003",
#   "message": "アクセス権限がありません",
#   "detail": "このリソースへのアクセス権限がありません"
# }
# HTTPステータス: 403
```

## 3. リソースエラーの確認（認証付き）

### RES_001: リソース未発見
```bash
# 存在しないタスクIDを取得
curl -X GET http://localhost:8080/api/task/999999 \
  -H "Authorization: Bearer $TOKEN" \
  -v

# 期待されるレスポンス:
# {
#   "code": "RES_001",
#   "message": "リソースが見つかりません",
#   "detail": "指定されたタスクが見つかりません"
# }
# HTTPステータス: 404
```

### メモリーの未発見エラー
```bash
# 存在しないメモリーIDを取得
curl -X GET http://localhost:8080/api/memory/999999 \
  -H "Authorization: Bearer $TOKEN" \
  -v

# 期待されるレスポンス:
# {
#   "code": "RES_001",
#   "message": "リソースが見つかりません",
#   "detail": "指定されたメモリーが見つかりません"
# }
# HTTPステータス: 404
```

## 4. 正常系の確認（認証付き）

### タスク作成成功
```bash
curl -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "正常なタスク",
    "description": "エラーコードテスト用",
    "status": "pending"
  }' \
  -v

# 期待されるレスポンス:
# {
#   "success": true,
#   "message": "タスクが正常に作成されました",
#   "data": {
#     "task": {
#       "id": 1,
#       "title": "正常なタスク",
#       "description": "エラーコードテスト用",
#       "status": "pending"
#     }
#   }
# }
# HTTPステータス: 200
```

### メモリー作成成功
```bash
curl -X POST http://localhost:8080/api/memory \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "テストメモリー",
    "content": "メモリー内容",
    "tags": ["test", "error-handling"]
  }' \
  -v

# 期待されるレスポンス:
# {
#   "success": true,
#   "message": "メモリーが正常に作成されました",
#   "data": {
#     "memory": {...}
#   }
# }
# HTTPステータス: 200
```

### ユーザープロファイル取得
```bash
curl -X GET http://localhost:8080/api/user/profile \
  -H "Authorization: Bearer $TOKEN" \
  -v

# 期待されるレスポンス:
# {
#   "success": true,
#   "data": {
#     "user": {
#       "id": 1,
#       "username": "testuser",
#       "email": "test@example.com"
#     }
#   }
# }
# HTTPステータス: 200
```

## 5. 更新・削除操作のエラー確認（認証付き）

### 存在しないリソースの更新
```bash
curl -X PUT http://localhost:8080/api/task/999999 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "id": 999999,
    "title": "更新タイトル",
    "description": "更新内容"
  }' \
  -v

# 期待されるレスポンス: RES_001エラー
```

### 存在しないリソースの削除
```bash
curl -X DELETE http://localhost:8080/api/task/999999 \
  -H "Authorization: Bearer $TOKEN" \
  -v

# 期待されるレスポンス: RES_001エラー
```

## 6. ログアウト

```bash
curl -X POST http://localhost:8080/api/logout \
  -H "Authorization: Bearer $TOKEN" \
  -v

# 期待されるレスポンス:
# {
#   "success": true,
#   "message": "ログアウトしました"
# }
# HTTPステータス: 200
```

## 7. レスポンス形式の確認

### jqを使った整形表示
```bash
# エラーレスポンスを整形して表示
curl -s -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{}' | jq '.'

# エラーコードだけを抽出
curl -s -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{}' | jq '.code'

# HTTPステータスコードとレスポンスボディを同時に確認
curl -w '\nHTTP Status: %{http_code}\n' \
  -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{}'
```

## 8. バッチテスト実行（認証対応版）

### 全エラーパターンを順次実行
```bash
#!/bin/bash

API_URL="http://localhost:8080"

echo "=== エラーハンドリングテスト開始 ==="

# ログインしてトークンを取得
echo -e "\n[SETUP] ログイン"
TOKEN=$(curl -s -X POST $API_URL/api/login \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "testpass123"}' | jq -r '.data.token')

echo "Token: ${TOKEN:0:20}..."

# テスト1: 必須フィールド不足
echo -e "\n[TEST 1] 必須フィールド不足"
curl -s -X POST $API_URL/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"description": "no title"}' | jq '.'

# テスト2: 不正なJSON
echo -e "\n[TEST 2] 不正なJSON"
curl -s -X POST $API_URL/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d 'bad json' | jq '.'

# テスト3: 認証なしアクセス
echo -e "\n[TEST 3] 認証なしアクセス"
curl -s -X GET $API_URL/api/task | jq '.'

# テスト4: 無効なトークン
echo -e "\n[TEST 4] 無効なトークン"
curl -s -X GET $API_URL/api/task \
  -H "Authorization: Bearer invalid_token" | jq '.'

# テスト5: 存在しないリソース
echo -e "\n[TEST 5] 存在しないリソース"
curl -s -X GET $API_URL/api/task/999999 \
  -H "Authorization: Bearer $TOKEN" | jq '.'

# テスト6: 正常系
echo -e "\n[TEST 6] 正常系"
curl -s -X POST $API_URL/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"title": "正常テスト", "description": "OK"}' | jq '.'

echo -e "\n=== テスト完了 ==="
```

## 9. ヘッダー情報を含む詳細確認

```bash
# 全ヘッダー情報を表示
curl -i -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"description": "test"}'

# デバッグ情報を含む詳細表示
curl -v -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"description": "test"}' 2>&1 | grep -E "(< HTTP|< |{|})"
```

## 使用上の注意

1. **サーバー起動確認**: テスト前に`http://localhost:8080`でサーバーが起動していることを確認
2. **jqインストール**: JSON整形に`jq`を使用（`brew install jq`または`apt-get install jq`）
3. **ポート番号**: 環境に応じてポート番号を変更（8080 → 実際のポート）
4. **トークンの有効期限**: JWTトークンには有効期限があるため、期限切れの場合は再ログインが必要
5. **環境変数の設定**: `export TOKEN="your_token_here"`でトークンを環境変数に設定

## トラブルシューティング

### Connection refused エラー
```bash
# サーバーが起動しているか確認
lsof -i :8080

# サーバー起動
cd backer/godotask && go run main.go
```

### CORS エラー
```bash
# Originヘッダーを追加
curl -X POST http://localhost:8080/api/task \
  -H "Content-Type: application/json" \
  -H "Origin: http://localhost:3000" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"title": "test"}'
```

### トークン期限切れエラー
```bash
# 再ログインしてトークンを再取得
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "testpass123"}'
```

### トークンのデコード（デバッグ用）
```bash
# JWTトークンの内容を確認（jwt.ioまたはjqを使用）
echo $TOKEN | cut -d. -f2 | base64 --decode 2>/dev/null | jq '.'
```