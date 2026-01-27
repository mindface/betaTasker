export interface InsightId {
  value: number;
}

export interface InsightTitle {
  value: string;
  isValid(): boolean;
}

export class InsightTitleValue implements InsightTitle {
  constructor(public readonly value: string) {
    if (!this.isValid()) {
      throw new Error('Insight title cannot be empty');
    }
  }

  isValid(): boolean {
    return this.value.trim().length > 0;
  }
}

export type InsightType = 'recommendation' | 'warning' | 'optimization' | 'pattern' | 'anomaly';

export interface InsightDescription {
  value: string;
  isValid(): boolean;
}

export class InsightDescriptionValue implements InsightDescription {
  constructor(public readonly value: string) {
    if (!this.isValid()) {
      throw new Error('Insight description cannot be empty');
    }
  }

  isValid(): boolean {
    return this.value.trim().length > 0;
  }
}

export interface InsightConfidence {
  value: number;
  isValid(): boolean;
}

export class InsightConfidenceValue implements InsightConfidence {
  constructor(public readonly value: number) {
    if (!this.isValid()) {
      throw new Error('Insight confidence must be between 0 and 1');
    }
  }

  isValid(): boolean {
    return this.value >= 0 && this.value <= 1;
  }
}

export type InsightImpact = 'low' | 'medium' | 'high' | 'critical';

export interface InsightSuggestion {
  value: string;
  priority: number;
}

export interface Insight {
  readonly id: InsightId;
  readonly title: InsightTitle;
  readonly type: InsightType;
  readonly description: InsightDescription;
  readonly confidence: InsightConfidence;
  readonly impact: InsightImpact;
  readonly suggestions: InsightSuggestion[];
  readonly isActive: boolean;
  readonly createdAt: Date;
  readonly updatedAt: Date;
}

export interface InsightRequest {
  readonly title: InsightTitle;
  readonly type: InsightType;
  readonly description: InsightDescription;
  readonly confidence: InsightConfidence;
  readonly impact: InsightImpact;
  readonly suggestions?: InsightSuggestion[];
}

export class InsightEntity implements Insight {
  constructor(
    public readonly id: InsightId,
    public readonly title: InsightTitle,
    public readonly type: InsightType,
    public readonly description: InsightDescription,
    public readonly confidence: InsightConfidence,
    public readonly impact: InsightImpact,
    public readonly suggestions: InsightSuggestion[],
    public readonly isActive: boolean,
    public readonly createdAt: Date,
    public readonly updatedAt: Date
  ) {}

  static create(request: InsightRequest): InsightEntity {
    return new InsightEntity(
      { value: 0 }, // Will be set by repository
      request.title,
      request.type,
      request.description,
      request.confidence,
      request.impact,
      request.suggestions || [],
      true,
      new Date(),
      new Date()
    );
  }

  isHighConfidence(threshold: number = 0.8): boolean {
    return this.confidence.value >= threshold;
  }

  isCritical(): boolean {
    return this.impact === 'critical';
  }

  isHigh(): boolean {
    return this.impact === 'high';
  }

  hasSuggestions(): boolean {
    return this.suggestions.length > 0;
  }

  getPrioritizedSuggestions(): InsightSuggestion[] {
    return [...this.suggestions].sort((a, b) => b.priority - a.priority);
  }

  deactivate(): InsightEntity {
    return new InsightEntity(
      this.id,
      this.title,
      this.type,
      this.description,
      this.confidence,
      this.impact,
      this.suggestions,
      false,
      this.createdAt,
      new Date()
    );
  }

  activate(): InsightEntity {
    return new InsightEntity(
      this.id,
      this.title,
      this.type,
      this.description,
      this.confidence,
      this.impact,
      this.suggestions,
      true,
      this.createdAt,
      new Date()
    );
  }

  updateConfidence(newConfidence: InsightConfidence): InsightEntity {
    return new InsightEntity(
      this.id,
      this.title,
      this.type,
      this.description,
      newConfidence,
      this.impact,
      this.suggestions,
      this.isActive,
      this.createdAt,
      new Date()
    );
  }

  addSuggestion(suggestion: InsightSuggestion): InsightEntity {
    return new InsightEntity(
      this.id,
      this.title,
      this.type,
      this.description,
      this.confidence,
      this.impact,
      [...this.suggestions, suggestion],
      this.isActive,
      this.createdAt,
      new Date()
    );
  }

  removeSuggestion(suggestionValue: string): InsightEntity {
    const filteredSuggestions = this.suggestions.filter(s => s.value !== suggestionValue);
    return new InsightEntity(
      this.id,
      this.title,
      this.type,
      this.description,
      this.confidence,
      this.impact,
      filteredSuggestions,
      this.isActive,
      this.createdAt,
      new Date()
    );
  }

  getImpactScore(): number {
    switch (this.impact) {
      case 'critical': return 4;
      case 'high': return 3;
      case 'medium': return 2;
      case 'low': return 1;
      default: return 0;
    }
  }

  getOverallPriority(): number {
    return this.getImpactScore() * this.confidence.value;
  }
}