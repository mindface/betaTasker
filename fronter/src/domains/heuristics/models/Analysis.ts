export interface AnalysisId {
  value: number;
}

export interface UserId {
  value: number;
}

export interface TaskId {
  value: number;
}

export type AnalysisType = 'performance' | 'behavior' | 'pattern' | 'cognitive' | 'efficiency';

export type AnalysisStatus = 'pending' | 'processing' | 'completed' | 'failed';

export interface AnalysisScore {
  value: number;
  min: number;
  max: number;
  isValid(): boolean;
}

export class AnalysisScoreValue implements AnalysisScore {
  constructor(
    public readonly value: number,
    public readonly min: number = 0,
    public readonly max: number = 100
  ) {
    if (!this.isValid()) {
      throw new Error(`Analysis score must be between ${min} and ${max}`);
    }
  }

  isValid(): boolean {
    return this.value >= this.min && this.value <= this.max;
  }
}

export interface AnalysisData {
  [key: string]: any;
}

export interface Analysis {
  readonly id: AnalysisId;
  readonly userId: UserId;
  readonly taskId?: TaskId;
  readonly analysisType: AnalysisType;
  readonly data: AnalysisData;
  readonly result: string;
  readonly score: AnalysisScore;
  readonly status: AnalysisStatus;
  readonly createdAt: Date;
  readonly updatedAt: Date;
}

export interface AnalysisRequest {
  readonly userId: UserId;
  readonly taskId?: TaskId;
  readonly analysisType: AnalysisType;
  readonly data: AnalysisData;
}

export class AnalysisEntity implements Analysis {
  constructor(
    public readonly id: AnalysisId,
    public readonly userId: UserId,
    public readonly analysisType: AnalysisType,
    public readonly data: AnalysisData,
    public readonly result: string,
    public readonly score: AnalysisScore,
    public readonly status: AnalysisStatus,
    public readonly createdAt: Date,
    public readonly updatedAt: Date,
    public readonly taskId?: TaskId
  ) {}

  static create(request: AnalysisRequest): AnalysisEntity {
    return new AnalysisEntity(
      { value: 0 }, // Will be set by repository
      request.userId,
      request.analysisType,
      request.data,
      '', // Will be populated by analysis service
      new AnalysisScoreValue(0),
      'pending',
      new Date(),
      new Date(),
      request.taskId
    );
  }

  isCompleted(): boolean {
    return this.status === 'completed';
  }

  isFailed(): boolean {
    return this.status === 'failed';
  }

  canRetry(): boolean {
    return this.status === 'failed';
  }

  updateStatus(status: AnalysisStatus): AnalysisEntity {
    return new AnalysisEntity(
      this.id,
      this.userId,
      this.analysisType,
      this.data,
      this.result,
      this.score,
      status,
      this.createdAt,
      new Date(),
      this.taskId
    );
  }

  updateResult(result: string, score: AnalysisScore): AnalysisEntity {
    return new AnalysisEntity(
      this.id,
      this.userId,
      this.analysisType,
      this.data,
      result,
      score,
      'completed',
      this.createdAt,
      new Date(),
      this.taskId
    );
  }
}