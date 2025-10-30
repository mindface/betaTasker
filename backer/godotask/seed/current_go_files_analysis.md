# 現状のGoファイル内容に基づくSeedデータ構成解析

## 📁 ファイル構成とデータ構造

### **主要Seedファイル**

```
seed/
├── main_seed.go          # 統合実行管理
├── seed.go              # MemoryContext中心のデータ
├── seedModel.go         # Book/Task/Memory/Assessment
├── heuristics_seed.go   # AI分析・学習データ
├── phenomenological_seed.go      # 現象学フレームワーク（新規）
├── data_accumulation.go         # 蓄積管理（新規）
└── migration_strategy.go       # 移行戦略（新規）
```

## 🏗️ **1. main_seed.go - 実行制御**

```go
func RunAllSeeds() error {
    // 実行順序（依存関係順）
    1. SeedMemoryContexts()     // 製造業知識基盤
    2. SeedBooksAndTasks()      // 書籍・タスク・評価
    3. SeedHeuristics(db)       # AI分析データ  
    4. SeedPhenomenologicalData(db) // 現象学フレームワーク
}

func CleanAndSeed() error {
    // テーブルクリーンアップ（逆順削除）
    - HeuristicsModel/Pattern/Insight/Tracking/Analysis
    - TRUNCATE または DELETE実行
    - RunAllSeeds()実行
}
```

**特徴:**
- **段階的実行**: 依存関係を考慮した順次実行
- **クリーンアップ**: 外部キー制約対応の逆順削除
- **エラーハンドリング**: 各段階での失敗時対応

---

## 🏭 **2. seed.go - 製造業ドメイン知識**

### **Level構造化された切削加工知識**

```go
// Level 1: 基本操作（L1-1 ～ L1-5）
level1Data := []struct {
    workTarget       string  // "[職務カテゴリ: 切削・初品確認] 対象工程 L1-1: 初品加工・基本寸法確認"
    changeFactor     string  // "新規ロット材導入（ロット番号: A-123）"
    goal            string  // "初品寸法公差内維持、不良率5%以下"
    toolSpec        string  // "TNMG160408 (汎用), 標準コーティング"
    concern         string  // "初回切削でのバリ発生（要因: 切削条件未調整）"
    countermeasure  string  // "メーカー推奨値での標準切削開始、目視での品質確認"
    learnedKnowledge string // "基本的な切削条件とバリ発生の関係を理解。目視確認の重要性を認識。"
}

// Level 2: 応用技術（L2-1 ～ L2-5）
// SUS304ステンレス鋼への材質変更、AI対応切削条件マップ導入まで
level2Data := []struct {
    workTarget: "[職務カテゴリ: 品質改善・条件最適化] L2-5: SUS304量産加工に向けた自動補正"
    toolSpec: "センサ付き工具（摩耗・振動計測）、AI対応切削条件マップ搭載"
    learnedKnowledge: "加工プロセスを静的ではなく動的・予測的に捉える視点が重要"
}
```

### **データ関係構造**

```go
// 1. MemoryContext（親テーブル）
MemoryContext {
    UserID: 1
    TaskID: i + 1
    Level: 1 or 2        # 技術レベル
    WorkTarget: string   # 作業対象の詳細
    Machine: "NC旋盤（Mazak QT-200）"
    MaterialSpec: string # 材料仕様
    ChangeFactor: string # 変更要因
    Goal: string        # 目標
}

// 2. TechnicalFactor（技術詳細）
TechnicalFactor {
    ContextID: 親ID
    ToolSpec: string           # 工具仕様
    EvalFactors: string        # 評価要因
    MeasurementMethod: string  # 測定方法
    Concern: string           # 懸念事項
}

// 3. KnowledgeTransformation（知識変換）
KnowledgeTransformation {
    ContextID: 親ID
    Transformation: string     # 変換内容
    Countermeasure: string    # 対策
    ModelFeedback: string     # モデルフィードバック
    LearnedKnowledge: string  # 学習内容
}
```

**特徴:**
- **階層的知識**: L1(基本) → L2(応用) → L3(高度) → L4(専門) → L5(指導)
- **実践的データ**: 実際の製造現場の課題と解決策
- **技術進歩**: 従来工具 → センサ付き工具 → AI統合システム

---

## 📚 **3. seedModel.go - 評価・学習システム**

### **評価階級システム**

