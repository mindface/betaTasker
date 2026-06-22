"use client"
import React, { Component, ErrorInfo, ReactNode } from 'react';
import { HeuristicsError } from '../../errors/heuristics/HeuristicsError';
import { AnalysisError } from '../../errors/heuristics/AnalysisError';

interface Props {
  children: ReactNode;
  fallback?: React.ComponentType<ErrorBoundaryFallbackProps>;
  onError?: (error: Error, errorInfo: ErrorInfo) => void;
  resetOnPropsChange?: boolean;
  resetKeys?: Array<string | number>;
}

interface State {
  hasError: boolean;
  error: Error | null;
  errorInfo: ErrorInfo | null;
  errorId: string;
  retryCount: number;
  lastResetTimeStamp: number;
}

export interface ErrorBoundaryFallbackProps {
  error: Error;
  errorInfo: ErrorInfo | null;
  resetError: () => void;
  retryError: () => void;
  canRetry: boolean;
  retryCount: number;
  errorId: string;
}

export class HeuristicsErrorBoundary extends Component<Props, State> {
  private resetTimeoutId: number | null = null;

  // エラー回復戦略
  private recoverStrategies = {
    network: () => this.retryWithBackoff(),
    validation: () => this.resetForm(),
    analysis: () => this.fallbackToCache(),
  };

  constructor(props: Props) {
    super(props);

    this.state = {
      hasError: false,
      error: null,
      errorInfo: null,
      errorId: '',
      retryCount: 0,
      lastResetTimeStamp: Date.now()
    };
  }

  static getDerivedStateFromError(error: Error): Partial<State> {
    return {
      hasError: true,
      error,
      errorId: `error_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`
    };
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    const enhancedError = this.enhanceError(error, errorInfo);
    
    this.setState({
      errorInfo
    });

    // Log the error
    this.logError(enhancedError, errorInfo);

    // Call the onError prop
    this.props.onError?.(error, errorInfo);

    // Auto-retry for retryable errors
    this.scheduleAutoRetry(error);
  }

  componentDidUpdate(prevProps: Props) {
    const { resetKeys, resetOnPropsChange } = this.props;
    const { hasError } = this.state;
    
    if (hasError && prevProps.resetKeys !== resetKeys) {
      if (resetKeys?.some((resetKey, idx) => resetKey !== prevProps.resetKeys?.[idx])) {
        this.resetError();
      }
    }

    if (hasError && resetOnPropsChange && prevProps.children !== this.props.children) {
      this.resetError();
    }
  }

  componentWillUnmount() {
    if (this.resetTimeoutId) {
      clearTimeout(this.resetTimeoutId);
    }
  }

  private enhanceError(error: Error, errorInfo: ErrorInfo): Error {
    // Add React component stack trace to error details
    if (error instanceof HeuristicsError) {
      return new (error.constructor as any)(
        error.message,
        error.code,
        {
          ...error.details,
          componentStack: errorInfo.componentStack
        }
      );
    }

    return error;
  }

  private logError(error: Error, errorInfo: ErrorInfo) {
    const errorDetails = {
      message: error.message,
      stack: error.stack,
      componentStack: errorInfo.componentStack,
      errorId: this.state.errorId,
      timestamp: new Date().toISOString(),
      retryCount: this.state.retryCount
    };

    // Log to console in development
    if (process.env.NODE_ENV === 'development') {
      console.error('HeuristicsErrorBoundary caught an error:', errorDetails);
    }

    // Send to monitoring service (implement based on your monitoring setup)
    this.sendToMonitoring(errorDetails);
  }

  private sendToMonitoring(errorDetails: any) {
    // Implement your error monitoring service integration here
    // For example: Sentry, LogRocket, etc.
    
    try {
      if (typeof window !== 'undefined' && (window as any).errorReportingService) {
        (window as any).errorReportingService.reportError({
          type: 'HeuristicsError',
          details: errorDetails
        });
      }
    } catch (e) {
      console.warn('Failed to send error to monitoring service:', e);
    }
  }

  private scheduleAutoRetry(error: Error) {
    if (error instanceof HeuristicsError && error.isRetryable()) {
      const delay = error.getRetryDelay();
      
      this.resetTimeoutId = window.setTimeout(() => {
        this.retryError();
      }, delay);
    }
  }

