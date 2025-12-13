import React, { useState, useEffect } from 'react';
import { useApiCall } from '../../hooks/useApiCall';
import { addTaskClient } from '../../client/taskApi';
import { AddTask, Task } from "../../model/task";
import { Memory } from "../../model/memory";
import CommonModal from './CommonModal';
import { formatDateTime } from "../../utils/dayApi";

interface TaskModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (taskData: AddTask | Task) => void;
  initialData?: AddTask | Task;
  memories: Memory[];
}

const TaskModal: React.FC<TaskModalProps> = ({ isOpen, onClose, onSave, initialData, memories }) => {
  const [formData, setFormData] = useState<AddTask | Task | undefined>();

  const { execute: saveTask } = useApiCall(
    addTaskClient,
    {
      onSuccess: () => {
        onClose();
      }
    }
  );

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
    await saveTask(formData);
  };

  const isTask = (data: AddTask | Task | undefined): data is Task => {
    console.log("isTask check", data);
    return data !== undefined && "id" in data; 
  };

  // 選択中のmemory_idに該当するMemoryを取得
  const selectedMemory = memories.find(m => m.id === Number(formData?.memory_id));

  return (
    <CommonModal
      isOpen={isOpen}
      onClose={onClose}
      title={initialData ? 'タスクを編集' : '新規タスク'}
    >
      <div className="modal-wrapper">
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
          <div className="form-group p-16">
            {isTask(formData) && formData.heuristics_analysis && formData.heuristics_analysis.length > 0 && (
              <div className="heuristics-analysis-section">
                <h4>ヒューリスティクス分析</h4>
                <ul>
                  {formData.heuristics_analysis.map(analysis => (
                    <li key={analysis.id}>
                      <p>analysis_type: {analysis.analysis_type}</p> 
                      <div className="result p-16">
                        <p className="p-b-8">confidence: {analysis.result.confidence}</p>
                        <p className="p-b-8">energy_saving: {analysis.result.energy_saving}</p>
                        <p>created_at: {formatDateTime(analysis.created_at,"YYYY/MM/DD HH:mm:ss")}</p> 
                      </div>
                    </li>
                  ))}
                </ul>
              </div>
            )}
          </div>
          <div className="form-group p-16">
            {isTask(formData) && formData.heuristics_patterns && formData.heuristics_patterns.length > 0 && (
              <div className="heuristics-analysis-section">
                <h4>ヒューリスティクス HeuristicsPattern</h4>
                <ul>
                  {formData.heuristics_patterns.map(analysis => (
                    <li key={analysis.id}>
                      <p>category: {analysis.category}</p> 
                      <div className="result p-16">
                        <p className="p-b-8">pattern: {analysis.pattern}</p>
                        <p className="p-b-8">frequency: {analysis.frequency}</p>
                        <p>created_at: {formatDateTime(analysis.created_at,"YYYY/MM/DD HH:mm:ss")}</p> 
                      </div>
                    </li>
                  ))}
                </ul>
              </div>
            )}
          </div>
          <div className="form-group p-16">
            {isTask(formData) && formData.language_optimizations && formData.language_optimizations.length > 0 && (
              <div className="heuristics-analysis-section">
                <h4>LanguageOptimization</h4>
                <ul>
                  {formData.language_optimizations.map(analysis => (
                    <li key={analysis.id}>
                      <div className="result p-16">
                        <p className="p-b-8">domain: {analysis.domain}</p>
                        <p className="p-b-8">optimized_text: {analysis.optimized_text}</p>
                        <p className="p-b-8">original_text: {analysis.original_text}</p> 
                        <p>created_at: {formatDateTime(analysis.created_at,"YYYY/MM/DD HH:mm:ss")}</p> 
                      </div>
                    </li>
                  ))}
                </ul> 
              </div>
            )}
          </div>
          {/* 選択中のメモ内容を表示 */}
          {selectedMemory && (
            <div className="selected-memory-info p-8" style={{margin: '1em 0', padding: '0.5em', background: '#f6f8fa', borderRadius: 6}}>
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
      </div>
    </CommonModal>
  );
};

export default TaskModal;
