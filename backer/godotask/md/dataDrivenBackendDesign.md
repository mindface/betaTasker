# Backend Seed Data から見たデータ駆動設計案

## 前提

現状の seed CSV は `backer/godotask/seed/data` 配下に配置されている。

この資料では、ドメイン駆動設計ではなく、データを中心に backend を設計する場合の構成案をまとめる。

データ駆動設計では、業務概念の境界よりも以下を重視する。

- どのデータを収集するか
- どのデータを正とするか
- どのデータを加工・評価・学習に使うか
- どのメトリクスで改善するか
- どのデータ品質を担保するか

## データ駆動設計の基本方針

この backend では、`task_id`, `user_id`, `context_id` を中心にデータをつなぎ、タスク実行から評価、知識化、最適化、ヒューリスティクス改善までを回す構成が自然。

```txt
入力データ
  ↓
タスク・作業文脈
  ↓
行動ログ・評価
  ↓
分析・知識化
  ↓
最適化・推薦
  ↓
次のタスク改善
```

## 推奨ディレクトリ構成

```txt
backer/godotask/
├── data/
│   ├── catalog/
│   │   ├── dataset_catalog.go
│   │   └── schema_registry.go
│   ├── quality/
│   │   ├── validator.go
│   │   ├── completeness.go
│   │   ├── consistency.go
│   │   └── confidence.go
│   ├── pipeline/
│   │   ├── ingest.go
│   │   ├── normalize.go
│   │   ├── enrich.go
│   │   ├── aggregate.go
│   │   └── export.go
│   ├── metrics/
│   │   ├── task_metrics.go
│   │   ├── assessment_metrics.go
│   │   ├── knowledge_metrics.go
│   │   └── heuristics_metrics.go
│   └── feature/
│       ├── task_features.go
│       ├── context_features.go
│       ├── robot_features.go
│       └── heuristic_features.go
├── seed/
│   ├── data/
│   └── importer/
├── infrastructure/
│   └── db/
│       ├── model/
│       └── repository/
├── usecase/
│   ├── data_pipeline_usecase.go
│   ├── data_quality_usecase.go
│   └── analytics_usecase.go
└── interface/
    └── controller/
        ├── data_quality/
        ├── analytics/
        └── dashboard/
```

DDD のように業務境界ごとに分けるより、データの流れ、品質、分析、特徴量、メトリクスを中心に構成する。

## データ分類

### 1. マスタデータ

変更頻度が低く、他データから参照されるデータ。

対象 CSV:

- `robot_specifications.csv`
- `optimization_models.csv`
- `phenomenological_frameworks.csv`

役割:

- ロボット仕様の参照
- 最適化モデルの参照
- 現象学的フレームワークの参照
- 分析や推薦時の候補データ

管理方針:

- ID の一意性を保証する
- 欠損値を極力許容しない
- version や status を持つデータは有効状態を明確にする
- seed から投入された値と運用中に追加された値を区別できるようにする

### 2. トランザクションデータ

ユーザー操作やタスク実行によって発生するデータ。

対象 CSV:

- `task.csv`
- `book.csv`
- `memories.csv`
- `memory_contexts.csv`
- `assessments.csv`
- `qualitative_label.csv`

役割:

- ユーザーの作業単位を記録する
- 作業文脈を保存する
- 実行結果や評価を保存する
- 分析の入力になる

管理方針:

- `user_id`, `task_id`, `memory_id`, `context_id` の整合性を確認する
- 作成日時と更新日時を必ず持つ
- status の定義を揃える
- 入力値の正規化を行う

### 3. 分析データ

トランザクションデータから生成される分析結果。

対象 CSV:

- `heuristics_analysis.csv`
- `heuristics_tracking.csv`
- `heuristics_insights.csv`
- `heuristics_patterns.csv`
- `knowledge_patterns.csv`
- `language_optimization.csv`
- `quantification_labels.csv`

役割:

- 行動ログからパターンを抽出する
- タスク結果から知識を抽出する
- 自然言語を定量化・正規化する
- 次の作業改善に使う

管理方針:

- 元データへの参照を保持する
- confidence, accuracy, coverage, consistency などの品質指標を持つ
- 分析ロジックの version を保存する
- 再計算可能なデータと手入力データを区別する

### 4. 学習・モデルデータ

分析結果を使って改善されるモデルやパターン。

対象 CSV:

- `heuristics_models.csv`
- `heuristics_patterns.csv`
- `optimization_models.csv`

役割:

- ユーザーやタスクごとの傾向を保持する
- 推薦・最適化・予測の入力にする
- 継続改善の基盤にする

管理方針:

- model_type, version, parameters, performance を必須扱いにする
- training data の範囲を記録する
- performance が一定以下のモデルは非 active にする
- 古い model と新しい model を比較可能にする

## 中心キー設計

データ駆動設計では、データを横断して分析できるキー設計が重要。

### user_id

ユーザー単位の傾向分析に使う。

関連データ:

- Task
- Assessment
- Memory
- MemoryContext
- HeuristicsAnalysis
- HeuristicsTracking
- HeuristicsInsight
- HeuristicsModel
- QualitativeLabel
- QuantificationLabel

### task_id

作業単位の中心キー。

関連データ:

- Task
- Book
- Assessment
- MemoryContext
- KnowledgePattern
- LanguageOptimization
- HeuristicsAnalysis
- HeuristicsTracking
- HeuristicsInsight
- HeuristicsModel
- HeuristicsPattern
- OptimizationModel
- PhenomenologicalFramework

### context_id

作業文脈と技術要因をつなぐキー。

関連データ:

- MemoryContext
- TechnicalFactor

### source_analysis_id

分析結果から生成された insight を追跡するキー。

関連データ:

- HeuristicsAnalysis
- HeuristicsInsight

## データパイプライン

### 1. Ingest

CSV、API、ユーザー入力からデータを取り込む。

対象:

- seed CSV
- ユーザー入力
- 行動ログ
- 分析結果

処理:

- CSV header の検証
- 型変換
- 必須値チェック
- 重複 ID チェック

### 2. Normalize

データ形式を揃える。

処理:

- status 表記の統一
- date / timestamp の形式統一
- domain 名の正規化
- unit の正規化
- JSON 文字列の構造検証

例:

```txt
mm, millimeter, ミリメートル -> mm
kg, kilogram, キログラム -> kg
active, Active, 有効 -> active
```

### 3. Enrich

既存データを使って情報を補完する。

処理:

- task_id から user_id を補完
- robot_id から payload や reach を付与
- context から domain を推定
- qualitative feedback から category を推定

### 4. Aggregate

分析しやすい形に集計する。

処理:

- user ごとの平均評価
- task type ごとの成功率
- domain ごとの knowledge pattern 数
- robot model ごとの適合スコア
- heuristic pattern ごとの frequency

### 5. Analyze

改善や推薦に使う分析を行う。

処理:

- 低 confidence データの抽出
- 高 impact pattern の抽出
- task ごとのボトルネック分析
- user ごとの作業傾向分析
- robot / optimization model の候補比較

### 6. Feedback

分析結果を次の入力に戻す。

処理:

- HeuristicsInsight を生成
- KnowledgePattern を更新
- OptimizationModel の performance を更新
- QuantificationLabel の validated を更新
- 次回 task の推薦条件に反映

## データ品質設計

### Completeness

必須項目が埋まっているか。

例:

- `task.title` が空ではない
- `assessment.effectiveness_score` が存在する
- `robot_specifications.payload_kg` が存在する
- `quantification_labels.value` と `unit` が揃っている

### Consistency

データ間に矛盾がないか。

例:

- 存在しない `task_id` を参照していない
- `min_range <= typical_value <= max_range` を満たす
- `confidence` が 0.0 から 1.0 の範囲にある
- `created_at <= updated_at` を満たす

### Accuracy

値が実態に近いか。

例:

- KnowledgePattern の `accuracy`
- HeuristicsPattern の `accuracy`
- RobotSpecification の仕様値

### Timeliness

データが古すぎないか。

例:

- `last_seen` が古い HeuristicsPattern を低優先度にする
- 古い OptimizationModel の performance を再評価する
- 長期間更新されていない RobotSpecification を確認対象にする

### Traceability

どの元データから生成されたか追跡できるか。

例:

- HeuristicsInsight は `source_analysis_id` を持つ
- KnowledgePattern は `task_id` を持つ
- QuantificationLabel は `source` を持つ

## メトリクス設計

### Task Metrics

- task 数
- status 別 task 数
- priority 別 task 数
- 完了率
- 平均完了時間

### Assessment Metrics

- effectiveness_score 平均
- effort_score 平均
- impact_score 平均
- qualitative feedback の category 分布

### Knowledge Metrics

- KnowledgePattern 数
- domain 別 pattern 数
- accuracy 平均
- coverage 平均
- consistency 平均
- abstract_level 分布

### Heuristics Metrics

- analysis 数
- tracking event 数
- insight 数
- active insight 数
- pattern frequency
- confidence 平均
- error_count 平均

### Optimization Metrics

- model type 別 performance
- convergence_rate 平均
- iteration_count 平均
- domain 別 optimization model 数

### Robotics Metrics

- robot model 数
- payload range
- reach range
- repeat accuracy 分布
- vision system 搭載率
- force sensor 搭載率

