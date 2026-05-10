# Backend Seed Data から見たプロダクトドメイン駆動設計案

## 前提

現状の seed CSV は `backer/godotask/seed/data` 配下に配置されている。

この資料では、ドメイン駆動設計やデータ駆動設計とは別に、プロダクトとしての価値提供を中心に backend を設計する場合の構成案をまとめる。

プロダクトドメイン駆動では、以下を重視する。

- ユーザーが何を達成したいか
- どの体験がプロダクト価値になるか
- どの業務フローを最短で支援するか
- どのデータが価値体験に直結するか
- どの機能単位で成長させるか

## プロダクトの中心価値

この backend の seed data から見ると、プロダクトの中心価値は以下に置ける。

```txt
作業タスクを記録する
  ↓
作業文脈と技術条件を整理する
  ↓
評価・行動・知識を蓄積する
  ↓
暗黙知を形式知化する
  ↓
次の作業改善や最適化に活かす
```

単なるタスク管理ではなく、作業経験を知識化し、改善提案や最適化につなげるプロダクトとして捉える。

## 推奨プロダクトドメイン

```txt
Product Domains
├── Workspace
├── Task Management
├── Work Context
├── Knowledge Capture
├── Evaluation
├── Heuristics Intelligence
├── Optimization Support
├── Robotics Reference
└── Insight Dashboard
```

DDD のように純粋な業務概念で切るのではなく、ユーザー体験上の機能領域で切る。

## 推奨ディレクトリ構成

```txt
backer/godotask/
├── product/
│   ├── workspace/
│   │   ├── workspace_usecase.go
│   │   ├── workspace_query.go
│   │   └── workspace_summary.go
│   ├── task_management/
│   │   ├── create_task.go
│   │   ├── update_task_status.go
│   │   ├── task_board.go
│   │   └── task_detail.go
│   ├── work_context/
│   │   ├── context_capture.go
│   │   ├── technical_factor_capture.go
│   │   └── context_summary.go
│   ├── knowledge_capture/
│   │   ├── knowledge_conversion.go
│   │   ├── quantification.go
│   │   ├── language_optimization.go
│   │   └── knowledge_search.go
│   ├── evaluation/
│   │   ├── task_assessment.go
│   │   ├── qualitative_feedback.go
│   │   └── evaluation_summary.go
│   ├── heuristics_intelligence/
│   │   ├── tracking.go
│   │   ├── analysis.go
│   │   ├── insight_generation.go
│   │   └── pattern_detection.go
│   ├── optimization_support/
│   │   ├── model_selection.go
│   │   ├── framework_matching.go
│   │   └── recommendation.go
│   ├── robotics_reference/
│   │   ├── robot_catalog.go
│   │   └── robot_matching.go
│   └── insight_dashboard/
│       ├── overview.go
│       ├── task_insights.go
│       ├── knowledge_insights.go
│       └── improvement_actions.go
├── domain/
├── usecase/
├── infrastructure/
└── interface/
```

`product/` は必須ではないが、プロダクト体験に直結する orchestration を置く場所として有効。

既存の `domain`, `usecase`, `infrastructure`, `interface` を残しながら、ユーザー価値に沿った機能単位を上位に置く考え方。

## プロダクトドメイン別の役割

### Workspace

ユーザーが最初に見る作業空間。

役割:

- 自分のタスク一覧を見る
- 直近の作業状況を見る
- 未評価タスクを見る
- 重要な insight を見る
- 改善アクションを確認する

関連データ:

- `task.csv`
- `assessments.csv`
- `heuristics_insights.csv`
- `memory_contexts.csv`

主な API:

```txt
GET /workspace/summary
GET /workspace/tasks
GET /workspace/insights
GET /workspace/actions
```

### Task Management

タスクの作成、状態管理、詳細表示を扱う。

役割:

- タスクを作成する
- タスクの status を更新する
- priority を管理する
- task detail に関連データを集約する

関連データ:

- `task.csv`
- `book.csv`
- `memories.csv`

主な API:

```txt
POST /tasks
GET /tasks
GET /tasks/{id}
PATCH /tasks/{id}/status
PATCH /tasks/{id}/priority
```

プロダクト上の重要点:

- タスク詳細画面に、評価、文脈、知識、insight を集約して表示できること
- タスクは単なる TODO ではなく、改善サイクルの起点として扱うこと

