import { useState, useCallback, useEffect } from 'react';
import { HeuristicsError } from '../../errors/heuristics/HeuristicsError';
import { AnalysisError } from '../../errors/heuristics/AnalysisError';

export interface ErrorInfo {
  error: Error;
  errorInfo?: {
    componentStack: string;
  };
  timestamp: Date;
  errorId: string;
  context?: Record<string, any>;
}

export interface ErrorBoundaryState {
  hasError: boolean;
  error: ErrorInfo | null;
  retryCount: number;
  isRetrying: boolean;
  lastErrorTime: Date | null;
}

export interface ErrorBoundaryActions {
  captureError: (error: Error, errorInfo?: { componentStack: string }, context?: Record<string, any>) => void;
  retry: () => Promise<void>;
  reset: () => void;
  clearError: () => void;
}

export interface ErrorBoundaryConfig {
  maxRetries?: number;
  retryDelay?: number;
  autoRetry?: boolean;
  onError?: (errorInfo: ErrorInfo) => void;
  onRetry?: (retryCount: number) => void;
  onReset?: () => void;
  fallbackComponent?: React.ComponentType<{ error: ErrorInfo; retry: () => void; reset: () => void }>;
}

export const useErrorBoundary = (config: ErrorBoundaryConfig = {}): [ErrorBoundaryState, ErrorBoundaryActions] => {
  const {
    maxRetries = 3,
    retryDelay = 1000,
    autoRetry = false,
    onError,
    onRetry,
    onReset
  } = config;

  const [state, setState] = useState<ErrorBoundaryState>({
    hasError: false,
    error: null,
    retryCount: 0,
    isRetrying: false,
    lastErrorTime: null
  });

  const generateErrorId = useCallback((): string => {
    return `error_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`;
  }, []);

  const captureError = useCallback((
    error: Error,
    errorInfo?: { componentStack: string },
    context?: Record<string, any>
  ) => {
    const errorData: ErrorInfo = {
      error,
      errorInfo,
      timestamp: new Date(),
      errorId: generateErrorId(),
      context
    };

    setState(prevState => ({
      ...prevState,
      hasError: true,
      error: errorData,
      lastErrorTime: new Date()
    }));

    // Log error for monitoring
    console.error('Error captured by useErrorBoundary:', {
      message: error.message,
      stack: error.stack,
      errorInfo,
      context,
      timestamp: errorData.timestamp.toISOString()
    });

    // Call error handler
    onError?.(errorData);

    // Send to error reporting service (if configured)
    if (typeof window !== 'undefined' && (window as any).errorReportingService) {
      (window as any).errorReportingService.reportError(errorData);
    }

    // Auto retry for retryable errors
    if (autoRetry && error instanceof HeuristicsError && error.isRetryable()) {
      setTimeout(() => {
        retry();
      }, error.getRetryDelay());
    }
  }, [generateErrorId, onError, autoRetry, maxRetries]);

  const retry = useCallback(async () => {
    if (state.retryCount >= maxRetries) {
      console.warn('Max retries exceeded, cannot retry');
      return;
    }

    if (!state.error) {
      console.warn('No error to retry');
      return;
    }

    setState(prevState => ({
      ...prevState,
      isRetrying: true
    }));

    try {
      // Determine retry delay
      let delay = retryDelay;
      if (state.error.error instanceof HeuristicsError) {
        delay = state.error.error.getRetryDelay();
      }

      // Wait for retry delay
      if (delay > 0) {
        await new Promise(resolve => setTimeout(resolve, delay));
      }

      // Increment retry count
      setState(prevState => ({
        ...prevState,
        retryCount: prevState.retryCount + 1,
        isRetrying: false
      }));

      onRetry?.(state.retryCount + 1);

      // The actual retry logic should be handled by the component
      // This hook only manages the state
      console.log(`Retry attempt ${state.retryCount + 1} for error: ${state.error.error.message}`);

    } catch (retryError) {
      console.error('Error during retry:', retryError);
      setState(prevState => ({
        ...prevState,
        isRetrying: false
      }));
    }
  }, [state.retryCount, state.error, maxRetries, retryDelay, onRetry]);

  const reset = useCallback(() => {
    setState({
      hasError: false,
      error: null,
      retryCount: 0,
      isRetrying: false,
      lastErrorTime: null
    });

    onReset?.();
  }, [onReset]);

  const clearError = useCallback(() => {
    setState(prevState => ({
      ...prevState,
      hasError: false,
      error: null,
      isRetrying: false
    }));
  }, []);

  // Auto-reset after successful operation
  useEffect(() => {
    if (state.hasError && state.retryCount > 0 && !state.isRetrying) {
      // If we've been retrying and no new errors for a while, consider it resolved
      const timeSinceLastError = state.lastErrorTime 
        ? Date.now() - state.lastErrorTime.getTime()
        : 0;

      if (timeSinceLastError > 30000) { // 30 seconds
        reset();
      }
    }
  }, [state.hasError, state.retryCount, state.isRetrying, state.lastErrorTime, reset]);

  const actions: ErrorBoundaryActions = {
    captureError,
    retry,
    reset,
    clearError
  };

  return [state, actions];
};

