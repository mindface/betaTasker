# アーキテクチャ移行 まとめ

## 現状

- ✅ **Auth機能**: クリーンアーキテクチャで実装済み
- ❌ **その他の機能**: 旧構造（Controller → Service → Repository → Model）

## 移行目標

すべての機能をAuth機能と同じクリーンアーキテクチャに統一する。

## アーキテクチャ構造

```
domain/              # ビジネスロジック層（フレームワーク非依存）
├── entity/          # エンティティ
├── repository/      # リポジトリインターフェース
└── service/        # ドメインサービスインターフェース

infrastructure/      # インフラ層（実装）
├── repository/      # リポジトリ実装（GORM等）
└── security/       # セキュリティ実装

usecase/            # ユースケース層（ビジネスロジック）

interface/http/      # HTTP層
├── controller/      # HTTPハンドラー
└── middleware/      # HTTPミドルウェア

model/              # GORMモデル（infrastructure層でのみ使用）
```

## 移行の流れ

1. **Domain層の作成** - EntityとRepositoryインターフェース
2. **Infrastructure層の移行** - Repository実装をdomain依存に変更
3. **Usecase層の作成** - ビジネスロジックの実装
4. **Interface層の移行** - Controllerをusecase依存に変更
5. **DI設定の更新** - server/init.goとrouter.goの更新

## 移行順序

1. **Task** (最優先)
2. **Memory**
3. **Assessment**
4. **Heuristics**
5. **その他**

## 参考ドキュメント

- [ARCHITECTURE_MIGRATION_PLAN.md](./ARCHITECTURE_MIGRATION_PLAN.md) - 詳細な移行計画
- [MIGRATION_EXAMPLE_TASK.md](./MIGRATION_EXAMPLE_TASK.md) - Task機能の実装例

## 次のアクション

1. Task機能の移行を開始
2. 移行パターンを確立
3. 他の機能に適用