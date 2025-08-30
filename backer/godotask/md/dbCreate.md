``` dockerを含む ```
docker exec -it dbgodotask bash
```

``` psqlでDBに接続 ```
psql -U dbgodotask -d dbgodotask
```


```
CREATE TABLE memory_contexts (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  task_id INTEGER NOT NULL,
  level INTEGER NOT NULL,
  work_target TEXT NOT NULL,
  machine TEXT NOT NULL,
  material_spec TEXT NOT NULL,
  change_factor TEXT NOT NULL,
  goal TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL
);

-- technical_factors テーブル
CREATE TABLE technical_factors (
  id SERIAL PRIMARY KEY,
  context_id INTEGER NOT NULL REFERENCES memory_contexts(id) ON DELETE CASCADE,
  tool_spec TEXT NOT NULL,
  eval_factors TEXT NOT NULL,
  measurement_method TEXT NOT NULL,
  concern TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL
);
ALTER TABLE technical_factors ADD COLUMN measurement_method TEXT NOT NULL DEFAULT '';

-- knowledge_transformations テーブル
CREATE TABLE knowledge_transformations (
  id SERIAL PRIMARY KEY,
  context_id INTEGER NOT NULL REFERENCES memory_contexts(id) ON DELETE CASCADE,
  transformation TEXT NOT NULL,
  countermeasure TEXT NOT NULL,
  model_feedback TEXT NOT NULL,
  learned_knowledge TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL
);
```
