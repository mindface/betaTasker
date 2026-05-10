# Product Domain と DDD を接続する構造化仕様

## 前提

この資料では、以下の 2 つの設計を接続する。

- `productDomainDrivenBackendDesign.md`
- `dddBackendSeedDesign.md`

ここでの目的は、プロダクト体験を中心にした `product/` 層と、業務ルールを中心にした `domain/` 層を分離しつつ、backend 全体として一貫した構造にすること。

また、その構造をプロジェクトへ反映するための shell script 仕様も定義する。

この資料では、ユーザーの「反省させる sh ファイル」という表現は「反映させる sh ファイル」と解釈する。

## 設計方針

### Product Domain

ユーザー体験、画面、操作フローを中心にする。

例:

- Workspace
- Task Detail
- Complete Task Flow
- Insight Feed
- Knowledge Library
- Optimization Recommendation

### DDD Domain

業務概念、集約、エンティティ、値オブジェクト、ドメインルールを中心にする。

例:

- Task
- Memory
- Assessment
- Knowledge
- Heuristics
- Optimization
- Robotics

### 接続方針

`product/` は複数の `domain/` を組み合わせて、ユーザー体験として意味のある usecase を作る。

```txt
interface/controller
        ↓
product/*
        ↓
usecase/*
        ↓
domain/*
        ↓
infrastructure/*
```

`product/` は業務ルールそのものを持たない。

業務ルールは `domain/` に置く。

`product/` は画面や体験に必要な情報を集約し、複数 usecase を orchestration する。

## 推奨レイヤ構造

```txt
backer/godotask/
├── product/
│   ├── workspace/
│   ├── task_detail/
│   ├── complete_task_flow/
│   ├── insight_feed/
│   ├── knowledge_library/
│   └── optimization_recommendation/
├── domain/
│   ├── task/
│   ├── memory/
│   ├── assessment/
│   ├── knowledge/
│   ├── heuristics/
│   ├── optimization/
│   └── robotics/
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
│   └── seed/
└── interface/
    └── controller/
```

## Product Domain と DDD Domain の対応表

| Product Domain | 利用する DDD Domain | 目的 |
| --- | --- | --- |
| Workspace | Task, Assessment, Heuristics, Knowledge | ユーザーの現在状態と次アクションを表示する |
| Task Detail | Task, Memory, Assessment, Knowledge, Heuristics, Optimization | タスクに関係する情報を一画面に集約する |
| Complete Task Flow | Task, Assessment, Knowledge, Heuristics | タスク完了、評価、分析準備を一連の体験にする |
| Insight Feed | Heuristics, Task, Assessment, Knowledge | 改善提案を一覧化し、採用可否を扱う |
| Knowledge Library | Knowledge, Task, Assessment | 蓄積された知識を検索・再利用する |
| Optimization Recommendation | Optimization, Robotics, Memory, Task | 作業条件に合う改善候補を提示する |

## 依存ルール

### 許可する依存

```txt
interface -> product
interface -> usecase
product -> usecase
product -> domain
usecase -> domain
usecase -> domain/repository
infrastructure -> domain
infrastructure -> domain/repository
```

### 禁止する依存

```txt
domain -> product
domain -> usecase
domain -> infrastructure
domain -> interface
usecase -> product
infrastructure -> product
```

### 理由

`domain/` は最も安定した業務ルールの層にする。

`product/` は画面や体験の変化に合わせて変わりやすい。

そのため、`domain/` が `product/` を知ってはいけない。

## Product 層の責務

`product/` は以下を行う。

- 画面単位のデータ集約
- 複数 usecase の呼び出し順序制御
- 体験単位の入力 DTO / 出力 DTO 定義
- dashboard や detail 画面向けの summary 作成
- ユーザーの次アクションに必要な情報整理

`product/` は以下を行わない。

- DB 直接アクセス
- GORM model の直接操作
- CSV 読み込み
- 業務ルールの実装
- 永続化ルールの実装

