import React, { useState, useEffect } from 'react';
import { AddTask, Task } from "../../model/task";
import { Memory } from "../../model/memory";
import CommonModal from "./CommonModal";

interface TaskModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (taskData: AddTask | Task) => void;
  initialData?: AddTask | Task;
  memories: Memory[];
}

const TaskModal: React.FC<TaskModalProps> = ({ isOpen, onClose, onSave, initialData, memories }) => {
  const [formData, setFormData] = useState<AddTask | Task | undefined>();

  useEffect(() => {
    if (initialData) {
      setFormData(initialData);
    } else {
      setFormData({
        user_id: 0,
        title: '',
        description: '',
        status: 'todo',
        priority: 1,
      });
    }
  }, [initialData]);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>
  ) => {
    const { name, value } = e.target;
    setFormData(prev => prev ? { ...prev, [name]: value } : prev);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!formData) return;
    await onSave(formData);
  };

  // 選択中のmemory_idに該当するMemoryを取得
  const selectedMemory = memories.find(m => m.id === Number(formData?.memory_id));

  if (!isOpen) return null;

  return (
    <CommonModal isOpen={isOpen} onClose={onClose} title={initialData ? 'タスクを編集' : '新規タスク'}>
      <form onSubmit={handleSubmit} className="task-form">
        <div className="form-group">
          <label htmlFor="title">タイトル</label>
          <input
            type="text"
            id="title"
            name="title"
            value={formData?.title || ''}
            onChange={handleChange}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="description">説明</label>
          <textarea
            id="description"
            name="description"
            value={formData?.description || ''}
            onChange={handleChange}
            rows={3}
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="status">ステータス</label>
          <select
            id="status"
            name="status"
            value={formData?.status || 'todo'}
            onChange={handleChange}
          >
            <option value="todo">未着手</option>
            <option value="in_progress">進行中</option>
            <option value="completed">完了</option>
          </select>
        </div>
        <div className="form-group">
          <label htmlFor="priority">優先度</label>
          <input
            type="number"
            id="priority"
            name="priority"
            value={formData?.priority || 1}
            onChange={handleChange}
            min={1}
            max={5}
          />
        </div>
        <div className="form-group">
          <label htmlFor="memory_id">関連メモ</label>
          <select
            id="memory_id"
            name="memory_id"
            value={formData?.memory_id ?? ''}
            onChange={handleChange}
          >
            <option value="">選択してください</option>
            {memories.map(memory => (
              <option key={memory.id} value={memory.id}>
                {memory.title}
              </option>
            ))}
          </select>
        </div>
        {/* 選択中のメモ内容を表示 */}
        {selectedMemory && (
          <div className="selected-memory-info" style={{margin: '1em 0', padding: '0.5em', background: '#f6f8fa', borderRadius: 6}}>
            <div><b>タイトル:</b> {selectedMemory.title}</div>
            <div><b>内容:</b> {selectedMemory.notes}</div>
            <div><b>タグ:</b> {selectedMemory.tags}</div>
            <div><b>ステータス:</b> {selectedMemory.read_status}</div>
          </div>
        )}
        <div className="form-actions">
          <button type="button" onClick={onClose} className="btn btn-secondary">キャンセル</button>
          <button type="submit" className="btn btn-primary">保存</button>
        </div>
      </form>
    </CommonModal>
  );
};

export default TaskModal;
