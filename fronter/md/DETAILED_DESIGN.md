# Memory & Assessment System 詳細設計書


## 1. コンポーネント詳細設計

### 1.1 SectionMemory コンポーネント

#### 概要
メモリ管理の中核コンポーネント。CRUD操作とMemory Aidの統合表示を提供。

#### Props
```typescript
interface SectionMemoryProps {
  userId?: number;
  initialFilter?: MemoryFilter;
  onMemorySelect?: (memory: Memory) => void;
  showAidSection?: boolean;
}
```

#### State管理
```typescript
const SectionMemory: React.FC<SectionMemoryProps> = () => {
  // Redux State
  const memories = useSelector((state: RootState) => state.memory.memories);
  const loading = useSelector((state: RootState) => state.memory.memoryLoading);
  const error = useSelector((state: RootState) => state.memory.memoryError);
  
  // Local State
  const [selectedMemory, setSelectedMemory] = useState<Memory | null>(null);
  const [showModal, setShowModal] = useState(false);
  const [filterCriteria, setFilterCriteria] = useState<MemoryFilter>({
    source_type: '',
    tags: [],
    read_status: '',
    date_range: null,
  });
  const [sortBy, setSortBy] = useState<'date' | 'title' | 'author'>('date');
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('desc');
};
```

#### メソッド仕様
```typescript
// メモリ取得
const fetchMemories = async (): Promise<void> => {
  try {
    await dispatch(loadMemories());
  } catch (error) {
    console.error('Failed to load memories:', error);
    showErrorNotification('メモリの読み込みに失敗しました');
  }
};

// メモリ作成
const handleCreateMemory = async (data: MemoryFormData): Promise<void> => {
  try {
    const newMemory = await dispatch(createMemory(data)).unwrap();
    setShowModal(false);
    showSuccessNotification('メモリを作成しました');
    onMemorySelect?.(newMemory);
  } catch (error) {
    showErrorNotification('メモリの作成に失敗しました');
  }
};

// フィルタリング
const filterMemories = (memories: Memory[]): Memory[] => {
  return memories.filter(memory => {
    if (filterCriteria.source_type && memory.source_type !== filterCriteria.source_type) {
      return false;
    }
    if (filterCriteria.tags.length > 0) {
      const memoryTags = memory.tags.split(',').map(t => t.trim());
      if (!filterCriteria.tags.some(tag => memoryTags.includes(tag))) {
        return false;
      }
    }
    if (filterCriteria.read_status && memory.read_status !== filterCriteria.read_status) {
      return false;
    }
    if (filterCriteria.date_range) {
      const memoryDate = new Date(memory.created_at);
      if (memoryDate < filterCriteria.date_range.start || 
          memoryDate > filterCriteria.date_range.end) {
        return false;
      }
    }
    return true;
  });
};

// ソート
const sortMemories = (memories: Memory[]): Memory[] => {
  return [...memories].sort((a, b) => {
    let compareValue = 0;
    switch (sortBy) {
      case 'date':
        compareValue = new Date(a.created_at).getTime() - new Date(b.created_at).getTime();
        break;
      case 'title':
        compareValue = a.title.localeCompare(b.title);
        break;
      case 'author':
        compareValue = a.author.localeCompare(b.author);
        break;
    }
    return sortOrder === 'asc' ? compareValue : -compareValue;
  });
};
```

#### レンダリング構造
```tsx
return (
  <div className={styles.container}>
    {/* ヘッダー部 */}
    <div className={styles.header}>
      <h2>メモリ管理</h2>
      <button onClick={() => setShowModal(true)}>新規作成</button>
    </div>
    
    {/* フィルター部 */}
    <MemoryFilter 
      criteria={filterCriteria}
      onChange={setFilterCriteria}
    />
    
    {/* ソート部 */}
    <MemorySort
      sortBy={sortBy}
      sortOrder={sortOrder}
      onSortChange={(by, order) => {
        setSortBy(by);
        setSortOrder(order);
      }}
    />
    
    {/* メモリ一覧 */}
    <div className={styles.memoryList}>
      {sortedAndFilteredMemories.map(memory => (
        <MemoryCard
          key={memory.id}
          memory={memory}
          onSelect={() => setSelectedMemory(memory)}
          onEdit={() => handleEditMemory(memory)}
          onDelete={() => handleDeleteMemory(memory.id)}
        />
      ))}
    </div>
    
    {/* Memory Aid セクション */}
    {showAidSection && selectedMemory && (
      <MemoryAidList memoryId={selectedMemory.id} />
    )}
    
    {/* モーダル */}
    {showModal && (
      <MemoryModal
        memory={selectedMemory}
        onSave={handleSaveMemory}
        onClose={() => setShowModal(false)}
      />
    )}
  </div>
);
```

