import { HeuristicsError } from './HeuristicsError';

export class AnalysisError extends HeuristicsError {
  getCategory(): string {
    return 'analysis';
  }

  getSeverity(): 'low' | 'medium' | 'high' | 'critical' {
    switch (this.code) {
      case 'ANALYSIS_VALIDATION_FAILED':
      case 'ANALYSIS_INVALID_DATA':
        return 'medium';
      case 'ANALYSIS_EXECUTION_FAILED':
      case 'ANALYSIS_TIMEOUT':
        return 'high';
      case 'ANALYSIS_CRITICAL_FAILURE':
        return 'critical';
      default:
        return 'low';
    }
  }

  isRetryable(): boolean {
    const nonRetryableCodes = [
      'ANALYSIS_VALIDATION_FAILED',
      'ANALYSIS_INVALID_DATA',
      'ANALYSIS_UNSUPPORTED_TYPE'
    ];
    return !nonRetryableCodes.includes(this.code);
  }
}

export class AnalysisValidationError extends AnalysisError {
  constructor(message: string, details?: Record<string, any>) {
    super(message, 'ANALYSIS_VALIDATION_FAILED', details);
  }
}

export class AnalysisExecutionError extends AnalysisError {
  constructor(message: string, details?: Record<string, any>) {
    super(message, 'ANALYSIS_EXECUTION_FAILED', details);
  }
}

export class AnalysisTimeoutError extends AnalysisError {
  constructor(message: string, timeoutMs: number) {
    super(message, 'ANALYSIS_TIMEOUT', { timeoutMs });
  }
}

export class AnalysisNotFoundError extends AnalysisError {
  constructor(analysisId: number) {
    super(`Analysis with ID ${analysisId} not found`, 'ANALYSIS_NOT_FOUND', { analysisId });
  }

  getSeverity(): 'low' | 'medium' | 'high' | 'critical' {
    return 'medium';
  }

  isRetryable(): boolean {
    return false;
  }
}

export class AnalysisInvalidDataError extends AnalysisError {
  constructor(message: string, invalidFields: string[]) {
    super(message, 'ANALYSIS_INVALID_DATA', { invalidFields });
  }

  isRetryable(): boolean {
    return false;
  }
}

export class AnalysisUnsupportedTypeError extends AnalysisError {
  constructor(analysisType: string) {
    super(
      `Unsupported analysis type: ${analysisType}`,
      'ANALYSIS_UNSUPPORTED_TYPE',
      { analysisType }
    );
  }

  isRetryable(): boolean {
    return false;
  }
}

export class AnalysisCriticalFailureError extends AnalysisError {
  constructor(message: string, originalError?: Error) {
    super(message, 'ANALYSIS_CRITICAL_FAILURE', {
      originalError: originalError?.message,
      originalStack: originalError?.stack
    });
  }

  getSeverity(): 'low' | 'medium' | 'high' | 'critical' {
    return 'critical';
  }
}