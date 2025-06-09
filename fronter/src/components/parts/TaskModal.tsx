import React, { useState, useEffect } from 'react';
import { AddTask, Task } from "../../model/task";

interface TaskModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (taskData: AddTask | Task) => void;
  initialData?: AddTask | Task;
}

const TaskModal: React.FC<TaskModalProps> = ({ isOpen, onClose, onSave, initialData }) => {
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

  if (!isOpen) return null;

  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <h2>{initialData ? 'タスクを編集' : '新規タスク'}</h2>
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
          <div className="form-actions">
            <button type="button" onClick={onClose} className="btn btn-secondary">キャンセル</button>
            <button type="submit" className="btn btn-primary">保存</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default TaskModal;