### 1.2 SectionAssessmentRelation コンポーネント

#### 概要
Task-Memory-Assessment間の関連性を管理し、コンテキストベースの評価を実現。

#### Props
```typescript
interface SectionAssessmentRelationProps {
  taskId?: number;
  userId?: number;
  onAssessmentCreate?: (assessment: Assessment) => void;
}
```

#### 関連性ロジック
```typescript
const SectionAssessmentRelation: React.FC = () => {
  // 関連データの取得
  const [relatedData, setRelatedData] = useState<RelationData>({
    task: null,
    memory: null,
    assessments: [],
  });

  // タスク選択時の処理
  const handleTaskSelect = async (taskId: number) => {
    try {
      // タスク取得
      const task = await dispatch(getTaskById(taskId)).unwrap();
      
      // 関連メモリ取得
      let memory = null;
      if (task.memory_id) {
        memory = await dispatch(getMemory(task.memory_id)).unwrap();
      }
      
      // 関連評価取得
      const assessments = await dispatch(
        getAssessmentsForTaskUser({
          task_id: taskId,
          user_id: userId || getCurrentUserId(),
        })
      ).unwrap();
      
      setRelatedData({ task, memory, assessments });
    } catch (error) {
      console.error('Failed to load related data:', error);
    }
  };

  // 評価作成時の関連付け
  const createAssessmentWithRelations = async (data: AssessmentFormData) => {
    if (!relatedData.task) {
      throw new Error('Task must be selected');
    }
    
    const assessmentData = {
      ...data,
      task_id: relatedData.task.id,
      user_id: userId || getCurrentUserId(),
      // メモリからのコンテキスト情報を追加
      context: relatedData.memory ? {
        memory_title: relatedData.memory.title,
        memory_notes: relatedData.memory.notes,
        memory_factor: relatedData.memory.factor,
      } : null,
    };
    
    const assessment = await dispatch(createAssessment(assessmentData)).unwrap();
    onAssessmentCreate?.(assessment);
    
    // 関連データを再取得
    await handleTaskSelect(relatedData.task.id);
  };
};
```

### 1.3 AssessmentModal コンポーネント

#### 概要
評価の作成・編集を行うモーダルコンポーネント。スコアリングとフィードバック入力を提供。

#### Props & State
```typescript
interface AssessmentModalProps {
  assessment?: Assessment | null;
  taskId: number;
  userId: number;
  relatedMemory?: Memory | null;
  onSave: (data: AssessmentFormData) => Promise<void>;
  onClose: () => void;
}

interface AssessmentFormData {
  effectiveness_score: number;
  effort_score: number;
  impact_score: number;
  qualitative_feedback: string;
}
```

#### バリデーションロジック
```typescript
const validateForm = (data: AssessmentFormData): ValidationErrors => {
  const errors: ValidationErrors = {};
  
  // スコアの範囲チェック
  if (data.effectiveness_score < 0 || data.effectiveness_score > 100) {
    errors.effectiveness_score = 'スコアは0〜100の範囲で入力してください';
  }
  
  if (data.effort_score < 0 || data.effort_score > 100) {
    errors.effort_score = 'スコアは0〜100の範囲で入力してください';
  }
  
  if (data.impact_score < 0 || data.impact_score > 100) {
    errors.impact_score = 'スコアは0〜100の範囲で入力してください';
  }
  
  // フィードバックの文字数チェック
  if (data.qualitative_feedback.length < 10) {
    errors.qualitative_feedback = 'フィードバックは10文字以上入力してください';
  }
  
  if (data.qualitative_feedback.length > 1000) {
    errors.qualitative_feedback = 'フィードバックは1000文字以内で入力してください';
  }
  
  return errors;
};
```

#### スコア計算ロジック
```typescript
const calculateOverallScore = (
  effectiveness: number,
  effort: number,
  impact: number
): number => {
  // 重み付け平均を計算
  const weights = {
    effectiveness: 0.4,
    effort: 0.3,
    impact: 0.3,
  };
  
  return Math.round(
    effectiveness * weights.effectiveness +
    effort * weights.effort +
    impact * weights.impact
  );
};

const getScoreGrade = (score: number): string => {
  if (score >= 90) return 'S';
  if (score >= 80) return 'A';
  if (score >= 70) return 'B';
  if (score >= 60) return 'C';
  if (score >= 50) return 'D';
  return 'F';
};

const getScoreColor = (score: number): string => {
  if (score >= 90) return '#4CAF50'; // Green
  if (score >= 70) return '#2196F3'; // Blue
  if (score >= 50) return '#FF9800'; // Orange
  return '#F44336'; // Red
};
```