```go
func classifyScore(score int) string {
    case score >= 95: return "s_plus"  // 卓越（非常に優れている）
    case score >= 90: return "s"       // 優秀（文句なし）
    case score >= 85: return "a_plus"  // 良好（高評価）
    case score >= 80: return "a"       // 実用性あり（改善余地あり）
    case score >= 75: return "b_plus"  // 満足（条件付きで応用可能）
    // ... e: 再考・再検証が必要
}
```

### **動的コンテンツ生成**

```go
// 評価クラスに基づく自動生成
generateNoteText(scoreClass, tags)     // 評価コメント
generateTaskDescription(scoreClass)    // タスク説明
generateTaskTitle(scoreClass, title)   // タスクタイトル
generateFactor(scoreClass)            // 要因分析
generateProcess(scoreClass)           // プロセス段階
generateEvaluationAxis(scoreClass)    // 評価軸
generateInformationAmount(scoreClass) // 情報量
```

### **書籍・記憶・タスクの関連**

```go
// 1. Books（技術書籍）
books := []Book{
    {Title: "Advanced Metal Printing Techniques", Name: "高性能金属プリント技術"}
    {Title: "Understanding Titanium Alloys", Name: "チタン合金の基礎と応用"}
    // ... 8冊の専門技術書
}

// 2. Memory（学習記録）
for each book {
    Memory {
        UserID: 1
        SourceType: "book"
        Title: book.Title
        Tags: "3D,素材,評価"
        ReadStatus: "finished"
        Factor: generateFactor(scoreClass)      # 自動生成
        Process: generateProcess(scoreClass)    # 自動生成
        EvaluationAxis: generateEvaluationAxis(scoreClass)
        InformationAmount: generateInformationAmount(scoreClass)
    }
}

// 3. Task（実行タスク）
for each memory {
    Task {
        UserID: 1
        MemoryID: memory.ID
        Title: generateTaskTitle(scoreClass, memory.Title)
        Description: generateTaskDescription(scoreClass)
        Priority: based on scoreClass
    }
}

// 4. Assessment（評価）
for each task {
    Assessment {
        TaskID: task.ID
        UserID: 1
        EffectivenessScore: randomized based on priority
        EffortScore: randomized
        ImpactScore: randomized
        QualitativeFeedback: generated comments
    }
}
```

**特徴:**
- **自動化された関連付け**: 評価クラスに基づく一貫した生成
- **実践的な評価体系**: S+からEまでの11段階評価
- **完全なトレーサビリティ**: Book → Memory → Task → Assessment

---

## 🧠 **4. heuristics_seed.go - AI分析データ**

### **5層のヒューリスティクス構造**

```go
// 1. HeuristicsAnalysis（分析）
seedHeuristicsAnalysis(db) {
    analyses := []model.HeuristicsAnalysis{
        {
            UserID: 1, TaskID: 1
            AnalysisType: "performance"
            Result: {
                "completion_rate": 0.85,
                "accuracy": 0.92,
                "speed": "fast"
            }
        }
    }
}

// 2. HeuristicsTracking（行動追跡）
seedHeuristicsTracking(db) {
    tracking := []model.HeuristicsTracking{
        {
            UserID: 1
            ActionType: "task_start"
            ActionData: {"task_id": 1, "timestamp": "2024-01-01T09:00:00Z"}
        }
    }
}

// 3. HeuristicsInsight（洞察）
seedHeuristicsInsights(db) {
    insights := []model.HeuristicsInsight{
        {
            Title: "作業効率パターンの発見"
            Description: "午前中の作業効率が20%向上"
            Category: "performance"
            ConfidenceScore: 0.87
        }
    }
}

// 4. HeuristicsPattern（パターン）
seedHeuristicsPatterns(db) {
    patterns := []model.HeuristicsPattern{
        {
            PatternType: "learning_curve"
            PatternData: {"improvement_rate": 0.15, "plateau_point": 10}
            ConfidenceScore: 0.92
        }
    }
}

// 5. HeuristicsModel（モデル）
seedHeuristicsModels(db) {
    models := []model.HeuristicsModel{
        {
            ModelName: "タスク完了予測モデル"
            ModelType: "regression"
            Parameters: {"learning_rate": 0.001, "epochs": 100}
            Accuracy: 0.89
        }
    }
}
```

**特徴:**
- **包括的AI分析**: パフォーマンス・行動・学習・意思決定分析
- **実時間データ**: ユーザー行動の詳細追跡
- **予測モデル**: 機械学習による将来予測

---

## 🌟 **5. 新規追加ファイル群**

