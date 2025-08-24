#!/bin/bash

# エラーコード実装のテスト用curlコマンド集
# 使用方法: bash test_error_codes.sh

API_BASE_URL="http://localhost:8080"
TOKEN=""

echo "================================"
echo "エラーコード実装テストケース"
echo "================================"
echo ""

# ログインしてトークンを取得する関数
echo "認証トークンを取得中..."

# まずユーザー登録を試みる（既に存在する場合はエラーになるが無視）
curl -s -X POST $API_BASE_URL/api/register \
    -H "Content-Type: application/json" \
    -d '{
        "username": "testuser",
        "password": "testpass123",
        "email": "test@example.com",
        "role": "user"
    }' > /dev/null 2>&1

# ログインしてトークンを取得
RESPONSE=$(curl -s -X POST $API_BASE_URL/api/login \
    -H "Content-Type: application/json" \
    -d '{
        "username": "testuser",
        "password": "testpass123",
        "email": "test@example.com"
    }')

# トークンを抽出（レスポンス形式に応じて調整が必要）
TOKEN=$(echo $RESPONSE | jq -r '.token // .data.token // .access_token // empty')

if [ -z "$TOKEN" ]; then
    echo "警告: トークンの取得に失敗しました。認証なしでテストを続行します。"
    echo "レスポンス: $RESPONSE"
else
    echo "トークン取得成功: ${TOKEN:0:20}..."
fi
echo "--------------------------------"

# 1. バリデーションエラー：必須フィールド不足
echo "1. 必須フィールド不足のテスト (VAL_002)"
echo "期待結果: エラーコード VAL_002 が返される"
echo "実行コマンド:"
echo 'curl -X POST '$API_BASE_URL'/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer '$TOKEN'" \
  -d "{\"description\": \"タイトルなしタスク\"}"'
echo ""
curl -X POST $API_BASE_URL/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"description": "タイトルなしタスク"}' \
  -w "\nHTTPステータス: %{http_code}\n"
echo ""
echo "--------------------------------"

# 2. バリデーションエラー：不正な入力形式
echo "2. 不正な入力形式のテスト (VAL_001)"
echo "期待結果: エラーコード VAL_001 が返される"
echo "実行コマンド:"
echo 'curl -X POST '$API_BASE_URL'/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer '$TOKEN'" \
  -d "invalid json"'
echo ""
curl -X POST $API_BASE_URL/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d 'invalid json' \
  -w "\nHTTPステータス: %{http_code}\n"
echo ""
echo "--------------------------------"

# 3. リソース未発見エラー
echo "3. 存在しないリソースの取得テスト (RES_001)"
echo "期待結果: エラーコード RES_001 が返される"
echo "実行コマンド:"
echo 'curl -X GET '$API_BASE_URL'/api/task/999999 \
  -H "Authorization: Bearer '$TOKEN'"'
echo ""
curl -X GET $API_BASE_URL/api/task/999999 \
  -H "Authorization: Bearer $TOKEN" \
  -w "\nHTTPステータス: %{http_code}\n"
echo ""
echo "--------------------------------"

# 4. 正常なタスク作成
echo "4. 正常なタスク作成のテスト"
echo "期待結果: 成功レスポンスが返される"
echo "実行コマンド:"
echo 'curl -X POST '$API_BASE_URL'/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer '$TOKEN'" \
  -d "{
    \"title\": \"テストタスク\",
    \"description\": \"エラーコードテスト用タスク\",
    \"status\": \"pending\"
  }"'
echo ""
curl -X POST $API_BASE_URL/api/task \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "テストタスク",
    "description": "エラーコードテスト用タスク",
    "status": "pending"
  }' \
  -w "\nHTTPステータス: %{http_code}\n"
echo ""
echo "--------------------------------"

# 5. 認証エラーのテスト（トークンなし）
echo "5. 認証なしでの保護されたエンドポイントへのアクセステスト (AUTH_001)"
echo "期待結果: エラーコード AUTH_001 または AUTH_003 が返される"
echo "実行コマンド:"
echo 'curl -X GET '$API_BASE_URL'/api/user/profile'
echo ""
curl -X GET $API_BASE_URL/api/user/profile \
  -w "\nHTTPステータス: %{http_code}\n"
