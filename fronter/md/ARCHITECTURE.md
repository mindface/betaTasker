# Memory & Assessment System 設計書

## 1. システム概要

### 目的
Memory & Assessmentシステムは、ユーザーの学習・作業記録（Memory）と、タスクに対する評価（Assessment）を統合管理し、知識の構造化と学習効果の可視化を実現するシステムです。

### 主要機能
- **Memory管理**: 学習内容・参考資料の記録と管理
- **Memory Aid**: 技術的要因と知識変換プロセスの構造化
- **Assessment**: タスクパフォーマンスの多角的評価
- **関連性管理**: Task-Memory-Assessment間の関連付け

## 2. アーキテクチャ概要

### レイヤー構成
```
┌─────────────────────────────────────────┐
│         UI Components (React)            │
├─────────────────────────────────────────┤
│         Redux Store & Slices             │
├─────────────────────────────────────────┤
│         Service Layer (APIs)             │
├─────────────────────────────────────────┤
│      Next.js API Routes (Proxy)          │
├─────────────────────────────────────────┤
│      Backend Server (Go - :8080)         │
├─────────────────────────────────────────┤
│            Database (DB)                 │
└─────────────────────────────────────────┘
```

### データフロー
1. **UI層**: ユーザー操作を受け付け、Reduxアクションを発行
2. **State層**: Redux Thunksがサービス層を呼び出し
3. **Service層**: HTTP通信を管理、エラーハンドリング
4. **API Routes層**: 認証とプロキシ処理
5. **Backend層**: ビジネスロジックとデータ永続化

## 3. データモデル

### Memory
```typescript
interface Memory {
  id: number;
  user_id: number;
  source_type: string;        // 情報源タイプ
  title: string;               // タイトル
  author: string;              // 著者
  notes: string;               // メモ
  factor: string;              // 要因
  process: string;             // プロセス
  evaluation_axis: string;     // 評価軸
  information_amount: string;  // 情報量
  tags: string;                // タグ
  read_status: string;         // 読了状態
  read_date: string | null;    // 読了日
  created_at: string;
  updated_at: string;
}
```

### MemoryContext (Memory Aid)
```typescript
interface MemoryContext {
  id: number;
  user_id: number;
  task_id: number;
  level: number;                              // レベル
  work_target: string;                        // 作業対象
  machine: string;                            // 機械・ツール
  material_spec: string;                      // 材料仕様
  change_factor: string;                      // 変更要因
  goal: string;                               // 目標
  technical_factors: TechnicalFactor[];       // 技術的要因
  knowledge_transformations: KnowledgeTransformation[]; // 知識変換
}

interface TechnicalFactor {
  domain: string;      // 領域
  subdomain: string;   // サブドメイン
  element: string;     // 要素
}

interface KnowledgeTransformation {
  from: string;        // 変換元
  to: string;          // 変換先
  description: string; // 説明
}
```

### Assessment
```typescript
interface Assessment {
  id: number;
  task_id: number;             // 関連タスクID
  user_id: number;             // ユーザーID
  effectiveness_score: number; // 有効性スコア (0-100)
  effort_score: number;        // 努力スコア (0-100)
  impact_score: number;        // 影響度スコア (0-100)
  qualitative_feedback: string;// 定性的フィードバック
  created_at: string;
  updated_at: string;
}
```

### Task (関連モデル)
```typescript
interface Task {
  id: number;
  user_id: number;
  memory_id?: number | null;  // 関連メモリID
  title: string;
  description: string;
  date?: string | null;
  status: string;
  priority: number;
  created_at: string;
  updated_at: string;
}
```

## 4. API仕様

### Memory API

#### エンドポイント一覧
| Method | Endpoint | 説明 |
|--------|----------|------|
| GET | `/api/memory` | メモリ一覧取得 |
| POST | `/api/memory` | メモリ作成 |
| GET | `/api/memory/:id` | 個別メモリ取得 |
| PUT | `/api/memory/:id` | メモリ更新 |
| DELETE | `/api/memory/:id` | メモリ削除 |

#### リクエスト/レスポンス例
```typescript
// POST /api/memory
Request: {
  user_id: number;
  source_type: string;
  title: string;
  author: string;
  notes: string;
  factor: string;
  process: string;
  evaluation_axis: string;
  information_amount: string;
  tags: string;
  read_status: string;
  read_date?: string;
}

Response: {
  data: Memory;
  message: string;
}
```

### Memory Aid API

#### エンドポイント
| Method | Endpoint | 説明 |
|--------|----------|------|
| GET | `/api/memoryAid?code={code}` | コード別Memory Aid取得 |
| GET | `/api/memory/context/:code` | コンテキスト取得 |
| GET | `/api/memory/aid/:code` | Aid情報取得 |

