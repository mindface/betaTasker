fronter/md/ARCHITECTURE.md
fronter/md/DETAILED_DESIGN.md
これをベースにコードを変更します。

fronter/src/components/heuristics
fronter/src/features/heuristics
fronter/src/services/heuristicsApi.ts
fronter/src/services/heuristicsDiscovery.ts
など含めて
fronter/md/heuristics.md
を対象としてリファクタリングする案を出してください。

//////////////////////////
1.0.98 (Claude Code)
//////////////////////////


1. コンポーネント: 5つの主要コンポーネント（Dashboard,
  Analysis, Tracking, Patterns, Insights）
  2. Redux State: heuristicsSlice.tsで統一的な状態管理
  3. API Services: heuristicsApi.tsでAPI通信
  4. Discovery Service:
  高度なパターン検出・ヒューリスティクス発見機能

  問題点

  - 分析結果一覧API (GET /api/heuristics/analyze) 未実装
  - Go側とフロント側でパラメータ名の不整合 (type vs data_type)
  - エラーハンドリングの不統一
  - heuristicsDiscovery.tsの複雑性とメンテナンス性の課題
  - コンポーネント間のコード重複

  リファクタリング提案

  1. アーキテクチャレベルの改善

  1.1 レイヤー分離の強化

  ┌─────────────────────────────────────────┐
  │         Presentation Layer              │
  │  (Components + Custom Hooks)           │
  ├─────────────────────────────────────────┤
  │         Business Logic Layer            │
  │    (Services + Domain Models)          │
  ├─────────────────────────────────────────┤
  │         Data Access Layer               │
  │      (API Clients + Transformers)      │
  ├─────────────────────────────────────────┤
  │         Infrastructure Layer            │
  │   (Error Handlers + Utils + Worker)    │
  └─────────────────────────────────────────┘

  1.2 Domain-Driven Design適用

  // fronter/src/domains/heuristics/
  ├── models/
  │   ├── Analysis.ts
  │   ├── Tracking.ts
  │   ├── Pattern.ts
  │   └── Insight.ts
  ├── services/
  │   ├── AnalysisService.ts
  │   ├── TrackingService.ts
  │   └── PatternDiscoveryService.ts
  ├── repositories/
  │   ├── HeuristicsRepository.ts
  │   └── interfaces/
  └── valueObjects/
      ├── AnalysisType.ts
      └── PatternMetrics.ts


  2.2 カスタムフックによるロジック分離

  // fronter/src/hooks/heuristics/
  ├── useHeuristicsAnalysis.ts     // 既存
  ├── useHeuristicsTracking.ts     // 既存
  ├── useHeuristicsPatterns.ts     // 既存
  ├── useHeuristicsInsights.ts     // 既存
  ├── usePagination.ts             // 新規
  ├── useDataVisualization.ts      // 新規
  ├── useRealTimeUpdates.ts        // 新規
  └── useErrorBoundary.ts          // 新規

  3. サービス層の再設計

  3.1 API Client統一

  // fronter/src/services/heuristics/
  ├── clients/
  │   ├── HeuristicsApiClient.ts   // Base HTTP client
  │   ├── AnalysisClient.ts
  │   ├── TrackingClient.ts
  │   ├── PatternClient.ts
  │   └── InsightClient.ts
  ├── transformers/
  │   ├── ResponseTransformer.ts
  │   └── RequestTransformer.ts
  ├── cache/
  │   ├── CacheManager.ts
  │   └── strategies/
  └── websocket/
      ├── RealtimeClient.ts
      └── EventHandler.ts

  3.2 Discovery Service簡素化

  // fronter/src/services/heuristics/discovery/
  ├── core/
  │   ├── PatternDetector.ts       // 核心ロジック
  │   ├── HeuristicInference.ts    // ヒューリスティクス推論
  │   └── ModelUpdater.ts          // モデル更新
  ├── workers/
  │   ├── PatternWorker.ts         // Web Worker抽象化
  │   └── workerScript.ts          // Worker実装
  ├── algorithms/
  │   ├── NGramAnalyzer.ts         // N-gram分析
  │   ├── TemporalClustering.ts    // 時間的クラスタリング
  │   └── CognitiveLoadEstimator.ts // 認知負荷推定
  └── interfaces/
      ├── IPatternDetector.ts
      └── IHeuristicEngine.ts

  4. Redux Store再構成

  4.1 State正規化

  // fronter/src/store/heuristics/
  ├── slices/
  │   ├── analysisSlice.ts         // 分析専用
  │   ├── trackingSlice.ts         // トラッキング専用
  │   ├── patternsSlice.ts         // パターン専用
  │   ├── insightsSlice.ts         // インサイト専用
  │   └── uiSlice.ts              // UI状態管理
  ├── selectors/
  │   ├── analysisSelectors.ts
  │   ├── trackingSelectors.ts
  │   ├── patternsSelectors.ts
  │   └── combinedSelectors.ts     // 複合セレクタ
  ├── middleware/
  │   ├── errorLoggingMiddleware.ts
  │   ├── cacheMiddleware.ts
  │   └── realtimeMiddleware.ts
  └── entities/
      ├── analysis.ts              // エンティティ正規化
      ├── tracking.ts
      └── patterns.ts


  4.2 正規化されたState構造
```ts
  interface HeuristicsState {
    entities: {
      analyses: Record<string, HeuristicsAnalysis>;
      tracking: Record<string, HeuristicsTracking>;
      patterns: Record<string, HeuristicsPattern>;
      insights: Record<string, HeuristicsInsight>;
    };

    collections: {
      analyses: {
        ids: string[];
        pagination: PaginationState;
        filters: AnalysisFilters;
        loading: boolean;
        error: string | null;
      };
      // 他のコレクションも同様
    };

    ui: {
      activeTab: TabType;
      modals: ModalState;
      notifications: NotificationState;
    };

    realtime: {
      connected: boolean;
      subscriptions: string[];
      lastUpdate: number;
    };
  }
```

  5. 型安全性の強化

  5.1 厳密な型定義