### **A. phenomenological_seed.go - 現象学フレームワーク**

```go
// 現象学的フレームワーク
PhenomenologicalFramework {
    ID: "robot_precision_framework"
    Goal: "G: 位置決め精度±0.01mm達成"
    Scope: "A: 6軸ロボットアームの動作範囲全体"  
    Process: {
        "Pa": "キャリブレーション→測定→補正→検証の反復プロセス"
        "steps": ["初期測定", "誤差解析", "補正値計算", "適用", "再測定"]
    }
    GoalFunction: "minimize(abs(measured_position - target_position))"
}

// 知識パターン（暗黙知→形式知変換）
KnowledgePattern {
    TacitKnowledge: "熟練工の『しっくりくる』感覚"
    ExplicitForm: "力覚センサ値: Fx<0.5N, Fy<0.5N, Tz<0.1Nm"
    ConversionPath: {
        "SECI": ["共同化", "表出化", "連結化", "内面化"]
        "method": "力覚データ記録→パターン分析→閾値設定"
    }
}
```

### **B. data_accumulation.go - データ蓄積管理**

```go
// 実運用データ収集
func (da *DataAccumulator) CollectProductionData() {
    // confidence > 0.8 AND validated = true の高品質データを抽出
    db.Where("confidence > ? AND validated = ?", 0.8, true).Find(&labels)
}

// 学習データ自動生成
func (da *DataAccumulator) GenerateLearningData() {
    // 既存パターンからバリエーション生成
    variations := generatePatternVariations(pattern)
}

// 差分バックアップ
func (da *DataAccumulator) CreateDifferentialBackup() {
    // 最後のバックアップ以降の変更データを抽出
    query := db.Where("updated_at > ?", lastBackup)
}
```

### **C. migration_strategy.go - 移行戦略**

```go
// 既存データ→新フレームワーク移行
func (dmm *DataMigrationManager) MigrateExistingToNewFramework() {
    1. migrateMemoryContextToFramework()     // MemoryContext → PhenomenologicalFramework
    2. migrateTechnicalFactorToLabel()       // TechnicalFactor → QuantificationLabel  
    3. migrateHeuristicsToKnowledge()        // HeuristicsPattern → KnowledgePattern
    4. migrateAssessmentToOptimization()     // Assessment → ProcessOptimization
}
```

---

## 🔗 **データ関係の実装詳細**

### **既存システムの関係図**

```
User(1) ──┬── Task(N) ── Assessment(1)
          │
          ├── Memory(N) ── MemoryContext(1) ──┬── TechnicalFactor(N)
          │                                   └── KnowledgeTransformation(N)
          │
          ├── HeuristicsAnalysis(N)
          ├── HeuristicsTracking(N)  
          ├── HeuristicsInsight(N)
          ├── HeuristicsPattern(N)
          └── HeuristicsModel(N)
```

### **新規システムとの統合**

```
[既存] MemoryContext ──┐
[既存] TechnicalFactor ──┼── [移行] ──┐
[既存] HeuristicsPattern ──┘           │
                                      ▼
                        [新規] PhenomenologicalFramework
                               KnowledgePattern
                               OptimizationModel
                               QuantificationLabel
```

## 🎯 **実装の特徴と優位性**

### **1. データ駆動型設計**
- **自動生成**: 評価クラスベースの一貫したコンテンツ生成
- **関連付け**: Book → Memory → Task → Assessment の完全なトレース
- **品質管理**: 11段階評価による詳細な品質分類

### **2. 段階的技術発展**
- **L1-L5の階層**: 基本操作から技術指導まで
- **技術進歩**: 従来工具 → AI統合システム
- **実践的知識**: 実際の製造現場のノウハウ

### **3. AI統合設計**  
- **包括的分析**: Performance/Behavior/Learning/Decision
- **予測モデル**: 機械学習による将来予測
- **リアルタイム**: 継続的な学習と改善

### **4. 現象学的統合**
- **Goal-Scope-Process**: G-A-Pa構造による明確化
- **暗黙知の形式知化**: SECI モデルによる知識変換
- **最適化関数**: goalFn()による数学的最適化

### **5. スケーラブル蓄積**
- **増分シード**: バージョン管理による段階的追加
- **品質向上**: 実運用データからの自動学習
- **移行戦略**: 既存データの有効活用

この構造により、製造業の実践的知識を現象学的フレームワークに統合し、AI支援による継続的改善が実現されます。