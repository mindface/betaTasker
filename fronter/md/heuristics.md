# Heuristics Components 仕様書

## 概要
ヒューリスティクス機能は、ユーザーの行動分析、パターン検出、インサイト生成を行うための統合システムです。

## コンポーネント構成

### 1. HeuristicsDashboard
**役割**: ヒューリスティクス機能の統合ダッシュボード  
**パス**: `/fronter/src/components/heuristics/HeuristicsDashboard.tsx`

### 2. HeuristicsAnalysis
**役割**: 分析の実行と結果表示  
**パス**: `/fronter/src/components/heuristics/HeuristicsAnalysis.tsx`
**機能**:
- 新規分析の実行
- 個別分析結果の取得（ID指定）
- 分析結果の詳細表示

### 3. HeuristicsTracking
**役割**: ユーザー行動のトラッキングと表示  
**パス**: `/fronter/src/components/heuristics/HeuristicsTracking.tsx`
**機能**:
- 行動データの記録
- トラッキングデータの一覧表示
- ユーザーIDによるデータ取得

### 4. HeuristicsPatterns
**役割**: パターン検出と表示  
**パス**: `/fronter/src/components/heuristics/HeuristicsPatterns.tsx`
**機能**:
- パターンの検出と表示
- フィルタリング（ユーザーID、データタイプ、期間）
- モデルトレーニング

### 5. HeuristicsInsights
**役割**: インサイトの生成と表示  
**パス**: `/fronter/src/components/heuristics/HeuristicsInsights.tsx`
**機能**:
- インサイト一覧の表示
- ページネーション
- 個別インサイトの詳細表示

## API仕様

### 基本設定
- **ベースURL**: `/api/heuristics`
- **認証**: クッキーベース（`credentials: 'include'`）
- **エラーハンドリング**: ApplicationErrorクラスによる統一エラー処理

### エンドポイント一覧

#### 1. 分析（Analysis）

##### POST /api/heuristics/analyze
**説明**: 新規分析の実行  
**リクエスト**:
```typescript
{
  user_id: number;
  task_id?: number;
  analysis_type: 'performance' | 'behavior' | 'pattern' | 'cognitive' | 'efficiency';
  data: object;
}
```
**レスポンス**:
```typescript
{
  success: boolean;
  message: string;
  data: {
    analysis: HeuristicsAnalysis;
  }
}
```

##### GET /api/heuristics/analyze/:id
**説明**: 個別分析結果の取得  
**レスポンス**:
```typescript
{
  success: boolean;
  data: {
    analysis: HeuristicsAnalysis;
  }
}
```

#### 2. トラッキング（Tracking）

##### POST /api/heuristics/track
**説明**: 行動データの記録  
**リクエスト**:
```typescript
{
  user_id: number;
  action: string;
  context: string;
  session_id?: string;
  timestamp?: string;
}
```
**レスポンス**:
```typescript
{
  success: boolean;
  message: string;
  data: {
    tracking_id: number;
  }
}
```

##### GET /api/heuristics/track/:user_id
**説明**: ユーザーのトラッキングデータ取得  
**レスポンス**:
```typescript
{
  success: boolean;
  data: {
    tracking_data: HeuristicsTracking[];
  }
}
```

#### 3. パターン（Patterns）

##### GET /api/heuristics/patterns
**説明**: パターンの検出と取得  
**クエリパラメータ**:
- `user_id`: ユーザーID（オプション）
- `data_type`: データタイプ（オプション、デフォルト: 'all'）
- `period`: 期間（オプション、デフォルト: 'week'）

**レスポンス**:
```typescript
{
  success: boolean;
  message: string;
  data: {
    patterns: HeuristicsPattern[];
    metadata: {
      user_id: string;
      data_type: string;
      period: string;
    }
  }
}
```

##### POST /api/heuristics/patterns/train
**説明**: モデルのトレーニング  
**リクエスト**:
```typescript
{
  model_type: 'pattern_detection' | 'behavior_prediction' | 'anomaly_detection' | 'recommendation';
  parameters: object;
  data_source: string;
  training_data: any[];
}
```
**レスポンス**:
```typescript
{
  success: boolean;
  message: string;
  data: {
    training_id: number;
    status: string;
    model_type: string;
  }
}
```

#### 4. インサイト（Insights）