```
  // fronter/src/types/heuristics/
  ├── api.ts                       // API型定義
  ├── domain.ts                    // ドメインモデル
  ├── ui.ts                       // UI関連型
  ├── events.ts                   // イベント型
  └── guards/                     // Type Guards
      ├── analysisGuards.ts
      ├── trackingGuards.ts
      └── patternGuards.ts
```

  5.2 Runtime Type Validation
```ts
  import { z } from 'zod';

  // Zodスキーマでランタイム検証
  const HeuristicsAnalysisSchema = z.object({
    id: z.number(),
    user_id: z.number(),
    task_id: z.number().optional(),
    analysis_type: z.enum(['performance', 'behavior',
  'pattern', 'cognitive', 'efficiency']),
    result: z.string(),
    score: z.number().min(0).max(100),
    status: z.string(),
    created_at: z.string(),
    updated_at: z.string()
  });
```

  6. エラーハンドリング統一

  6.1 エラー階層化

  // fronter/src/errors/heuristics/
  ├── HeuristicsError.ts           // Base error
  ├── AnalysisError.ts            // 分析エラー
  ├── TrackingError.ts            // トラッキングエラー
  ├── PatternError.ts             // パターンエラー
  ├── NetworkError.ts             // ネットワークエラー
  └── ValidationError.ts          // バリデーションエラー

```ts
  // fronter/src/components/heuristics/ErrorBoundary.tsx
  interface ErrorBoundaryState {
    hasError: boolean;
    error: Error | null;
    errorInfo: ErrorInfo | null;
    errorId: string;
  }

  class HeuristicsErrorBoundary extends Component<Props, 
  ErrorBoundaryState> {
    // エラー回復戦略
    private recoverStrategies = {
      network: () => this.retryWithBackoff(),
      validation: () => this.resetForm(),
      analysis: () => this.fallbackToCache(),
    };
  }
```

  7. パフォーマンス最適化

  7.1 Code Splitting
```ts
  // 動的インポートによるコード分割
  const HeuristicsAnalysis = lazy(() =>
    import('./HeuristicsAnalysis').then(m => ({ default:
  m.HeuristicsAnalysis }))
  );

  const HeuristicsPatterns = lazy(() =>
    import('./HeuristicsPatterns')
  );
```

  7.2 仮想化リスト
```ts
  // react-windowを使用した大量データ表示
  import { FixedSizeList } from 'react-window';

  const VirtualizedAnalysisList: FC<Props> = ({ items }) => {
    const Row = ({ index, style }: ListChildComponentProps) =>
   (
      <div style={style}>
        <AnalysisCard analysis={items[index]} />
      </div>
    );

    return (
      <FixedSizeList
        height={600}
        itemCount={items.length}
        itemSize={150}
      >
        {Row}
      </FixedSizeList>
    );
  };
```

  8. テスト戦略

  8.1 テスト構造
```
  fronter/src/__tests__/heuristics/
  ├── components/
  │   ├── __snapshots__/
  │   ├── HeuristicsDashboard.test.tsx
  │   ├── HeuristicsAnalysis.test.tsx
  │   └── shared/
  ├── hooks/
  │   ├── useHeuristicsAnalysis.test.ts
  │   └── useHeuristicsPatterns.test.ts
  ├── services/
  │   ├── heuristicsApi.test.ts
  │   └── discovery/
  ├── store/
  │   ├── heuristicsSlice.test.ts
  │   └── selectors/
  └── integration/
      ├── analysisFlow.test.tsx
      └── patternDetection.test.tsx
```

  8.2 Testing Strategy

  - 単体テスト: Jest + React Testing Library
  - 統合テスト: MSW (Mock Service Worker)
  - E2Eテスト: Playwright
  - Visual Regression: Chromatic

  9. 実装段階

  Phase 1: Core Infrastructure (2週間)

  1. エラーハンドリングシステム構築
  2. 型定義・スキーマ整備
  3. Base APIクライアント実装
  4. Redux Store再設計

  Phase 2: Service Layer (2週間)

  1. Domain Services実装
  2. Discovery Service簡素化
  3. キャッシング戦略実装
  4. WebSocket統合

  Phase 3: UI Components (3週間)

  1. 共通コンポーネント実装
  2. カスタムフック移行
  3. 既存コンポーネントリファクタリング
  4. パフォーマンス最適化

  Phase 4: Testing & Documentation (1週間)

  1. テストカバレッジ向上
  2. ドキュメント更新
  3. 型安全性検証
  4. パフォーマンステスト


⏺ 優先順位と推奨アクション

  最優先 (即座に実施)

  1. API仕様統一: Go側とフロント側のパラメータ名統一
  2. 分析結果一覧API実装: GET /api/heuristics/analyze
  3. エラーハンドリング統一: ApplicationErrorクラス活用

  高優先 (1-2週間以内)

  1. Discovery Service簡素化:
  複雑すぎるheuristicsDiscovery.tsを分割
  2. Redux State正規化: エンティティベースのState設計
  3. 共通コンポーネント抽出: コード重複解消

  中優先 (1-2ヶ月以内)

  1. Domain-Driven Design適用: ビジネスロジック整理
  2. WebSocket統合: リアルタイム更新機能
  3. パフォーマンス最適化: 仮想化・メモ化

  この提案により、保守性・拡張性・パフォーマンスが大幅に向上し
  、チーム開発効率も改善されると期待できます。

