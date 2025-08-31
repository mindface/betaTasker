"use client"
import React, { useState } from 'react';
import HeuristicsInsights from './HeuristicsInsights';
import HeuristicsTracking from './HeuristicsTracking';
import HeuristicsAnalysis from './HeuristicsAnalysis';
import HeuristicsPatterns from './HeuristicsPatterns';
import styles from './HeuristicsDashboard.module.scss';

type TabType = 'insights' | 'tracking' | 'analysis' | 'patterns';

export default function HeuristicsDashboard() {
  const [activeTab, setActiveTab] = useState<TabType>('insights');

  const tabs = [
    { id: 'insights', label: 'インサイト', icon: '💡' },
    { id: 'tracking', label: 'トラッキング', icon: '📊' },
    { id: 'analysis', label: '分析', icon: '🔍' },
    { id: 'patterns', label: 'パターン', icon: '🎯' },
  ];

  const renderTabContent = () => {
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
  };

  return (
    <div className={styles.dashboard}>
      <div className={styles.header}>
        <h1>ヒューリスティック分析ダッシュボード</h1>
        <p>ユーザー行動の分析とインサイトの発見</p>
      </div>

      <div className={styles.tabNavigation}>
        {tabs.map((tab) => (
          <button
            key={tab.id}
            className={`${styles.tab} ${activeTab === tab.id ? styles.active : ''}`}
            onClick={() => setActiveTab(tab.id as TabType)}
          >
            <span className={styles.icon}>{tab.icon}</span>
            <span className={styles.label}>{tab.label}</span>
          </button>
        ))}
      </div>

      <div className={styles.tabContent}>
        {renderTabContent()}
      </div>
    </div>
  );
}