  private retryWithBackoff = async () => {
    const backoffDelay = Math.min(1000 * Math.pow(2, this.state.retryCount), 10000);
    
    return new Promise<void>((resolve) => {
      setTimeout(() => {
        this.retryError();
        resolve();
      }, backoffDelay);
    });
  };

  private resetForm = () => {
    // Reset form state if applicable
    // This would be implemented based on your specific form handling
    this.resetError();
  };

  private fallbackToCache = () => {
    // Implement cache fallback logic
    // This could involve switching to cached data
    console.log('Falling back to cached data...');
    this.resetError();
  };

  private resetError = () => {
    if (this.resetTimeoutId) {
      clearTimeout(this.resetTimeoutId);
      this.resetTimeoutId = null;
    }

    this.setState({
      hasError: false,
      error: null,
      errorInfo: null,
      errorId: '',
      retryCount: 0,
      lastResetTimeStamp: Date.now()
    });
  };

  private retryError = () => {
    if (this.state.retryCount >= 3) {
      console.warn('Maximum retry attempts reached');
      return;
    }

    this.setState(prevState => ({
      hasError: false,
      error: null,
      errorInfo: null,
      retryCount: prevState.retryCount + 1,
      lastResetTimeStamp: Date.now()
    }));
  };

  private canRetry(): boolean {
    const { error, retryCount } = this.state;
    
    if (retryCount >= 3) return false;
    if (!error) return false;
    
    if (error instanceof HeuristicsError) {
      return error.isRetryable();
    }
    
    return true;
  }

  render() {
    const { hasError, error, errorInfo, errorId, retryCount } = this.state;
    const { children, fallback: FallbackComponent } = this.props;

    if (hasError && error) {
      const fallbackProps: ErrorBoundaryFallbackProps = {
        error,
        errorInfo,
        resetError: this.resetError,
        retryError: this.retryError,
        canRetry: this.canRetry(),
        retryCount,
        errorId
      };

      if (FallbackComponent) {
        return <FallbackComponent {...fallbackProps} />;
      }

      return <DefaultErrorFallback {...fallbackProps} />;
    }

    return children;
  }
}

