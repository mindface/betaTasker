#!/bin/sh

# 保護したいブランチのリスト（スペース区切り）
protected_branches="main master develop"

# 現在のブランチ名を取得
current_branch=$(git rev-parse --abbrev-ref HEAD)

# 実行中のブランチが保護対象に含まれているかチェック
for branch in $protected_branches; do
    if [ "$current_branch" = "$branch" ]; then
        echo "========================================================"
        echo "エラー: '$branch' ブランチへの直接 push は禁止されています。"
        echo "プルリクエストを作成してください。"
        echo "========================================================"
        exit 1
    fi
done

exit 0