## 2. API層詳細設計

### 2.1 Next.js API Routes

#### Memory API Route実装
```typescript
// /app/api/memory/route.ts
import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import jwt from 'jsonwebtoken';

// 認証ミドルウェア
const authenticate = async (request: NextRequest): Promise<UserInfo | null> => {
  const cookieStore = cookies();
  const userCookie = cookieStore.get('user_info');
  
  if (!userCookie) {
    return null;
  }
  
  try {
    const decoded = jwt.verify(userCookie.value, process.env.JWT_SECRET!);
    return decoded as UserInfo;
  } catch (error) {
    console.error('JWT verification failed:', error);
    return null;
  }
};

// GET: メモリ一覧取得
export async function GET(request: NextRequest) {
  const user = await authenticate(request);
  if (!user) {
    return NextResponse.json(
      { error: 'Unauthorized' },
      { status: 401 }
    );
  }
  
  try {
    // クエリパラメータの取得
    const { searchParams } = new URL(request.url);
    const params = new URLSearchParams();
    
    // フィルタパラメータの追加
    if (searchParams.get('source_type')) {
      params.append('source_type', searchParams.get('source_type')!);
    }
    if (searchParams.get('tags')) {
      params.append('tags', searchParams.get('tags')!);
    }
    if (searchParams.get('read_status')) {
      params.append('read_status', searchParams.get('read_status')!);
    }
    
    // ページネーション
    const page = searchParams.get('page') || '1';
    const limit = searchParams.get('limit') || '20';
    params.append('page', page);
    params.append('limit', limit);
    
    // バックエンドAPIコール
    const response = await fetch(
      `${process.env.BACKEND_URL}/api/memory?${params}`,
      {
        headers: {
          'Authorization': `Bearer ${user.token}`,
          'Content-Type': 'application/json',
        },
      }
    );
    
    if (!response.ok) {
      throw new Error(`Backend API error: ${response.status}`);
    }
    
    const data = await response.json();
    
    // レスポンスの整形
    return NextResponse.json({
      data: data.memories || [],
      pagination: {
        page: parseInt(page),
        limit: parseInt(limit),
        total: data.total || 0,
        totalPages: Math.ceil((data.total || 0) / parseInt(limit)),
      },
      message: 'Memories fetched successfully',
    });
    
  } catch (error) {
    console.error('Memory fetch error:', error);
    return NextResponse.json(
      { error: 'Failed to fetch memories' },
      { status: 500 }
    );
  }
}

// POST: メモリ作成
export async function POST(request: NextRequest) {
  const user = await authenticate(request);
  if (!user) {
    return NextResponse.json(
      { error: 'Unauthorized' },
      { status: 401 }
    );
  }
  
  try {
    const body = await request.json();
    
    // バリデーション
    const validationError = validateMemoryInput(body);
    if (validationError) {
      return NextResponse.json(
        { error: validationError },
        { status: 400 }
      );
    }
    
    // ユーザーIDの追加
    const memoryData = {
      ...body,
      user_id: user.id,
    };
    
    // バックエンドAPIコール
    const response = await fetch(
      `${process.env.BACKEND_URL}/api/memory`,
      {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${user.token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(memoryData),
      }
    );

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to create memory');
    }

    const data = await response.json();

    return NextResponse.json({
      data: data,
      message: 'Memory created successfully',
    }, { status: 201 });
    
  } catch (error) {
    console.error('Memory creation error:', error);
    return NextResponse.json(
      { error: error.message || 'Failed to create memory' },
      { status: 500 }
    );
  }
}

// バリデーション関数
function validateMemoryInput(data: any): string | null {
  if (!data.title || data.title.trim().length === 0) {
    return 'Title is required';
  }
  
  if (data.title.length > 200) {
    return 'Title must be less than 200 characters';
  }
  
  if (!data.source_type) {
    return 'Source type is required';
  }
  
  const validSourceTypes = ['book', 'article', 'video', 'course', 'other'];
  if (!validSourceTypes.includes(data.source_type)) {
    return 'Invalid source type';
  }
  
  if (data.notes && data.notes.length > 5000) {
    return 'Notes must be less than 5000 characters';
  }

  return null;
}
```