// Default fallback component
const DefaultErrorFallback: React.FC<ErrorBoundaryFallbackProps> = ({
  error,
  resetError,
  retryError,
  canRetry,
  retryCount,
  errorId
}) => {
  const isHeuristicsError = error instanceof HeuristicsError;
  const severity = isHeuristicsError ? error.getSeverity() : 'medium';
  
  const getSeverityColor = (severity: string) => {
    switch (severity) {
      case 'critical': return '#d63031';
      case 'high': return '#e17055';
      case 'medium': return '#fdcb6e';
      case 'low': return '#55a3ff';
      default: return '#636e72';
    }
  };

  const getSeverityLabel = (severity: string) => {
    switch (severity) {
      case 'critical': return '重大';
      case 'high': return '高';
      case 'medium': return '中';
      case 'low': return '低';
      default: return '不明';
    }
  };

  return (
    <div style={{
      padding: '24px',
      border: `2px solid ${getSeverityColor(severity)}`,
      borderRadius: '8px',
      backgroundColor: '#fff',
      margin: '16px',
      boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)'
    }}>
      <div style={{ display: 'flex', alignItems: 'center', marginBottom: '16px' }}>
        <span style={{
          fontSize: '24px',
          marginRight: '12px'
        }}>⚠️</span>
        <div>
          <h3 style={{ 
            margin: '0',
            color: getSeverityColor(severity),
            fontSize: '18px'
          }}>
            エラーが発生しました
          </h3>
          <div style={{ 
            fontSize: '12px', 
            color: '#666',
            marginTop: '4px'
          }}>
            重要度: {getSeverityLabel(severity)} | エラーID: {errorId}
          </div>
        </div>
      </div>

      <div style={{
        backgroundColor: '#f8f9fa',
        padding: '12px',
        borderRadius: '4px',
        marginBottom: '16px',
        fontSize: '14px',
        color: '#495057'
      }}>
        <strong>エラー内容:</strong><br />
        {error.message}
      </div>

      {isHeuristicsError && (
        <div style={{
          backgroundColor: '#e3f2fd',
          padding: '12px',
          borderRadius: '4px',
          marginBottom: '16px',
          fontSize: '13px',
          color: '#1976d2'
        }}>
          <strong>カテゴリ:</strong> {error.getCategory()}<br />
          <strong>再試行可能:</strong> {error.isRetryable() ? 'はい' : 'いいえ'}
        </div>
      )}

      {retryCount > 0 && (
        <div style={{
          backgroundColor: '#fff3cd',
          padding: '8px 12px',
          borderRadius: '4px',
          marginBottom: '16px',
          fontSize: '12px',
          color: '#856404'
        }}>
          再試行回数: {retryCount}/3
        </div>
      )}

      <div style={{ display: 'flex', gap: '12px', flexWrap: 'wrap' }}>
        {canRetry && (
          <button
            onClick={retryError}
            style={{
              padding: '8px 16px',
              backgroundColor: '#28a745',
              color: 'white',
              border: 'none',
              borderRadius: '4px',
              cursor: 'pointer',
              fontSize: '14px'
            }}
            onMouseOver={(e) => e.currentTarget.style.backgroundColor = '#218838'}
            onMouseOut={(e) => e.currentTarget.style.backgroundColor = '#28a745'}
          >
            🔄 再試行 ({3 - retryCount}回まで)
          </button>
        )}
        
        <button
          onClick={resetError}
          style={{
            padding: '8px 16px',
            backgroundColor: '#6c757d',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '14px'
          }}
          onMouseOver={(e) => e.currentTarget.style.backgroundColor = '#5a6268'}
          onMouseOut={(e) => e.currentTarget.style.backgroundColor = '#6c757d'}
        >
          🔄 リセット
        </button>

        <button
          onClick={() => window.location.reload()}
          style={{
            padding: '8px 16px',
            backgroundColor: '#17a2b8',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
            fontSize: '14px'
          }}
          onMouseOver={(e) => e.currentTarget.style.backgroundColor = '#138496'}
          onMouseOut={(e) => e.currentTarget.style.backgroundColor = '#17a2b8'}
        >
          🔄 ページ再読み込み
        </button>
      </div>

      {process.env.NODE_ENV === 'development' && (
        <details style={{ marginTop: '16px' }}>
          <summary style={{ 
            cursor: 'pointer', 
            fontSize: '12px',
            color: '#666'
          }}>
            開発者情報を表示
          </summary>
          <pre style={{
            backgroundColor: '#f8f9fa',
            padding: '12px',
            borderRadius: '4px',
            fontSize: '11px',
            overflow: 'auto',
            marginTop: '8px'
          }}>
            {error.stack}
          </pre>
        </details>
      )}
    </div>
  );
};

// Specialized error boundaries for different areas
export const AnalysisErrorBoundary: React.FC<{ children: ReactNode }> = ({ children }) => (
  <HeuristicsErrorBoundary
    onError={(error, errorInfo) => {
      console.log('Analysis error caught:', error.message);
    }}
  >
    {children}
  </HeuristicsErrorBoundary>
);

export const TrackingErrorBoundary: React.FC<{ children: ReactNode }> = ({ children }) => (
  <HeuristicsErrorBoundary
    onError={(error, errorInfo) => {
      console.log('Tracking error caught:', error.message);
    }}
  >
    {children}
  </HeuristicsErrorBoundary>
);

export const PatternErrorBoundary: React.FC<{ children: ReactNode }> = ({ children }) => (
  <HeuristicsErrorBoundary
    onError={(error, errorInfo) => {
      console.log('Pattern error caught:', error.message);
    }}
  >
    {children}
  </HeuristicsErrorBoundary>
);

export const InsightErrorBoundary: React.FC<{ children: ReactNode }> = ({ children }) => (
  <HeuristicsErrorBoundary
    onError={(error, errorInfo) => {
      console.log('Insight error caught:', error.message);
    }}
  >
    {children}
  </HeuristicsErrorBoundary>
);

export default HeuristicsErrorBoundary;