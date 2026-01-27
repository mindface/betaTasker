import { 
  Analysis, 
  AnalysisEntity, 
  AnalysisId, 
  AnalysisRequest, 
  UserId, 
  TaskId,
  AnalysisScoreValue 
} from '../models/Analysis';
import { AnalysisType } from '../valueObjects/AnalysisType';
import { IAnalysisRepository } from '../repositories/interfaces/IAnalysisRepository';

export interface AnalysisResult {
  success: boolean;
  analysis?: Analysis;
  error?: string;
}

export interface BatchAnalysisResult {
  success: boolean;
  analyses: Analysis[];
  failed: Array<{ request: AnalysisRequest; error: string }>;
}

export class AnalysisService {
  constructor(
    private readonly analysisRepository: IAnalysisRepository
  ) {}

  /**
   * 新しい分析を実行する
   */
  async executeAnalysis(request: AnalysisRequest): Promise<AnalysisResult> {
    try {
      // リクエストの検証
      const validationResult = this.validateAnalysisRequest(request);
      if (!validationResult.isValid) {
        return {
          success: false,
          error: validationResult.errors.join(', ')
        };
      }

      // 重複チェック
      const duplicates = await this.analysisRepository.findDuplicates(request);
      if (duplicates.length > 0) {
        // 最新の重複分析を返す
        return {
          success: true,
          analysis: duplicates[0]
        };
      }

      // 分析エンティティを作成
      const analysisEntity = AnalysisEntity.create(request);
      const pendingAnalysis = await this.analysisRepository.save(analysisEntity);

      // 分析を実行（非同期）
      this.performAnalysis(pendingAnalysis.id)
        .catch(error => console.error('Analysis failed:', error));

      return {
        success: true,
        analysis: pendingAnalysis
      };
    } catch (error) {
      console.error('Failed to execute analysis:', error);
      return {
        success: false,
        error: error instanceof Error ? error.message : '分析の実行に失敗しました'
      };
    }
  }

  /**
   * 分析結果を取得する
   */
  async getAnalysis(id: AnalysisId): Promise<Analysis | null> {
    return this.analysisRepository.findById(id);
  }

  /**
   * ユーザーの分析履歴を取得する
   */
  async getUserAnalysisHistory(userId: UserId, limit: number = 20): Promise<Analysis[]> {
    const result = await this.analysisRepository.findByUserId(userId, { limit, offset: 0 });
    return result.analyses;
  }

  /**
   * 分析タイプ別の結果を取得する
   */
  async getAnalysesByType(analysisType: AnalysisType, limit: number = 20): Promise<Analysis[]> {
    const result = await this.analysisRepository.findByAnalysisType(analysisType, { limit, offset: 0 });
    return result.analyses;
  }

  /**
   * バッチ分析を実行する
   */
  async executeBatchAnalysis(requests: AnalysisRequest[]): Promise<BatchAnalysisResult> {
    const results: Analysis[] = [];
    const failed: Array<{ request: AnalysisRequest; error: string }> = [];

    for (const request of requests) {
      const result = await this.executeAnalysis(request);
      if (result.success && result.analysis) {
        results.push(result.analysis);
      } else {
        failed.push({
          request,
          error: result.error || '不明なエラー'
        });
      }
    }

    return {
      success: true,
      analyses: results,
      failed
    };
  }

  /**
   * 分析統計を取得する
   */
  async getAnalysisStatistics(userId?: UserId): Promise<{
    total: number;
    completed: number;
    pending: number;
    failed: number;
    averageScore: number;
    byType: Record<string, number>;
  }> {
    if (userId) {
      const stats = await this.analysisRepository.getStatsByUserId(userId);
      return {
        total: stats.total,
        completed: stats.byStatus['completed'] || 0,
        pending: stats.byStatus['pending'] || 0,
        failed: stats.byStatus['failed'] || 0,
        averageScore: stats.averageScore,
        byType: stats.byType
      };
    }

    // 全体統計の実装（リポジトリに追加メソッドが必要）
    return {
      total: 0,
      completed: 0,
      pending: 0,
      failed: 0,
      averageScore: 0,
      byType: {}
    };
  }

  /**
   * 分析を再実行する
   */
  async retryAnalysis(id: AnalysisId): Promise<AnalysisResult> {
    try {
      const analysis = await this.analysisRepository.findById(id);
      if (!analysis) {
        return {
          success: false,
          error: '分析が見つかりません'
        };
      }

      if (!analysis.canRetry()) {
        return {
          success: false,
          error: 'この分析は再実行できません'
        };
      }

      // ステータスをpendingに更新
      const updatedAnalysis = analysis.updateStatus('pending');
      await this.analysisRepository.update(updatedAnalysis as AnalysisEntity);

      // 分析を実行（非同期）
      this.performAnalysis(id)
        .catch(error => console.error('Analysis retry failed:', error));

      return {
        success: true,
        analysis: updatedAnalysis
      };
    } catch (error) {
      console.error('Failed to retry analysis:', error);
      return {
        success: false,
        error: error instanceof Error ? error.message : '分析の再実行に失敗しました'
      };
    }
  }

