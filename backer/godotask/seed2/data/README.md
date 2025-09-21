# Seed Data Directory

## 概要
このディレクトリは、システムで使用する様々なseedデータを管理します。

## ディレクトリ構造

```
seed/
├── data/                           # JSONデータファイル
│   ├── phenomenological_frameworks.json
│   ├── knowledge_patterns.json
│   ├── optimization_models.json
│   └── robot_specifications.json
├── import/                        # インポート用データ
│   ├── csv/
│   └── json/
├── export/                        # エクスポート先
│   ├── high_quality_seeds_*.json
│   └── backups/
├── backup/                        # バックアップデータ
│   ├── 20240902_120000/
│   └── metadata/
└── templates/                     # データテンプレート
    ├── quantification_label_template.json
    └── knowledge_pattern_template.json
```

## データ蓄積戦略

### 1. 実運用データ収集
- **目的**: 高品質な実データを蓄積
- **頻度**: 日次
- **対象**: confidence > 0.8 かつ validated = true のデータ

### 2. 学習データ生成
- **目的**: 既存パターンからバリエーション生成
- **手法**: ドメイン横展開、パラメータ調整
- **適用**: 暗黙知→形式知変換の拡充

### 3. 差分バックアップ
- **目的**: データの変更履歴管理
- **頻度**: 週次
- **形式**: JSON差分ファイル

### 4. データ品質向上
- **目的**: 低品質データの自動改善
- **手法**: 類似高品質データからの推定
- **閾値**: confidence < 0.5 のデータを対象

## 使用方法

### シードデータの実行
```go
// 全シードデータの実行
seed.RunAllSeeds()

// 増分シード
seed.IncrementalSeed(db)

// データ蓄積
accumulator := seed.NewDataAccumulator(db)
accumulator.CollectProductionData()
```

### CSVインポート
```go
// CSVファイルからのインポート
accumulator.ImportFromCSV("robot_data.csv", "quantification_label")
```

### データエクスポート
```go
// JSONエクスポート
seed.ExportSeedToJSON(db)

// 高品質データの抽出
accumulator.CollectProductionData()
```

## データ品質基準

### 定量化ラベル
- **信頼度**: >= 0.8
- **検証済み**: validated = true
- **完全性**: 必須フィールドすべて入力済み

### 知識パターン
- **精度**: accuracy >= 0.85
- **網羅性**: coverage >= 0.75
- **一貫性**: consistency >= 0.90

### 現象学的フレームワーク
- **目的関数**: 明確に定義された goalFunction
- **抽象レベル**: L0-L3 の適切な分類
- **フィードバック**: 継続的な改善プロセス

## バージョン管理

### シードバージョン
- v1.0.0: 初期リリース
- v1.0.1: ロボットアーム基本パターン追加
- v1.0.2: 現象学的フレームワーク統合

### マイグレーション
```go
// バージョン間マイグレーション
accumulator.MigrateData("1.0.0", "1.1.0")
```

## データ形式

### JSON Schema
各データタイプには対応するJSONスキーマが定義されています：

- `phenomenological_framework.schema.json`
- `knowledge_pattern.schema.json`
- `quantification_label.schema.json`

### CSV形式
CSVインポート時のヘッダ形式：
```csv
id,original_text,normalized_text,category,domain,value,unit,confidence
```

## 自動化

### 定期実行
- **日次**: 実運用データ収集
- **週次**: データ品質向上処理
- **月次**: 差分バックアップ作成

### トリガー
- 新規データ追加時の自動精製
- 閾値を下回る品質データの自動改善提案
- 異常データの検出・隔離

## 監視・メトリクス

### データ品質メトリクス
- 全体信頼度の推移
- ドメイン別データ分布
- 未検証データの割合

### 蓄積メトリクス
- 日次データ増加量
- 重複データの検出率
- バックアップサイズの推移

## トラブルシューティング

### よくある問題
1. **JSON解析エラー**: フォーマット確認
2. **重複データ**: ID重複チェック
3. **メモリ不足**: バッチサイズ調整

### 復旧手順
1. バックアップからの復元
2. データ整合性チェック
3. 増分データの再適用