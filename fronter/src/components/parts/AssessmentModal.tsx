import React, { useState, useEffect } from 'react';
import { AddAssessment, Assessment } from "../../model/assessment";
import Cookies from 'js-cookie';
import { Memory } from "../../model/memory";
import { Task } from "../../model/task";
import CommonModal from './CommonModal';


interface AssessmentModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (assessmentData: AddAssessment | Assessment) => void;
  initialData?: AddAssessment | Assessment;
  tasks: Task[];
  memories?: Memory[];
}

const setCheker = ['user_id', 'task_id', 'effectiveness_score', 'effort_score', 'impact_score'];
const AssessmentModal: React.FC<AssessmentModalProps> = ({ isOpen, onClose, onSave, initialData, tasks, memories }) => {
  const [formData, setFormData] = useState<AddAssessment | Assessment | undefined>();
  // メモリー詳細表示用のstate
  const [openMemoryId, setOpenMemoryId] = useState<number | null>(null);

  useEffect(() => {
    if (initialData) {
      setFormData(initialData);
    } else {
      setFormData({
        task_id: 0,
        user_id: 0,
        effectiveness_score: 0,
        effort_score: 0,
        impact_score: 0,
        qualitative_feedback: '',
      });
    }
  }, [initialData]);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>
  ) => {
    const { name, value } = e.target;
    setFormData(prev => prev ? {
      ...prev,
      [name]: setCheker.includes(name) ? Number(value) : value
    } : prev);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!formData) return;
    // cookieからuser_id取得
    const userIdStr = Cookies.get('user_id');
    const user_id = userIdStr ? Number(userIdStr) : 0;
    await onSave({ ...formData, user_id });
  };

  // 選択中のタスクに紐づくメモリーを取得
  const selectedTask = (tasks ?? []).find(t => t.id === Number(formData?.task_id));
  const relatedMemory = (memories ?? []).find(m => m.id === selectedTask?.memory_id);

  return (
    <CommonModal
      isOpen={isOpen}
      onClose={onClose}
      title={initialData ? 'アセスメントを編集' : '新規アセスメント'}
    >
      <div className="assessment-modal-content">
        {/* 全メモリー一覧を表示（タイトルクリックで詳細トグル） */}
        {memories && memories.length > 0 && (
          <div className="all-memories-list" style={{margin: '1em 0', padding: '0.5em', background: '#f0f4fa', borderRadius: 6}}>
            <div style={{fontWeight: 'bold', marginBottom: 4}}>全メモリー一覧</div>
            <ul style={{margin: 0, padding: 0, listStyle: 'none'}}>
              {memories.map(memory => (
                <li key={memory.id} style={{marginBottom: 6, borderBottom: '1px solid #e0e0e0', paddingBottom: 4}}>
                  <div>
                    <span
                      style={{ cursor: 'pointer', color: '#2563eb', textDecoration: 'underline' }}
                      onClick={() => setOpenMemoryId(openMemoryId === memory.id ? null : memory.id)}
                    >
                      {memory.title}
                    </span>
                  </div>
                  {openMemoryId === memory.id && (
                    <div style={{marginTop: 4, paddingLeft: 8}}>
                      <div><b>内容:</b> {memory.notes}</div>
                      <div><b>タグ:</b> {memory.tags}</div>
                      <div><b>ステータス:</b> {memory.read_status}</div>
                    </div>
                  )}
                </li>
              ))}
            </ul>
          </div>
        )}
        {/* 紐づくメモリー情報を表示 */}
        {relatedMemory && (
          <div className="related-memory-info" style={{margin: '1em 0', padding: '0.5em', background: '#f6f8fa', borderRadius: 6}}>
            <div><b>関連メモ:</b> {relatedMemory.title}</div>
            <div><b>内容:</b> {relatedMemory.notes}</div>
            <div><b>タグ:</b> {relatedMemory.tags}</div>
            <div><b>ステータス:</b> {relatedMemory.read_status}</div>
          </div>
        )}
        <form onSubmit={handleSubmit} className="assessment-form">
          <div className="form-group">
            <label htmlFor="task_id">タスクID</label>
            <input
              type="number"
              id="task_id"
              name="task_id"
              value={formData?.task_id || 0}
              onChange={handleChange}
              required
            />
          </div>
          {/* ユーザーID入力欄は削除 */}
          <div className="form-group">
            <label htmlFor="effectiveness_score">効果スコア</label>
            <input
              type="number"
              id="effectiveness_score"
              name="effectiveness_score"
              value={formData?.effectiveness_score || 0}
              onChange={handleChange}
              min={0}
              max={100}
            />
          </div>
          <div className="form-group">
            <label htmlFor="effort_score">努力スコア</label>
            <input
              type="number"
              id="effort_score"
              name="effort_score"
              value={formData?.effort_score || 0}
              onChange={handleChange}
              min={0}
              max={100}
            />
          </div>
          <div className="form-group">
            <label htmlFor="impact_score">インパクトスコア</label>
            <input
              type="number"
              id="impact_score"
              name="impact_score"
              value={formData?.impact_score || 0}
              onChange={handleChange}
              min={0}
              max={100}
            />
          </div>
          <div className="form-group">
            <label htmlFor="qualitative_feedback">定性的フィードバック</label>
            <textarea
              id="qualitative_feedback"
              name="qualitative_feedback"
              value={formData?.qualitative_feedback || ''}
              onChange={handleChange}
              rows={3}
            />
          </div>
          <div className="form-actions">
            <button type="button" onClick={onClose} className="btn btn-secondary">キャンセル</button>
            <button type="submit" className="btn btn-primary">保存</button>
          </div>
        </form>
      </div>
    </CommonModal>
  );
};

export default AssessmentModal;
