"use client"
import React, { useEffect, useState } from 'react';
import { useHeuristicsInsights } from '../../hooks/useHeuristics';
import styles from './HeuristicsInsights.module.scss';

export default function HeuristicsInsights() {
  const { insights, total, loading, error, getInsights, clearError } = useHeuristicsInsights();
  const [currentPage, setCurrentPage] = useState(1);
  const [userId, setUserId] = useState<string>('');
  const itemsPerPage = 10;

  useEffect(() => {
    loadInsights();
  }, [currentPage, userId]);

  const loadInsights = () => {
    const params = {
      limit: itemsPerPage,
      offset: (currentPage - 1) * itemsPerPage,
      ...(userId && { user_id: userId }),
    };
    getInsights(params);
  };

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  const handleUserFilter = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUserId(e.target.value);
    setCurrentPage(1);
  };

  const totalPages = Math.ceil(total / itemsPerPage);

  if (loading) {
    return (
      <div className={styles.container}>
        <div className={styles.loading}>インサイトを読み込み中...</div>
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
        <h2>ヒューリスティックインサイト</h2>
        <div className={styles.filters}>
          <input
            type="text"
            placeholder="ユーザーIDでフィルター"
            value={userId}
            onChange={handleUserFilter}
            className={styles.filterInput}
          />
          <button onClick={loadInsights} className={styles.refreshButton}>
            更新
          </button>
        </div>
      </div>

      <div className={styles.insightsList}>
        {insights.length === 0 ? (
          <div className={styles.noData}>インサイトがありません</div>
        ) : (
          insights.map((insight) => (
            <div key={insight.id} className={styles.insightCard}>
              <div className={styles.insightHeader}>
                <h3>{insight.title}</h3>
                <span className={styles.insightType}>{insight.type}</span>
              </div>
              <p className={styles.description}>{insight.description}</p>
              <div className={styles.metadata}>
                <div className={styles.confidence}>
                  信頼度: <span>{(insight.confidence * 100).toFixed(1)}%</span>
                </div>
                <div className={styles.status}>
                  ステータス: {insight.is_active ? '有効' : '無効'}
                </div>
                <div className={styles.date}>
                  作成日: {new Date(insight.created_at).toLocaleDateString('ja-JP')}
                </div>
              </div>
            </div>
          ))
        )}
      </div>

      {totalPages > 1 && (
        <div className={styles.pagination}>
          <button
            onClick={() => handlePageChange(currentPage - 1)}
            disabled={currentPage === 1}
            className={styles.pageButton}
          >
            前へ
          </button>
          <span className={styles.pageInfo}>
            {currentPage} / {totalPages}
          </span>
          <button
            onClick={() => handlePageChange(currentPage + 1)}
            disabled={currentPage === totalPages}
            className={styles.pageButton}
          >
            次へ
          </button>
        </div>
      )}
    </div>
  );
}