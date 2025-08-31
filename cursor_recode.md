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
Version: 1.5.7 (Universal)
VSCode Version: 1.99.3
Commit: 6aa7b3af0d578b9a3aa3ab443571e1a51ebb4e80
Date: 2025-08-27T22:43:32.248Z
Electron: 34.5.8
Chromium: 132.0.6834.210
Node.js: 20.19.1
V8: 13.2.152.41-electron.0
OS: Darwin arm64 22.6.0
//////////////////////////


ヒューリスティクス機能のリファクタリング案

1. アーキテクチャの統一化

1.1 Service層の統合と標準化
現在のheuristicsApi.tsとheuristicsDiscovery.tsを統合し、詳細設計書のパターンに合わせて標準化します。

```ts
// 新しい統合サービス構造
/services/heuristics/
├── heuristicsApi.ts          // 基本API呼び出し
├── heuristicsDiscovery.ts    // パターン発見ロジック
├── heuristicsAnalytics.ts    // 分析処理
└── heuristicsCache.ts        // キャッシュ管理
```

キャッシュ管理
1.2 エラーハンドリングの統一
詳細設計書のApplicationErrorパターンを適用し、一貫したエラー処理を実装します。

2. データモデルの改善
2.1 型定義の強化
```ts
// より厳密な型定義
export interface HeuristicsAnalysis {
  id: number;
  user_id: number;
  task_id?: number;
  analysis_type: 'performance' | 'behavior' | 'pattern' | 'cognitive' | 'efficiency';
  result: AnalysisResult;
  score: number;
  status: 'pending' | 'completed' | 'failed';
  metadata: AnalysisMetadata;
  created_at: string;
  updated_at: string;
}

export interface AnalysisResult {
  insights: string[];
  patterns: Pattern[];
  recommendations: Recommendation[];
  confidence: number;
}

export interface AnalysisMetadata {
  execution_time: number;
  data_points: number;
  algorithm_version: string;
}
```

2.2 バリデーション強化
```ts
// 入力バリデーション
export const validateAnalysisRequest = (data: HeuristicsAnalysisRequest): ValidationResult => {
  const errors: string[] = [];
  
  if (!data.user_id || data.user_id <= 0) {
    errors.push('有効なユーザーIDが必要です');
  }
  
  if (!data.analysis_type || !VALID_ANALYSIS_TYPES.includes(data.analysis_type)) {
    errors.push('有効な分析タイプを指定してください');
  }
  
  return {
    isValid: errors.length === 0,
    errors
  };
};
```


3. Redux State管理の最適化
3.1 状態管理の階層化
```ts
// より細かい状態管理
interface HeuristicsState {
  // 分析
  analysis: {
    items: HeuristicsAnalysis[];
    current: HeuristicsAnalysis | null;
    loading: boolean;
    error: string | null;
    filters: AnalysisFilters;
    pagination: PaginationState;
  };
  
  // トラッキング
  tracking: {
    data: HeuristicsTracking[];
    loading: boolean;
    error: string | null;
    realtime: boolean;
    sessionId: string | null;
  };
  
  // パターン
  patterns: {
    items: HeuristicsPattern[];
    loading: boolean;
    error: string | null;
    categories: string[];
    confidenceThreshold: number;
  };
  
  // インサイト
  insights: {
    items: HeuristicsInsight[];
    current: HeuristicsInsight | null;
    loading: boolean;
    error: string | null;
    total: number;
    filters: InsightFilters;
  };
}
```

3.2 非同期処理の改善
```ts
// より堅牢な非同期処理
export const analyzeData = createAsyncThunk(
  'heuristics/analyze',
  async (request: HeuristicsAnalysisRequest, { rejectWithValue, dispatch }) => {
    try {
      // 分析開始の状態更新
      dispatch(setAnalysisStatus('pending'));
      
      const response = await heuristicsApi.analyzeData(request);
      
      if ('error' in response) {
        return rejectWithValue({
          code: 'ANALYSIS_FAILED',
          message: response.error,
          details: response
        });
      }
      
      return response as HeuristicsAnalysis;
      
    } catch (error) {
      // エラーログ記録
      console.error('Analysis failed:', error);
      
      return rejectWithValue({
        code: 'SYSTEM_ERROR',
        message: '分析処理中にエラーが発生しました',
        details: error
      });
    }
  }
);
```

