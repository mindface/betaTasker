/**
 * ヒューリスティクス発見サービス
 * ユーザー行動パターンから認知的ヒューリスティクスを抽出
 */

interface UserAction {
  timestamp: number;
  actionType: string;
  elementId: string;
  context: Record<string, any>;
  duration?: number;
  sequence?: number;
}

interface Pattern {
  id: string;
  name: string;
  frequency: number;
  confidence: number;
  actions: UserAction[];
  heuristic?: string;
}

interface HeuristicModel {
  patterns: Pattern[];
  accuracy: number;
  lastUpdated: Date;
}

export class HeuristicsDiscoveryClient {
  private actionBuffer: UserAction[] = [];
  private patterns: Map<string, Pattern> = new Map();
  private modelVersion: number = 0;
  private cycleInterval: number = 5000; // 5秒ごとに分析
  private worker: Worker | null = null;

  constructor() {
    this.initializeWorker();
    this.startDiscoveryCycle();
  }

  /**
   * Web Workerで並列処理
   */
  private initializeWorker() {
    if (typeof Worker !== 'undefined') {
      // パターン分析用ワーカー作成
      const workerCode = `
        self.addEventListener('message', (e) => {
          const { actions, threshold } = e.data;
          const patterns = extractPatterns(actions, threshold);
          self.postMessage({ patterns });
        });

        function extractPatterns(actions, threshold) {
          const sequences = {};
          
          // N-gramベースのパターン抽出
          for (let n = 2; n <= 5; n++) {
            for (let i = 0; i <= actions.length - n; i++) {
              const sequence = actions.slice(i, i + n)
                .map(a => a.actionType)
                .join('-');
              
              sequences[sequence] = (sequences[sequence] || 0) + 1;
            }
          }

          // 頻度閾値を超えるパターンのみ返す
          return Object.entries(sequences)
            .filter(([_, count]) => count >= threshold)
            .map(([seq, count]) => ({
              sequence: seq,
              frequency: count,
              confidence: count / actions.length
            }));
        }
      `;

      const blob = new Blob([workerCode], { type: 'application/javascript' });
      this.worker = new Worker(URL.createObjectURL(blob));
    }
  }

  /**
   * ユーザーアクションを記録
   */
  public recordAction(action: UserAction): void {
    // タイムスタンプ追加
    action.timestamp = action.timestamp || Date.now();
    
    // バッファに追加（最大1000件保持）
    this.actionBuffer.push(action);
    if (this.actionBuffer.length > 1000) {
      this.actionBuffer.shift();
    }

    // リアルタイム分析トリガー
    this.analyzeRecentActions();
  }

  /**
   * 最近のアクションを分析
   */
  private analyzeRecentActions(): void {
    const recentActions = this.actionBuffer.slice(-20);
    
    // 短期記憶パターン検出
    const shortTermPatterns = this.detectShortTermPatterns(recentActions);
    
    // 認知負荷推定
    const cognitiveLoad = this.estimateCognitiveLoad(recentActions);
    
    // ヒューリスティクス適用
    this.applyHeuristics(shortTermPatterns, cognitiveLoad);
  }

  /**
   * 短期パターン検出
   */
  private detectShortTermPatterns(actions: UserAction[]): Pattern[] {
    const patterns: Pattern[] = [];
    
    // 繰り返しパターン検出
    const repetitions = this.findRepetitions(actions);
    
    // 時間的クラスタリング
    const temporalClusters = this.temporalClustering(actions);
    
    // パターンマージ
    patterns.push(...repetitions, ...temporalClusters);
    
    return patterns;
  }

  /**
   * 繰り返しパターン検出
   */
  private findRepetitions(actions: UserAction[]): Pattern[] {
    const patterns: Pattern[] = [];
    const actionTypes = actions.map(a => a.actionType);
    
    // 同一アクションの繰り返し
    let currentType = '';
    let count = 0;
    
    for (const type of actionTypes) {
      if (type === currentType) {
        count++;
      } else {
        if (count >= 3) {
          patterns.push({
            id: `repetition-${currentType}`,
            name: `Repetitive ${currentType}`,
            frequency: count,
            confidence: count / actions.length,
            actions: actions.filter(a => a.actionType === currentType),
            heuristic: 'repetition_bias'
          });
        }
        currentType = type;
        count = 1;
      }
    }
    
    return patterns;
  }

