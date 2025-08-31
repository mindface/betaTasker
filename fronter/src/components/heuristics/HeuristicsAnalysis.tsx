"use client"
import React, { useState, useEffect } from 'react';
import { useHeuristicsAnalysis } from '../../hooks/useHeuristics';
import { HeuristicsAnalysisRequest } from '../../model/heuristics';
import styles from './HeuristicsAnalysis.module.scss';

export default function HeuristicsAnalysis() {
  const { analyses, currentAnalysis, loading, error, analyze, getAnalysis, clearError } = useHeuristicsAnalysis();
  const [showAnalysisForm, setShowAnalysisForm] = useState(false);
  const [analysisForm, setAnalysisForm] = useState<HeuristicsAnalysisRequest>({
    user_id: 1,
    task_id: undefined,
    analysis_type: 'performance',
    data: {},
  });
  const [loadingAnalyses, setLoadingAnalyses] = useState(false);
  const [analysisIdInput, setAnalysisIdInput] = useState('');
  const [showFetchForm, setShowFetchForm] = useState(false);

  useEffect(() => {
    loadAnalyses();
  }, []);

  const loadAnalyses = async () => {
    setLoadingAnalyses(true);
    try {
      // TODO: 過去の分析結果を取得するAPIエンドポイントを実装後、ここで呼び出す
      // 現時点では個別の分析結果取得のみ実装されている
      console.log('Loading analyses list...');
    } catch (err) {
      console.error('Failed to load analyses:', err);
    } finally {
      setLoadingAnalyses(false);
    }
  };

  const handleAnalyze = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await analyze(analysisForm);
      setShowAnalysisForm(false);
      setAnalysisForm({
        user_id: 1,
        task_id: undefined,
        analysis_type: 'performance',
        data: {},
      });
    } catch (err) {
      console.error('Analysis failed:', err);
    }
  };

  const handleFormChange = (field: string, value: any) => {
    setAnalysisForm(prev => ({
      ...prev,
      [field]: value,
    }));
  };

  const handleGetAnalysis = (id: number) => {
    getAnalysis(id.toString());
  };

  const handleFetchAnalysis = async (e: React.FormEvent) => {
    e.preventDefault();
    if (analysisIdInput) {
      await getAnalysis(analysisIdInput);
      setAnalysisIdInput('');
      setShowFetchForm(false);
    }
  };

  const analysisTypes = [
    { value: 'performance', label: 'パフォーマンス分析' },
    { value: 'behavior', label: '行動分析' },
    { value: 'pattern', label: 'パターン分析' },
    { value: 'cognitive', label: '認知分析' },
    { value: 'efficiency', label: '効率性分析' },
  ];

  if (loading || loadingAnalyses) {
    return (
      <div className={styles.container}>
        <div className={styles.loading}>
          {loadingAnalyses ? '分析結果を読み込み中...' : '分析を実行中...'}
        </div>
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
        <h2>ヒューリスティック分析</h2>
        <div className={styles.headerButtons}>
          <button 
            onClick={() => setShowAnalysisForm(!showAnalysisForm)} 
            className={styles.analyzeButton}
          >
            新しい分析
          </button>
          <button 
            onClick={() => setShowFetchForm(!showFetchForm)} 
            className={styles.fetchButton}
          >
            分析結果取得
          </button>
        </div>
      </div>

      {showFetchForm && (
        <div className={styles.fetchForm}>
          <h3>分析結果を取得</h3>
          <form onSubmit={handleFetchAnalysis}>
            <div className={styles.formGroup}>
              <label>分析ID:</label>
              <input
                type="text"
                value={analysisIdInput}
                onChange={(e) => setAnalysisIdInput(e.target.value)}
                placeholder="分析IDを入力"
                required
              />
            </div>
            <div className={styles.formActions}>
              <button type="submit" className={styles.submitButton}>
                取得
              </button>
              <button 
                type="button" 
                onClick={() => {
                  setShowFetchForm(false);
                  setAnalysisIdInput('');
                }}
                className={styles.cancelButton}
              >
                キャンセル
              </button>
            </div>
          </form>
        </div>
      )}

      {showAnalysisForm && (
        <div className={styles.analysisForm}>
          <h3>新しい分析を実行</h3>
          <form onSubmit={handleAnalyze}>
            <div className={styles.formRow}>
              <div className={styles.formGroup}>
                <label>ユーザーID:</label>
                <input
                  type="number"
                  value={analysisForm.user_id}
                  onChange={(e) => handleFormChange('user_id', parseInt(e.target.value))}
                  required
                />
              </div>
              <div className={styles.formGroup}>
                <label>タスクID (オプション):</label>
                <input
                  type="number"
                  value={analysisForm.task_id || ''}
                  onChange={(e) => handleFormChange('task_id', e.target.value ? parseInt(e.target.value) : undefined)}
                  placeholder="タスクID"
                />
              </div>
            </div>
            
            <div className={styles.formGroup}>
              <label>分析タイプ:</label>
              <select
                value={analysisForm.analysis_type}
                onChange={(e) => handleFormChange('analysis_type', e.target.value)}
                required
              >
                {analysisTypes.map(type => (
                  <option key={type.value} value={type.value}>
                    {type.label}
                  </option>
                ))}
              </select>
            </div>
            
            <div className={styles.formGroup}>
              <label>分析データ (JSON):</label>
              <textarea
                value={JSON.stringify(analysisForm.data, null, 2)}
                onChange={(e) => {
                  try {
                    handleFormChange('data', JSON.parse(e.target.value));
                  } catch {
                    // Invalid JSON, keep as string for now
                  }
                }}
                placeholder='{"metric": "completion_time", "value": 120}'
                rows={4}
              />
            </div>
            
            <div className={styles.formActions}>
              <button type="submit" className={styles.submitButton}>
                分析実行
              </button>
              <button 
                type="button" 
                onClick={() => setShowAnalysisForm(false)}
                className={styles.cancelButton}
              >
                キャンセル
              </button>
            </div>
          </form>
        </div>
      )}

      {currentAnalysis && (
        <div className={styles.currentAnalysis}>
          <h3>現在の分析結果</h3>
          <div className={styles.analysisCard}>
            <div className={styles.analysisHeader}>
              <span className={styles.analysisType}>{currentAnalysis.analysis_type}</span>
              <span className={styles.score}>スコア: {currentAnalysis.score}</span>
              <span className={styles.status}>{currentAnalysis.status}</span>
            </div>
            <div className={styles.analysisDetails}>
              <div className={styles.metadata}>
                <div>ユーザーID: {currentAnalysis.user_id}</div>
                {currentAnalysis.task_id && <div>タスクID: {currentAnalysis.task_id}</div>}
                <div>作成日: {new Date(currentAnalysis.created_at).toLocaleString('ja-JP')}</div>
              </div>
              <div className={styles.result}>
                <strong>結果:</strong>
                <pre>{currentAnalysis.result}</pre>
              </div>
            </div>
          </div>
        </div>
      )}

      <div className={styles.analysesList}>
        <h3>過去の分析結果</h3>
        {analyses.length === 0 ? (
          <div className={styles.noData}>分析結果がありません</div>
        ) : (
          <div className={styles.analysisGrid}>
            {analyses.map((analysis) => (
              <div key={analysis.id} className={styles.analysisCard}>
                <div className={styles.analysisHeader}>
                  <span className={styles.analysisType}>{analysis.analysis_type}</span>
                  <span className={styles.score}>スコア: {analysis.score}</span>
                </div>
                <div className={styles.analysisInfo}>
                  <div className={styles.metadata}>
                    <div>ID: {analysis.id}</div>
                    <div>ユーザー: {analysis.user_id}</div>
                    <div>ステータス: {analysis.status}</div>
                  </div>
                  <div className={styles.date}>
                    {new Date(analysis.created_at).toLocaleDateString('ja-JP')}
                  </div>
                </div>
                <button 
                  onClick={() => handleGetAnalysis(analysis.id)}
                  className={styles.viewButton}
                >
                  詳細表示
                </button>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}