## Domain 層の責務

`domain/` は以下を行う。

- entity 定義
- value object 定義
- aggregate root 定義
- domain service 定義
- repository interface 定義
- status 遷移などの業務ルール
- score や confidence の妥当性検証

`domain/` は以下を行わない。

- HTTP request / response の処理
- DB 接続
- CSV 読み込み
- 画面都合の summary 作成
- 複数画面にまたがる UI orchestration

## Product Usecase 詳細

### Workspace

目的:

ユーザーが現在取り組むべきことを把握する。

利用 domain:

- Task
- Assessment
- Heuristics
- Knowledge

出力例:

```txt
WorkspaceSummary
├── ActiveTasks
├── RecentlyCompletedTasks
├── PendingAssessments
├── ActiveInsights
└── ImprovementActions
```

### Task Detail

目的:

タスクに関連する情報を一画面に集約する。

利用 domain:

- Task
- Memory
- Assessment
- Knowledge
- Heuristics
- Optimization

出力例:

```txt
TaskDetail
├── Task
├── Books
├── MemoryContexts
├── TechnicalFactors
├── Assessments
├── QualitativeLabels
├── KnowledgePatterns
├── ActiveInsights
└── OptimizationRecommendations
```

### Complete Task Flow

目的:

タスク完了から評価、分析準備までを一つの体験にする。

処理:

```txt
1. Task status を completed にする
2. Assessment を作成する
3. QualitativeLabel を作成する
4. HeuristicsAnalysis の生成を予約する
5. Insight generation の対象にする
```

### Insight Feed

目的:

改善提案をユーザーが確認・採用・無効化できるようにする。

利用 domain:

- Heuristics
- Task
- Assessment
- Knowledge

出力例:

```txt
InsightFeed
├── Insight
├── RelatedTask
├── SourceAnalysis
├── ExpectedImpact
├── Confidence
└── Recommendation
```

### Knowledge Library

目的:

蓄積された知識を検索・再利用する。

利用 domain:

- Knowledge
- Task
- Assessment

出力例:

```txt
KnowledgeLibrary
├── KnowledgePatterns
├── QuantificationLabels
├── LanguageOptimizations
├── DomainFilters
└── QualityScores
```

### Optimization Recommendation

目的:

タスク文脈に対して、最適化モデル、現象学的フレームワーク、ロボット候補を提示する。

利用 domain:

- Optimization
- Robotics
- Memory
- Task

出力例:

```txt
OptimizationRecommendation
├── CandidateModels
├── CandidateFrameworks
├── CandidateRobots
├── MatchingReasons
└── Constraints
```

## Product 層のファイル雛形

各 product domain は以下の構成を基本にする。

```txt
product/{product_domain}/
├── dto.go
├── query.go
├── command.go
├── service.go
└── README.md
```

### dto.go

画面や体験単位の入出力構造を置く。

### query.go

参照系の product usecase を置く。

### command.go

更新系の product usecase を置く。

### service.go

query / command から共通利用される orchestration を置く。

### README.md

その product domain が接続する DDD domain、API、画面、責務を記述する。

## Domain 層のファイル雛形

各 DDD domain は以下の構成を基本にする。

```txt
domain/{domain_name}/
├── entity.go
├── value_object.go
├── repository.go
├── service.go
└── README.md
```

### entity.go

集約ルート、エンティティを置く。

### value_object.go

score、status、confidence、range、unit などを置く。

### repository.go

repository interface を置く。

### service.go

複数 entity にまたがる domain service を置く。

### README.md

集約、業務ルール、禁止依存、関連 CSV を記述する。

## Shell Script 仕様

## スクリプト名

```txt
backer/godotask/scripts/apply_product_ddd_structure.sh
```

## 目的

Product Domain と DDD Domain を接続するためのディレクトリ、README、Go 雛形ファイルを生成する。

この script は構造の初期反映を目的にする。

