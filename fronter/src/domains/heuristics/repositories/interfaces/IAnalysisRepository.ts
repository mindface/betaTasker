import { Analysis, AnalysisEntity, AnalysisId, AnalysisRequest, UserId, TaskId, AnalysisType } from '../../models/Analysis';

export interface AnalysisFilters {
  userId?: UserId;
  taskId?: TaskId;
  analysisType?: AnalysisType;
  status?: string;
  dateFrom?: Date;
  dateTo?: Date;
}

export interface AnalysisPagination {
  limit: number;
  offset: number;
}

export interface AnalysisSearchResult {
  analyses: Analysis[];
  total: number;
  hasMore: boolean;
}

export interface IAnalysisRepository {
  /**
   * 分析結果を保存する
   */
  save(analysis: AnalysisEntity): Promise<Analysis>;

  /**
   * IDで分析結果を取得する
   */
  findById(id: AnalysisId): Promise<Analysis | null>;

  /**
   * 複数のIDで分析結果を取得する
   */
  findByIds(ids: AnalysisId[]): Promise<Analysis[]>;

  /**
   * フィルタ条件に基づいて分析結果を検索する
   */
  findByFilters(
    filters: AnalysisFilters,
    pagination?: AnalysisPagination
  ): Promise<AnalysisSearchResult>;

  /**
   * ユーザーの分析結果を取得する
   */
  findByUserId(userId: UserId, pagination?: AnalysisPagination): Promise<AnalysisSearchResult>;

  /**
   * タスクの分析結果を取得する
   */
  findByTaskId(taskId: TaskId, pagination?: AnalysisPagination): Promise<AnalysisSearchResult>;

  /**
   * 分析タイプ別の分析結果を取得する
   */
  findByAnalysisType(
    analysisType: AnalysisType,
    pagination?: AnalysisPagination
  ): Promise<AnalysisSearchResult>;

  /**
   * 最新の分析結果を取得する
   */
  findLatest(count: number): Promise<Analysis[]>;

  /**
   * ユーザーの最新の分析結果を取得する
   */
  findLatestByUserId(userId: UserId, count: number): Promise<Analysis[]>;

  /**
   * 分析結果を更新する
   */
  update(analysis: AnalysisEntity): Promise<Analysis>;

  /**
   * 分析結果を削除する
   */
  delete(id: AnalysisId): Promise<boolean>;

  /**
   * 複数の分析結果を削除する
   */
  deleteByIds(ids: AnalysisId[]): Promise<number>;

  /**
   * ユーザーの分析結果の統計を取得する
   */
  getStatsByUserId(userId: UserId): Promise<{
    total: number;
    byType: Record<string, number>;
    byStatus: Record<string, number>;
    averageScore: number;
    lastAnalysisDate: Date | null;
  }>;

  /**
   * 分析結果の総数を取得する
   */
  count(filters?: AnalysisFilters): Promise<number>;

  /**
   * 分析結果が存在するかチェックする
   */
  exists(id: AnalysisId): Promise<boolean>;

  /**
   * 重複する分析結果を検索する
   */
  findDuplicates(analysis: AnalysisRequest): Promise<Analysis[]>;

  /**
   * バッチで分析結果を保存する
   */
  saveBatch(analyses: AnalysisEntity[]): Promise<Analysis[]>;
}