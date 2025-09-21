# 🔗 データ関係性詳細マップ

betaTaskerプロジェクトにおけるseedデータの詳細な関係性と依存関係を説明します。

## 📊 全体データ構造図

```
基礎データ層
├── RobotSpecification (ロボット仕様 - 21種類)
├── OptimizationModel (最適化モデル - 21種類)
├── PhenomenologicalFramework (現象学的フレームワーク - 21種類)
└── QuantificationLabel (定量化ラベル)

ユーザー・タスク層
├── User (ユーザー)
├── Task (タスク)
└── MemoryContext (メモリコンテキスト L1-L5)

評価・分析層
├── StateEvaluation (状態評価 - 5サンプル)
├── ToolMatchingResult (ツールマッチング結果 - 3サンプル)
├── ProcessMonitoring (プロセス監視 - 2サンプル)
└── LearningPattern (学習パターン - 3+CSV)

ヒューリスティクス層
├── HeuristicsAnalysis (ヒューリスティクス分析)
├── HeuristicsTracking (行動追跡)
└── HeuristicsInsight (インサイト)
```

## 🏗️ テーブル間関係詳細

### 1. StateEvaluation（状態評価）の中心的役割

```sql
-- StateEvaluationテーブルの外部キー関係
CREATE TABLE state_evaluations (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,          -- Users.id への参照
    task_id INT NOT NULL,                   -- Tasks.id への参照  
    framework VARCHAR(255),                 -- phenomenological_frameworks.id への参照
    -- ... その他のカラム
);

-- 関連テーブル
CREATE TABLE tool_matching_results (
    state_evaluation_id VARCHAR(255),      -- state_evaluations.id への外部キー
    robot_id VARCHAR(255),                 -- robot_specifications.id への参照
    optimization_model_id VARCHAR(255),    -- optimization_models.id への参照
    -- ...
);

CREATE TABLE process_monitoring (
    state_evaluation_id VARCHAR(255),      -- state_evaluations.id への外部キー
    -- ...
);
```

### 2. 参照整合性マップ

| 参照元テーブル | 外部キー | 参照先テーブル | 関係性 |
|----------------|----------|----------------|--------|
| StateEvaluation | user_id | User | N:1 |
| StateEvaluation | task_id | Task | N:1 |
| StateEvaluation | framework | PhenomenologicalFramework | N:1 |
| ToolMatchingResult | state_evaluation_id | StateEvaluation | N:1 |
| ToolMatchingResult | robot_id | RobotSpecification | N:1 |
| ToolMatchingResult | optimization_model_id | OptimizationModel | N:1 |
| ProcessMonitoring | state_evaluation_id | StateEvaluation | N:1 |
| LearningPattern | user_id | User | N:1 |

## 📋 Seedデータの依存順序

### Phase 1: 基礎データ（依存関係なし）
```go
// 1. ユーザー・基本データ
SeedUsers()                    // ユーザーデータ
SeedTasks()                    // タスクデータ
SeedBooks()                    // 書籍データ

// 2. 機器・モデルデータ  
SeedRobotSpecifications()      // ロボット仕様（21種類）
SeedOptimizationModels()       // 最適化モデル（21種類）
SeedPhenomenologicalFrameworks() // 現象学的フレームワーク（21種類）
SeedQuantificationLabels()     // 定量化ラベル

// 3. メモリ・知識データ
SeedMemoryContexts()           // メモリコンテキスト（L1-L5）
SeedHeuristics()              // ヒューリスティクス分析
```

### Phase 2: 評価データ（Phase 1に依存）
```go
// 4. 状態評価データ（Users, Tasks, PhenomenologicalFrameworksを参照）
SeedStateEvaluations()         // 状態評価（5サンプル）
```

### Phase 3: 結果データ（Phase 2に依存）
```go
// 5. マッチング・監視データ（StateEvaluationsを参照）
SeedToolMatchingResults()      // ツールマッチング結果（3サンプル）
SeedProcessMonitoring()        // プロセス監視（2サンプル）
```

### Phase 4: 学習データ（Phase 1に依存）
```go
// 6. 学習パターンデータ（Usersを参照）
SeedLearningPatterns()         // 学習パターン（3+CSV）
```

## 🎯 具体的なデータ関連例

### 1. 状態評価からツールマッチングへの流れ

