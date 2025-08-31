export interface PatternId {
  value: number;
}

export interface PatternName {
  value: string;
  isValid(): boolean;
}

export class PatternNameValue implements PatternName {
  constructor(public readonly value: string) {
    if (!this.isValid()) {
      throw new Error('Pattern name cannot be empty');
    }
  }

  isValid(): boolean {
    return this.value.trim().length > 0;
  }
}

export type PatternCategory = 'behavioral' | 'cognitive' | 'temporal' | 'sequential' | 'frequency';

export interface PatternFrequency {
  value: number;
  isValid(): boolean;
}

export class PatternFrequencyValue implements PatternFrequency {
  constructor(public readonly value: number) {
    if (!this.isValid()) {
      throw new Error('Pattern frequency must be non-negative');
    }
  }

  isValid(): boolean {
    return this.value >= 0;
  }
}

export interface PatternAccuracy {
  value: number;
  isValid(): boolean;
}

export class PatternAccuracyValue implements PatternAccuracy {
  constructor(public readonly value: number) {
    if (!this.isValid()) {
      throw new Error('Pattern accuracy must be between 0 and 1');
    }
  }

  isValid(): boolean {
    return this.value >= 0 && this.value <= 1;
  }
}

export interface PatternData {
  [key: string]: any;
}

export interface Pattern {
  readonly id: PatternId;
  readonly name: PatternName;
  readonly category: PatternCategory;
  readonly frequency: PatternFrequency;
  readonly accuracy: PatternAccuracy;
  readonly pattern: PatternData;
  readonly lastSeen: Date;
  readonly createdAt: Date;
}

export interface PatternRequest {
  readonly name: PatternName;
  readonly category: PatternCategory;
  readonly pattern: PatternData;
}

export class PatternEntity implements Pattern {
  constructor(
    public readonly id: PatternId,
    public readonly name: PatternName,
    public readonly category: PatternCategory,
    public readonly frequency: PatternFrequency,
    public readonly accuracy: PatternAccuracy,
    public readonly pattern: PatternData,
    public readonly lastSeen: Date,
    public readonly createdAt: Date
  ) {}

  static create(request: PatternRequest): PatternEntity {
    return new PatternEntity(
      { value: 0 }, // Will be set by repository
      request.name,
      request.category,
      new PatternFrequencyValue(1),
      new PatternAccuracyValue(0.5),
      request.pattern,
      new Date(),
      new Date()
    );
  }

  isHighConfidence(threshold: number = 0.8): boolean {
    return this.accuracy.value >= threshold;
  }

  isFrequent(threshold: number = 10): boolean {
    return this.frequency.value >= threshold;
  }

  incrementFrequency(): PatternEntity {
    return new PatternEntity(
      this.id,
      this.name,
      this.category,
      new PatternFrequencyValue(this.frequency.value + 1),
      this.accuracy,
      this.pattern,
      new Date(),
      this.createdAt
    );
  }

  updateAccuracy(newAccuracy: PatternAccuracy): PatternEntity {
    return new PatternEntity(
      this.id,
      this.name,
      this.category,
      this.frequency,
      newAccuracy,
      this.pattern,
      new Date(),
      this.createdAt
    );
  }

  updateLastSeen(): PatternEntity {
    return new PatternEntity(
      this.id,
      this.name,
      this.category,
      this.frequency,
      this.accuracy,
      this.pattern,
      new Date(),
      this.createdAt
    );
  }

  getPatternValue<T>(key: string): T | undefined {
    return this.pattern[key] as T;
  }

  isSimilarTo(other: Pattern, threshold: number = 0.7): boolean {
    // Simple similarity check based on category and name
    if (this.category !== other.category) return false;
    
    const nameA = this.name.value.toLowerCase();
    const nameB = other.name.value.toLowerCase();
    
    // Simple string similarity (Jaccard similarity)
    const setA = new Set(nameA.split(''));
    const setB = new Set(nameB.split(''));
    const intersection = new Set([...setA].filter(x => setB.has(x)));
    const union = new Set([...setA, ...setB]);
    
    return intersection.size / union.size >= threshold;
  }
}