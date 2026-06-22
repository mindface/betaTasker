import { Tracking, TrackingEntity, TrackingId, TrackingRequest, UserId, SessionId, TrackingAction } from '../../models/Tracking';

export interface TrackingFilters {
  userId?: UserId;
  sessionId?: SessionId;
  action?: TrackingAction;
  dateFrom?: Date;
  dateTo?: Date;
  minDuration?: number;
  maxDuration?: number;
}

export interface TrackingPagination {
  limit: number;
  offset: number;
}

export interface TrackingSearchResult {
  trackingData: Tracking[];
  total: number;
  hasMore: boolean;
}

export interface TrackingStats {
  totalActions: number;
  uniqueActions: number;
  averageDuration: number;
  totalDuration: number;
  sessionCount: number;
  mostFrequentActions: Array<{ action: string; count: number }>;
  dailyStats: Array<{ date: string; count: number }>;
}

export interface ITrackingRepository {
  /**
   * トラッキングデータを保存する
   */
  save(tracking: TrackingEntity): Promise<Tracking>;

  /**
   * IDでトラッキングデータを取得する
   */
  findById(id: TrackingId): Promise<Tracking | null>;

  /**
   * 複数のIDでトラッキングデータを取得する
   */
  findByIds(ids: TrackingId[]): Promise<Tracking[]>;

  /**
   * フィルタ条件に基づいてトラッキングデータを検索する
   */
  findByFilters(
    filters: TrackingFilters,
    pagination?: TrackingPagination
  ): Promise<TrackingSearchResult>;

  /**
   * ユーザーのトラッキングデータを取得する
   */
  findByUserId(userId: UserId, pagination?: TrackingPagination): Promise<TrackingSearchResult>;

  /**
   * セッションのトラッキングデータを取得する
   */
  findBySessionId(sessionId: SessionId, pagination?: TrackingPagination): Promise<TrackingSearchResult>;

  /**
   * アクション別のトラッキングデータを取得する
   */
  findByAction(action: TrackingAction, pagination?: TrackingPagination): Promise<TrackingSearchResult>;

  /**
   * 最新のトラッキングデータを取得する
   */
  findLatest(count: number): Promise<Tracking[]>;

  /**
   * ユーザーの最新のトラッキングデータを取得する
   */
  findLatestByUserId(userId: UserId, count: number): Promise<Tracking[]>;

  /**
   * 期間内のトラッキングデータを取得する
   */
  findByDateRange(
    dateFrom: Date,
    dateTo: Date,
    pagination?: TrackingPagination
  ): Promise<TrackingSearchResult>;

  /**
   * トラッキングデータを更新する
   */
  update(tracking: TrackingEntity): Promise<Tracking>;

  /**
   * トラッキングデータを削除する
   */
  delete(id: TrackingId): Promise<boolean>;

  /**
   * 複数のトラッキングデータを削除する
   */
  deleteByIds(ids: TrackingId[]): Promise<number>;

  /**
   * ユーザーのトラッキングデータ統計を取得する
   */
  getStatsByUserId(userId: UserId, dateRange?: { from: Date; to: Date }): Promise<TrackingStats>;

  /**
   * セッションのトラッキングデータ統計を取得する
   */
  getStatsBySessionId(sessionId: SessionId): Promise<TrackingStats>;

  /**
   * 全体のトラッキングデータ統計を取得する
   */
  getGlobalStats(dateRange?: { from: Date; to: Date }): Promise<TrackingStats>;

  /**
   * トラッキングデータの総数を取得する
   */
  count(filters?: TrackingFilters): Promise<number>;

  /**
   * トラッキングデータが存在するかチェックする
   */
  exists(id: TrackingId): Promise<boolean>;

  /**
   * ユニークなアクションを取得する
   */
  getUniqueActions(userId?: UserId): Promise<string[]>;

  /**
   * ユニークなセッションIDを取得する
   */
  getUniqueSessionIds(userId?: UserId): Promise<string[]>;

  /**
   * バッチでトラッキングデータを保存する
   */
  saveBatch(trackingData: TrackingEntity[]): Promise<Tracking[]>;

  /**
   * 古いトラッキングデータを削除する（データ保持期間管理）
   */
  deleteOldData(olderThan: Date): Promise<number>;

  /**
   * ユーザーのアクションシーケンスを取得する
   */
  getActionSequence(
    userId: UserId,
    sessionId?: SessionId,
    limit?: number
  ): Promise<Tracking[]>;

  /**
   * トラッキングデータのサマリーを取得する
   */
  getSummary(
    filters: TrackingFilters,
    groupBy: 'day' | 'hour' | 'action' | 'user'
  ): Promise<Array<{ group: string; count: number; avgDuration: number }>>;
}