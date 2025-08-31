"use client"
import React, { useEffect, useState } from 'react';
import { useHeuristicsPatterns } from '../../hooks/useHeuristics';
import { useHeuristics } from '../../hooks/useHeuristics';
import { HeuristicsTrainRequest } from '../../model/heuristics';
import styles from './HeuristicsPatterns.module.scss';

export default function HeuristicsPatterns() {
  const { patterns, loading, error, getPatterns, clearError } = useHeuristicsPatterns();
  const { currentModel, modelLoading, modelError, trainModel } = useHeuristics();
  const [filters, setFilters] = useState({
    user_id: '',
    data_type: 'all',
    period: 'week',
  });
  const [showTrainForm, setShowTrainForm] = useState(false);
  const [trainForm, setTrainForm] = useState<HeuristicsTrainRequest>({
    model_type: 'pattern_detection',
    parameters: {},
    data_source: 'user_behavior',
    training_data: [],
  });

  useEffect(() => {
    loadPatterns();
  }, []);

  const loadPatterns = () => {
    const params = {
      ...(filters.user_id && { user_id: filters.user_id }),
      ...(filters.data_type !== 'all' && { data_type: filters.data_type }),
      ...(filters.period && { period: filters.period }),
    };
    console.log("Loading patterns with params:", params);
    getPatterns(params);
  };

  const handleFilterChange = (field: string, value: string) => {
    setFilters(prev => ({
      ...prev,
      [field]: value,
    }));
  };

  const handleTrainModel = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await trainModel(trainForm);
      setShowTrainForm(false);
      setTrainForm({
        model_type: 'pattern_detection',
        parameters: {},
        data_source: 'user_behavior',
        training_data: [],
      });
    } catch (err) {
      console.error('Model training failed:', err);
    }
  };

  const handleTrainFormChange = (field: string, value: any) => {
    setTrainForm(prev => ({
      ...prev,
      [field]: value,
    }));
  };

  const dataTypes = [
    { value: 'all', label: 'すべて' },
    { value: 'task', label: 'タスク関連' },
    { value: 'navigation', label: 'ナビゲーション' },
    { value: 'interaction', label: 'インタラクション' },
    { value: 'performance', label: 'パフォーマンス' },
  ];

  const periods = [
    { value: 'day', label: '今日' },
    { value: 'week', label: '今週' },
    { value: 'month', label: '今月' },
    { value: 'year', label: '今年' },
  ];

  const modelTypes = [
    { value: 'pattern_detection', label: 'パターン検出' },
    { value: 'behavior_prediction', label: '行動予測' },
    { value: 'anomaly_detection', label: '異常検出' },
    { value: 'recommendation', label: '推薦システム' },
  ];

  if (loading) {
    return (
      <div className={styles.container}>
        <div className={styles.loading}>パターンを検出中...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className={styles.container}>
        <div className={styles.error}>
          <p>エラーが発生しました: {error}</p>
          <button onClick={clearError} className={styles.retryButton}>
            再試行
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}>
        <h2>パターン検出</h2>
        <div className={styles.controls}>
          <button onClick={() => setShowTrainForm(!showTrainForm)} className={styles.trainButton}>
            モデルトレーニング
          </button>
        </div>
      </div>

      <div className={styles.filters}>
        <div className={styles.filterGroup}>
          <label>ユーザーID:</label>
          <input
            type="text"
            value={filters.user_id}
            onChange={(e) => handleFilterChange('user_id', e.target.value)}
            placeholder="ユーザーID"
            className={styles.filterInput}
          />
        </div>
        <div className={styles.filterGroup}>
          <label>データタイプ:</label>
          <select
            value={filters.data_type}
            onChange={(e) => handleFilterChange('data_type', e.target.value)}
            className={styles.filterSelect}
          >
            {dataTypes.map(type => (
              <option key={type.value} value={type.value}>
                {type.label}
              </option>
            ))}
          </select>
        </div>
        <div className={styles.filterGroup}>
          <label>期間:</label>
          <select
            value={filters.period}
            onChange={(e) => handleFilterChange('period', e.target.value)}
            className={styles.filterSelect}
          >
            {periods.map(period => (
              <option key={period.value} value={period.value}>
                {period.label}
              </option>
            ))}
          </select>
        </div>
        <button onClick={loadPatterns} className={styles.applyButton}>
          フィルター適用
        </button>
      </div>

      {showTrainForm && (
        <div className={styles.trainForm}>
          <h3>新しいモデルをトレーニング</h3>
          <form onSubmit={handleTrainModel}>
            <div className={styles.formRow}>
              <div className={styles.formGroup}>
                <label>モデルタイプ:</label>
                <select
                  value={trainForm.model_type}
                  onChange={(e) => handleTrainFormChange('model_type', e.target.value)}
                  required
                >
                  {modelTypes.map(type => (
                    <option key={type.value} value={type.value}>
                      {type.label}
                    </option>
                  ))}
                </select>
              </div>
              <div className={styles.formGroup}>
                <label>データソース:</label>
                <input
                  type="text"
                  value={trainForm.data_source}
                  onChange={(e) => handleTrainFormChange('data_source', e.target.value)}
                  placeholder="例: user_behavior, task_completion"
                  required
                />
              </div>
            </div>
            
            <div className={styles.formGroup}>
              <label>パラメータ (JSON):</label>
              <textarea
                value={JSON.stringify(trainForm.parameters, null, 2)}
                onChange={(e) => {
                  try {
                    handleTrainFormChange('parameters', JSON.parse(e.target.value));
                  } catch {
                    // Invalid JSON, keep as string for now
                  }
                }}
                placeholder='{"learning_rate": 0.01, "epochs": 100}'
                rows={3}
              />
            </div>
            
            <div className={styles.formActions}>
              <button type="submit" className={styles.submitButton} disabled={modelLoading}>
                {modelLoading ? 'トレーニング中...' : 'トレーニング開始'}
              </button>
              <button 
                type="button" 
                onClick={() => setShowTrainForm(false)}
                className={styles.cancelButton}
              >
                キャンセル
              </button>
            </div>
          </form>
        </div>
      )}

      {currentModel && (
        <div className={styles.modelInfo}>
          <h3>現在のモデル</h3>
          <div className={styles.modelCard}>
            <div className={styles.modelHeader}>
              <span className={styles.modelType}>{currentModel.model_type}</span>
              <span className={styles.version}>v{currentModel.version}</span>
              <span className={`${styles.status} ${styles[currentModel.status]}`}>
                {currentModel.status}
              </span>
            </div>
            <div className={styles.modelDetails}>
              <div>トレーニング日: {new Date(currentModel.trained_at).toLocaleString('ja-JP')}</div>
              <div>作成日: {new Date(currentModel.created_at).toLocaleString('ja-JP')}</div>
            </div>
          </div>
        </div>
      )}

      <div className={styles.patternsList}>
        {patterns.length === 0 ? (
          <div className={styles.noData}>パターンが検出されていません</div>
        ) : (
          <>
            <div className={styles.stats}>
              <div className={styles.statCard}>
                <h4>検出パターン数</h4>
                <span>{patterns.length}</span>
              </div>
              <div className={styles.statCard}>
                <h4>平均精度</h4>
                <span>
                  {patterns.length > 0 
                    ? (patterns.reduce((sum, p) => sum + p.accuracy, 0) / patterns.length * 100).toFixed(1)
                    : 0}%
                </span>
              </div>
              <div className={styles.statCard}>
                <h4>最高頻度</h4>
                <span>
                  {patterns.length > 0 
                    ? Math.max(...patterns.map(p => p.frequency))
                    : 0}
                </span>
              </div>
            </div>

            <div className={styles.patternsGrid}>
              {patterns.map((pattern) => (
                <div key={pattern.id} className={styles.patternCard}>
                  <div className={styles.patternHeader}>
                    <h4>{pattern.name}</h4>
                    <span className={styles.category}>{pattern.category}</span>
                  </div>
                  <div className={styles.patternMetrics}>
                    <div className={styles.metric}>
                      <span className={styles.label}>頻度:</span>
                      <span className={styles.value}>{pattern.frequency}</span>
                    </div>
                    <div className={styles.metric}>
                      <span className={styles.label}>精度:</span>
                      <span className={styles.value}>{(pattern.accuracy * 100).toFixed(1)}%</span>
                    </div>
                    <div className={styles.metric}>
                      <span className={styles.label}>最終検出:</span>
                      <span className={styles.value}>
                        {new Date(pattern.last_seen).toLocaleDateString('ja-JP')}
                      </span>
                    </div>
                  </div>
                  <div className={styles.patternData}>
                    <strong>パターンデータ:</strong>
                    <pre>{JSON.stringify(JSON.parse(pattern.pattern), null, 2)}</pre>
                  </div>
                </div>
              ))}
            </div>
          </>
        )}
      </div>
    </div>
  );
}