### 2.2 Service層実装

#### Memory Service
```typescript
// /services/memoryApi.ts
import { ApplicationError, ErrorCode } from '../errors/errorCodes';

class MemoryService {
  private baseUrl = '/api/memory';
  
  // 一覧取得（ページネーション対応）
  async getMemories(params?: MemoryQueryParams): Promise<PaginatedResponse<Memory>> {
    try {
      const queryParams = new URLSearchParams();
      
      if (params) {
        Object.entries(params).forEach(([key, value]) => {
          if (value !== undefined && value !== null) {
            queryParams.append(key, String(value));
          }
        });
      }
      
      const response = await fetch(
        `${this.baseUrl}?${queryParams}`,
        {
          method: 'GET',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json',
          },
        }
      );
      
      if (!response.ok) {
        await this.handleError(response);
      }
      
      const data = await response.json();
      return data;
      
    } catch (error) {
      if (error instanceof ApplicationError) {
        throw error;
      }
      throw new ApplicationError(
        ErrorCode.SYS_INTERNAL_ERROR,
        'Failed to fetch memories',
        error.message
      );
    }
  }
  
  // 個別取得
  async getMemoryById(id: number): Promise<Memory> {
    try {
      const response = await fetch(
        `${this.baseUrl}/${id}`,
        {
          method: 'GET',
          credentials: 'include',
        }
      );
      
      if (!response.ok) {
        await this.handleError(response);
      }
      
      const data = await response.json();
      return data.data;
      
    } catch (error) {
      if (error instanceof ApplicationError) {
        throw error;
      }
      throw new ApplicationError(
        ErrorCode.SYS_INTERNAL_ERROR,
        'Failed to fetch memory',
        error.message
      );
    }
  }
  
  // 作成
  async createMemory(memory: MemoryCreateInput): Promise<Memory> {
    try {
      const response = await fetch(
        this.baseUrl,
        {
          method: 'POST',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(memory),
        }
      );
      
      if (!response.ok) {
        await this.handleError(response);
      }
      
      const data = await response.json();
      return data.data;
      
    } catch (error) {
      if (error instanceof ApplicationError) {
        throw error;
      }
      throw new ApplicationError(
        ErrorCode.SYS_INTERNAL_ERROR,
        'Failed to create memory',
        error.message
      );
    }
  }
  
  // 更新
  async updateMemory(id: number, updates: Partial<Memory>): Promise<Memory> {
    try {
      const response = await fetch(
        `${this.baseUrl}/${id}`,
        {
          method: 'PUT',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(updates),
        }
      );
      
      if (!response.ok) {
        await this.handleError(response);
      }
      
      const data = await response.json();
      return data.data;
      
    } catch (error) {
      if (error instanceof ApplicationError) {
        throw error;
      }
      throw new ApplicationError(
        ErrorCode.SYS_INTERNAL_ERROR,
        'Failed to update memory',
        error.message
      );
    }
  }
  
  // 削除
  async deleteMemory(id: number): Promise<void> {
    try {
      const response = await fetch(
        `${this.baseUrl}/${id}`,
        {
          method: 'DELETE',
          credentials: 'include',
        }
      );
      
      if (!response.ok) {
        await this.handleError(response);
      }
      
    } catch (error) {
      if (error instanceof ApplicationError) {
        throw error;
      }
      throw new ApplicationError(
        ErrorCode.SYS_INTERNAL_ERROR,
        'Failed to delete memory',
        error.message
      );
    }
  }
  
  // エラーハンドリング
  private async handleError(response: Response): Promise<never> {
    let errorData: any;
    
    try {
      errorData = await response.json();
    } catch {
      errorData = { message: 'Unknown error' };
    }
    
    switch (response.status) {
      case 400:
        throw new ApplicationError(
          ErrorCode.VAL_INVALID_INPUT,
          errorData.message || 'Invalid input'
        );
      case 401:
        throw new ApplicationError(
          ErrorCode.AUTH_INVALID_CREDENTIALS,
          'Authentication required'
        );
      case 403:
        throw new ApplicationError(
          ErrorCode.AUTH_UNAUTHORIZED,
          'Access denied'
        );
      case 404:
        throw new ApplicationError(
          ErrorCode.RES_NOT_FOUND,
          'Resource not found'
        );
      case 409:
        throw new ApplicationError(
          ErrorCode.RES_CONFLICT,
          errorData.message || 'Resource conflict'
        );
      default:
        throw new ApplicationError(
          ErrorCode.SYS_INTERNAL_ERROR,
          errorData.message || 'Server error'
        );
    }
  }
}

export const memoryService = new MemoryService();
```