## API / Usecase の設計例

### Data Quality API

```txt
GET /data-quality/summary
GET /data-quality/datasets/{name}
GET /data-quality/errors
POST /data-quality/validate
```

用途:

- CSV 取り込み前の検証
- DB 内データの品質確認
- 欠損や不整合の一覧化

### Analytics API

```txt
GET /analytics/tasks
GET /analytics/assessments
GET /analytics/knowledge
GET /analytics/heuristics
GET /analytics/optimization
```

用途:

- dashboard 表示
- 改善対象の抽出
- domain 別の傾向確認

### Pipeline API

```txt
POST /pipeline/ingest
POST /pipeline/normalize
POST /pipeline/enrich
POST /pipeline/analyze
POST /pipeline/feedback
```

用途:

- seed CSV の投入
- データ再処理
- 分析結果の再生成

## CSV ごとの設計観点

| CSV | 主な用途 | 品質チェック |
| --- | --- | --- |
| `task.csv` | 作業単位 | title, status, priority, user_id |
| `book.csv` | 作業資料 | task_id, title, status |
| `memories.csv` | 記憶・メモ | user_id, book_id, title |
| `memory_contexts.csv` | 作業文脈 | user_id, task_id, level, goal |
| `technical_factors.csv` | 技術条件 | context_id, tool_spec, eval_factors |
| `assessments.csv` | 評価 | score 範囲、task_id, user_id |
| `qualitative_label.csv` | 定性分類 | category, content |
| `knowledge_patterns.csv` | 知識化 | accuracy, coverage, consistency |
| `quantification_labels.csv` | 定量化 | value, unit, range, confidence |
| `language_optimization.csv` | 言語最適化 | original_text, optimized_text, score |
| `heuristics_tracking.csv` | 行動ログ | session_id, timestamp, duration |
| `heuristics_analysis.csv` | 分析結果 | score, confidence, status |
| `heuristics_insights.csv` | 改善提案 | source_analysis_id, confidence, is_active |
| `heuristics_models.csv` | 学習モデル | model_type, version, performance |
| `heuristics_patterns.csv` | 発見パターン | frequency, accuracy, impact_score |
| `optimization_models.csv` | 最適化モデル | objective_function, constraints, performance |
| `phenomenological_frameworks.csv` | 評価フレーム | goal, scope, process, result |
| `robot_specifications.csv` | ロボット仕様 | dof, reach, payload, accuracy |

## DDD との違い

DDD では業務概念、集約、責務境界を中心に設計する。

データ駆動設計では、データの収集、品質、加工、分析、フィードバックを中心に設計する。

```txt
DDD:
  Task, Memory, Knowledge などの業務概念を中心に分ける

Data Driven:
  Ingest, Normalize, Quality, Metrics, Analytics などのデータ処理を中心に分ける
```

どちらか一方だけに寄せる必要はない。実際には以下のように併用するのが現実的。

```txt
domain/
  業務ルール

data/
  データ品質、分析、集計、特徴量

infrastructure/
  DB、CSV、外部連携

usecase/
  業務操作とデータ処理の orchestration
```

## 最初に作るべきもの

データ駆動設計として最初に作るなら、以下の順がよい。

### 1. Dataset Catalog

各 CSV / DB table の定義を管理する。

持つ情報:

- dataset name
- source file
- primary key
- foreign keys
- required columns
- metric columns
- quality rules

### 2. Data Quality Validator

CSV と DB データの品質を検証する。

検証対象:

- 必須値
- 型
- ID 重複
- 外部キー存在
- score 範囲
- confidence 範囲
- timestamp 整合性

### 3. Data Pipeline

取り込みから分析までの処理を段階化する。

段階:

- ingest
- normalize
- validate
- enrich
- aggregate
- analyze
- feedback

### 4. Metrics Layer

dashboard や改善判断に使う指標を定義する。

最初に作る指標:

- task 完了率
- assessment score 平均
- knowledge accuracy 平均
- heuristics confidence 平均
- active insight 数
- low quality data 数

## まとめ

この backend をデータ駆動で設計する場合、中心は `task_id`, `user_id`, `context_id`, `source_analysis_id` になる。

DDD のように業務概念を中心に分けるのではなく、以下の流れを明確にする。

```txt
収集
→ 正規化
→ 品質検証
→ 補完
→ 集計
→ 分析
→ 改善提案
→ 次の入力へ反映
```

最初に取り組むべきは、`Dataset Catalog`, `Data Quality Validator`, `Data Pipeline`, `Metrics Layer` の 4 つ。

これにより、seed CSV が単なる初期データではなく、継続的に改善される backend のデータ基盤として扱えるようになる。