4. コンポーネントの改善
4.1 パフォーマンス最適化
```ts
// メモ化と仮想化
export const HeuristicsPatterns: React.FC = () => {
  const patterns = useSelector(selectFilteredPatterns);
  const loading = useSelector(selectPatternsLoading);
  
  // 高コストな計算のメモ化
  const categorizedPatterns = useMemo(() => {
    return groupPatternsByCategory(patterns);
  }, [patterns]);
  
  // 仮想スクロール（大量データ対応）
  const VirtualizedPatternList = useCallback(({ patterns }: { patterns: HeuristicsPattern[] }) => {
    return (
      <FixedSizeList
        height={600}
        itemCount={patterns.length}
        itemSize={120}
        width="100%"
      >
        {({ index, style }) => (
          <PatternCard
            key={patterns[index].id}
            pattern={patterns[index]}
            style={style}
          />
        )}
      </FixedSizeList>
    );
  }, []);
  
  return (
    <div className={styles.container}>
      <PatternFilters />
      {loading ? (
        <LoadingSpinner />
      ) : (
        <VirtualizedPatternList patterns={categorizedPatterns} />
      )}
    </div>
  );
};
```

4.2 リアルタイム更新対応
```ts
// WebSocket統合
export const useRealtimeUpdates = () => {
  const dispatch = useDispatch();
  const [socket, setSocket] = useState<WebSocket | null>(null);
  
  useEffect(() => {
    const ws = new WebSocket(process.env.NEXT_PUBLIC_WS_URL!);
    
    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      
      switch (data.type) {
        case 'PATTERN_DETECTED':
          dispatch(addPattern(data.pattern));
          break;
        case 'INSIGHT_GENERATED':
          dispatch(addInsight(data.insight));
          break;
        case 'ANALYSIS_COMPLETED':
          dispatch(updateAnalysis(data.analysis));
          break;
      }
    };
    
    setSocket(ws);
    
    return () => ws.close();
  }, [dispatch]);
  
  return socket;
};
```

5. API層の改善
5.1 レスポンス処理の統一
```ts
// 統一されたレスポンス処理
class HeuristicsApiService {
  private baseUrl = '/api/heuristics';
  
  private async handleResponse<T>(response: Response): Promise<T> {
    if (!response.ok) {
      await this.handleError(response);
    }
    
    const data = await response.json();
    
    // レスポンス構造の正規化
    if (data.success === false) {
      throw new ApplicationError(
        ErrorCode.API_ERROR,
        data.message || 'API呼び出しに失敗しました'
      );
    }
    
    return data.data || data;
  }
  
  private async handleError(response: Response): Promise<never> {
    let errorData: any;
    
    try {
      errorData = await response.json();
    } catch {
      errorData = { message: 'Unknown error' };
    }
    
    const errorCode = this.mapHttpStatusToErrorCode(response.status);
    throw new ApplicationError(errorCode, errorData.message);
  }
  
  private mapHttpStatusToErrorCode(status: number): ErrorCode {
    switch (status) {
      case 400: return ErrorCode.VAL_INVALID_INPUT;
      case 401: return ErrorCode.AUTH_INVALID_CREDENTIALS;
      case 403: return ErrorCode.AUTH_UNAUTHORIZED;
      case 404: return ErrorCode.RES_NOT_FOUND;
      case 500: return ErrorCode.SYS_INTERNAL_ERROR;
      default: return ErrorCode.SYS_INTERNAL_ERROR;
    }
  }
}
```

5.2 キャッシュ戦略
```ts
// インメモリキャッシュ
class HeuristicsCache {
  private cache = new Map<string, { data: any; timestamp: number; ttl: number }>();
  
  set(key: string, data: any, ttl: number = 5 * 60 * 1000): void {
    this.cache.set(key, {
      data,
      timestamp: Date.now(),
      ttl
    });
  }
  
  get(key: string): any | null {
    const item = this.cache.get(key);
    
    if (!item) return null;
    
    if (Date.now() - item.timestamp > item.ttl) {
      this.cache.delete(key);
      return null;
    }
    
    return item.data;
  }
  
  invalidate(pattern: string): void {
    for (const key of this.cache.keys()) {
      if (key.includes(pattern)) {
        this.cache.delete(key);
      }
    }
  }
}
```

6. 実装優先順位

Phase 1: 基盤整備（1-2週間）
データモデルの型定義強化
エラーハンドリングの統一
基本的なバリデーション実装

Phase 2: Service層改善（1週間）
APIサービスの統合
キャッシュ戦略の実装
レスポンス処理の標準化

Phase 3: Redux最適化（1週間）
状態管理の階層化
非同期処理の改善
パフォーマンス最適化

Phase 4: コンポーネント改善（1-2週間）
メモ化と仮想化
リアルタイム更新対応
エラーバウンダリの実装

Phase 5: テストとドキュメント（1週間）
単体テストの追加
統合テストの実装
API仕様書の更新

このリファクタリングにより、詳細設計書のパターンに準拠し、保守性、拡張性、パフォーマンスが大幅に向上します。特に、エラーハンドリングの統一、型安全性の向上、パフォーマンス最適化が重要な改善点となります。