## 3. Redux State詳細設計

### 3.1 Memory Slice実装
```typescript
// /features/memory/memorySlice.ts
import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { memoryService } from '../../services/memoryApi';

// State定義
interface MemoryState {
  memories: Memory[];
  memoryItem: Memory | null;
  memoryLoading: boolean;
  memoryError: string | null;
  pagination: {
    page: number;
    limit: number;
    total: number;
    totalPages: number;
  };
  filters: MemoryFilter;
}

const initialState: MemoryState = {
  memories: [],
  memoryItem: null,
  memoryLoading: false,
  memoryError: null,
  pagination: {
    page: 1,
    limit: 20,
    total: 0,
    totalPages: 0,
  },
  filters: {
    source_type: '',
    tags: [],
    read_status: '',
    date_range: null,
  },
};

// Async Thunks
export const loadMemories = createAsyncThunk(
  'memory/loadMemories',
  async (params?: MemoryQueryParams) => {
    const response = await memoryService.getMemories(params);
    return response;
  }
);

export const getMemory = createAsyncThunk(
  'memory/getMemory',
  async (id: number) => {
    const memory = await memoryService.getMemoryById(id);
    return memory;
  }
);

export const createMemory = createAsyncThunk(
  'memory/createMemory',
  async (memory: MemoryCreateInput) => {
    const newMemory = await memoryService.createMemory(memory);
    return newMemory;
  }
);

export const updateMemory = createAsyncThunk(
  'memory/updateMemory',
  async ({ id, updates }: { id: number; updates: Partial<Memory> }) => {
    const updatedMemory = await memoryService.updateMemory(id, updates);
    return updatedMemory;
  }
);

export const removeMemory = createAsyncThunk(
  'memory/removeMemory',
  async (id: number) => {
    await memoryService.deleteMemory(id);
    return id;
  }
);

// Slice定義
const memorySlice = createSlice({
  name: 'memory',
  initialState,
  reducers: {
    setFilters: (state, action: PayloadAction<MemoryFilter>) => {
      state.filters = action.payload;
    },
    clearError: (state) => {
      state.memoryError = null;
    },
    resetMemories: (state) => {
      state.memories = [];
      state.memoryItem = null;
      state.memoryError = null;
      state.pagination = initialState.pagination;
    },
  },
  extraReducers: (builder) => {
    // Load Memories
    builder
      .addCase(loadMemories.pending, (state) => {
        state.memoryLoading = true;
        state.memoryError = null;
      })
      .addCase(loadMemories.fulfilled, (state, action) => {
        state.memoryLoading = false;
        state.memories = action.payload.data;
        state.pagination = action.payload.pagination;
      })
      .addCase(loadMemories.rejected, (state, action) => {
        state.memoryLoading = false;
        state.memoryError = action.error.message || 'Failed to load memories';
      });
    
    // Get Memory
    builder
      .addCase(getMemory.pending, (state) => {
        state.memoryLoading = true;
        state.memoryError = null;
      })
      .addCase(getMemory.fulfilled, (state, action) => {
        state.memoryLoading = false;
        state.memoryItem = action.payload;
      })
      .addCase(getMemory.rejected, (state, action) => {
        state.memoryLoading = false;
        state.memoryError = action.error.message || 'Failed to get memory';
      });
    
    // Create Memory
    builder
      .addCase(createMemory.pending, (state) => {
        state.memoryLoading = true;
        state.memoryError = null;
      })
      .addCase(createMemory.fulfilled, (state, action) => {
        state.memoryLoading = false;
        state.memories.unshift(action.payload);
        state.pagination.total += 1;
      })
      .addCase(createMemory.rejected, (state, action) => {
        state.memoryLoading = false;
        state.memoryError = action.error.message || 'Failed to create memory';
      });
    
    // Update Memory
    builder
      .addCase(updateMemory.pending, (state) => {
        state.memoryLoading = true;
        state.memoryError = null;
      })
      .addCase(updateMemory.fulfilled, (state, action) => {
        state.memoryLoading = false;
        const index = state.memories.findIndex(m => m.id === action.payload.id);
        if (index !== -1) {
          state.memories[index] = action.payload;
        }
        if (state.memoryItem?.id === action.payload.id) {
          state.memoryItem = action.payload;
        }
      })
      .addCase(updateMemory.rejected, (state, action) => {
        state.memoryLoading = false;
        state.memoryError = action.error.message || 'Failed to update memory';
      });
    
    // Remove Memory
    builder
      .addCase(removeMemory.pending, (state) => {
        state.memoryLoading = true;
        state.memoryError = null;
      })
      .addCase(removeMemory.fulfilled, (state, action) => {
        state.memoryLoading = false;
        state.memories = state.memories.filter(m => m.id !== action.payload);
        if (state.memoryItem?.id === action.payload) {
          state.memoryItem = null;
        }
        state.pagination.total -= 1;
      })
      .addCase(removeMemory.rejected, (state, action) => {
        state.memoryLoading = false;
        state.memoryError = action.error.message || 'Failed to delete memory';
      });
  },
});

export const { setFilters, clearError, resetMemories } = memorySlice.actions;
export default memorySlice.reducer;
```

