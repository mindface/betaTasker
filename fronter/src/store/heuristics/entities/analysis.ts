import { createEntityAdapter, EntityState } from '@reduxjs/toolkit';
import { HeuristicsAnalysis } from '../../../model/heuristics';

export interface AnalysisEntity extends HeuristicsAnalysis {
  // 追加のメタデータ
  _metadata?: {
    lastUpdated: string;
    cacheExpiry?: string;
    version: number;
  };
}

// Entity Adapter の作成
export const analysisAdapter = createEntityAdapter<AnalysisEntity>({
  selectId: (analysis) => analysis.id,
  sortComparer: (a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime()
});

// 初期状態の定義
export interface AnalysisEntityState extends EntityState<AnalysisEntity> {
  // 追加のメタデータ
  lastFetchTime: string | null;
  totalCount: number;
  hasMore: boolean;
}

export const initialAnalysisEntityState: AnalysisEntityState = analysisAdapter.getInitialState({
  lastFetchTime: null,
  totalCount: 0,
  hasMore: true
});

// 基本的なセレクタをエクスポート
export const {
  selectIds: selectAnalysisIds,
  selectEntities: selectAnalysisEntities,
  selectAll: selectAllAnalyses,
  selectTotal: selectAnalysisTotal,
  selectById: selectAnalysisById
} = analysisAdapter.getSelectors();

// カスタムセレクタ
export const analysisSelectors = {
  // 基本セレクタ
  ...analysisAdapter.getSelectors(),
  
  // カスタムセレクタ
  selectByUserId: (entities: Record<string, AnalysisEntity>, userId: number) =>
    Object.values(entities).filter(analysis => analysis.user_id === userId),
    
  selectByTaskId: (entities: Record<string, AnalysisEntity>, taskId: number) =>
    Object.values(entities).filter(analysis => analysis.task_id === taskId),
    
  selectByType: (entities: Record<string, AnalysisEntity>, analysisType: string) =>
    Object.values(entities).filter(analysis => analysis.analysis_type === analysisType),
    
  selectByStatus: (entities: Record<string, AnalysisEntity>, status: string) =>
    Object.values(entities).filter(analysis => analysis.status === status),
    
  selectRecent: (entities: Record<string, AnalysisEntity>, limit: number = 10) =>
    Object.values(entities)
      .sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime())
      .slice(0, limit),
      
  selectCompletedAnalyses: (entities: Record<string, AnalysisEntity>) =>
    Object.values(entities).filter(analysis => analysis.status === 'completed'),
    
  selectPendingAnalyses: (entities: Record<string, AnalysisEntity>) =>
    Object.values(entities).filter(analysis => analysis.status === 'pending'),
    
  selectFailedAnalyses: (entities: Record<string, AnalysisEntity>) =>
    Object.values(entities).filter(analysis => analysis.status === 'failed'),
    
  selectHighScoreAnalyses: (entities: Record<string, AnalysisEntity>, threshold: number = 80) =>
    Object.values(entities).filter(analysis => analysis.score >= threshold),
    
  selectAnalysisStats: (entities: Record<string, AnalysisEntity>) => {
    const analyses = Object.values(entities);
    const total = analyses.length;
    const completed = analyses.filter(a => a.status === 'completed').length;
    const pending = analyses.filter(a => a.status === 'pending').length;
    const failed = analyses.filter(a => a.status === 'failed').length;
    
    const completedAnalyses = analyses.filter(a => a.status === 'completed');
    const averageScore = completedAnalyses.length > 0
      ? completedAnalyses.reduce((sum, a) => sum + a.score, 0) / completedAnalyses.length
      : 0;
    
    const byType = analyses.reduce((acc, analysis) => {
      acc[analysis.analysis_type] = (acc[analysis.analysis_type] || 0) + 1;
      return acc;
    }, {} as Record<string, number>);
    
    return {
      total,
      completed,
      pending,
      failed,
      averageScore,
      byType,
      completionRate: total > 0 ? (completed / total) * 100 : 0
    };
  }
};

// エンティティ操作ヘルパー
export const analysisEntityHelpers = {
  // エンティティを正規化して追加
  addOne: analysisAdapter.addOne,
  addMany: analysisAdapter.addMany,
  
  // エンティティを更新
  updateOne: analysisAdapter.updateOne,
  updateMany: analysisAdapter.updateMany,
  
  // エンティティを削除
  removeOne: analysisAdapter.removeOne,
  removeMany: analysisAdapter.removeMany,
  removeAll: analysisAdapter.removeAll,
  
  // エンティティを置換
  setOne: analysisAdapter.setOne,
  setMany: analysisAdapter.setMany,
  setAll: analysisAdapter.setAll,
  
  // Upsert操作
  upsertOne: analysisAdapter.upsertOne,
  upsertMany: analysisAdapter.upsertMany,
  
  // カスタムヘルパー
  updateAnalysisStatus: (state: AnalysisEntityState, id: number, status: string) => {
    analysisAdapter.updateOne(state, {
      id,
      changes: {
        status,
        updated_at: new Date().toISOString(),
        _metadata: {
          lastUpdated: new Date().toISOString(),
          version: (state.entities[id]?._metadata?.version || 0) + 1
        }
      }
    });
  },
  
  updateAnalysisResult: (state: AnalysisEntityState, id: number, result: string, score: number) => {
    analysisAdapter.updateOne(state, {
      id,
      changes: {
        result,
        score,
        status: 'completed',
        updated_at: new Date().toISOString(),
        _metadata: {
          lastUpdated: new Date().toISOString(),
          version: (state.entities[id]?._metadata?.version || 0) + 1
        }
      }
    });
  },
  
  markAsExpired: (state: AnalysisEntityState, ids: number[]) => {
    const now = new Date();
    const expiry = new Date(now.getTime() + 30 * 60 * 1000); // 30分後
    
    analysisAdapter.updateMany(state, ids.map(id => ({
      id,
      changes: {
        _metadata: {
          ...state.entities[id]?._metadata,
          cacheExpiry: expiry.toISOString(),
          lastUpdated: now.toISOString(),
          version: (state.entities[id]?._metadata?.version || 0) + 1
        }
      }
    })));
  },
  
  removeExpired: (state: AnalysisEntityState) => {
    const now = new Date();
    const expiredIds = Object.entries(state.entities)
      .filter(([_, entity]) => {
        if (!entity?._metadata?.cacheExpiry) return false;
        return new Date(entity._metadata.cacheExpiry) < now;
      })
      .map(([id]) => Number(id));
    
    analysisAdapter.removeMany(state, expiredIds);
  }
};