// Specialized hook for different types of errors
export const useAnalysisErrorBoundary = (config: ErrorBoundaryConfig = {}) => {
  const [state, actions] = useErrorBoundary({
    ...config,
    onError: (errorInfo) => {
      // Specialized handling for analysis errors
      if (errorInfo.error instanceof AnalysisError) {
        console.log(`Analysis error category: ${errorInfo.error.getCategory()}`);
        console.log(`Analysis error severity: ${errorInfo.error.getSeverity()}`);
        console.log(`Analysis error retryable: ${errorInfo.error.isRetryable()}`);
      }
      
      config.onError?.(errorInfo);
    }
  });

  return [state, actions];
};

// Hook for error recovery strategies
export interface ErrorRecoveryStrategy {
  network: () => Promise<void>;
  validation: () => Promise<void>;
  analysis: () => Promise<void>;
  fallback: () => Promise<void>;
}

export const useErrorRecovery = (strategies: Partial<ErrorRecoveryStrategy>) => {
  const [state, actions] = useErrorBoundary();

  const recoverFromError = useCallback(async (error: Error) => {
    if (error instanceof HeuristicsError) {
      const category = error.getCategory();
      
      switch (category) {
        case 'network':
          await strategies.network?.();
          break;
        case 'validation':
          await strategies.validation?.();
          break;
        case 'analysis':
          await strategies.analysis?.();
          break;
        default:
          await strategies.fallback?.();
      }
    } else {
      await strategies.fallback?.();
    }
    
    actions.clearError();
  }, [strategies, actions]);

  return {
    ...state,
    ...actions,
    recoverFromError
  };
};

// Error boundary component factory
export const createErrorBoundaryComponent = (config: ErrorBoundaryConfig) => {
  return function ErrorBoundary({ 
    children, 
    fallback 
  }: { 
    children: React.ReactNode; 
    fallback?: React.ComponentType<any> 
  }) {
    const [state, actions] = useErrorBoundary(config);

    if (state.hasError && state.error) {
      const FallbackComponent = fallback || config.fallbackComponent;
      
      if (FallbackComponent) {
        return <FallbackComponent 
          error={state.error} 
          retry={actions.retry} 
          reset={actions.reset} 
        />;
      }

      // Default fallback UI
      return (
        <div style={{ padding: '20px', border: '1px solid #ff6b6b', borderRadius: '4px', backgroundColor: '#fff5f5' }}>
          <h3 style={{ color: '#c92a2a', margin: '0 0 10px 0' }}>エラーが発生しました</h3>
          <p style={{ margin: '0 0 15px 0' }}>{state.error.error.message}</p>
          <div style={{ display: 'flex', gap: '10px' }}>
            <button 
              onClick={actions.retry} 
              disabled={state.isRetrying || state.retryCount >= 3}
              style={{ 
                padding: '8px 16px', 
                backgroundColor: state.isRetrying ? '#ccc' : '#228be6',
                color: 'white',
                border: 'none',
                borderRadius: '4px',
                cursor: state.isRetrying ? 'not-allowed' : 'pointer'
              }}
            >
              {state.isRetrying ? '再試行中...' : '再試行'}
            </button>
            <button 
              onClick={actions.reset}
              style={{ 
                padding: '8px 16px', 
                backgroundColor: '#868e96',
                color: 'white',
                border: 'none',
                borderRadius: '4px',
                cursor: 'pointer'
              }}
            >
              リセット
            </button>
          </div>
        </div>
      );
    }

    return <>{children}</>;
  };
};