  /**
   * 時間的クラスタリング
   */
  private temporalClustering(actions: UserAction[]): Pattern[] {
    const patterns: Pattern[] = [];
    const timeThreshold = 2000; // 2秒以内のアクションをクラスタ化
    
    let cluster: UserAction[] = [];
    let lastTime = 0;
    
    for (const action of actions) {
      if (action.timestamp - lastTime <= timeThreshold) {
        cluster.push(action);
      } else {
        if (cluster.length >= 2) {
          patterns.push({
            id: `cluster-${Date.now()}`,
            name: 'Temporal Cluster',
            frequency: cluster.length,
            confidence: 0.7,
            actions: cluster,
            heuristic: 'temporal_proximity'
          });
        }
        cluster = [action];
      }
      lastTime = action.timestamp;
    }
    
    return patterns;
  }

  /**
   * 認知負荷推定
   */
  private estimateCognitiveLoad(actions: UserAction[]): number {
    // アクションの多様性
    const uniqueTypes = new Set(actions.map(a => a.actionType)).size;
    const diversity = uniqueTypes / actions.length;
    
    // アクション頻度
    const frequency = actions.length / (
      actions[actions.length - 1]?.timestamp - actions[0]?.timestamp || 1
    ) * 1000;
    
    // 認知負荷スコア（0-1）
    const load = Math.min(1, (diversity * 0.5 + frequency * 0.5));
    
    return load;
  }

  /**
   * ヒューリスティクス適用
   */
  private applyHeuristics(patterns: Pattern[], cognitiveLoad: number): void {
    // 認知負荷が高い場合
    if (cognitiveLoad > 0.7) {
      this.applySatisficingHeuristic();
    }
    
    // 繰り返しパターンがある場合
    const repetitivePattern = patterns.find(p => p.heuristic === 'repetition_bias');
    if (repetitivePattern) {
      this.applyAnchroingHeuristic(repetitivePattern);
    }
    
    // 時間的近接パターン
    const temporalPattern = patterns.find(p => p.heuristic === 'temporal_proximity');
    if (temporalPattern) {
      this.applyAvailabilityHeuristic(temporalPattern);
    }
  }

  /**
   * 満足化ヒューリスティクス
   * 完璧な選択より十分な選択を優先
   */
  private applySatisficingHeuristic(): void {
    console.log('Applying satisficing heuristic - reducing options');
    // UIに選択肢削減を通知
    this.notifyUI('reduce_options', { maxOptions: 3 });
  }

  /**
   * アンカリングヒューリスティクス
   * 最初の情報に強く影響される
   */
  private applyAnchroingHeuristic(pattern: Pattern): void {
    console.log('Applying anchoring heuristic', pattern);
    // 最初のアクションを重視
    this.notifyUI('emphasize_first', { pattern });
  }

  /**
   * 利用可能性ヒューリスティクス
   * 思い出しやすい情報を重視
   */
  private applyAvailabilityHeuristic(pattern: Pattern): void {
    console.log('Applying availability heuristic', pattern);
    // 最近使用した機能を前面に
    this.notifyUI('prioritize_recent', { pattern });
  }

  /**
   * 発見サイクル開始
   */
  private startDiscoveryCycle(): void {
    setInterval(() => {
      this.runDiscoveryCycle();
    }, this.cycleInterval);
  }

  /**
   * 発見サイクル実行
   */
  private async runDiscoveryCycle(): Promise<void> {
    if (this.actionBuffer.length < 10) return;

    // モデルテスト入力
    const testInput = this.prepareModelTestInput();

    // パターン抽出（Web Worker使用）
    const patterns = await this.extractPatternsAsync(testInput);

    // ヒューリスティクス発見
    const heuristics = this.discoverHeuristics(patterns);

    // モデル更新
    this.updateModel(heuristics);

    // バックエンドに送信
    this.syncWithBackend(heuristics);
  }

  /**
   * モデルテスト入力準備
   */
  private prepareModelTestInput(): any {
    return {
      actions: this.actionBuffer.slice(-100),
      timestamp: Date.now(),
      context: {
        timeOfDay: new Date().getHours(),
        dayOfWeek: new Date().getDay(),
        sessionDuration: this.getSessionDuration()
      }
    };
  }

  /**
   * 非同期パターン抽出
   */
  private extractPatternsAsync(input: any): Promise<Pattern[]> {
    return new Promise((resolve) => {
      if (this.worker) {
        this.worker.onmessage = (e) => {
          resolve(e.data.patterns);
        };
        this.worker.postMessage({
          actions: input.actions,
          threshold: 3
        });
      } else {
        // フォールバック: メインスレッドで実行
        resolve(this.extractPatterns(input.actions));
      }
    });
  }

