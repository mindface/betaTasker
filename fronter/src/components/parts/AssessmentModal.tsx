import React, { useState, useEffect } from 'react';
import { AddAssessment, Assessment } from "../../model/assessment";

interface AssessmentModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (assessmentData: AddAssessment | Assessment) => void;
  initialData?: AddAssessment | Assessment;
}

const AssessmentModal: React.FC<AssessmentModalProps> = ({ isOpen, onClose, onSave, initialData }) => {
  const [formData, setFormData] = useState<AddAssessment | Assessment | undefined>();

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
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
      setFormData(prev => prev ? {
        ...prev,
        [name]: ['user_id', 'task_id', 'effectiveness_score', 'effort_score', 'impact_score'].includes(name)
          ? Number(value)
          : value
      } : prev);
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
        <h2>{initialData ? 'アセスメントを編集' : '新規アセスメント'}</h2>
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
          <div className="form-group">
            <label htmlFor="user_id">ユーザーID</label>
            <input
              type="number"
              id="user_id"
              name="user_id"
              value={formData?.user_id || 0}
              onChange={handleChange}
              required
            />
          </div>
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
    </div>
  );
};

export default AssessmentModal;
