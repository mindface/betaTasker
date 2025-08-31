/**
 * ヒューリスティクス機能用キャッシュサービス
 * インメモリキャッシュとローカルストレージキャッシュを提供
 */

interface CacheItem<T> {
  data: T;
  timestamp: number;
  ttl: number;
}

interface CacheOptions {
  ttl?: number; // ミリ秒
  useLocalStorage?: boolean;
  keyPrefix?: string;
}

export class HeuristicsCache {
  private memoryCache = new Map<string, CacheItem<any>>();
  private readonly defaultTTL = 5 * 60 * 1000; // 5分
  private readonly keyPrefix: string;
  private readonly useLocalStorage: boolean;

  constructor(options: CacheOptions = {}) {
    this.keyPrefix = options.keyPrefix || 'heuristics_';
    this.useLocalStorage = options.useLocalStorage ?? true;
  }

  /**
   * キャッシュにデータを保存
   */
  set<T>(key: string, data: T, ttl: number = this.defaultTTL): void {
    const cacheKey = this.getFullKey(key);
    const item: CacheItem<T> = {
      data,
      timestamp: Date.now(),
      ttl
    };

    // メモリキャッシュに保存
    this.memoryCache.set(cacheKey, item);

    // ローカルストレージにも保存（オプション）
    if (this.useLocalStorage && typeof window !== 'undefined') {
      try {
        localStorage.setItem(cacheKey, JSON.stringify(item));
      } catch (error) {
        console.warn('Failed to save to localStorage:', error);
      }
    }
  }

  /**
   * キャッシュからデータを取得
   */
  get<T>(key: string): T | null {
    const cacheKey = this.getFullKey(key);
    
    // メモリキャッシュから取得を試行
    let item = this.memoryCache.get(cacheKey);
    
    // メモリキャッシュにない場合はローカルストレージから取得
    if (!item && this.useLocalStorage && typeof window !== 'undefined') {
      try {
        const stored = localStorage.getItem(cacheKey);
        if (stored) {
          item = JSON.parse(stored);
          // メモリキャッシュにも復元
          this.memoryCache.set(cacheKey, item);
        }
      } catch (error) {
        console.warn('Failed to load from localStorage:', error);
      }
    }

    if (!item) return null;

    // TTLチェック
    if (Date.now() - item.timestamp > item.ttl) {
      this.delete(key);
      return null;
    }

    return item.data;
  }

  /**
   * キャッシュからデータを削除
   */
  delete(key: string): void {
    const cacheKey = this.getFullKey(key);
    
    // メモリキャッシュから削除
    this.memoryCache.delete(cacheKey);

    // ローカルストレージからも削除
    if (this.useLocalStorage && typeof window !== 'undefined') {
      try {
        localStorage.removeItem(cacheKey);
      } catch (error) {
        console.warn('Failed to remove from localStorage:', error);
      }
    }
  }

  /**
   * パターンに一致するキーを削除
   */
  invalidate(pattern: string): void {
    const keysToDelete: string[] = [];

    // メモリキャッシュからパターンに一致するキーを検索
    for (const key of this.memoryCache.keys()) {
      if (key.includes(pattern)) {
        keysToDelete.push(key);
      }
    }

    // 一致するキーを削除
    keysToDelete.forEach(key => {
      this.memoryCache.delete(key);
    });

    // ローカルストレージからも削除
    if (this.useLocalStorage && typeof window !== 'undefined') {
      try {
        for (const key of keysToDelete) {
          localStorage.removeItem(key);
        }
      } catch (error) {
        console.warn('Failed to remove from localStorage:', error);
      }
    }
  }

  /**
   * キャッシュをクリア
   */
  clear(): void {
    // メモリキャッシュをクリア
    this.memoryCache.clear();

    // ローカルストレージからも削除
    if (this.useLocalStorage && typeof window !== 'undefined') {
      try {
        const keysToDelete: string[] = [];
        for (let i = 0; i < localStorage.length; i++) {
          const key = localStorage.key(i);
          if (key && key.startsWith(this.keyPrefix)) {
            keysToDelete.push(key);
          }
        }
        keysToDelete.forEach(key => localStorage.removeItem(key));
      } catch (error) {
        console.warn('Failed to clear localStorage:', error);
      }
    }
  }

  /**
   * キャッシュの統計情報を取得
   */
  getStats(): {
    memorySize: number;
    localStorageSize: number;
    totalKeys: number;
  } {
    let localStorageSize = 0;
    
    if (this.useLocalStorage && typeof window !== 'undefined') {
      try {
        for (let i = 0; i < localStorage.length; i++) {
          const key = localStorage.key(i);
          if (key && key.startsWith(this.keyPrefix)) {
            const value = localStorage.getItem(key);
            if (value) {
              localStorageSize += key.length + value.length;
            }
          }
        }
      } catch (error) {
        console.warn('Failed to calculate localStorage size:', error);
      }
    }

    return {
      memorySize: this.memoryCache.size,
      localStorageSize,
      totalKeys: this.memoryCache.size
    };
  }

  /**
   * 期限切れのキャッシュアイテムをクリーンアップ
   */
  cleanup(): void {
    const now = Date.now();
    const keysToDelete: string[] = [];

    // 期限切れのアイテムを検索
    for (const [key, item] of this.memoryCache.entries()) {
      if (now - item.timestamp > item.ttl) {
        keysToDelete.push(key);
      }
    }

    // 期限切れのアイテムを削除
    keysToDelete.forEach(key => {
      this.memoryCache.delete(key);
    });

    // ローカルストレージからも削除
    if (this.useLocalStorage && typeof window !== 'undefined') {
      try {
        keysToDelete.forEach(key => {
          localStorage.removeItem(key);
        });
      } catch (error) {
        console.warn('Failed to cleanup localStorage:', error);
      }
    }
  }

  /**
   * フルキーを生成
   */
  private getFullKey(key: string): string {
    return `${this.keyPrefix}${key}`;
  }
}

// デフォルトのキャッシュインスタンス
export const heuristicsCache = new HeuristicsCache({
  ttl: 5 * 60 * 1000, // 5分
  useLocalStorage: true,
  keyPrefix: 'heuristics_'
});

// 定期的なクリーンアップ（5分ごと）
if (typeof window !== 'undefined') {
  setInterval(() => {
    heuristicsCache.cleanup();
  }, 5 * 60 * 1000);
}

// キャッシュキーの定数
export const CACHE_KEYS = {
  ANALYSES: 'analyses',
  ANALYSIS: 'analysis',
  PATTERNS: 'patterns',
  INSIGHTS: 'insights',
  INSIGHT: 'insight',
  TRACKING: 'tracking',
  MODELS: 'models',
  MODEL: 'model'
} as const;
