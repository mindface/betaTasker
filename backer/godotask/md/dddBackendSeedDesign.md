# Backend Seed Data から見た DDD 構成案

## 前提

現状の seed CSV は `backer/godotask/seed/data` 配下に配置されている。

ユーザーの指定では `backend/seed/data` という表現だったが、リポジトリ上では `backer/godotask/seed/data` が該当する。

この資料では、現状の CSV データからドメイン駆動設計に寄せる場合の構成案をまとめる。コード変更は前提にしない。

## 現状から見える構造

既存 backend には以下のようなレイヤ構成がある。

```txt
backer/godotask/
├── domain/
├── usecase/
├── infrastructure/
├── interface/
└── seed/
```

すでに `domain / usecase / infrastructure / interface` の骨格は存在している。

ただし現状の usecase service では `infrastructure/db/model` の GORM model を直接扱っている箇所がある。DDD として整理する場合は、ドメインモデルと永続化モデルを分離し、repository で変換する構成が望ましい。

## 推奨ドメイン構成

```txt
backer/godotask/
├── domain/
│   ├── user/
│   │   ├── entity.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── task/
│   │   ├── entity.go
│   │   ├── value_object.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── memory/
│   │   ├── memory.go
│   │   ├── memory_context.go
│   │   ├── technical_factor.go
│   │   └── repository.go
│   ├── assessment/
│   │   ├── assessment.go
│   │   ├── qualitative_label.go
│   │   └── repository.go
│   ├── knowledge/
│   │   ├── knowledge_pattern.go
│   │   ├── quantification_label.go
│   │   ├── language_optimization.go
│   │   └── repository.go
│   ├── heuristics/
│   │   ├── analysis.go
│   │   ├── tracking.go
│   │   ├── insight.go
│   │   ├── model.go
│   │   ├── pattern.go
│   │   └── repository.go
│   ├── optimization/
│   │   ├── optimization_model.go
│   │   ├── phenomenological_framework.go
│   │   └── repository.go
│   └── robotics/
│       ├── robot_specification.go
│       └── repository.go
├── usecase/
│   ├── task/
│   ├── memory/
│   ├── assessment/
│   ├── knowledge/
│   ├── heuristics/
│   ├── optimization/
│   └── robotics/
├── infrastructure/
│   ├── db/
│   │   ├── model/
│   │   └── repository/
│   └── seed/
│       ├── csv_reader.go
│       └── importer/
└── interface/
    └── controller/
```

## CSV とドメインの対応

| ドメイン | 対応 CSV / seed | 役割 |
| --- | --- | --- |
| user | users_seed.go 由来 | ユーザー、所有者、認証主体 |
| task | `task.csv`, `book.csv` | タスク、作業資料、参照情報 |
| memory | `memories.csv`, `memory_contexts.csv`, `technical_factors.csv` | 記憶、作業文脈、技術条件 |
| assessment | `assessments.csv`, `qualitative_label.csv` | 評価、定性的フィードバック、ラベル |
| knowledge | `knowledge_patterns.csv`, `quantification_labels.csv`, `language_optimization.csv` | 暗黙知の形式知化、定量化、言語最適化 |
| heuristics | `heuristics_analysis.csv`, `heuristics_tracking.csv`, `heuristics_insights.csv`, `heuristics_models.csv`, `heuristics_patterns.csv` | 行動ログ、分析、インサイト、学習済みパターン |
| optimization | `optimization_models.csv`, `phenomenological_frameworks.csv` | 最適化モデル、目的関数、制約、現象学的フレームワーク |
| robotics | `robot_specifications.csv` | ロボット仕様マスタ |

## 境界づけられたコンテキスト

### Task Context

タスク実行そのものを扱う中心コンテキスト。

対象 CSV:

- `task.csv`
- `book.csv`
- `assessments.csv`
- `qualitative_label.csv`

主な概念:

- Task
- Book
- Assessment
- QualitativeLabel

`Task` は多くの CSV から `task_id` で参照されているため、システム全体の中心集約として扱いやすい。

### Memory Context

作業文脈、記憶、条件を扱うコンテキスト。

対象 CSV:

- `memories.csv`
- `memory_contexts.csv`
- `technical_factors.csv`

主な概念:

- Memory
- MemoryContext
- TechnicalFactor

`memory_contexts.csv` は `work_target`, `machine`, `material_spec`, `change_factor`, `goal` を持っており、作業条件を表すドメインモデルとして分離する価値がある。

### Knowledge Context

暗黙知、形式知、定量化、言語変換を扱うコンテキスト。

対象 CSV:

- `knowledge_patterns.csv`
- `quantification_labels.csv`
- `language_optimization.csv`

主な概念:

- KnowledgePattern
- QuantificationLabel
- LanguageOptimization

このコンテキストは、現場知識を構造化して再利用可能にする役割を持つ。

### Heuristics Context

ユーザーの行動、分析結果、学習パターン、インサイトを扱うコンテキスト。

対象 CSV:

- `heuristics_analysis.csv`
- `heuristics_tracking.csv`
- `heuristics_insights.csv`
- `heuristics_models.csv`
- `heuristics_patterns.csv`

主な概念:

- HeuristicsAnalysis
- HeuristicsTracking
- HeuristicsInsight
- HeuristicsModel
- HeuristicsPattern

行動ログと分析結果を分けて扱い、分析結果から insight や pattern を生成する流れにすると責務が明確になる。

