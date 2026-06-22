import { 
  Insight, 
  InsightEntity, 
  InsightId, 
  InsightRequest, 
  InsightType, 
  InsightImpact,
  InsightTitle 
} from '../../models/Insight';

export interface InsightFilters {
  type?: InsightType;
  impact?: InsightImpact;
  isActive?: boolean;
  minConfidence?: number;
  dateFrom?: Date;
  dateTo?: Date;
  hasActions?: boolean;
}

export interface InsightPagination {
  limit: number;
  offset: number;
}

export interface InsightSearchResult {
  insights: Insight[];
  total: number;
  hasMore: boolean;
}

export interface InsightStats {
  totalInsights: number;
  activeInsights: number;
  byType: Record<InsightType, number>;
  byImpact: Record<InsightImpact, number>;
  averageConfidence: number;
  highConfidenceInsights: number;
  criticalInsights: number;
  implementedInsights: number;
}

export interface IInsightRepository {
  /**
   * インサイトを保存する
   */
  save(insight: InsightEntity): Promise<Insight>;

  /**
   * IDでインサイトを取得する
   */
  findById(id: InsightId): Promise<Insight | null>;

  /**
   * 複数のIDでインサイトを取得する
   */
  findByIds(ids: InsightId[]): Promise<Insight[]>;

  /**
   * フィルタ条件に基づいてインサイトを検索する
   */
  findByFilters(
    filters: InsightFilters,
    pagination?: InsightPagination
  ): Promise<InsightSearchResult>;

  /**
   * タイトルでインサイトを検索する
   */
  findByTitle(title: InsightTitle): Promise<Insight | null>;

  /**
   * タイトルの部分一致でインサイトを検索する
   */
  findByTitleContaining(
    titlePattern: string,
    pagination?: InsightPagination
  ): Promise<InsightSearchResult>;

  /**
   * タイプ別のインサイトを取得する
   */
  findByType(
    type: InsightType,
    pagination?: InsightPagination
  ): Promise<InsightSearchResult>;

  /**
   * インパクト別のインサイトを取得する
   */
  findByImpact(
    impact: InsightImpact,
    pagination?: InsightPagination
  ): Promise<InsightSearchResult>;

  /**
   * 高信頼度のインサイトを取得する
   */
  findHighConfidenceInsights(
    threshold?: number,
    pagination?: InsightPagination
  ): Promise<InsightSearchResult>;

  /**
   * クリティカルなインサイトを取得する
   */
  findCriticalInsights(pagination?: InsightPagination): Promise<InsightSearchResult>;

  /**
   * アクティブなインサイトを取得する
   */
  findActive(pagination?: InsightPagination): Promise<InsightSearchResult>;

  /**
   * 最新のインサイトを取得する
   */
  findLatest(count: number): Promise<Insight[]>;

  /**
   * 優先度の高いインサイトを取得する
   */
  findByPriority(
    limit?: number
  ): Promise<Insight[]>;

  /**
   * 未実装のインサイトを取得する
   */
  findUnimplemented(pagination?: InsightPagination): Promise<InsightSearchResult>;

  /**
   * インサイトを更新する
   */
  update(insight: InsightEntity): Promise<Insight>;

  /**
   * インサイトを削除する
   */
  delete(id: InsightId): Promise<boolean>;

  /**
   * 複数のインサイトを削除する
   */
  deleteByIds(ids: InsightId[]): Promise<number>;

  /**
   * インサイトの統計を取得する
   */
  getStats(filters?: InsightFilters): Promise<InsightStats>;

  /**
   * タイプ別の統計を取得する
   */
  getStatsByType(): Promise<Record<InsightType, InsightStats>>;

  /**
   * インパクト別の統計を取得する
   */
  getStatsByImpact(): Promise<Record<InsightImpact, InsightStats>>;

  /**
   * インサイトの総数を取得する
   */
  count(filters?: InsightFilters): Promise<number>;

  /**
   * インサイトが存在するかチェックする
   */
  exists(id: InsightId): Promise<boolean>;

  /**
   * タイトルが存在するかチェックする
   */
  existsByTitle(title: InsightTitle): Promise<boolean>;

  /**
   * 重複するインサイトを検索する
   */
  findDuplicates(insight: InsightRequest): Promise<Insight[]>;

  /**
   * バッチでインサイトを保存する
   */
  saveBatch(insights: InsightEntity[]): Promise<Insight[]>;

  /**
   * インサイトをアクティブ化する
   */
  activate(id: InsightId): Promise<Insight>;

  /**
   * インサイトを非アクティブ化する
   */
  deactivate(id: InsightId): Promise<Insight>;

  /**
   * インサイトの信頼度を更新する
   */
  updateConfidence(id: InsightId, confidence: number): Promise<Insight>;

  /**
   * インサイトにアクションを追加する
   */
  addAction(id: InsightId, action: string, priority?: number): Promise<Insight>;

  /**
   * インサイトからアクションを削除する
   */
  removeAction(id: InsightId, action: string): Promise<Insight>;

  /**
   * インサイトを実装済みとしてマークする
   */
  markAsImplemented(id: InsightId): Promise<Insight>;

  /**
   * インサイトを未実装としてマークする
   */
  markAsUnimplemented(id: InsightId): Promise<Insight>;

  /**
   * 古いインサイトを削除する（データ保持期間管理）
   */
  deleteOldInsights(olderThan: Date, keepCritical?: boolean): Promise<number>;

  /**
   * インサイトのランキングを取得する
   */
  getRanking(
    criteria: 'confidence' | 'impact' | 'priority' | 'created',
    limit?: number
  ): Promise<Insight[]>;

  /**
   * 関連するインサイトを検索する
   */
  findRelated(
    insight: Insight,
    limit?: number
  ): Promise<Insight[]>;

  /**
   * インサイトの時系列データを取得する
   */
  getTimeSeriesData(
    filters: InsightFilters,
    interval: 'hour' | 'day' | 'week'
  ): Promise<Array<{ timestamp: Date; count: number; avgConfidence: number }>>;

  /**
   * インサイトのダッシュボードデータを取得する
   */
  getDashboardData(): Promise<{
    totalInsights: number;
    newInsights: number;
    criticalInsights: number;
    implementedInsights: number;
    topInsights: Insight[];
    trends: Array<{ date: string; count: number }>;
  }>;
}