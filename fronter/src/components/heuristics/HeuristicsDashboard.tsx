"use client"
import React, { useState, useCallback, useMemo } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { RootState } from '../../store';
import { clearAllErrors } from '../../features/heuristics/heuristicsSlice';
import HeuristicsInsights from './HeuristicsInsights';
import HeuristicsTracking from './HeuristicsTracking';
import HeuristicsAnalysis from './HeuristicsAnalysis';
import HeuristicsPatterns from './HeuristicsPatterns';
import styles from './HeuristicsDashboard.module.scss';

type TabType = 'insights' | 'tracking' | 'analysis' | 'patterns';

export default function HeuristicsDashboard() {
  const dispatch = useDispatch();
  const [activeTab, setActiveTab] = useState<TabType>('insights');

  // Redux状態の取得
  const analysisError = useSelector((state: RootState) => state.heuristics.analysis.error);
  const trackingError = useSelector((state: RootState) => state.heuristics.tracking.error);
  const patternsError = useSelector((state: RootState) => state.heuristics.patterns.error);
  const insightsError = useSelector((state: RootState) => state.heuristics.insights.error);

  // エラーがあるかどうかを判定
  const hasErrors = useMemo(() => {
    return !!(analysisError || trackingError || patternsError || insightsError);
  }, [analysisError, trackingError, patternsError, insightsError]);

  // タブの定義
  const tabs = useMemo(() => [
    { 
      id: 'insights' as TabType, 
      label: 'インサイト', 
      icon: '💡',
      error: insightsError
    },
    { 
      id: 'tracking' as TabType, 
      label: 'トラッキング', 
      icon: '📊',
      error: trackingError
    },
    { 
      id: 'analysis' as TabType, 
      label: '分析', 
      icon: '🔍',
      error: analysisError
    },
    { 
      id: 'patterns' as TabType, 
      label: 'パターン', 
      icon: '🎯',
      error: patternsError
    },
  ], [insightsError, trackingError, analysisError, patternsError]);

  // タブ切り替えのハンドラー
  const handleTabChange = useCallback((tabId: TabType) => {
    setActiveTab(tabId);
    // エラーをクリア
    dispatch(clearAllErrors());
  }, [dispatch]);

  // タブコンテンツのレンダリング
  const renderTabContent = useCallback(() => {
    switch (activeTab) {
      case 'insights':
        return <HeuristicsInsights />;
      case 'tracking':
        return <HeuristicsTracking />;
      case 'analysis':
        return <HeuristicsAnalysis />;
      case 'patterns':
        return <HeuristicsPatterns />;
      default:
        return <HeuristicsInsights />;
    }
  }, [activeTab]);

  // エラーメッセージの表示
  const renderErrorMessage = useCallback(() => {
    if (!hasErrors) return null;

    const errorMessages = [];
    if (analysisError) errorMessages.push(`分析: ${analysisError}`);
    if (trackingError) errorMessages.push(`トラッキング: ${trackingError}`);
    if (patternsError) errorMessages.push(`パターン: ${patternsError}`);
    if (insightsError) errorMessages.push(`インサイト: ${insightsError}`);

    return (
      <div className={styles.errorBanner}>
        <div className={styles.errorContent}>
          <span className={styles.errorIcon}>⚠️</span>
          <div className={styles.errorText}>
            <strong>エラーが発生しました:</strong>
            <ul>
              {errorMessages.map((msg, index) => (
                <li key={index}>{msg}</li>
              ))}
            </ul>
          </div>
          <button 
            className={styles.errorClose}
            onClick={() => dispatch(clearAllErrors())}
            aria-label="エラーを閉じる"
          >
            ✕
          </button>
        </div>
      </div>
    );
  }, [hasErrors, analysisError, trackingError, patternsError, insightsError, dispatch]);

  return (
    <div className={styles.dashboard}>
      <div className={styles.header}>
        <h1>ヒューリスティック分析ダッシュボード</h1>
        <p>ユーザー行動の分析とインサイトの発見</p>
      </div>

      {/* エラーメッセージ */}
      {renderErrorMessage()}

      <div className={styles.tabNavigation}>
        {tabs.map((tab) => (
          <button
            key={tab.id}
            className={`${styles.tab} ${activeTab === tab.id ? styles.active : ''} ${tab.error ? styles.hasError : ''}`}
            onClick={() => handleTabChange(tab.id)}
            aria-pressed={activeTab === tab.id}
            aria-describedby={tab.error ? `error-${tab.id}` : undefined}
          >
            <span className={styles.icon}>{tab.icon}</span>
            <span className={styles.label}>{tab.label}</span>
            {tab.error && (
              <span className={styles.errorIndicator} aria-hidden="true">●</span>
            )}
          </button>
        ))}
      </div>

      <div className={styles.tabContent}>
        {renderTabContent()}
      </div>

      {/* アクセシビリティ: エラーの詳細説明 */}
      {tabs.map(tab => 
        tab.error && (
          <div key={`error-${tab.id}`} id={`error-${tab.id}`} className={styles.srOnly}>
            {tab.label}タブでエラーが発生しています: {tab.error}
          </div>
        )
      )}
    </div>
  );
}