#### コード体系
- **MA-C-01**: Memory Aid - Category - 01
- **MA-Q-02**: Memory Aid - Quality - 02
- **PM-P-03**: Project Management - Process - 03

### Assessment API

#### エンドポイント一覧
| Method | Endpoint | 説明 |
|--------|----------|------|
| GET | `/api/assessment` | 評価一覧取得 |
| POST | `/api/assessment` | 評価作成 |
| GET | `/api/assessment/:id` | 個別評価取得 |
| PUT | `/api/assessment/:id` | 評価更新 |
| DELETE | `/api/assessment/:id` | 評価削除 |
| POST | `/api/assessmentsForTaskUser` | タスク・ユーザー別評価取得 |

#### 特殊エンドポイント
```typescript
// POST /api/assessmentsForTaskUser
Request: {
  task_id: number;
  user_id: number;
}

Response: {
  data: Assessment[];
  message: string;
}
```

## 5. コンポーネント構成

### 主要コンポーネント

#### Memory関連
- **SectionMemory** (`/components/SectionMemory.tsx`)
  - メモリ管理のメインインターフェース
  - CRUD操作の統合
  - MemoryAidListの表示

- **MemoryAidList** (`/components/MemoryAidList.tsx`)
  - Memory Aidの一覧表示
  - コード別フィルタリング
  - 詳細モーダル表示

#### Assessment関連
- **SectionAssessment** (`/components/SectionAssessment.tsx`)
  - 評価管理のメインインターフェース
  - ヒューリスティクスダッシュボード統合
  - 学習構造データ表示

- **SectionAssessmentRelation** (`/components/SectionAssessmentRelation.tsx`)
  - Task-Memory-Assessment関連付けUI
  - タスク選択による評価フィルタリング
  - 関連メモリの表示

#### 共通コンポーネント
- **AssessmentModal** (`/components/parts/AssessmentModal.tsx`)
  - 評価の作成/編集フォーム
  - 関連メモリの参照表示
  - スコアリング入力

- **ItemAssessment** (`/components/parts/ItemAssessment.tsx`)
  - 個別評価の表示カード
  - スコアの視覚化
  - 編集/削除アクション

## 6. State管理

### Redux Store構造
```typescript
interface RootState {
  memory: {
    memories: Memory[];
    memoryItem: Memory | null;
    memoryLoading: boolean;
    memoryError: string | null;
  };
  
  memoryAid: {
    contexts: MemoryContext[];
    loading: boolean;
    error: string | null;
  };
  
  assessment: {
    assessments: Assessment[];
    assessmentLoading: boolean;
    assessmentError: string | null;
  };
  
  task: {
    tasks: Task[];
    taskLoading: boolean;
    taskError: string | null;
  };
}
```

### Slice構成

#### memorySlice
- **Actions**: 
  - `loadMemories`: 一覧取得
  - `createMemory`: 新規作成
  - `updateMemory`: 更新
  - `removeMemory`: 削除
  - `getMemory`: 個別取得

#### memoryAidSlice
- **Actions**:
  - `loadMemoryAidsByCode`: コード別取得

#### assessmentSlice
- **Actions**:
  - `loadAssessments`: 一覧取得
  - `createAssessment`: 新規作成
  - `updateAssessment`: 更新
  - `removeAssessment`: 削除
  - `getAssessmentsForTaskUser`: タスク・ユーザー別取得

## 7. サービス層

### APIStrategy Pattern
```typescript
// /services/apiStrategy.ts
interface ApiStrategy<T> {
  getAll(): Promise<T[]>;
  getById(id: string): Promise<T>;
  create(item: Omit<T, 'id'>): Promise<T>;
  update(id: string, item: Partial<T>): Promise<T>;
  delete(id: string): Promise<void>;
}
```

### Service実装
- **memoryApi.ts**: Memory CRUD操作
- **memoryAidApi.ts**: Memory Aid取得
- **assessmentApi.ts**: Assessment CRUD操作

## 8. Hooks

### useItemOperations
```typescript
// 汎用CRUD操作Hook
const useItemOperations = <T extends BaseItem>(
  entityType: EntityType
) => {
  const operations = {
    loadItems: () => void;
    createItem: (item: T) => Promise<void>;
    updateItem: (id: string, item: Partial<T>) => Promise<void>;
    deleteItem: (id: string) => Promise<void>;
  };
  
  return operations;
};
```

## 9. 認証とセキュリティ

