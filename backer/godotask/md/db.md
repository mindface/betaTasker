``` dockerを含む ```
docker exec -it dbgodotask bash

``` psqlでDBに接続 ```
psql -U dbgodotask -d dbgodotask

```

CREATE TYPE source_type_enum AS ENUM ('book', 'article', 'video', 'lecture', 'other');
CREATE TYPE read_status_enum AS ENUM ('unread', 'reading', 'finished');


CREATE TABLE memory (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  source_type source_type_enum DEFAULT 'book',
  title VARCHAR(255) NOT NULL,
  author VARCHAR(255),
  notes TEXT,
  tags VARCHAR(255),
  read_status read_status_enum DEFAULT 'unread',
  read_date TIMESTAMP,
  factor VARCHAR(255),
  process VARCHAR(255),
  evaluation_axis VARCHAR(255),
  information_amount VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE task_status_enum AS ENUM ('todo', 'in_progress', 'completed');

CREATE TABLE task (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  memory_id INT,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  date TIMESTAMP,
  status task_status_enum DEFAULT 'todo',
  priority INT DEFAULT 3,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE assessment (
  id SERIAL PRIMARY KEY,
  task_id INT NOT NULL,
  user_id INT NOT NULL,
  effectiveness_score SMALLINT,
  effort_score SMALLINT,
  impact_score SMALLINT,
  qualitative_feedback TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  role VARCHAR(50) DEFAULT 'user',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  is_active BOOLEAN DEFAULT TRUE
);

DROP TABLE IF EXISTS book CASCADE;

CREATE TABLE book (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255),
  name VARCHAR(255),
  text TEXT,
  disc TEXT,
  img_path VARCHAR(1024),
  status VARCHAR(255)
);
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
