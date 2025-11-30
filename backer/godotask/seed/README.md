# Seed Data Documentation

このディレクトリには、betaTaskerアプリケーションのデータベースを初期化するためのseedデータが含まれています。


## 📁 構成

### Seed ファイル
- `main_seed.go` - メインのseed実行関数
- `state_evaluation_seed.go` - 状態評価、ツールマッチング、プロセス監視、学習パターンのseedデータ
- `csv_loader.go` - CSVファイルからのデータ読み込み
- `seed.go` - 既存のbooks、tasksデータ
- `heuristics_seed.go` - ヒューリスティクス分析データ
- `phenomenological_seed.go` - 現象学的フレームワークデータ


### CSVデータファイル (`data/`)
- `robot_specifications.csv` - ロボット仕様データ（21種類）
- `optimization_models.csv` - 最適化モデルデータ（21種類）
- `phenomenological_frameworks.csv` - 現象学的フレームワーク（21種類）
- `memory_contexts.csv` - メモリコンテキスト（製造技能L1-L5）
- `knowledge_patterns.csv` - 知識パターン（21種類）
- `quantification_labels.csv` - 定量化ラベル


## 🚀 使用方法

### 1. 基本のseed実行
```bash
go run main.go seed
```

### 2. クリーンアップ後のseed実行
```bash
go run main.go clean-seed
```

### 3. 通常のアプリケーション起動
```bash
go run main.go
```

## 📊 Seedデータの内容

### 状態評価システム
- **StateEvaluation**: 5つのサンプル評価データ（L1-L5レベル）
- **ToolMatchingResult**: 3つのツールマッチング結果
- **ProcessMonitoring**: 2つの監視レコード（組み立て・溶接）
- **LearningPattern**: 3つのデフォルト学習パターン + CSV読み込み

### ロボット・最適化
- **RobotSpecification**: 21種類のロボット仕様
- **OptimizationModel**: 21種類の最適化モデル

### フレームワーク・パターン
- **PhenomenologicalFramework**: 21種類のG-A-Pa構造フレームワーク
- **QuantificationLabel**: 定量化ラベルデータ

## 🔄 データの依存関係

1. **基礎データ**: RobotSpecification, OptimizationModel, PhenomenologicalFramework
2. **評価データ**: StateEvaluation（基礎データを参照）
3. **結果データ**: ToolMatchingResult, ProcessMonitoring（評価データを参照）
4. **学習データ**: LearningPattern（独立）


## 🎯 特徴

### リアルな製造データ
- 実際の製造現場のタスク構成（L1-L5）
- 現実的な切削条件・材料データ
- 段階的な技能習得プロセス

### 多様なロボット仕様
- 教示フリーロボットから協働ロボット
- 精密加工用から重負荷用まで
- AI機能・安全機能の組み合わせ

### 包括的な最適化モデル
- 制御理論ベース〜ML・AIベース
- エネルギー最適化から品質管理まで
- 収束率・反復回数の実データ

## 📝 カスタマイズ

### 新しいseedデータの追加
1. `seed/data/` 配下にCSVファイルを追加
2. `csv_loader.go` に読み込み関数を実装
3. `main_seed.go` の `RunAllSeeds()` に呼び出しを追加

### 既存データの修正
- CSVファイルを直接編集
- または各seedファイル内のデフォルトデータを修正

## 🔧 トラブルシューティング

### CSVファイルが見つからない場合
- デフォルトデータが使用されます
- ログにWarningが出力されますが、処理は続行されます

### 外部キー制約エラー
- `CleanAndSeed()` 使用時はテーブル削除順序が重要
- `main_seed.go` の `tables` 配列の順序を確認

### データの重複
- IDが重複している場合、挿入時にエラーが発生
- `clean-seed` オプションで全データをクリアしてから実行

## 📈 拡張性

この seed システムは以下の拡張に対応：
- 新しいモデルタイプの追加
- CSV以外のデータソース（JSON、XML等）
- 環境別の seed データ切り替え
- バッチ処理でのデータ更新