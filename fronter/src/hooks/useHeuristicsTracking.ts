/**
 * ヒューリスティクストラッキングフック
 * コンポーネントでユーザー行動を自動記録
 * 利用用途がなれば削除する
 */

import { useEffect, useCallback, useRef } from 'react';
import { heuristicsDiscovery } from '../client/heuristicsDiscovery';

interface TrackingOptions {
  trackClicks?: boolean;
  trackScroll?: boolean;
  trackFocus?: boolean;
  trackHover?: boolean;
  trackInput?: boolean;
  debounceMs?: number;
}

export function useHeuristicsTracking(
  elementId: string,
  options: TrackingOptions = {}
) {
  const {
    trackClicks = true,
    trackScroll = true,
    trackFocus = true,
    trackHover = false,
    trackInput = true,
    debounceMs = 100
  } = options;
  console.warn("AAA HeuristicsTracking component is deprecated and may be removed in future versions.");

  const elementRef = useRef<HTMLElement | null>(null);
  const lastActionTime = useRef<number>(0);
  const hoverStartTime = useRef<number>(0);
  const scrollStartPos = useRef<number>(0);

  // デバウンス処理
  const debounce = useCallback((fn: Function) => {
    return (...args: any[]) => {
      const now = Date.now();
      if (now - lastActionTime.current > debounceMs) {
        lastActionTime.current = now;
        fn(...args);
      }
    };
  }, [debounceMs]);

  // クリックトラッキング
  const handleClick = useCallback((e: MouseEvent) => {
    const target = e.target as HTMLElement;

    heuristicsDiscovery.recordAction({
      timestamp: Date.now(),
      actionType: 'click',
      elementId,
      context: {
        targetTag: target.tagName,
        targetClass: target.className,
        targetId: target.id,
        x: e.clientX,
        y: e.clientY,
        ctrlKey: e.ctrlKey,
        shiftKey: e.shiftKey
      }
    });
  }, [elementId]);

  // スクロールトラッキング
  const handleScroll = useCallback(
    debounce((e: Event) => {
      const element = e.target as HTMLElement;
      const scrollDistance = Math.abs(element.scrollTop - scrollStartPos.current);
      
      heuristicsDiscovery.recordAction({
        timestamp: Date.now(),
        actionType: 'scroll',
        elementId,
        context: {
          scrollTop: element.scrollTop,
          scrollHeight: element.scrollHeight,
          scrollDistance,
          direction: element.scrollTop > scrollStartPos.current ? 'down' : 'up',
          velocity: scrollDistance / debounceMs
        }
      });
      
      scrollStartPos.current = element.scrollTop;
    }),
    [elementId, debounceMs]
  );

  // フォーカストラッキング
  const handleFocus = useCallback((e: FocusEvent) => {
    const target = e.target as HTMLElement;
    
    heuristicsDiscovery.recordAction({
      timestamp: Date.now(),
      actionType: 'focus',
      elementId,
      context: {
        targetTag: target.tagName,
        targetType: (target as HTMLInputElement).type || null,
        targetName: (target as HTMLInputElement).name || null
      }
    });
  }, [elementId]);

  // ブラートラッキング（フォーカス失う）
  const handleBlur = useCallback((e: FocusEvent) => {
    const target = e.target as HTMLElement;
    const focusDuration = Date.now() - lastActionTime.current;
    
    heuristicsDiscovery.recordAction({
      timestamp: Date.now(),
      actionType: 'blur',
      elementId,
      duration: focusDuration,
      context: {
        targetTag: target.tagName,
        targetValue: (target as HTMLInputElement).value || null,
        hasValue: !!(target as HTMLInputElement).value
      }
    });
  }, [elementId]);

  // ホバートラッキング
  const handleMouseEnter = useCallback(() => {
    hoverStartTime.current = Date.now();
    
    heuristicsDiscovery.recordAction({
      timestamp: Date.now(),
      actionType: 'hover_start',
      elementId,
      context: {}
    });
  }, [elementId]);

  const handleMouseLeave = useCallback(() => {
    const hoverDuration = Date.now() - hoverStartTime.current;
    
    heuristicsDiscovery.recordAction({
      timestamp: Date.now(),
      actionType: 'hover_end',
      elementId,
      duration: hoverDuration,
      context: {
        longHover: hoverDuration > 2000
      }
    });
  }, [elementId]);

  // 入力トラッキング
  const handleInput = useCallback(
    debounce((e: Event) => {
      const target = e.target as HTMLInputElement;
      
      heuristicsDiscovery.recordAction({
        timestamp: Date.now(),
        actionType: 'input',
        elementId,
        context: {
          inputLength: target.value.length,
          inputType: target.type,
          hasValue: !!target.value,
          isValid: target.validity.valid
        }
      });
    }),
    [elementId, debounceMs]
  );

  // エレメント参照を設定
  const setRef = useCallback((element: HTMLElement | null) => {
    elementRef.current = element;
  }, []);

  // イベントリスナー設定
  useEffect(() => {
    const element = elementRef.current;
    if (!element) return;

    if (trackClicks) {
      element.addEventListener('click', handleClick);
    }
    
    if (trackScroll) {
      element.addEventListener('scroll', handleScroll);
    }
    
    if (trackFocus) {
      element.addEventListener('focus', handleFocus, true);
      element.addEventListener('blur', handleBlur, true);
    }
    
    if (trackHover) {
      element.addEventListener('mouseenter', handleMouseEnter);
      element.addEventListener('mouseleave', handleMouseLeave);
    }
    
    if (trackInput) {
      element.addEventListener('input', handleInput);
    }

    // クリーンアップ
    return () => {
      if (trackClicks) {
        element.removeEventListener('click', handleClick);
      }
      
      if (trackScroll) {
        element.removeEventListener('scroll', handleScroll);
      }
      
      if (trackFocus) {
        element.removeEventListener('focus', handleFocus, true);
        element.removeEventListener('blur', handleBlur, true);
      }
      
      if (trackHover) {
        element.removeEventListener('mouseenter', handleMouseEnter);
        element.removeEventListener('mouseleave', handleMouseLeave);
      }
      
      if (trackInput) {
        element.removeEventListener('input', handleInput);
      }
    };
  }, [
    trackClicks,
    trackScroll,
    trackFocus,
    trackHover,
    trackInput,
    handleClick,
    handleScroll,
    handleFocus,
    handleBlur,
    handleMouseEnter,
    handleMouseLeave,
    handleInput
  ]);

  return { ref: setRef };
}