## 4. データベース設計

### 4.1 テーブル定義

#### memories テーブル
```sql
CREATE TABLE memories (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users(id),
  source_type VARCHAR(50) NOT NULL,
  title VARCHAR(200) NOT NULL,
  author VARCHAR(100),
  notes TEXT,
  factor TEXT,
  process TEXT,
  evaluation_axis TEXT,
  information_amount TEXT,
  tags TEXT,
  read_status VARCHAR(20) DEFAULT 'unread',
  read_date TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  INDEX idx_user_id (user_id),
  INDEX idx_source_type (source_type),
  INDEX idx_read_status (read_status),
  INDEX idx_created_at (created_at)
);
```

#### memory_contexts テーブル
```sql
CREATE TABLE memory_contexts (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users(id),
  task_id INTEGER REFERENCES tasks(id),
  code VARCHAR(20) NOT NULL UNIQUE,
  level INTEGER NOT NULL,
  work_target TEXT,
  machine TEXT,
  material_spec TEXT,
  change_factor TEXT,
  goal TEXT,
  technical_factors JSON,
  knowledge_transformations JSON,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  INDEX idx_code (code),
  INDEX idx_user_task (user_id, task_id)
);
```

#### assessments テーブル
```sql
CREATE TABLE assessments (
  id SERIAL PRIMARY KEY,
  task_id INTEGER NOT NULL REFERENCES tasks(id),
  user_id INTEGER NOT NULL REFERENCES users(id),
  effectiveness_score INTEGER NOT NULL CHECK (effectiveness_score >= 0 AND effectiveness_score <= 100),
  effort_score INTEGER NOT NULL CHECK (effort_score >= 0 AND effort_score <= 100),
  impact_score INTEGER NOT NULL CHECK (impact_score >= 0 AND impact_score <= 100),
  qualitative_feedback TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  INDEX idx_task_user (task_id, user_id),
  INDEX idx_created_at (created_at)
);
```

### 4.2 インデックス戦略
```sql
-- 複合インデックス
CREATE INDEX idx_memory_user_status_date 
ON memories(user_id, read_status, created_at DESC);

CREATE INDEX idx_assessment_scores 
ON assessments(effectiveness_score, effort_score, impact_score);

-- 全文検索インデックス
CREATE FULLTEXT INDEX idx_memory_search 
ON memories(title, notes, tags);

-- JSONパスインデックス
CREATE INDEX idx_technical_factors_domain 
ON memory_contexts((technical_factors->>'$.domain'));
```

## 5. パフォーマンス最適化詳細

### 5.1 フロントエンド最適化

#### メモ化戦略
```typescript
// 高コストな計算のメモ化
const memoizedFilteredMemories = useMemo(() => {
  return filterMemories(memories);
}, [memories, filterCriteria]);

const memoizedSortedMemories = useMemo(() => {
  return sortMemories(memoizedFilteredMemories);
}, [memoizedFilteredMemories, sortBy, sortOrder]);

// コールバックのメモ化
const handleMemoryUpdate = useCallback(
  async (id: number, updates: Partial<Memory>) => {
    await dispatch(updateMemory({ id, updates }));
  },
  [dispatch]
);
```

#### 仮想スクロール実装
```typescript
import { FixedSizeList } from 'react-window';

const VirtualMemoryList: React.FC<{ memories: Memory[] }> = ({ memories }) => {
  const Row = ({ index, style }) => (
    <div style={style}>
      <MemoryCard memory={memories[index]} />
    </div>
  );
  
  return (
    <FixedSizeList
      height={600}
      itemCount={memories.length}
      itemSize={120}
      width="100%"
    >
      {Row}
    </FixedSizeList>
  );
};
```