### 認証フロー
1. **Cookie認証**: `user_info`クッキーからユーザー情報取得
2. **Token管理**: JWTトークンによるAPI認証
3. **プロキシ層**: Next.js API Routesでの認証チェック

### セキュリティ対策
- CSRFトークン検証
- API Rate Limiting
- 入力値検証とサニタイゼーション
- SQLインジェクション対策

## 10. 関連性管理

### データ関連図
```
┌──────────┐     ┌──────────┐     ┌──────────────┐
│   Task   │────▶│  Memory  │────▶│  Assessment  │
└──────────┘     └──────────┘     └──────────────┘
     │                │                    │
     │           ┌────▼────┐              │
     └──────────▶│MemoryAid│◀─────────────┘
                 └──────────┘
```

### 関連性ルール
1. **Task → Memory**: 1対0..1関係（オプショナル）
2. **Task → Assessment**: 1対多関係
3. **Memory → MemoryAid**: コードベースの関連付け
4. **User → All Entities**: すべてのエンティティはユーザーに紐付く

## 11. エラーハンドリング

### エラー階層
```typescript
class ApplicationError extends Error {
  constructor(
    public code: ErrorCode,
    public message: string,
    public detail?: string
  ) {}
}

enum ErrorCode {
  // 認証エラー
  AUTH_INVALID_CREDENTIALS = 'AUTH001',
  AUTH_UNAUTHORIZED = 'AUTH002',
  
  // バリデーションエラー
  VAL_INVALID_INPUT = 'VAL001',
  VAL_MISSING_FIELD = 'VAL002',
  
  // リソースエラー
  RES_NOT_FOUND = 'RES001',
  RES_CONFLICT = 'RES002',
  
  // システムエラー
  SYS_INTERNAL_ERROR = 'SYS001',
  SYS_EXTERNAL_SERVICE = 'SYS002',
}
```

## 12. パフォーマンス最適化

### 実装済み最適化
- **遅延ローディング**: コンポーネントの動的インポート
- **メモ化**: `useMemo`と`useCallback`の活用
- **バッチ処理**: 複数APIコールの最適化
- **キャッシング**: Redux永続化

### 推奨最適化
- **仮想スクロール**: 大量データ表示時
- **デバウンス**: 検索・フィルタリング
- **楽観的更新**: UX改善
- **プリフェッチ**: 予測的データ取得

## 13. テスト戦略

### テストレベル
1. **単体テスト**: サービス層、Reduxロジック
2. **統合テスト**: API Routes、Redux統合
3. **E2Eテスト**: ユーザーフロー全体

### テストツール
- Jest: 単体・統合テスト
- React Testing Library: コンポーネントテスト
- Cypress/Playwright: E2Eテスト

## 14. デプロイメント

### 環境構成
- **開発環境**: localhost:3000 (Next.js) + localhost:8080 (Go)
- **ステージング**: Vercel/Netlify + Cloud Run
- **本番環境**: Vercel/Netlify + Cloud Run + Cloud SQL

### CI/CDパイプライン
1. GitHubプッシュ
2. GitHub Actions実行
3. テスト実行
4. ビルド
5. デプロイ

## 15. 今後の拡張計画

### Phase 1: 基盤強化 (現在)
- ✅ 基本CRUD機能
- ✅ Redux統合
- ✅ 認証システム
- ⬜ エラーハンドリング統一

### Phase 2: 機能拡張
- ⬜ 検索・フィルタリング強化
- ⬜ バルク操作
- ⬜ インポート/エクスポート
- ⬜ テンプレート機能

### Phase 3: 分析・可視化
- ⬜ 学習進捗ダッシュボード
- ⬜ パフォーマンストレンド
- ⬜ レコメンデーション
- ⬜ レポート生成

### Phase 4: AI統合
- ⬜ 自動評価提案
- ⬜ 知識グラフ生成
- ⬜ 学習パス最適化
- ⬜ 自然言語処理

### Phase 5: エンタープライズ
- ⬜ チーム機能
- ⬜ 権限管理
- ⬜ 監査ログ
- ⬜ API公開

## 16. 運用・保守

### モニタリング
- アプリケーションログ
- パフォーマンスメトリクス
- エラー追跡
- ユーザー行動分析

### バックアップ
- データベース: 日次バックアップ
- 設定ファイル: Git管理
- メディアファイル: オブジェクトストレージ

### ドキュメント
- API仕様書: OpenAPI/Swagger
- コンポーネントカタログ: Storybook
- ユーザーマニュアル: Markdown
- 開発者ガイド: このドキュメント

---

最終更新日: 2025-08-31
バージョン: 1.0.0