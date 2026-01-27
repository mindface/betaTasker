export interface TrackingId {
  value: number;
}

export interface SessionId {
  value: string;
}

export interface TrackingAction {
  value: string;
  isValid(): boolean;
}

export class TrackingActionValue implements TrackingAction {
  constructor(public readonly value: string) {
    if (!this.isValid()) {
      throw new Error('Tracking action cannot be empty');
    }
  }

  isValid(): boolean {
    return this.value.trim().length > 0;
  }
}

export interface TrackingContext {
  [key: string]: any;
}

export interface TrackingDuration {
  value: number;
  isValid(): boolean;
}

export class TrackingDurationValue implements TrackingDuration {
  constructor(public readonly value: number) {
    if (!this.isValid()) {
      throw new Error('Tracking duration must be non-negative');
    }
  }

  isValid(): boolean {
    return this.value >= 0;
  }
}

export interface Tracking {
  readonly id: TrackingId;
  readonly userId: UserId;
  readonly action: TrackingAction;
  readonly context: TrackingContext;
  readonly sessionId?: SessionId;
  readonly duration: TrackingDuration;
  readonly timestamp: Date;
}

export interface TrackingRequest {
  readonly userId: UserId;
  readonly action: TrackingAction;
  readonly context: TrackingContext;
  readonly sessionId?: SessionId;
  readonly duration: TrackingDuration;
}

export class TrackingEntity implements Tracking {
  constructor(
    public readonly id: TrackingId,
    public readonly userId: UserId,
    public readonly action: TrackingAction,
    public readonly context: TrackingContext,
    public readonly duration: TrackingDuration,
    public readonly timestamp: Date,
    public readonly sessionId?: SessionId
  ) {}

  static create(request: TrackingRequest): TrackingEntity {
    return new TrackingEntity(
      { value: 0 }, // Will be set by repository
      request.userId,
      request.action,
      request.context,
      request.duration,
      new Date(),
      request.sessionId
    );
  }

  isSameSession(other: Tracking): boolean {
    return this.sessionId?.value === other.sessionId?.value;
  }

  isRecentAction(thresholdMs: number = 5000): boolean {
    const now = new Date();
    return (now.getTime() - this.timestamp.getTime()) <= thresholdMs;
  }

  getContextValue<T>(key: string): T | undefined {
    return this.context[key] as T;
  }

  withUpdatedContext(newContext: TrackingContext): TrackingEntity {
    return new TrackingEntity(
      this.id,
      this.userId,
      this.action,
      { ...this.context, ...newContext },
      this.duration,
      this.timestamp,
      this.sessionId
    );
  }
}