#### デバウンス実装
```typescript
const useDebounce = <T,>(value: T, delay: number): T => {
  const [debouncedValue, setDebouncedValue] = useState(value);
  
  useEffect(() => {
    const handler = setTimeout(() => {
      setDebouncedValue(value);
    }, delay);
    
    return () => {
      clearTimeout(handler);
    };
  }, [value, delay]);
  
  return debouncedValue;
};

// 使用例
const SearchMemories: React.FC = () => {
  const [searchTerm, setSearchTerm] = useState('');
  const debouncedSearchTerm = useDebounce(searchTerm, 300);
  
  useEffect(() => {
    if (debouncedSearchTerm) {
      dispatch(loadMemories({ search: debouncedSearchTerm }));
    }
  }, [debouncedSearchTerm]);
};
```

### 5.2 バックエンド最適化

#### クエリ最適化
```go
// N+1問題の回避
func GetMemoriesWithRelations(userID int) ([]Memory, error) {
    var memories []Memory
    
    // JOINを使用して関連データを一度に取得
    query := `
        SELECT 
            m.*,
            t.id as task_id,
            t.title as task_title,
            a.id as assessment_id,
            a.effectiveness_score,
            a.effort_score,
            a.impact_score
        FROM memories m
        LEFT JOIN tasks t ON m.id = t.memory_id
        LEFT JOIN assessments a ON t.id = a.task_id
        WHERE m.user_id = ?
        ORDER BY m.created_at DESC
    `
    
    rows, err := db.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    // 結果のマッピング
    memoryMap := make(map[int]*Memory)
    for rows.Next() {
        // データマッピング処理
    }
    
    return memories, nil
}
```

#### キャッシング戦略
```go
import "github.com/go-redis/redis/v8"

type MemoryCache struct {
    client *redis.Client
}

func (c *MemoryCache) GetMemory(id int) (*Memory, error) {
    key := fmt.Sprintf("memory:%d", id)
    
    // キャッシュから取得
    val, err := c.client.Get(ctx, key).Result()
    if err == nil {
        var memory Memory
        json.Unmarshal([]byte(val), &memory)
        return &memory, nil
    }
    
    // DBから取得
    memory, err := db.GetMemory(id)
    if err != nil {
        return nil, err
    }
    
    // キャッシュに保存（TTL: 1時間）
    data, _ := json.Marshal(memory)
    c.client.Set(ctx, key, data, time.Hour)
    
    return memory, nil
}

func (c *MemoryCache) InvalidateMemory(id int) {
    key := fmt.Sprintf("memory:%d", id)
    c.client.Del(ctx, key)
}
```

## 6. セキュリティ実装詳細

### 6.1 入力検証
```typescript
// XSS対策
import DOMPurify from 'dompurify';

const sanitizeInput = (input: string): string => {
  return DOMPurify.sanitize(input, {
    ALLOWED_TAGS: ['b', 'i', 'em', 'strong', 'a'],
    ALLOWED_ATTR: ['href'],
  });
};

// SQLインジェクション対策（パラメータ化クエリ）
const safeQuery = async (userId: number, searchTerm: string) => {
  const query = 'SELECT * FROM memories WHERE user_id = $1 AND title ILIKE $2';
  const values = [userId, `%${searchTerm}%`];
  
  return await db.query(query, values);
};
```

### 6.2 認証・認可
```typescript
// JWTトークン検証
import { verify } from 'jsonwebtoken';

const verifyToken = (token: string): TokenPayload => {
  try {
    const payload = verify(token, process.env.JWT_SECRET!) as TokenPayload;
    
    // 有効期限チェック
    if (payload.exp < Date.now() / 1000) {
      throw new Error('Token expired');
    }
    
    return payload;
  } catch (error) {
    throw new ApplicationError(
      ErrorCode.AUTH_INVALID_CREDENTIALS,
      'Invalid token'
    )
  }
};

// リソースアクセス権限チェック
const checkResourceAccess = async (
  userId: number,
  resourceId: number,
  resourceType: 'memory' | 'assessment'
): Promise<boolean> => {
  switch (resourceType) {
    case 'memory':
      const memory = await db.getMemory(resourceId);
      return memory.user_id === userId;
    
    case 'assessment':
      const assessment = await db.getAssessment(resourceId);
      return assessment.user_id === userId;
    
    default:
      return false;
  }
};
```

## 7. エラーハンドリング詳細