### Optimization Context

目的関数、制約、最適化モデル、現象学的フレームワークを扱うコンテキスト。

対象 CSV:

- `optimization_models.csv`
- `phenomenological_frameworks.csv`

主な概念:

- OptimizationModel
- PhenomenologicalFramework

最適化モデルは task に紐づく場合もあるが、マスタデータとしても使えるため、Task Context から独立させるのがよい。

### Robotics Context

ロボット仕様を扱うマスタ系コンテキスト。

対象 CSV:

- `robot_specifications.csv`

主な概念:

- RobotSpecification

ロボットの DOF、リーチ、可搬重量、精度、速度、制御方式、センサーなどを値として保持する。

## 集約ルート候補

```txt
Task
├── Book
├── Assessment
└── QualitativeLabel

MemoryContext
└── TechnicalFactor

KnowledgePattern
├── QuantificationLabel
└── LanguageOptimization

HeuristicsAnalysis
├── HeuristicsInsight
└── HeuristicsPattern

HeuristicsTracking

OptimizationModel

PhenomenologicalFramework

RobotSpecification
```

`Task` は多くのデータの起点になるが、すべてを Task 集約に入れると肥大化する。`Knowledge`, `Heuristics`, `Optimization`, `Robotics` は別コンテキストにして、`task_id` は参照 ID として持つ程度に留めるのがよい。

## レイヤ別責務

### domain

業務ルールを置く。

例:

- Task の status 遷移
- priority の範囲
- Assessment score の妥当性
- QuantificationLabel の value / unit / confidence の整合性
- RobotSpecification の仕様値の妥当性
- HeuristicsInsight の active 判定

### usecase

アプリケーション操作を表す。

例:

- タスクを作成する
- タスクに評価を追加する
- memory context から technical factor を登録する
- heuristics analysis から insight を生成する
- knowledge pattern を検索する

usecase は DB model ではなく domain model を扱うのが望ましい。

### infrastructure

DB、CSV、外部 API などの技術詳細を置く。

例:

- GORM model
- repository 実装
- CSV reader
- seed importer
- domain model と DB model の mapper

### interface

HTTP controller、request / response DTO を置く。

controller は usecase を呼び出すだけにし、業務判断を持たせない。

## DDD 化する場合の移行順

### 1. Task Context から始める

`task.csv`, `book.csv`, `assessments.csv` は既存サービスにも近く、最初の分離対象として扱いやすい。

最初にやること:

- `domain/task` を作る
- `Task` entity を定義する
- `TaskRepository` interface を domain 側に置く
- GORM repository は infrastructure 側に置く
- usecase service が `model.Task` ではなく `domain/task.Task` を扱うようにする

### 2. Memory Context を分ける

`memory_contexts.csv` と `technical_factors.csv` は作業文脈としてまとまりがある。

最初にやること:

- `domain/memory` を作る
- `MemoryContext` を集約ルートにする
- `TechnicalFactor` を `MemoryContext` に紐づく entity または value object として扱う

### 3. Knowledge Context を分ける

暗黙知、定量化、言語最適化は一つの業務領域として扱う。

最初にやること:

- `domain/knowledge` を作る
- `KnowledgePattern`, `QuantificationLabel`, `LanguageOptimization` を定義する
- confidence, accuracy, coverage, consistency などの品質値を value object 化する

### 4. Heuristics Context を分ける

分析・追跡・インサイト・モデル・パターンは責務が増えやすいため、Task から独立させる。

最初にやること:

- `domain/heuristics` を作る
- tracking はログ系、analysis は結果系、insight は提案系として分ける
- analysis から insight を生成する domain service を検討する

### 5. Optimization / Robotics をマスタ系として整理する

最適化モデル、現象学的フレームワーク、ロボット仕様は参照される側のデータとして扱いやすい。

最初にやること:

- `domain/optimization` を作る
- `domain/robotics` を作る
- Task や Heuristics からは ID 参照にする

## 実装時の注意点

### DB model を domain model にしない

GORM tag や nullable の都合を domain に持ち込むと、業務ルールが DB 都合に引きずられる。

```txt
domain model        業務ルール中心
infrastructure model DB 保存中心
DTO                 API 入出力中心
```

この 3 つは分ける。

### CSV seed は infrastructure 扱いにする

CSV は初期データ投入の手段であり、ドメインそのものではない。

```txt
seed/data/*.csv
        ↓
infrastructure/seed importer
        ↓
domain model
        ↓
repository
        ↓
DB model
```

### task_id で全部を強集約にしない

多くの CSV は `task_id` を持っているが、すべてを Task 集約配下にすると Task が肥大化する。

以下のように分ける。

```txt
Task が所有するもの:
- Task
- Book
- Assessment
- QualitativeLabel

Task を参照するもの:
- KnowledgePattern
- HeuristicsAnalysis
- HeuristicsTracking
- HeuristicsInsight
- OptimizationModel
- PhenomenologicalFramework
```

## まとめ

現状の CSV から見ると、DDD の切り方は以下が自然。

```txt
user
task
memory
assessment
knowledge
heuristics
optimization
robotics
```

最初から全体を大きく作り直すより、まずは `Task Context` を小さく分離し、usecase が DB model を直接扱わない形に寄せるのが現実的。

その後、`Memory`, `Knowledge`, `Heuristics` の順で分離すると、CSV の構造と既存 backend の責務を崩さずに DDD 化しやすい。
