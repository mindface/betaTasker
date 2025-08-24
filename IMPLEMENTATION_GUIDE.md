# エラーコード実装ガイド

## 実装済みファイル

### バックエンド (Go)
1. **エラーコード定義**: `backer/godotask/errors/error_codes.go`
2. **コントローラー統合**: 
   - `backer/godotask/controller/task/add.go` (更新済み)
   - `backer/godotask/controller/memory_controller.go` (更新済み)
3. **ミドルウェア**: `backer/godotask/middleware/error_handler.go`
4. **テスト**: `backer/godotask/errors/error_codes_test.go`

### フロントエンド (TypeScript)
1. **エラーコード定義**: `fronter/src/errors/errorCodes.ts`
2. **API統合**: `fronter/src/services/taskApi.ts` (更新済み)
3. **テスト**: `fronter/src/errors/errorCodes.test.ts`

### テスト
- **統合テスト**: `test_error_codes.sh`

## 使用方法

### 1. バックエンド実装例

```go
// 既存のコントローラーに統合
import "github.com/godotask/errors"

// エラー作成
appErr := errors.NewAppError(
    errors.VAL_MISSING_FIELD,
    errors.GetErrorMessage(errors.VAL_MISSING_FIELD),
    "タイトルは必須項目です",
)

// JSONレスポンス
c.JSON(appErr.HTTPStatus, gin.H{
    "code":    appErr.Code,
    "message": appErr.Message,
    "detail":  appErr.Detail,
})
```

### 2. フロントエンド実装例

```typescript
import { ApplicationError, ErrorCode } from '../errors/errorCodes';

// エラー作成
throw new ApplicationError(
    ErrorCode.VAL_MISSING_FIELD,
    'タイトルは必須項目です'
);

// APIエラーハンドリング
try {
    const result = await addTaskService(task);
    return result;
} catch (error) {
    const appError = parseErrorResponse(error);
    console.error(`[${appError.code}] ${appError.message}`);
}
```

### 3. ミドルウェア適用

main.goまたはrouter.goに以下を追加：

```go
import "github.com/godotask/middleware"

router := gin.Default()

// ミドルウェアを適用
router.Use(middleware.ErrorHandlerMiddleware())
router.Use(middleware.CORSMiddleware())
router.Use(middleware.RequestValidationMiddleware())
router.Use(middleware.LoggingMiddleware())

// 404ハンドラー
router.NoRoute(middleware.NotFoundMiddleware())
```

## テスト実行

### バックエンドテスト
```bash
cd backer/godotask
go test ./errors -v
```

### フロントエンドテスト
```bash
cd fronter
npm test -- --testPathPattern=errorCodes
```

### 統合テスト
```bash
# サーバー起動後
bash test_error_codes.sh
```

## エラーコード一覧

### 認証・認可 (AUTH_xxx)
- `AUTH_001`: 認証情報無効
- `AUTH_002`: トークン期限切れ
- `AUTH_003`: 無効トークン
- `AUTH_004`: 権限なし
- `AUTH_005`: アカウント無効

### バリデーション (VAL_xxx)
- `VAL_001`: 入力値無効
- `VAL_002`: 必須項目未入力
- `VAL_003`: 入力形式不正
- `VAL_004`: 重複エントリ
- `VAL_005`: 制約違反

### リソース (RES_xxx)
- `RES_001`: リソース未発見
- `RES_002`: リソース既存
- `RES_003`: アクセス拒否
- `RES_004`: リソースロック

### システム (SYS_xxx)
- `SYS_001`: 内部エラー
- `SYS_002`: サービス利用不可
- `SYS_003`: タイムアウト
- `SYS_004`: レート制限超過

## 導入手順

1. **エラーコード定義の確認**
2. **既存コントローラーの更新**
3. **フロントエンドAPI統合**
4. **ミドルウェア適用**
5. **テスト実行と確認**

各ステップの詳細は上記ファイルを参照してください。