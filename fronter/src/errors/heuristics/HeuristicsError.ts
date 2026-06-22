export abstract class HeuristicsError extends Error {
  public readonly code: string;
  public readonly timestamp: Date;
  public readonly details?: Record<string, any>;

  constructor(
    message: string,
    code: string,
    details?: Record<string, any>
  ) {
    super(message);
    this.name = this.constructor.name;
    this.code = code;
    this.timestamp = new Date();
    this.details = details;

    // Error.captureStackTrace がある場合のみ呼び出す
    if (Error.captureStackTrace) {
      Error.captureStackTrace(this, this.constructor);
    }
  }

  toJSON() {
    return {
      name: this.name,
      message: this.message,
      code: this.code,
      timestamp: this.timestamp.toISOString(),
      details: this.details,
      stack: this.stack
    };
  }

  toString(): string {
    return `${this.name}: ${this.message} (Code: ${this.code})`;
  }

  abstract getCategory(): string;
  abstract getSeverity(): 'low' | 'medium' | 'high' | 'critical';
  abstract isRetryable(): boolean;

  getRetryDelay(): number {
    if (!this.isRetryable()) return 0;
    
    switch (this.getSeverity()) {
      case 'low': return 1000;
      case 'medium': return 3000;
      case 'high': return 5000;
      case 'critical': return 10000;
    }
  }

  shouldNotifyUser(): boolean {
    return this.getSeverity() === 'high' || this.getSeverity() === 'critical';
  }

  getErrorContext(): Record<string, any> {
    return {
      category: this.getCategory(),
      severity: this.getSeverity(),
      retryable: this.isRetryable(),
      timestamp: this.timestamp,
      ...this.details
    };
  }
}