### Work Context

作業対象、機械、材料、変化要因、目標、技術条件を記録する領域。

役割:

- 作業条件を登録する
- 技術要因を整理する
- タスクに対する作業文脈を明確にする
- 後続の分析や知識化の入力にする

関連データ:

- `memory_contexts.csv`
- `technical_factors.csv`

主な API:

```txt
POST /tasks/{task_id}/contexts
GET /tasks/{task_id}/contexts
POST /contexts/{context_id}/technical-factors
GET /contexts/{context_id}/summary
```

プロダクト上の重要点:

- ユーザーが作業前後に短時間で文脈を登録できること
- 分析のために詳細すぎる入力を強制しないこと
- 入力された文脈が後で insight や知識化につながること

### Knowledge Capture

暗黙知を形式知化し、再利用可能な知識として蓄積する領域。

役割:

- 暗黙知を明文化する
- 自然言語を正規化する
- 定量化ラベルを作成する
- domain ごとに知識パターンを検索する

関連データ:

- `knowledge_patterns.csv`
- `quantification_labels.csv`
- `language_optimization.csv`

主な API:

```txt
POST /knowledge/patterns
GET /knowledge/patterns
POST /knowledge/quantify
POST /knowledge/optimize-language
GET /knowledge/search
```

プロダクト上の重要点:

- ユーザーの曖昧な表現を責めず、あとから構造化できること
- `confidence`, `accuracy`, `coverage`, `consistency` を使い、知識品質を見える化すること
- 知識が task detail や insight dashboard に戻ってくること

### Evaluation

タスクの結果、効果、労力、影響、定性的フィードバックを扱う領域。

役割:

- タスク完了後の評価を記録する
- qualitative feedback を分類する
- タスク改善に使える評価サマリを作る

関連データ:

- `assessments.csv`
- `qualitative_label.csv`

主な API:

```txt
POST /tasks/{task_id}/assessments
GET /tasks/{task_id}/assessments
POST /tasks/{task_id}/qualitative-labels
GET /evaluation/summary
```

プロダクト上の重要点:

- 評価入力は軽くする
- 評価結果は dashboard や insight に反映する
- スコアだけでなく、定性的な違和感や改善点も蓄積する

### Heuristics Intelligence

ユーザー行動、分析、インサイト、パターン検出を扱うプロダクトの知能部分。

役割:

- 行動ログを記録する
- タスク実行を分析する
- insight を生成する
- recurring pattern を検出する
- 改善提案を出す

関連データ:

- `heuristics_tracking.csv`
- `heuristics_analysis.csv`
- `heuristics_insights.csv`
- `heuristics_models.csv`
- `heuristics_patterns.csv`

主な API:

```txt
POST /heuristics/tracking
POST /heuristics/analyze
GET /heuristics/insights
PATCH /heuristics/insights/{id}/active
GET /heuristics/patterns
```

プロダクト上の重要点:

- insight は出すだけでなく、ユーザーが使ったかどうかを追跡する
- confidence が低い提案は UI 上で控えめに扱う
- pattern はユーザー別、タスク別、domain 別に比較できること

### Optimization Support

最適化モデルや現象学的フレームワークを使って、タスク改善を支援する領域。

役割:

- タスクに適した最適化モデルを探す
- 目的関数と制約を整理する
- 現象学的フレームワークと照合する
- 改善候補を提示する

関連データ:

- `optimization_models.csv`
- `phenomenological_frameworks.csv`
- `task.csv`
- `memory_contexts.csv`

主な API:

```txt
GET /optimization/models
POST /optimization/match
POST /optimization/recommend
GET /frameworks
POST /frameworks/match
```

プロダクト上の重要点:

- ユーザーに数理最適化の詳細を強制しない
- タスク文脈から候補モデルを提示する
- 推薦理由を説明できる形で返す

### Robotics Reference

ロボット仕様を参照し、作業条件や最適化候補と照合する領域。

役割:

- ロボット仕様を一覧する
- 可搬重量、リーチ、精度などで絞り込む
- タスク条件に合うロボット候補を提示する

関連データ:

- `robot_specifications.csv`
- `technical_factors.csv`
- `optimization_models.csv`

主な API:

```txt
GET /robots
GET /robots/{id}
POST /robots/match
GET /robots/capabilities
```

