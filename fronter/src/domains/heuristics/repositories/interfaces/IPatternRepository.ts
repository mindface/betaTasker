import { Pattern, PatternEntity, PatternId, PatternRequest, PatternName, PatternCategory } from '../../models/Pattern';
import { PatternMetrics } from '../../valueObjects/PatternMetrics';

export interface PatternFilters {
  category?: PatternCategory;
  minFrequency?: number;
  minAccuracy?: number;
  dateFrom?: Date;
  dateTo?: Date;
  isActive?: boolean;
}

export interface PatternPagination {
  limit: number;
  offset: number;
}

export interface PatternSearchResult {
  patterns: Pattern[];
  total: number;
  hasMore: boolean;
}

export interface PatternSimilarity {
  pattern: Pattern;
  similarity: number;
}

export interface PatternStats {
  totalPatterns: number;
  byCategory: Record<PatternCategory, number>;
  averageAccuracy: number;
  averageFrequency: number;
  highQualityPatterns: number;
  activePatterns: number;
  topPatterns: Pattern[];
}

export interface IPatternRepository {
  /**
   * パターンを保存する
   */
  save(pattern: PatternEntity): Promise<Pattern>;

  /**
   * IDでパターンを取得する
   */
  findById(id: PatternId): Promise<Pattern | null>;

  /**
   * 複数のIDでパターンを取得する
   */
  findByIds(ids: PatternId[]): Promise<Pattern[]>;

  /**
   * フィルタ条件に基づいてパターンを検索する
   */
  findByFilters(
    filters: PatternFilters,
    pagination?: PatternPagination
  ): Promise<PatternSearchResult>;

  /**
   * 名前でパターンを検索する
   */
  findByName(name: PatternName): Promise<Pattern | null>;

  /**
   * 名前の部分一致でパターンを検索する
   */
  findByNameContaining(
    namePattern: string,
    pagination?: PatternPagination
  ): Promise<PatternSearchResult>;

  /**
   * カテゴリ別のパターンを取得する
   */
  findByCategory(
    category: PatternCategory,
    pagination?: PatternPagination
  ): Promise<PatternSearchResult>;

  /**
   * 高品質なパターンを取得する
   */
  findHighQualityPatterns(
    threshold?: number,
    pagination?: PatternPagination
  ): Promise<PatternSearchResult>;

  /**
   * 頻出パターンを取得する
   */
  findFrequentPatterns(
    threshold?: number,
    pagination?: PatternPagination
  ): Promise<PatternSearchResult>;

  /**
   * 最新のパターンを取得する
   */
  findLatest(count: number): Promise<Pattern[]>;

  /**
   * 類似したパターンを検索する
   */
  findSimilarPatterns(
    pattern: Pattern,
    similarityThreshold?: number,
    limit?: number
  ): Promise<PatternSimilarity[]>;

  /**
   * パターンを更新する
   */
  update(pattern: PatternEntity): Promise<Pattern>;

  /**
   * パターンを削除する
   */
  delete(id: PatternId): Promise<boolean>;

  /**
   * 複数のパターンを削除する
   */
  deleteByIds(ids: PatternId[]): Promise<number>;

  /**
   * パターンの統計を取得する
   */
  getStats(filters?: PatternFilters): Promise<PatternStats>;

  /**
   * カテゴリ別の統計を取得する
   */
  getStatsByCategory(): Promise<Record<PatternCategory, PatternStats>>;

  /**
   * パターンの総数を取得する
   */
  count(filters?: PatternFilters): Promise<number>;

  /**
   * パターンが存在するかチェックする
   */
  exists(id: PatternId): Promise<boolean>;

  /**
   * 名前が存在するかチェックする
   */
  existsByName(name: PatternName): Promise<boolean>;

  /**
   * 重複するパターンを検索する
   */
  findDuplicates(pattern: PatternRequest): Promise<Pattern[]>;

  /**
   * バッチでパターンを保存する
   */
  saveBatch(patterns: PatternEntity[]): Promise<Pattern[]>;

  /**
   * パターンのメトリクスを更新する
   */
  updateMetrics(id: PatternId, metrics: PatternMetrics): Promise<Pattern>;

  /**
   * パターンの頻度を増加させる
   */
  incrementFrequency(id: PatternId): Promise<Pattern>;

  /**
   * 最後に見た日時を更新する
   */
  updateLastSeen(id: PatternId, timestamp?: Date): Promise<Pattern>;

  /**
   * アクティブなパターンを取得する
   */
  findActive(pagination?: PatternPagination): Promise<PatternSearchResult>;

  /**
   * 非アクティブなパターンを取得する
   */
  findInactive(pagination?: PatternPagination): Promise<PatternSearchResult>;

  /**
   * パターンをアクティブ化する
   */
  activate(id: PatternId): Promise<Pattern>;

  /**
   * パターンを非アクティブ化する
   */
  deactivate(id: PatternId): Promise<Pattern>;

  /**
   * 古いパターンを削除する（データ保持期間管理）
   */
  deleteOldPatterns(olderThan: Date, keepHighQuality?: boolean): Promise<number>;

  /**
   * パターンのランキングを取得する
   */
  getRanking(
    criteria: 'frequency' | 'accuracy' | 'overall',
    limit?: number
  ): Promise<Pattern[]>;

  /**
   * パターンの時系列データを取得する
   */
  getTimeSeriesData(
    id: PatternId,
    dateRange: { from: Date; to: Date },
    interval: 'hour' | 'day' | 'week'
  ): Promise<Array<{ timestamp: Date; frequency: number; accuracy: number }>>;
}