##### GET /api/heuristics/insights
**説明**: インサイト一覧の取得  
**クエリパラメータ**:
- `limit`: 取得件数（オプション）
- `offset`: オフセット（オプション）
- `user_id`: ユーザーID（オプション）

**レスポンス**:
```typescript
{
  success: boolean;
  data: {
    insights: HeuristicsInsight[];
    total: number;
    limit: number;
    offset: number;
  }
}
```

##### GET /api/heuristics/insights/:id
**説明**: 個別インサイトの取得  
**レスポンス**:
```typescript
{
  success: boolean;
  data: {
    insight: HeuristicsInsight;
  }
}
```

## データモデル

### HeuristicsAnalysis
```typescript
interface HeuristicsAnalysis {
  id: number;
  user_id: number;
  task_id?: number;
  analysis_type: string;
  result: string;
  score: number;
  status: string;
  created_at: string;
  updated_at: string;
}
```

### HeuristicsTracking
```typescript
interface HeuristicsTracking {
  id: number;
  user_id: number;
  action: string;
  context: string;
  session_id: string;
  timestamp: string;
  duration: number;
  created_at: string;
  updated_at: string;
}
```

### HeuristicsPattern
```typescript
interface HeuristicsPattern {
  id: number;
  name: string;
  category: string;
  pattern: string; // JSON文字列
  frequency: number;
  accuracy: number;
  last_seen: string;
  created_at: string;
  updated_at: string;
}
```

### HeuristicsInsight
```typescript
interface HeuristicsInsight {
  id: number;
  user_id: number;
  type: string;
  title: string;
  description: string;
  confidence: number;
  impact: string;
  suggestions: string[];
  created_at: string;
  updated_at: string;
}
```

### HeuristicsModel
```typescript
interface HeuristicsModel {
  id: number;
  model_type: string;
  version: string;
  status: 'training' | 'ready' | 'failed';
  trained_at: string;
  created_at: string;
}
```

## 既知の問題と改善点

### 実装済み
- ✅ 個別分析結果の取得UI
- ✅ パターン検出でのgetPatterns関数の呼び出し
- ✅ トラッキングデータの正しいレスポンス構造対応

### 未実装・要改善
1. **分析結果一覧API**: `GET /api/heuristics/analyze`エンドポイントが未実装
2. **Go側のクエリパラメータ不一致**: patterns.goで`type`を使用しているが、フロント側は`data_type`
3. **メタデータの活用**: patternsのmetadataが取得されているが未活用
4. **リアルタイムデータ更新**: WebSocketによるリアルタイム更新は未実装
5. **エラーハンドリングの統一**: 一部のAPIでtry-catchが欠けている

## 使用方法

### 1. 分析の実行
```typescript
const { analyze } = useHeuristicsAnalysis();
await analyze({
  user_id: 1,
  analysis_type: 'performance',
  data: { metric: 'completion_time', value: 120 }
});
```

### 2. パターンの取得
```typescript
const { getPatterns } = useHeuristicsPatterns();
getPatterns({
  user_id: '1',
  data_type: 'task',
  period: 'month'
});
```

### 3. トラッキングデータの記録
```typescript
const { track } = useHeuristicsTracking();
await track({
  user_id: 1,
  action: 'task_completed',
  context: JSON.stringify({ task_id: 123 })
});
```

## Redux統合

### Store構造
```typescript
state.heuristics = {
  // 分析
  analyses: HeuristicsAnalysis[];
  currentAnalysis: HeuristicsAnalysis | null;
  analysisLoading: boolean;
  analysisError: string | null;
  
  // トラッキング
  trackingData: HeuristicsTracking[];
  trackingLoading: boolean;
  trackingError: string | null;
  
  // パターン
  patterns: HeuristicsPattern[];
  patternsLoading: boolean;
  patternsError: string | null;
  
  // インサイト
  insights: HeuristicsInsight[];
  currentInsight: HeuristicsInsight | null;
  insightsTotal: number;
  insightsLoading: boolean;
  insightsError: string | null;
  
  // モデル
  currentModel: HeuristicsModel | null;
  modelLoading: boolean;
  modelError: string | null;
}
```

## 今後の開発計画

1. **Phase 1**: 分析結果一覧APIの実装
2. **Phase 2**: WebSocketによるリアルタイム更新
3. **Phase 3**: ビジュアライゼーション（グラフ、チャート）の追加
4. **Phase 4**: 機械学習モデルの統合強化
5. **Phase 5**: エクスポート機能（CSV、PDF）の追加