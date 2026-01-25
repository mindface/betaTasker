# 🚀 Seed Data 使用例ガイド

betaTaskerのseedデータを実際に使用する具体的な例を示します。

## 📋 目次

1. [基本的な使用方法](#基本的な使用方法)
2. [開発シナリオ別例](#開発シナリオ別例)
3. [API統合例](#api統合例)
4. [データ分析例](#データ分析例)
5. [カスタマイズ例](#カスタマイズ例)

## 基本的な使用方法

### 1. Seed実行コマンド

```bash
# プロジェクトルートで実行

# 初回セットアップ（全データをクリアしてからseed）
go run main.go clean-seed

# 追加データのみseed（既存データ保持）
go run main.go seed

# 通常のアプリケーション起動
go run main.go
```

### 2. 実行結果例

```bash
$ go run main.go clean-seed

2024/01/15 10:00:00 Cleaning database tables...
2024/01/15 10:00:01 ✓ Database cleaned
2024/01/15 10:00:01 Starting database seeding...
2024/01/15 10:00:01 Seeding memory contexts...
2024/01/15 10:00:01 ✓ Memory contexts seeded successfully
2024/01/15 10:00:01 Seeding books and tasks...
2024/01/15 10:00:02 ✓ Books and tasks seeded successfully
2024/01/15 10:00:02 Seeding heuristics data...
2024/01/15 10:00:02 ✓ Heuristics data seeded successfully
2024/01/15 10:00:02 Seeding phenomenological framework data...
2024/01/15 10:00:02 ✓ Phenomenological framework data seeded successfully
2024/01/15 10:00:02 Seeding data from CSV files...
2024/01/15 10:00:03 ✓ Successfully seeded 21 robot specifications
2024/01/15 10:00:03 ✓ Successfully seeded 21 optimization models
2024/01/15 10:00:03 ✓ Successfully seeded 21 phenomenological frameworks
2024/01/15 10:00:03 ✓ Successfully seeded 12 quantification labels
2024/01/15 10:00:03 ✓ CSV data seeded successfully
2024/01/15 10:00:03 Seeding state evaluation data...
2024/01/15 10:00:03 ✓ Successfully seeded 5 state evaluations
2024/01/15 10:00:03 ✓ State evaluation data seeded successfully
2024/01/15 10:00:03 Seeding tool matching results...
2024/01/15 10:00:03 ✓ Successfully seeded 3 tool matching results
2024/01/15 10:00:03 ✓ Tool matching results seeded successfully
2024/01/15 10:00:03 Seeding process monitoring data...
2024/01/15 10:00:03 ✓ Successfully seeded 2 process monitoring records
2024/01/15 10:00:03 ✓ Process monitoring data seeded successfully
2024/01/15 10:00:04 Seeding learning patterns...
2024/01/15 10:00:04 ✓ Successfully seeded 21 learning patterns
2024/01/15 10:00:04 ✓ Learning patterns seeded successfully
2024/01/15 10:00:04 Database seeding completed successfully!
```

## 開発シナリオ別例

### シナリオ 1: 新人開発者の環境構築

```bash
# 1. プロジェクトをクローン
git clone https://github.com/yourorg/betaTasker.git
cd betaTasker/backer/godotask

# 2. 環境変数設定
cp .env.example .env
# DATABASE_DSNを適切に設定

# 3. 依存関係インストール
go mod tidy

# 4. データベースセットアップ
go run main.go clean-seed

# 5. アプリケーション起動
go run main.go

# これで以下のデータが利用可能：
# - 21種類のロボット仕様
# - 21種類の最適化モデル
# - 5つの状態評価サンプル
# - 3つのツールマッチング結果
# - リアルな製造現場データ
```

### シナリオ 2: 機能テスト用データ準備

```go
// test_data_setup.go
package main

import (
    "log"
    "github.com/godotask/infrastructure/db/model"
    "github.com/godotask/seed"
)

func setupTestData() {
    // テスト用データベース初期化
    model.InitDB()
    
    // 必要な基礎データのみseed
    if err := seed.SeedRobotSpecifications(model.DB); err != nil {
        log.Fatalf("Robot seeding failed: %v", err)
    }
    
    if err := seed.SeedStateEvaluations(model.DB); err != nil {
        log.Fatalf("State evaluation seeding failed: %v", err)
    }
    
    log.Println("Test data setup completed")
}

func TestToolMatching(t *testing.T) {
    setupTestData()
    
    // テスト実行
    // ...
}
```

### シナリオ 3: デモ環境のセットアップ

```bash
# デモ用の特別なデータセットを使用
go run main.go clean-seed

# デモ用の追加データを投入
curl -X POST http://localhost:8080/api/state-evaluations \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "demo_user",
    "task_id": 999,
    "level": 3,
    "work_target": "[デモ] 高精度組立作業",
    "current_state": {
      "accuracy": 0.85,
      "efficiency": 0.80,
      "consistency": 0.88,
      "innovation": 0.75
    },
    "target_state": {
      "accuracy": 0.92,
      "efficiency": 0.87,
      "consistency": 0.93,
      "innovation": 0.82
    }
  }'
```

## API統合例

### 1. 状態評価システムの利用

#### 評価データの作成
```bash
curl -X POST http://localhost:8080/api/state-evaluations \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user_001",
    "task_id": 1,
    "level": 2,
    "work_target": "[MA-Q-02] 材料硬度変動への対応",
    "current_state": {
      "accuracy": 0.82,
      "efficiency": 0.71,
      "consistency": 0.79,
      "innovation": 0.63
    },
    "target_state": {
      "accuracy": 0.90,
      "efficiency": 0.80,
      "consistency": 0.85,
      "innovation": 0.70
    }
  }'
```

#### レスポンス例
```json
{
  "status": "success",
  "data": {
    "id": "eval_12345",
    "user_id": "user_001",
    "task_id": 1,
    "level": 2,
    "work_target": "[MA-Q-02] 材料硬度変動への対応",
    "evaluation_score": 73.8,
    "framework": "force_control_framework",
    "status": "completed",
    "created_at": "2024-01-15T10:30:00Z"
  }
}
```

### 2. ツールマッチングの実行

```bash
curl -X POST http://localhost:8080/api/tool-matching \
  -H "Content-Type: application/json" \
  -d '{
    "state_evaluation_id": "eval_12345",
    "requirements": {
      "payload": 5.0,
      "reach": 800.0,
      "precision": 0.02,
      "speed": 1000.0
    },
    "constraints": {
      "budget": 500000,
      "space": "limited"
    }
  }'
```

#### レスポンス例
```json
{
  "status": "success",
  "data": {
    "id": "match_67890",
    "state_evaluation_id": "eval_12345",
    "robot_id": "teaching_free_arm_v1",
    "optimization_model_id": "trajectory_optimization",
    "matching_score": 0.87,
    "recommendations": {
      "robot": {
        "model": "TF-ARM-001",
        "recommended_use": "高精度作業に最適、AI学習機能搭載"
      },
      "optimization": {
        "model_name": "軌道最適化",
        "expected_improvement": "25%の改善が期待できます"
      }
    }
  }
}
```

### 3. プロセス監視の開始

```bash
curl -X POST http://localhost:8080/api/process-monitoring/start \
  -H "Content-Type: application/json" \
  -d '{
    "state_evaluation_id": "eval_12345",
    "process_type": "robot_assembly",
    "initial_data": {
      "target_cycle_time": 25.0,
      "quality_threshold": 0.95
    }
  }'
```

### 4. WebSocket監視の接続

```javascript
// JavaScript例
const ws = new WebSocket('ws://localhost:8080/api/process-monitoring/monitor_123/ws');

ws.onmessage = function(event) {
  const data = JSON.parse(event.data);
  console.log('Monitoring data:', data);
  
  // リアルタイムデータの表示
  updateDashboard({
    timestamp: data.timestamp,
    metrics: data.metrics,
    anomalies: data.anomalies,
    performance: data.performance
  });
};

function updateDashboard(data) {
  // ダッシュボードの更新ロジック
  document.getElementById('force-x').textContent = data.metrics.force_x;
  document.getElementById('efficiency').textContent = 
    `${(data.performance.efficiency * 100).toFixed(1)}%`;
  
  // 異常の表示
  if (data.anomalies && data.anomalies.length > 0) {
    showAlert(data.anomalies);
  }
}
```

## データ分析例

### 1. PostgreSQLでの分析クエリ

```sql
-- レベル別技能向上分析
WITH level_progression AS (
  SELECT 
    user_id,
    level,
    AVG(evaluation_score) as avg_score,
    COUNT(*) as attempts,
    MAX(created_at) as latest_attempt
  FROM state_evaluations 
  WHERE status = 'completed'
  GROUP BY user_id, level
)
SELECT 
  user_id,
  level,
  avg_score,
  attempts,
  LEAD(avg_score) OVER (PARTITION BY user_id ORDER BY level) - avg_score as improvement
FROM level_progression
ORDER BY user_id, level;

-- 最適なロボット・モデル組み合わせ分析
SELECT 
  rs.model_name as robot,
  om.name as optimization_model,
  AVG(tmr.matching_score) as avg_score,
  COUNT(*) as usage_count,
  AVG(CAST(JSON_EXTRACT(tmr.expected_performance, '$.predicted_score') AS DECIMAL)) as predicted_performance
FROM tool_matching_results tmr
JOIN robot_specifications rs ON tmr.robot_id = rs.id
JOIN optimization_models om ON tmr.optimization_model_id = om.id
GROUP BY rs.model_name, om.name
HAVING usage_count >= 2
ORDER BY avg_score DESC;

-- 異常検知パターン分析
SELECT 
  process_type,
  JSON_EXTRACT(anomalies, '$[0].type') as anomaly_type,
  COUNT(*) as occurrence_count,
  AVG(JSON_EXTRACT(anomalies, '$[0].value')) as avg_value,
  AVG(JSON_EXTRACT(anomalies, '$[0].threshold')) as avg_threshold
FROM process_monitoring 
WHERE JSON_LENGTH(anomalies) > 0
GROUP BY process_type, JSON_EXTRACT(anomalies, '$[0].type')
ORDER BY occurrence_count DESC;
```

### 2. Go言語での分析コード

```go
package analysis

import (
    "github.com/godotask/infrastructure/db/model"
    "gorm.io/gorm"
)

type AnalysisService struct {
    db *gorm.DB
}

// レベル別評価スコア分析
func (s *AnalysisService) AnalyzeLevelProgression(userID string) (map[int]float64, error) {
    var results []struct {
        Level int
        AvgScore float64
    }

    err := s.db.Model(&model.StateEvaluation{}).
        Select("level, AVG(evaluation_score) as avg_score").
        Where("user_id = ? AND status = 'completed'", userID).
        Group("level").
        Order("level").
        Scan(&results).Error

    if err != nil {
        return nil, err
    }

    progression := make(map[int]float64)
    for _, result := range results {
        progression[result.Level] = result.AvgScore
    }

    return progression, nil
}

// ロボット性能分析
func (s *AnalysisService) AnalyzeRobotPerformance() ([]RobotAnalysis, error) {
    var results []RobotAnalysis

    err := s.db.Table("tool_matching_results tmr").
        Select(`rs.model_name as robot_name, 
                AVG(tmr.matching_score) as avg_matching_score,
                COUNT(*) as usage_count`).
        Joins("JOIN robot_specifications rs ON tmr.robot_id = rs.id").
        Group("rs.model_name").
        Having("COUNT(*) >= 2").
        Order("avg_matching_score DESC").
        Scan(&results).Error

    return results, err
}

type RobotAnalysis struct {
    RobotName        string  `json:"robot_name"`
    AvgMatchingScore float64 `json:"avg_matching_score"`
    UsageCount       int     `json:"usage_count"`
}
```

## カスタマイズ例

### 1. 新しい業界向けカスタマイズ

```go
// automotive_seed.go - 自動車業界向け
func SeedAutomotiveData(db *gorm.DB) error {
    automotiveRobots := []model.RobotSpecification{
        {
            ID:              "automotive_welder_v1",
            ModelName:       "自動車溶接ロボット",
            DOF:            6,
            ReachMm:        2500.0,
            PayloadKg:      50.0,
            RepeatAccuracyMm: 0.1,
            MaxSpeedMmS:    1500.0,
            WorkEnvelopeShape: "rectangular",
            TeachingMethod: "offline_programming",
            ControlType:    "position",
            VisionSystem:   &model.NullString{String: "laser_tracker", Valid: true},
            SafetyFeatures: &model.NullString{String: "automotive_safety_standard", Valid: true},
        },
        // ... 他の自動車業界特化ロボット
    }

    for _, robot := range automotiveRobots {
        if err := db.Create(&robot).Error; err != nil {
            return err
        }
    }

    return nil
}

// 使用方法
func main() {
    model.InitDB()
    
    // 基本seedデータ
    if err := seed.RunAllSeeds(); err != nil {
        log.Fatal(err)
    }
    
    // 自動車業界特化データ
    if err := SeedAutomotiveData(model.DB); err != nil {
        log.Fatal(err)
    }
}
```

### 2. 環境別設定

```go
// config/seed_config.go
type SeedConfig struct {
    Environment string
    DataSets    []string
}

func GetSeedConfig() SeedConfig {
    env := os.Getenv("APP_ENV")
    
    switch env {
    case "development":
        return SeedConfig{
            Environment: "development",
            DataSets:    []string{"basic", "sample", "test"},
        }
    case "staging":
        return SeedConfig{
            Environment: "staging",
            DataSets:    []string{"basic", "demo"},
        }
    case "production":
        return SeedConfig{
            Environment: "production",
            DataSets:    []string{"basic"},
        }
    default:
        return SeedConfig{
            Environment: "development",
            DataSets:    []string{"basic", "sample"},
        }
    }
}

// 環境別seed実行
func RunEnvironmentSeeds() error {
    config := GetSeedConfig()
    
    for _, dataSet := range config.DataSets {
        switch dataSet {
        case "basic":
            if err := seed.SeedBasicData(model.DB); err != nil {
                return err
            }
        case "sample":
            if err := seed.SeedStateEvaluations(model.DB); err != nil {
                return err
            }
        case "demo":
            if err := seed.SeedDemoData(model.DB); err != nil {
                return err
            }
        }
    }
    
    return nil
}
```

### 3. 段階的データ投入

```bash
# 段階1: 基礎データのみ
go run main.go seed --phase=basic

# 段階2: 評価データ追加
go run main.go seed --phase=evaluation

# 段階3: 監視データ追加
go run main.go seed --phase=monitoring

# 全段階実行
go run main.go seed --phase=all
```

これらの例を参考に、betaTaskerのseedデータを効果的に活用し、開発・テスト・本番環境で適切なデータセットを構築できます。