既存の実装を移動したり、削除したり、上書きしたりしない。

## 基本方針

- 既存ファイルは上書きしない
- 存在しないディレクトリだけ作成する
- 存在しない雛形ファイルだけ作成する
- 生成対象を標準出力に表示する
- `--dry-run` で作成予定だけ確認できる
- `--check` で構造が存在するか検証できる

## 実行例

```sh
cd backer/godotask
./scripts/apply_product_ddd_structure.sh --dry-run
./scripts/apply_product_ddd_structure.sh --apply
./scripts/apply_product_ddd_structure.sh --check
```

## 引数仕様

| 引数 | 必須 | 説明 |
| --- | --- | --- |
| `--dry-run` | 任意 | 作成予定を表示する。ファイルは作成しない |
| `--apply` | 任意 | ディレクトリと雛形ファイルを作成する |
| `--check` | 任意 | 必須ディレクトリとファイルが存在するか検証する |
| `--force` | 任意 | 将来用。既存ファイル上書きには使わない。初期仕様では無効扱い |
| `--help` | 任意 | 使い方を表示する |

`--dry-run`, `--apply`, `--check` は同時指定不可。

未指定の場合は `--dry-run` と同じ動作にする。

## 生成対象 Product Domain

```txt
product/workspace
product/task_detail
product/complete_task_flow
product/insight_feed
product/knowledge_library
product/optimization_recommendation
```

各ディレクトリに以下を作成する。

```txt
dto.go
query.go
command.go
service.go
README.md
```

## 生成対象 DDD Domain

```txt
domain/task
domain/memory
domain/assessment
domain/knowledge
domain/heuristics
domain/optimization
domain/robotics
```

各ディレクトリに以下を作成する。

```txt
entity.go
value_object.go
repository.go
service.go
README.md
```

## 生成対象 Usecase

```txt
usecase/task
usecase/memory
usecase/assessment
usecase/knowledge
usecase/heuristics
usecase/optimization
usecase/robotics
```

各ディレクトリに以下を作成する。

```txt
service.go
dto.go
README.md
```

## 生成対象 Infrastructure

```txt
infrastructure/seed
infrastructure/db/mapper
```

各ディレクトリに以下を作成する。

```txt
README.md
```

`infrastructure/db/mapper` は domain model と DB model の変換責務を置く場所。

## 生成対象 Interface

```txt
interface/controller/workspace
interface/controller/task_detail
interface/controller/insight_feed
interface/controller/knowledge_library
interface/controller/optimization_recommendation
```

各ディレクトリに以下を作成する。

```txt
controller.go
README.md
```

## README の記載内容

各 README には最低限以下を入れる。

### Product README

```txt
# {Product Domain Name}

## Role

## Connected DDD Domains

## Main Usecases

## API Candidates

## Notes
```

### Domain README

```txt
# {DDD Domain Name}

## Role

## Aggregate Roots

## Entities

## Value Objects

## Repository Interfaces

## Related Seed CSV

## Rules
```

### Usecase README

```txt
# {Usecase Name}

## Role

## Input

## Output

## Dependencies
```

## Go 雛形ファイル仕様

### package 名

ディレクトリ名に合わせる。

例:

```txt
product/task_detail -> package task_detail
domain/task -> package task
usecase/heuristics -> package heuristics
```

### 最小内容

空ファイルではなく、package 宣言と TODO コメントを置く。

例:

```go
package task_detail

// TODO: Define product-level DTOs for Task Detail.
```

## check モード仕様

`--check` は以下を検証する。

- 必須 product directory が存在する
- 必須 domain directory が存在する
- 必須 usecase directory が存在する
- 必須 infrastructure directory が存在する
- 必須 interface controller directory が存在する
- 各 directory に README が存在する
- Go ファイルに package 宣言がある

検証結果は標準出力に表示する。

成功時:

```txt
[OK] product/workspace
[OK] domain/task
[OK] usecase/task
Structure check passed.
```