プロダクト上の重要点:

- 仕様表としてだけでなく、作業改善の選択肢として使えること
- payload, reach, repeat accuracy などを比較しやすくすること

### Insight Dashboard

プロダクト全体の価値を可視化する領域。

役割:

- タスクの状態を見る
- 評価傾向を見る
- 知識化の進捗を見る
- active insight を見る
- 改善アクションを提示する

関連データ:

- `task.csv`
- `assessments.csv`
- `knowledge_patterns.csv`
- `heuristics_insights.csv`
- `heuristics_patterns.csv`
- `optimization_models.csv`

主な API:

```txt
GET /dashboard/overview
GET /dashboard/tasks
GET /dashboard/knowledge
GET /dashboard/heuristics
GET /dashboard/improvements
```

プロダクト上の重要点:

- ユーザーが次に何をすべきか分かること
- 数値と提案を一緒に出すこと
- 単なる分析画面ではなく、行動につながる画面にすること

## プロダクト体験フロー

### 基本フロー

```txt
1. ユーザーが Task を作成する
2. Work Context を登録する
3. 作業を進める
4. Heuristics Tracking が行動を記録する
5. Task 完了後に Assessment を入力する
6. Knowledge Capture が知識化する
7. Heuristics Analysis が分析する
8. Insight が生成される
9. Optimization Support が改善候補を出す
10. Dashboard が次の行動を提示する
```

### 画面単位で見た backend 境界

```txt
Workspace
  - task summary
  - active insights
  - improvement actions

Task Detail
  - task
  - book
  - memory context
  - technical factors
  - assessments
  - related knowledge
  - related insights

Knowledge Library
  - knowledge patterns
  - quantification labels
  - language optimized texts

Insight Dashboard
  - heuristics analysis
  - insights
  - patterns
  - optimization recommendations

Robot Catalog
  - robot specifications
  - matching result
```

## プロダクト KPI

プロダクトドメイン駆動では、backend の設計も KPI と接続する。

### Activation

ユーザーが初回価値に到達したか。

候補 KPI:

- 初回 task 作成率
- 初回 memory context 登録率
- 初回 assessment 入力率
- 初回 insight 表示率

### Engagement

継続的に使われているか。

候補 KPI:

- 週次 task 作成数
- assessment 入力率
- knowledge pattern 登録数
- active insight 閲覧数
- dashboard 訪問数

### Knowledge Growth

知識資産が増えているか。

候補 KPI:

- knowledge pattern 数
- quantification label 数
- language optimization 数
- validated data 率
- confidence 平均

### Improvement

プロダクトが作業改善に貢献しているか。

候補 KPI:

- effectiveness_score 平均
- effort_score 改善
- impact_score 平均
- insight 採用率
- recurring error_count の減少

### Intelligence Quality

分析や提案が有効か。

候補 KPI:

- insight confidence 平均
- insight expected_impact 平均
- heuristics pattern accuracy
- model performance
- recommendation 採用率

## データとプロダクト機能の対応

| プロダクト機能 | 主な CSV | ユーザー価値 |
| --- | --- | --- |
| Task Board | `task.csv` | 作業を管理できる |
| Task Detail | `task.csv`, `book.csv`, `memory_contexts.csv` | 作業の文脈を一箇所で見られる |
| Context Capture | `memory_contexts.csv`, `technical_factors.csv` | 作業条件を整理できる |
| Evaluation Form | `assessments.csv`, `qualitative_label.csv` | 作業結果を振り返れる |
| Knowledge Library | `knowledge_patterns.csv`, `quantification_labels.csv` | 暗黙知を再利用できる |
| Language Optimizer | `language_optimization.csv` | 曖昧な表現を整理できる |
| Insight Feed | `heuristics_insights.csv` | 次の改善行動が分かる |
| Behavior Analysis | `heuristics_tracking.csv`, `heuristics_analysis.csv` | 作業傾向を把握できる |
| Pattern Detection | `heuristics_patterns.csv` | 繰り返し現れる成功・失敗パターンを見つけられる |
| Optimization Recommendation | `optimization_models.csv`, `phenomenological_frameworks.csv` | 改善候補を比較できる |
| Robot Matching | `robot_specifications.csv` | 作業条件に合う機器を検討できる |

## API 設計の考え方