/**
 * コンポーネント全体のトラッキング
 */
export function useGlobalHeuristicsTracking() {
  useEffect(() => {
    // ページ遷移トラッキング
    const handleNavigation = () => {
      heuristicsDiscovery.recordAction({
        timestamp: Date.now(),
        actionType: 'navigation',
        elementId: 'global',
        context: {
          url: window.location.href,
          pathname: window.location.pathname,
          referrer: document.referrer
        }
      });
    };

    // エラートラッキング
    const handleError = (e: ErrorEvent) => {
      heuristicsDiscovery.recordAction({
        timestamp: Date.now(),
        actionType: 'error',
        elementId: 'global',
        context: {
          message: e.message,
          filename: e.filename,
          lineno: e.lineno
        }
      });
    };

    // パフォーマンストラッキング
    const trackPerformance = () => {
      const perfData = performance.getEntriesByType('navigation')[0] as PerformanceNavigationTiming;
      
      if (perfData) {
        heuristicsDiscovery.recordAction({
          timestamp: Date.now(),
          actionType: 'performance',
          elementId: 'global',
          context: {
            loadTime: perfData.loadEventEnd - perfData.fetchStart,
            domReady: perfData.domContentLoadedEventEnd - perfData.fetchStart,
            firstPaint: performance.getEntriesByName('first-paint')[0]?.startTime || 0
          }
        });
      }
    };

    window.addEventListener('popstate', handleNavigation);
    window.addEventListener('error', handleError);
    
    // 初回ロード時のパフォーマンス計測
    if (document.readyState === 'complete') {
      trackPerformance();
    } else {
      window.addEventListener('load', trackPerformance);
    }

    return () => {
      window.removeEventListener('popstate', handleNavigation);
      window.removeEventListener('error', handleError);
      window.removeEventListener('load', trackPerformance);
    };
  }, []);

  // ヒューリスティクス適用イベントのリスニング
  useEffect(() => {
    const handleHeuristicApplied = (e: CustomEvent) => {
      console.log('Heuristic applied:', e.detail);
      // UIの調整などを実行
      applyUIOptimization(e.detail);
    };

    window.addEventListener('heuristic-applied', handleHeuristicApplied as EventListener);

    return () => {
      window.removeEventListener('heuristic-applied', handleHeuristicApplied as EventListener);
    };
  }, []);
}

/**
 * UI最適化の適用
 */
function applyUIOptimization(detail: { action: string; data: any }) {
  switch (detail.action) {
    case 'reduce_options':
      // 選択肢を削減
      document.querySelectorAll('.option-item').forEach((item, index) => {
        if (index >= detail.data.maxOptions) {
          (item as HTMLElement).style.display = 'none';
        }
      });
      break;
      
    case 'emphasize_first':
      // 最初の要素を強調
      const firstElement = document.querySelector('.task-item:first-child');
      if (firstElement) {
        (firstElement as HTMLElement).classList.add('emphasized');
      }
      break;
      
    case 'prioritize_recent':
      // 最近使用した要素を前面に
      const recentElements = document.querySelectorAll('[data-recent="true"]');
      recentElements.forEach(el => {
        (el as HTMLElement).classList.add('priority');
      });
      break;
  }
}