echo ""
echo "--------------------------------"

# 6. 無効なトークンでのアクセステスト
echo "6. 無効なトークンでのアクセステスト (AUTH_002)"
echo "期待結果: エラーコード AUTH_002 が返される"
echo "実行コマンド:"
echo 'curl -X GET '$API_BASE_URL'/api/user/profile \
  -H "Authorization: Bearer invalid_token_12345"'
echo ""
curl -X GET $API_BASE_URL/api/user/profile \
  -H "Authorization: Bearer invalid_token_12345" \
  -w "\nHTTPステータス: %{http_code}\n"
echo ""
echo "--------------------------------"

# 7. メモリー作成エラーテスト
echo "7. メモリー作成の必須フィールド不足テスト (VAL_002)"
echo "期待結果: エラーコード VAL_002 が返される"
echo "実行コマンド:"
echo 'curl -X POST '$API_BASE_URL'/api/memory \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer '$TOKEN'" \
  -d "{}"'
echo ""
curl -X POST $API_BASE_URL/api/memory \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{}' \
  -w "\nHTTPステータス: %{http_code}\n"
echo ""
echo "--------------------------------"

# 8. アセスメント取得エラーテスト
echo "8. 存在しないアセスメントの取得テスト (RES_001)"
echo "期待結果: エラーコード RES_001 が返される"
echo "実行コマンド:"
echo 'curl -X GET '$API_BASE_URL'/api/assessment/999999 \
  -H "Authorization: Bearer '$TOKEN'"'
echo ""
curl -X GET $API_BASE_URL/api/assessment/999999 \
  -H "Authorization: Bearer $TOKEN" \
  -w "\nHTTPステータス: %{http_code}\n"
echo ""
echo "--------------------------------"

# 9. 正常なプロファイル取得
echo "9. 認証済みユーザーのプロファイル取得テスト"
echo "期待結果: ユーザー情報が返される"
echo "実行コマンド:"
echo 'curl -X GET '$API_BASE_URL'/api/user/profile \
  -H "Authorization: Bearer '$TOKEN'"'
echo ""
curl -X GET $API_BASE_URL/api/user/profile \
  -H "Authorization: Bearer $TOKEN" \
  -w "\nHTTPステータス: %{http_code}\n"
echo ""

# ログアウト
echo "10. ログアウトテスト"
echo "実行コマンド:"
echo 'curl -X POST '$API_BASE_URL'/api/logout \
  -H "Authorization: Bearer '$TOKEN'"'
echo ""
curl -X POST $API_BASE_URL/api/logout \
  -H "Authorization: Bearer $TOKEN" \
  -w "\nHTTPステータス: %{http_code}\n"
echo ""

echo "================================"
echo "テスト完了"
echo "================================"
echo ""
echo "【確認ポイント】"
echo "1. 各エラーケースで適切なエラーコードが返されているか"
echo "2. HTTPステータスコードが適切か (400, 401, 404, 500など)"
echo "3. エラーメッセージが日本語で分かりやすく表示されているか"
echo "4. エラーの詳細情報(detail)が適切に含まれているか"
echo "5. 認証が必要なエンドポイントで適切に認証チェックが行われているか"
echo ""
echo "【バックエンドのテスト実行】"
echo "cd backer/godotask && go test ./errors -v"
echo "cd backer/godotask && go test ./middleware -v"
echo ""
echo "【フロントエンドのテスト実行】"
echo "cd fronter && npm test -- --testPathPattern=errorCodes"
echo ""
echo "【トークンの手動取得方法】"
echo "1. ユーザー登録:"
echo "   curl -X POST $API_BASE_URL/api/register -H \"Content-Type: application/json\" -d '{\"username\":\"test\",\"password\":\"pass123\",\"email\":\"test@example.com\"}'"
echo "2. ログイン:"
echo "   curl -X POST $API_BASE_URL/api/login -H \"Content-Type: application/json\" -d '{\"username\":\"test\",\"password\":\"pass123\"}'"