失敗時:

```txt
[NG] missing: product/workspace
[NG] missing: domain/task/README.md
Structure check failed.
```

失敗時の exit code は `1`。

成功時の exit code は `0`。

## dry-run モード仕様

`--dry-run` は作成予定だけを表示する。

例:

```txt
[DRY-RUN] mkdir product/workspace
[DRY-RUN] touch product/workspace/dto.go
[DRY-RUN] touch product/workspace/README.md
```

既に存在する場合:

```txt
[SKIP] exists: product/workspace
[SKIP] exists: product/workspace/README.md
```

## apply モード仕様

`--apply` は実際にディレクトリとファイルを作成する。

例:

```txt
[CREATE] directory: product/workspace
[CREATE] file: product/workspace/dto.go
[SKIP] exists: domain/entity/user.go
```

既存ファイルは上書きしない。

既存ディレクトリはそのまま使う。

## エラー仕様

### 不明な引数

```txt
Unknown option: --xxx
Use --help for usage.
```

exit code は `1`。

### 複数モード指定

```txt
Only one mode can be specified: --dry-run, --apply, --check.
```

exit code は `1`。

### 想定外の実行場所

script は `backer/godotask` 直下で実行される想定。

`go.mod` が存在しない場合は中断する。

```txt
This script must be run from backer/godotask.
go.mod not found.
```

exit code は `1`。

## 安全性仕様

この script は以下を禁止する。

- `rm`
- 既存ファイルの上書き
- 既存ディレクトリの削除
- `git mv`
- `git checkout`
- DB 操作
- seed 実行
- migration 実行

構造作成だけを行う。

## 実装時の推奨 shell 設定

```sh
#!/usr/bin/env bash
set -euo pipefail
```

## 内部データ定義イメージ

script 内では配列で生成対象を持つ。

```sh
PRODUCT_DOMAINS=(
  "workspace"
  "task_detail"
  "complete_task_flow"
  "insight_feed"
  "knowledge_library"
  "optimization_recommendation"
)

DDD_DOMAINS=(
  "task"
  "memory"
  "assessment"
  "knowledge"
  "heuristics"
  "optimization"
  "robotics"
)
```

## 生成後の構造イメージ

```txt
backer/godotask/
├── product/
│   ├── workspace/
│   ├── task_detail/
│   ├── complete_task_flow/
│   ├── insight_feed/
│   ├── knowledge_library/
│   └── optimization_recommendation/
├── domain/
│   ├── task/
│   ├── memory/
│   ├── assessment/
│   ├── knowledge/
│   ├── heuristics/
│   ├── optimization/
│   └── robotics/
├── usecase/
│   ├── task/
│   ├── memory/
│   ├── assessment/
│   ├── knowledge/
│   ├── heuristics/
│   ├── optimization/
│   └── robotics/
└── interface/
    └── controller/
        ├── workspace/
        ├── task_detail/
        ├── insight_feed/
        ├── knowledge_library/
        └── optimization_recommendation/
```

## 移行時の注意

既存の `domain/entity`, `domain/repository`, `usecase/service` はすぐに削除しない。

まず新しい構造を追加し、既存実装を段階的に移す。

移行順は以下を推奨する。

```txt
1. product/task_detail を作る
2. domain/task を作る
3. usecase/task を作る
4. infrastructure/db/mapper を作る
5. 既存 task_service を usecase/task に寄せる
6. product/task_detail から usecase/task を呼ぶ
7. controller を product usecase に接続する
```

## まとめ

Product Domain と DDD Domain は役割が異なる。

```txt
Product Domain:
  ユーザー体験を成立させる単位

DDD Domain:
  業務ルールを守る単位
```

`product/` は複数の `domain/` と `usecase/` を組み合わせる orchestration 層として扱う。

`domain/` は product の都合から独立させる。

`apply_product_ddd_structure.sh` は、この考え方を安全にプロジェクトへ反映するための構造生成 script として設計する。