  /**
   * パターン抽出（同期版）
   */
  private extractPatterns(actions: UserAction[]): Pattern[] {
    const patterns: Pattern[] = [];
    const sequenceMap = new Map<string, number>();
    
    // シーケンシャルパターン
    for (let i = 0; i < actions.length - 1; i++) {
      const seq = `${actions[i].actionType}->${actions[i + 1].actionType}`;
      sequenceMap.set(seq, (sequenceMap.get(seq) || 0) + 1);
    }
    
    // 頻出パターンを抽出
    sequenceMap.forEach((count, sequence) => {
      if (count >= 3) {
        patterns.push({
          id: `seq-${Date.now()}-${Math.random()}`,
          name: sequence,
          frequency: count,
          confidence: count / actions.length,
          actions: [],
          heuristic: this.inferHeuristic(sequence, count)
        });
      }
    });
    
    return patterns;
  }

  /**
   * ヒューリスティクス推論
   */
  private inferHeuristic(sequence: string, frequency: number): string {
    // パターンからヒューリスティクスを推論
    if (sequence.includes('->')) {
      const [from, to] = sequence.split('->');
      
      // 同じアクションの繰り返し
      if (from === to) return 'confirmation_bias';
      
      // 戻る動作
      if (to === 'back' || to === 'cancel') return 'loss_aversion';
      
      // 順次進行
      if (from === 'next' || to === 'next') return 'progressive_disclosure';
    }
    
    // 高頻度
    if (frequency > 10) return 'habit_formation';
    
    return 'general_pattern';
  }

  /**
   * ヒューリスティクス発見
   */
  private discoverHeuristics(patterns: Pattern[]): HeuristicModel {
    const heuristicGroups = new Map<string, Pattern[]>();
    
    // ヒューリスティクスごとにグループ化
    patterns.forEach(p => {
      const group = heuristicGroups.get(p.heuristic || '') || [];
      group.push(p);
      heuristicGroups.set(p.heuristic || '', group);
    });
    
    // 信頼度計算
    let totalConfidence = 0;
    patterns.forEach(p => {
      totalConfidence += p.confidence;
    });
    
    return {
      patterns,
      accuracy: totalConfidence / patterns.length || 0,
      lastUpdated: new Date()
    };
  }

  /**
   * モデル更新
   */
  private updateModel(heuristics: HeuristicModel): void {
    this.modelVersion++;
    
    // 既存パターンとマージ
    heuristics.patterns.forEach(p => {
      const existing = this.patterns.get(p.id);
      if (existing) {
        // 信頼度を更新
        existing.confidence = (existing.confidence + p.confidence) / 2;
        existing.frequency += p.frequency;
      } else {
        this.patterns.set(p.id, p);
      }
    });
    
    // 古いパターンを削除
    this.pruneOldPatterns();
  }

  /**
   * 古いパターンの削除
   */
  private pruneOldPatterns(): void {
    const threshold = 0.1;
    
    this.patterns.forEach((pattern, id) => {
      if (pattern.confidence < threshold) {
        this.patterns.delete(id);
      }
    });
  }

  /**
   * バックエンド同期
   */
  private async syncWithBackend(heuristics: HeuristicModel): Promise<void> {
    try {
      const response = await fetch('/api/heuristics/sync', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          model: heuristics,
          version: this.modelVersion,
          timestamp: Date.now()
        })
      });
      
      if (response.ok) {
        console.log('Heuristics synced with backend');
      }
    } catch (error) {
      console.error('Failed to sync heuristics:', error);
    }
  }

  /**
   * UIへの通知
   */
  private notifyUI(action: string, data: any): void {
    window.dispatchEvent(new CustomEvent('heuristic-applied', {
      detail: { action, data }
    }));
  }

  /**
   * セッション継続時間取得
   */
  private getSessionDuration(): number {
    if (this.actionBuffer.length === 0) return 0;
    return Date.now() - this.actionBuffer[0].timestamp;
  }

  /**
   * 公開API: 現在のヒューリスティクスを取得
   */
  public getCurrentHeuristics(): Pattern[] {
    return Array.from(this.patterns.values());
  }

  /**
   * 公開API: サイクル間隔を設定
   */
  public setCycleInterval(ms: number): void {
    this.cycleInterval = ms;
  }

  /**
   * クリーンアップ
   */
  public destroy(): void {
    if (this.worker) {
      this.worker.terminate();
    }
    this.actionBuffer = [];
    this.patterns.clear();
  }
}

// シングルトンインスタンス
export const heuristicsDiscovery = new HeuristicsDiscoveryClient();