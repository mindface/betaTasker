"use client"
import React, { useCallback, useEffect, useState } from 'react';
import { useHeuristicsTracking } from '../../hooks/useHeuristics';
import { HeuristicsTrackingData } from '../../model/heuristics';
import styles from './HeuristicsTracking.module.scss';

export default function HeuristicsTracking() {
  const { trackingData, loading, error, track, getTracking, clearError } = useHeuristicsTracking();
  const [userId, setUserId] = useState<string>('1');
  const [showTrackForm, setShowTrackForm] = useState(false);
  const [trackForm, setTrackForm] = useState<HeuristicsTrackingData>({
    user_id: 1,
    action: '',
    context: {},
    session_id: '',
    duration: 0,
  });

  const handleLoadTracking = useCallback(() => {
    if (userId) {
      getTracking(userId);
    }
  },[getTracking, userId]);

  const handleTrack = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await track(trackForm);
      setShowTrackForm(false);
      setTrackForm({
        user_id: parseInt(userId) || 1,
        action: '',
        context: {},
        session_id: '',
        duration: 0,
      });
      handleLoadTracking();
    } catch (err) {
      console.error('Tracking failed:', err);
    }
  };

  const handleFormChange = (field: string, value: any) => {
    setTrackForm(prev => ({
      ...prev,
      [field]: value,
    }));
  };

  useEffect(() => {
    if (userId) {
      handleLoadTracking();
    }
  }, [handleLoadTracking, userId]);

  if (loading) {
    return (
      <div className={styles.container}>
        <div className={styles.loading}>トラッキングデータを読み込み中...</div>
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
        <h2>ユーザー行動トラッキング</h2>
        <div className={styles.controls}>
          <input
            type="text"
            placeholder="ユーザーID"
            value={userId}
            onChange={(e) => setUserId(e.target.value)}
            className={styles.userInput}
          />
          <button onClick={handleLoadTracking} className={styles.loadButton}>
            読み込み
          </button>
          <button 
            onClick={() => setShowTrackForm(!showTrackForm)} 
            className={styles.trackButton}
          >
            新しいトラッキング
          </button>
        </div>
      </div>

      {showTrackForm && (
        <div className={styles.trackForm}>
          <h3>新しい行動を記録</h3>
          <form onSubmit={handleTrack}>
            <div className={styles.formGroup}>
              <label>ユーザーID:</label>
              <input
                type="number"
                value={trackForm.user_id}
                onChange={(e) => handleFormChange('user_id', parseInt(e.target.value))}
                required
              />
            </div>
            <div className={styles.formGroup}>
              <label>アクション:</label>
              <input
                type="text"
                value={trackForm.action}
                onChange={(e) => handleFormChange('action', e.target.value)}
                placeholder="例: task_completed, page_viewed"
                required
              />
            </div>
            <div className={styles.formGroup}>
              <label>セッションID:</label>
              <input
                type="text"
                value={trackForm.session_id}
                onChange={(e) => handleFormChange('session_id', e.target.value)}
                placeholder="セッションID（オプション）"
              />
            </div>
            <div className={styles.formGroup}>
              <label>継続時間 (ミリ秒):</label>
              <input
                type="number"
                value={trackForm.duration}
                onChange={(e) => handleFormChange('duration', parseInt(e.target.value))}
                placeholder="例: 5000"
              />
            </div>
            <div className={styles.formGroup}>
              <label>コンテキスト (JSON):</label>
              <textarea
                value={JSON.stringify(trackForm.context, null, 2)}
                onChange={(e) => {
                  try {
                    handleFormChange('context', JSON.parse(e.target.value));
                  } catch {
                    // Invalid JSON, keep as string for now
                  }
                }}
                placeholder='{"page": "dashboard", "feature": "task_list"}'
                rows={3}
              />
            </div>
            <div className={styles.formActions}>
              <button type="submit" className={styles.submitButton}>
                記録
              </button>
              <button 
                type="button" 
                onClick={() => setShowTrackForm(false)}
                className={styles.cancelButton}
              >
                キャンセル
              </button>
            </div>
          </form>
        </div>
      )}

      <div className={styles.trackingList}>
        {trackingData.length === 0 ? (
          <div className={styles.noData}>トラッキングデータがありません</div>
        ) : (
          <>
            <div className={styles.stats}>
              <div className={styles.statCard}>
                <h4>総記録数</h4>
                <span>{trackingData.length}</span>
              </div>
              <div className={styles.statCard}>
                <h4>ユニークアクション</h4>
                { trackingData.length > 0 ? <span>{new Set((trackingData ?? []).map(t => t.action)).size}</span> : <span>0</span>}
              </div>
              <div className={styles.statCard}>
                <h4>平均継続時間</h4>
                <span>
                  {trackingData.length > 0 
                    ? Math.round(trackingData.reduce((sum, t) => sum + t.duration, 0) / trackingData.length)
                    : 0}ms
                </span>
              </div>
            </div>

            <div className={styles.trackingItems}>
              {trackingData.length > 0 ? trackingData.map((item) => (
                <div key={item.id} className={styles.trackingCard}>
                  <div className={styles.trackingHeader}>
                    <h4>{item.action}</h4>
                    <span className={styles.timestamp}>
                      {new Date(item.timestamp).toLocaleString('ja-JP')}
                    </span>
                  </div>
                  <div className={styles.trackingDetails}>
                    <div className={styles.detail}>
                      <strong>継続時間:</strong> {item.duration}ms
                    </div>
                    {item.session_id && (
                      <div className={styles.detail}>
                        <strong>セッションID:</strong> {item.session_id}
                      </div>
                    )}
                    {item.context && (
                      <div className={styles.context}>
                        <strong>コンテキスト:</strong>
                        <pre>{JSON.stringify(JSON.parse(item.context), null, 2)}</pre>
                      </div>
                    )}
                  </div>
                </div>
              )) : (
                <div className={styles.noData}>トラッキングデータがありません</div>
              )}
            </div>
          </>
        )}
      </div>
    </div>
  );
}