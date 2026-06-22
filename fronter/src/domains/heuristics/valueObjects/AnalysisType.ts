export type AnalysisTypeValue = 'performance' | 'behavior' | 'pattern' | 'cognitive' | 'efficiency';

export class AnalysisType {
  private constructor(private readonly value: AnalysisTypeValue) {}

  static readonly PERFORMANCE = new AnalysisType('performance');
  static readonly BEHAVIOR = new AnalysisType('behavior');
  static readonly PATTERN = new AnalysisType('pattern');
  static readonly COGNITIVE = new AnalysisType('cognitive');
  static readonly EFFICIENCY = new AnalysisType('efficiency');

  static fromString(value: string): AnalysisType {
    switch (value.toLowerCase()) {
      case 'performance':
        return AnalysisType.PERFORMANCE;
      case 'behavior':
        return AnalysisType.BEHAVIOR;
      case 'pattern':
        return AnalysisType.PATTERN;
      case 'cognitive':
        return AnalysisType.COGNITIVE;
      case 'efficiency':
        return AnalysisType.EFFICIENCY;
      default:
        throw new Error(`Invalid analysis type: ${value}`);
    }
  }

  getValue(): AnalysisTypeValue {
    return this.value;
  }

  toString(): string {
    return this.value;
  }

  equals(other: AnalysisType): boolean {
    return this.value === other.value;
  }

  getDisplayName(): string {
    switch (this.value) {
      case 'performance':
        return 'パフォーマンス分析';
      case 'behavior':
        return '行動分析';
      case 'pattern':
        return 'パターン分析';
      case 'cognitive':
        return '認知分析';
      case 'efficiency':
        return '効率性分析';
    }
  }

  getDescription(): string {
    switch (this.value) {
      case 'performance':
        return 'タスクの実行速度や完了率を分析します';
      case 'behavior':
        return 'ユーザーの行動パターンを分析します';
      case 'pattern':
        return 'データ内のパターンを検出・分析します';
      case 'cognitive':
        return '認知的負荷や判断プロセスを分析します';
      case 'efficiency':
        return 'プロセスの効率性と最適化を分析します';
    }
  }

  getExpectedDuration(): number {
    switch (this.value) {
      case 'performance':
        return 30000; // 30 seconds
      case 'behavior':
        return 60000; // 1 minute
      case 'pattern':
        return 120000; // 2 minutes
      case 'cognitive':
        return 90000; // 1.5 minutes
      case 'efficiency':
        return 45000; // 45 seconds
    }
  }

  requiresTaskId(): boolean {
    return this.value === 'performance' || this.value === 'efficiency';
  }

  static getAllTypes(): AnalysisType[] {
    return [
      AnalysisType.PERFORMANCE,
      AnalysisType.BEHAVIOR,
      AnalysisType.PATTERN,
      AnalysisType.COGNITIVE,
      AnalysisType.EFFICIENCY
    ];
  }

  static getTypesByCategory(): Record<string, AnalysisType[]> {
    return {
      'user-focused': [AnalysisType.BEHAVIOR, AnalysisType.COGNITIVE],
      'task-focused': [AnalysisType.PERFORMANCE, AnalysisType.EFFICIENCY],
      'data-focused': [AnalysisType.PATTERN]
    };
  }
}