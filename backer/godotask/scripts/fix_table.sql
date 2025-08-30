-- technical_factorsテーブルを再作成
DROP TABLE IF EXISTS technical_factors CASCADE;

-- 再度マイグレーションを実行するか、手動で作成
CREATE TABLE technical_factors (
    id SERIAL PRIMARY KEY,
    context_id INTEGER NOT NULL,
    tool_spec TEXT,
    eval_factors TEXT,
    measurement_method TEXT,
    concern TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (context_id) REFERENCES memory_contexts(id) ON DELETE CASCADE
);