import React, { Component, ErrorInfo, ReactNode } from 'react';

interface Props {
  children: ReactNode;
  fallback?: ReactNode;
  onError?: (error: Error, errorInfo: ErrorInfo) => void;
}

interface State {
  hasError: boolean;
  error: Error | null;
  errorInfo: ErrorInfo | null;
}

export class ErrorBoundary extends Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      hasError: false,
      error: null,
      errorInfo: null
    };
  }

  static getDerivedStateFromError(error: Error): State {
    return {
      hasError: true,
      error,
      errorInfo: null
    };
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    this.setState({
      error,
      errorInfo
    });

    // エラーログの送信
    this.logErrorToService(error, errorInfo);

    // カスタムエラーハンドラー
    if (this.props.onError) {
      this.props.onError(error, errorInfo);
    }
  }

  private logErrorToService(error: Error, errorInfo: ErrorInfo) {
    // エラーログの送信処理
    console.error('ErrorBoundary caught an error:', error, errorInfo);
    
    // 本番環境では外部サービスにログを送信
    if (process.env.NODE_ENV === 'production') {
      // 例: Sentry, LogRocket, カスタムエラー追跡サービス
      try {
        // エラー情報の送信
        fetch('/api/error-log', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            error: {
              name: error.name,
              message: error.message,
              stack: error.stack,
            },
            errorInfo: {
              componentStack: errorInfo.componentStack,
            },
            timestamp: new Date().toISOString(),
            userAgent: navigator.userAgent,
            url: window.location.href,
          }),
        }).catch(console.error);
      } catch (logError) {
        console.error('Failed to send error log:', logError);
      }
    }
  }

  private handleReset = () => {
    this.setState({
      hasError: false,
      error: null,
      errorInfo: null
    });
  };

  private handleReportError = () => {
    if (this.state.error && this.state.errorInfo) {
      // エラー報告の処理
      const errorReport = {
        error: this.state.error,
        errorInfo: this.state.errorInfo,
        timestamp: new Date().toISOString(),
        userAgent: navigator.userAgent,
        url: window.location.href,
      };

      // クリップボードにコピー
      navigator.clipboard.writeText(JSON.stringify(errorReport, null, 2))
        .then(() => {
          alert('エラー情報がクリップボードにコピーされました。開発者にお知らせください。');
        })
        .catch(() => {
          // クリップボードAPIが利用できない場合
          console.log('Error Report:', errorReport);
          alert('エラー情報がコンソールに出力されました。開発者にお知らせください。');
        });
    }
  };

  render() {
    if (this.state.hasError) {
      // カスタムフォールバックUI
      if (this.props.fallback) {
        return this.props.fallback;
      }

      // デフォルトのエラーUI
      return (
        <div className="error-boundary">
          <div className="error-boundary__content">
            <div className="error-boundary__icon">⚠️</div>
            <h2 className="error-boundary__title">エラーが発生しました</h2>
            <p className="error-boundary__message">
              予期しないエラーが発生しました。ページを再読み込みするか、しばらく時間をおいてから再度お試しください。
            </p>
            
            <div className="error-boundary__actions">
              <button
                className="error-boundary__button error-boundary__button--primary"
                onClick={this.handleReset}
              >
                再試行
              </button>
              
              <button
                className="error-boundary__button error-boundary__button--secondary"
                onClick={this.handleReportError}
              >
                エラーを報告
              </button>
            </div>

            {process.env.NODE_ENV === 'development' && this.state.error && (
              <details className="error-boundary__details">
                <summary>エラーの詳細（開発モード）</summary>
                <div className="error-boundary__error-info">
                  <h4>エラー:</h4>
                  <pre>{this.state.error.toString()}</pre>
                  
                  {this.state.errorInfo && (
                    <>
                      <h4>コンポーネントスタック:</h4>
                      <pre>{this.state.errorInfo.componentStack}</pre>
                    </>
                  )}
                </div>
              </details>
            )}
          </div>
        </div>
      );
    }

    return this.props.children;
  }
}

// ヒューリスティクス機能専用のエラーバウンダリ
export const HeuristicsErrorBoundary: React.FC<{ children: ReactNode }> = ({ children }) => {
  const handleError = (error: Error, errorInfo: ErrorInfo) => {
    // ヒューリスティクス機能固有のエラーハンドリング
    console.error('Heuristics Error:', error, errorInfo);
    
    // エラーをReduxストアに記録
    // 必要に応じて実装
  };

  const fallback = (
    <div className="heuristics-error-fallback">
      <div className="heuristics-error-fallback__content">
        <h3>ヒューリスティクス機能でエラーが発生しました</h3>
        <p>分析機能の一部が利用できません。ページを再読み込みしてください。</p>
        <button onClick={() => window.location.reload()}>
          ページを再読み込み
        </button>
      </div>
    </div>
  );

  return (
    <ErrorBoundary
      fallback={fallback}
      onError={handleError}
    >
      {children}
    </ErrorBoundary>
  );
};

// 高階コンポーネントとしてエラーバウンダリを適用
export const withErrorBoundary = <P extends object>(
  Component: React.ComponentType<P>,
  fallback?: ReactNode
) => {
  const WrappedComponent = (props: P) => (
    <ErrorBoundary fallback={fallback}>
      <Component {...props} />
    </ErrorBoundary>
  );

  WrappedComponent.displayName = `withErrorBoundary(${Component.displayName || Component.name})`;
  
  return WrappedComponent;
};