  /**
   * 分析リクエストの検証
   */
  private validateAnalysisRequest(request: AnalysisRequest): {
    isValid: boolean;
    errors: string[];
  } {
    const errors: string[] = [];

    if (!request.userId || request.userId.value <= 0) {
      errors.push('有効なユーザーIDが必要です');
    }

    if (!request.analysisType) {
      errors.push('分析タイプが必要です');
    }

    if (!request.data || Object.keys(request.data).length === 0) {
      errors.push('分析データが必要です');
    }

    // 分析タイプ別の検証
    const analysisType = AnalysisType.fromString(request.analysisType);
    if (analysisType.requiresTaskId() && (!request.taskId || request.taskId.value <= 0)) {
      errors.push(`${analysisType.getDisplayName()}にはタスクIDが必要です`);
    }

    return {
      isValid: errors.length === 0,
      errors
    };
  }

  /**
   * 実際の分析処理を実行する
   */
  private async performAnalysis(analysisId: AnalysisId): Promise<void> {
    try {
      const analysis = await this.analysisRepository.findById(analysisId);
      if (!analysis) {
        throw new Error('Analysis not found');
      }

      // ステータスを処理中に更新
      const processingAnalysis = analysis.updateStatus('processing');
      await this.analysisRepository.update(processingAnalysis as AnalysisEntity);

      // 分析タイプに基づいて実際の分析を実行
      const analysisType = AnalysisType.fromString(analysis.analysisType);
      const result = await this.executeSpecificAnalysis(analysisType, analysis);

      // 結果を更新
      const completedAnalysis = analysis.updateResult(
        result.result,
        new AnalysisScoreValue(result.score)
      );
      await this.analysisRepository.update(completedAnalysis as AnalysisEntity);

    } catch (error) {
      console.error('Analysis execution failed:', error);
      
      // エラー状態に更新
      try {
        const analysis = await this.analysisRepository.findById(analysisId);
        if (analysis) {
          const failedAnalysis = analysis.updateStatus('failed');
          await this.analysisRepository.update(failedAnalysis as AnalysisEntity);
        }
      } catch (updateError) {
        console.error('Failed to update analysis status:', updateError);
      }
    }
  }

  /**
   * 特定の分析タイプの処理を実行する
   */
  private async executeSpecificAnalysis(
    analysisType: AnalysisType,
    analysis: Analysis
  ): Promise<{ result: string; score: number }> {
    // 実際の分析ロジックは分析タイプによって異なる
    // ここではシミュレーション
    
    const delay = analysisType.getExpectedDuration();
    await this.delay(Math.min(delay, 10000)); // 最大10秒

    switch (analysisType.getValue()) {
      case 'performance':
        return this.analyzePerformance(analysis);
      case 'behavior':
        return this.analyzeBehavior(analysis);
      case 'pattern':
        return this.analyzePattern(analysis);
      case 'cognitive':
        return this.analyzeCognitive(analysis);
      case 'efficiency':
        return this.analyzeEfficiency(analysis);
      default:
        throw new Error(`Unknown analysis type: ${analysisType.getValue()}`);
    }
  }

  private async analyzePerformance(analysis: Analysis): Promise<{ result: string; score: number }> {
    // パフォーマンス分析のロジック
    const mockScore = Math.random() * 100;
    return {
      result: `パフォーマンス分析が完了しました。総合スコア: ${mockScore.toFixed(1)}点`,
      score: mockScore
    };
  }

  private async analyzeBehavior(analysis: Analysis): Promise<{ result: string; score: number }> {
    // 行動分析のロジック
    const mockScore = Math.random() * 100;
    return {
      result: `行動分析が完了しました。行動パターンスコア: ${mockScore.toFixed(1)}点`,
      score: mockScore
    };
  }

  private async analyzePattern(analysis: Analysis): Promise<{ result: string; score: number }> {
    // パターン分析のロジック
    const mockScore = Math.random() * 100;
    return {
      result: `パターン分析が完了しました。パターン検出スコア: ${mockScore.toFixed(1)}点`,
      score: mockScore
    };
  }

  private async analyzeCognitive(analysis: Analysis): Promise<{ result: string; score: number }> {
    // 認知分析のロジック
    const mockScore = Math.random() * 100;
    return {
      result: `認知分析が完了しました。認知負荷スコア: ${mockScore.toFixed(1)}点`,
      score: mockScore
    };
  }

  private async analyzeEfficiency(analysis: Analysis): Promise<{ result: string; score: number }> {
    // 効率性分析のロジック
    const mockScore = Math.random() * 100;
    return {
      result: `効率性分析が完了しました。効率性スコア: ${mockScore.toFixed(1)}点`,
      score: mockScore
    };
  }

  private delay(ms: number): Promise<void> {
    return new Promise(resolve => setTimeout(resolve, ms));
  }
}