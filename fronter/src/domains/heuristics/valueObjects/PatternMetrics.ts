export interface PatternMetricsData {
  frequency: number;
  accuracy: number;
  confidence: number;
  coverage: number;
  stability: number;
}

export class PatternMetrics {
  private constructor(
    private readonly frequency: number,
    private readonly accuracy: number,
    private readonly confidence: number,
    private readonly coverage: number,
    private readonly stability: number
  ) {
    this.validate();
  }

  static create(data: PatternMetricsData): PatternMetrics {
    return new PatternMetrics(
      data.frequency,
      data.accuracy,
      data.confidence,
      data.coverage,
      data.stability
    );
  }

  static empty(): PatternMetrics {
    return new PatternMetrics(0, 0, 0, 0, 0);
  }

  private validate(): void {
    if (this.frequency < 0) {
      throw new Error('Frequency must be non-negative');
    }
    if (this.accuracy < 0 || this.accuracy > 1) {
      throw new Error('Accuracy must be between 0 and 1');
    }
    if (this.confidence < 0 || this.confidence > 1) {
      throw new Error('Confidence must be between 0 and 1');
    }
    if (this.coverage < 0 || this.coverage > 1) {
      throw new Error('Coverage must be between 0 and 1');
    }
    if (this.stability < 0 || this.stability > 1) {
      throw new Error('Stability must be between 0 and 1');
    }
  }

  getFrequency(): number {
    return this.frequency;
  }

  getAccuracy(): number {
    return this.accuracy;
  }

  getConfidence(): number {
    return this.confidence;
  }

  getCoverage(): number {
    return this.coverage;
  }

  getStability(): number {
    return this.stability;
  }

  getOverallScore(): number {
    // Weighted average of all metrics
    const weights = {
      accuracy: 0.3,
      confidence: 0.25,
      stability: 0.2,
      coverage: 0.15,
      frequency: 0.1
    };

    const normalizedFrequency = Math.min(this.frequency / 100, 1); // Normalize to 0-1

    return (
      this.accuracy * weights.accuracy +
      this.confidence * weights.confidence +
      this.stability * weights.stability +
      this.coverage * weights.coverage +
      normalizedFrequency * weights.frequency
    );
  }

  isHighQuality(threshold: number = 0.7): boolean {
    return this.getOverallScore() >= threshold;
  }

  isStable(threshold: number = 0.6): boolean {
    return this.stability >= threshold;
  }

  isReliable(): boolean {
    return this.accuracy >= 0.7 && this.confidence >= 0.6 && this.stability >= 0.5;
  }

  updateFrequency(newFrequency: number): PatternMetrics {
    return new PatternMetrics(
      newFrequency,
      this.accuracy,
      this.confidence,
      this.coverage,
      this.stability
    );
  }

  updateAccuracy(newAccuracy: number): PatternMetrics {
    return new PatternMetrics(
      this.frequency,
      newAccuracy,
      this.confidence,
      this.coverage,
      this.stability
    );
  }

  updateConfidence(newConfidence: number): PatternMetrics {
    return new PatternMetrics(
      this.frequency,
      this.accuracy,
      newConfidence,
      this.coverage,
      this.stability
    );
  }

  updateCoverage(newCoverage: number): PatternMetrics {
    return new PatternMetrics(
      this.frequency,
      this.accuracy,
      this.confidence,
      newCoverage,
      this.stability
    );
  }

  updateStability(newStability: number): PatternMetrics {
    return new PatternMetrics(
      this.frequency,
      this.accuracy,
      this.confidence,
      this.coverage,
      newStability
    );
  }

  combine(other: PatternMetrics): PatternMetrics {
    // Combine metrics using weighted averages
    const totalFreq = this.frequency + other.frequency;
    const weight1 = totalFreq > 0 ? this.frequency / totalFreq : 0.5;
    const weight2 = totalFreq > 0 ? other.frequency / totalFreq : 0.5;

    return new PatternMetrics(
      totalFreq,
      this.accuracy * weight1 + other.accuracy * weight2,
      this.confidence * weight1 + other.confidence * weight2,
      this.coverage * weight1 + other.coverage * weight2,
      this.stability * weight1 + other.stability * weight2
    );
  }

  toDisplayString(): string {
    return `Score: ${(this.getOverallScore() * 100).toFixed(1)}% | ` +
           `Frequency: ${this.frequency} | ` +
           `Accuracy: ${(this.accuracy * 100).toFixed(1)}% | ` +
           `Confidence: ${(this.confidence * 100).toFixed(1)}%`;
  }

  toJSON(): PatternMetricsData {
    return {
      frequency: this.frequency,
      accuracy: this.accuracy,
      confidence: this.confidence,
      coverage: this.coverage,
      stability: this.stability
    };
  }

  equals(other: PatternMetrics): boolean {
    const epsilon = 0.001; // Small tolerance for floating point comparison
    return Math.abs(this.frequency - other.frequency) < epsilon &&
           Math.abs(this.accuracy - other.accuracy) < epsilon &&
           Math.abs(this.confidence - other.confidence) < epsilon &&
           Math.abs(this.coverage - other.coverage) < epsilon &&
           Math.abs(this.stability - other.stability) < epsilon;
  }

  getQualityLevel(): 'excellent' | 'good' | 'fair' | 'poor' {
    const score = this.getOverallScore();
    if (score >= 0.8) return 'excellent';
    if (score >= 0.6) return 'good';
    if (score >= 0.4) return 'fair';
    return 'poor';
  }

  getRecommendations(): string[] {
    const recommendations: string[] = [];

    if (this.accuracy < 0.7) {
      recommendations.push('精度を向上させるために、より多くのトレーニングデータが必要です');
    }
    if (this.confidence < 0.6) {
      recommendations.push('信頼度を上げるために、パターンの検証を強化してください');
    }
    if (this.stability < 0.5) {
      recommendations.push('安定性を向上させるために、時系列での一貫性を確認してください');
    }
    if (this.coverage < 0.4) {
      recommendations.push('カバレッジを拡大するために、より幅広いデータセットでテストしてください');
    }
    if (this.frequency < 5) {
      recommendations.push('頻度が低いため、パターンの妥当性を再検証してください');
    }

    if (recommendations.length === 0) {
      recommendations.push('パターンの品質は良好です。継続的な監視を推奨します');
    }

    return recommendations;
  }
}