プロダクトドメイン駆動では、DB table に対応した CRUD API だけでは不十分。

画面やユーザー体験に合わせた API を用意する。

### よくない例

```txt
GET /tasks/{id}
GET /assessments?task_id=...
GET /memory_contexts?task_id=...
GET /heuristics_insights?task_id=...
GET /knowledge_patterns?task_id=...
```

frontend が複数 API を組み合わせる必要があり、プロダクト体験の責務が frontend に寄りすぎる。

### よい例

```txt
GET /task-detail/{task_id}
```

返す情報:

- task
- related books
- memory contexts
- technical factors
- assessments
- qualitative labels
- knowledge patterns
- active insights
- optimization recommendations

backend がプロダクト体験単位で集約して返す。

## Product Usecase の例

### GetTaskDetail

目的:

タスク詳細画面に必要な情報を一括取得する。

入力:

- user_id
- task_id

出力:

- task
- books
- contexts
- technical factors
- assessments
- knowledge patterns
- insights
- recommendations

### CompleteTaskWithAssessment

目的:

タスク完了と評価入力を一つの体験として扱う。

処理:

```txt
1. task status を completed にする
2. assessment を作成する
3. qualitative label を作成する
4. heuristics analysis を予約する
5. dashboard summary を更新する
```

### GenerateTaskInsights

目的:

タスクに関連するデータから改善 insight を生成する。

入力:

- task
- memory context
- assessment
- tracking logs
- knowledge patterns

出力:

- insight
- expected impact
- confidence
- recommendation

### RecommendOptimization

目的:

作業文脈に合う最適化モデルやフレームワークを提示する。

入力:

- task
- memory context
- technical factors
- robot specifications
- optimization models

出力:

- candidate models
- matching reason
- expected benefit
- constraints

## DDD / データ駆動との違い

### DDD

業務概念とルールを中心に設計する。

```txt
Task
Memory
Knowledge
Heuristics
Optimization
Robotics
```

### データ駆動

データの収集、品質、加工、分析を中心に設計する。

```txt
Ingest
Normalize
Validate
Aggregate
Analyze
Feedback
```

### プロダクトドメイン駆動

ユーザー体験と価値提供単位を中心に設計する。

```txt
Workspace
Task Detail
Context Capture
Knowledge Library
Insight Feed
Optimization Recommendation
Robot Matching
Dashboard
```

## 併用方針

実際の backend では、3 つを分けて考えるとよい。

```txt
product/
  ユーザー体験単位の usecase

domain/
  業務ルールとエンティティ

data/
  データ品質、分析、メトリクス

infrastructure/
  DB、CSV、外部連携

interface/
  HTTP controller、request / response
```

`product/` は `domain/` を直接置き換えるものではない。

`product/` は、複数の domain と data processing を組み合わせて、ユーザー体験として意味のある単位にまとめる層として扱う。

## 最初に作るべき Product Usecase

### 1. Workspace Summary

最初にユーザーが見る画面の backend。

集約するもの:

- active tasks
- recently completed tasks
- pending assessments
- active insights
- improvement actions

### 2. Task Detail

プロダクトの中心画面。

集約するもの:

- task
- context
- technical factors
- assessment
- knowledge
- insights
- recommendations

### 3. Complete Task Flow

タスク完了時の体験。

処理するもの:

- status 更新
- assessment 登録
- qualitative feedback 登録
- analysis trigger
- insight generation trigger

### 4. Insight Feed

改善提案を表示する体験。

集約するもの:

- heuristics insights
- expected impact
- confidence
- source analysis
- related task

### 5. Knowledge Library

蓄積された知識を再利用する体験。

集約するもの:

- knowledge patterns
- quantification labels
- language optimizations
- domain filters
- quality scores

## まとめ

プロダクトドメイン駆動で見ると、この backend は単なる CRUD や seed data 管理ではなく、以下の体験を支えるシステムとして設計する。

```txt
タスクを記録する
作業文脈を残す
評価する
知識化する
分析する
改善提案する
次の作業に活かす
```

最初に作るべき中心機能は以下。

- Workspace Summary
- Task Detail
- Complete Task Flow
- Insight Feed
- Knowledge Library

この 5 つを軸に backend API と usecase を整理すると、CSV に存在するデータがプロダクト価値に直結しやすくなる。