### 7.1 グローバルエラーハンドラー
```typescript
// エラーバウンダリコンポーネント
class ErrorBoundary extends React.Component<
  { children: React.ReactNode },
  { hasError: boolean; error: Error | null }
> {
  state = { hasError: false, error: null };
  
  static getDerivedStateFromError(error: Error) {
    return { hasError: true, error };
  }
  
  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    // エラーログ送信
    logErrorToService(error, errorInfo);
  }
  
  render() {
    if (this.state.hasError) {
      return (
        <ErrorFallback
          error={this.state.error}
          resetError={() => this.setState({ hasError: false, error: null })}
        />
      );
    }
    
    return this.props.children;
  }
}
```

### 7.2 非同期エラーハンドリング
```typescript
// Redux Thunkでのエラーハンドリング
const handleAsyncError = createAsyncThunk(
  'memory/operation',
  async (params, { rejectWithValue }) => {
    try {
      const result = await apiCall(params);
      return result;
    } catch (error) {
      // エラーの分類と処理
      if (error instanceof ApplicationError) {
        // アプリケーションエラー
        return rejectWithValue({
          code: error.code,
          message: error.message,
          detail: error.detail,
        });
      } else if (error instanceof NetworkError) {
        // ネットワークエラー
        return rejectWithValue({
          code: 'NETWORK_ERROR',
          message: 'ネットワーク接続を確認してください',
        });
      } else {
        // 予期しないエラー
        return rejectWithValue({
          code: 'UNKNOWN_ERROR',
          message: 'エラーが発生しました',
        });
      }
    }
  }
);
```

## 8. テスト実装詳細

### 8.1 単体テスト
```typescript
// Service層のテスト
describe('MemoryService', () => {
  let service: MemoryService;
  
  beforeEach(() => {
    service = new MemoryService();
    fetchMock.resetMocks();
  });
  
  describe('getMemories', () => {
    it('should fetch memories successfully', async () => {
      const mockMemories = [
        { id: 1, title: 'Test Memory 1' },
        { id: 2, title: 'Test Memory 2' },
      ];
      
      fetchMock.mockResponseOnce(
        JSON.stringify({ data: mockMemories })
      );
      
      const result = await service.getMemories();
      
      expect(result.data).toEqual(mockMemories);
      expect(fetchMock).toHaveBeenCalledWith(
        '/api/memory',
        expect.objectContaining({
          method: 'GET',
          credentials: 'include',
        })
      );
    });
    
    it('should handle errors properly', async () => {
      fetchMock.mockRejectOnce(new Error('Network error'));
      
      await expect(service.getMemories()).rejects.toThrow(
        ApplicationError
      );
    });
  });
});
```

### 8.2 統合テスト
```typescript
// Redux統合テスト
describe('Memory Redux Integration', () => {
  let store: AppStore;
  
  beforeEach(() => {
    store = configureStore({
      reducer: {
        memory: memoryReducer,
      },
    });
  });
  
  it('should handle memory CRUD flow', async () => {
    // Create
    const newMemory = {
      title: 'Test Memory',
      source_type: 'book',
    };
    
    await store.dispatch(createMemory(newMemory));
    expect(store.getState().memory.memories).toHaveLength(1);
    
    // Update
    const updates = { title: 'Updated Memory' };
    await store.dispatch(
      updateMemory({ id: 1, updates })
    );
    
    expect(store.getState().memory.memories[0].title).toBe(
      'Updated Memory'
    );
    
    // Delete
    await store.dispatch(removeMemory(1));
    expect(store.getState().memory.memories).toHaveLength(0);
  });
});
```

### 8.3 E2Eテスト
```typescript
// Cypress E2Eテスト
describe('Memory Management E2E', () => {
  beforeEach(() => {
    cy.login('testuser@example.com', 'password');
    cy.visit('/memories');
  });
  
  it('should complete memory creation flow', () => {
    // 新規作成ボタンクリック
    cy.get('[data-testid="create-memory-btn"]').click();
    
    // フォーム入力
    cy.get('[data-testid="memory-title"]').type('E2E Test Memory');
    cy.get('[data-testid="memory-source-type"]').select('book');
    cy.get('[data-testid="memory-author"]').type('Test Author');
    cy.get('[data-testid="memory-notes"]').type('Test notes content');
    
    // 保存
    cy.get('[data-testid="save-memory-btn"]').click();
    
    // 確認
    cy.get('[data-testid="memory-list"]')
      .should('contain', 'E2E Test Memory')
      .should('contain', 'Test Author');
    
    // 通知確認
    cy.get('[data-testid="success-notification"]')
      .should('be.visible')
      .should('contain', 'メモリを作成しました');
  });
});
```

---

最終更新日: 2025-08-31
バージョン: 1.0.0