```yaml
# Step 1: 状態評価データ
StateEvaluation:
  id: "eval_001"
  user_id: "user_001" 
  task_id: 1
  level: 2
  work_target: "[MA-Q-02] 材料硬度変動への対応"
  evaluation_score: 73.8
  framework: "force_control_framework"

# Step 2: ツールマッチング実行（eval_001を参照）
ToolMatchingResult:
  id: "match_001"
  state_evaluation_id: "eval_001"        # ← StateEvaluationへの参照
  robot_id: "collaborative_robot_v2"     # ← RobotSpecificationへの参照  
  optimization_model_id: "energy_optimization" # ← OptimizationModelへの参照
  matching_score: 0.92
```

### 2. プロセス監視データの生成

```yaml
# 状態評価に基づくプロセス監視開始
ProcessMonitoring:
  id: "monitor_001"
  state_evaluation_id: "eval_001"        # ← StateEvaluationへの参照
  process_type: "robot_assembly"
  status: "running"
  monitoring_data:
    force_x: 5.2
    force_y: 3.1
    cycle_time: 27.3
    success_rate: 0.96
```

### 3. 学習パターンの知識蓄積

```yaml
# ユーザーの学習パターン記録
LearningPattern:
  id: "pattern_001"
  user_id: "user_001"                    # ← Userへの参照
  pattern_type: "assembly_skill_pattern"
  domain: "robot_assembly"
  tacit_knowledge: "熟練工の『しっくりくる』感覚"
  explicit_form: "力覚センサ値: Fx<0.5N Fy<0.5N Tz<0.1Nm"
  seci_stage: "共同化→表出化→連結化→内面化"
```

## 🎮 インタラクティブなデータ探索

### SQLクエリ例

```sql
-- 1. レベル別評価スコア分析
SELECT 
    level,
    AVG(evaluation_score) as avg_score,
    MIN(evaluation_score) as min_score,
    MAX(evaluation_score) as max_score,
    COUNT(*) as count
FROM state_evaluations 
GROUP BY level 
ORDER BY level;

-- 2. ロボット-最適化モデルの組み合わせ分析  
SELECT 
    rs.model_name as robot,
    om.name as optimization_model,
    AVG(tmr.matching_score) as avg_matching_score,
    COUNT(*) as usage_count
FROM tool_matching_results tmr
JOIN robot_specifications rs ON tmr.robot_id = rs.id
JOIN optimization_models om ON tmr.optimization_model_id = om.id
GROUP BY rs.model_name, om.name
ORDER BY avg_matching_score DESC;

-- 3. 学習パターンの進捗分析
SELECT 
    domain,
    seci_stage,
    AVG(accuracy) as avg_accuracy,
    AVG(consistency) as avg_consistency,
    COUNT(*) as pattern_count
FROM learning_patterns
WHERE validated = true
GROUP BY domain, seci_stage
ORDER BY domain, avg_accuracy DESC;

-- 4. プロセス監視の異常検知状況
SELECT 
    process_type,
    status,
    COUNT(*) as count,
    AVG(JSON_EXTRACT(metrics, '$.overall')) as avg_overall_score
FROM process_monitoring
GROUP BY process_type, status;
```

## 🔧 カスタマイズ・拡張ポイント

### 1. 新しい業界ドメインの追加

```go
// 新しいドメイン：医療機器製造
type MedicalDeviceSpecification struct {
    ID                string
    DeviceType        string  // "surgical_robot", "prosthetic", etc.
    SterilizationMethod string
    FDAApprovalLevel  string
    BiocompatibilityRating string
}

// 関連するseed関数
func SeedMedicalDevices(db *gorm.DB) error {
    devices := []MedicalDeviceSpecification{
        // ... 医療機器データ
    }
    // ... 実装
}
```

### 2. リアルタイムデータ統合

```go
// 実機データとの同期
type RealTimeDataSync struct {
    StateEvaluationID string
    MachineID        string
    SensorData       JSON
    Timestamp        time.Time
    SyncStatus       string
}

func SyncRealTimeData() {
    // 実機からのデータ取得・同期ロジック
}
```

このデータ関係性マップにより、betaTaskerのseedデータの全体像と各データの役割・依存関係が明確になり、効率的な開発とカスタマイズが可能になります。