import React, { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../../store';
import CommonModal from './CommonModal';
import { AddMemory, Memory } from "../../model/memory";

interface MemoryModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (memoryData: (AddMemory|Memory)) => void;
  initialData?: (AddMemory|Memory);
}

const MemoryModal: React.FC<MemoryModalProps> = ({
  isOpen,
  onClose,
  onSave,
  initialData,
}) => {
  const [formData, setFormData] = useState<(AddMemory|Memory|undefined)>();
  const { loading, error } = useSelector((state: RootState) => state.memory);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>
  ) => {
    const { name, value } = e.target;
    if (formData) {
      setFormData(prev =>{
        if (prev) {
          return {
            ...prev,
            [name]: value,
          };
        }
        return prev;
      });
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!formData) return;
    await onSave(formData);
  };

  useEffect(() => {
    if (initialData) {
      setFormData(initialData);
    }
  }, [initialData]);

  return (
    <CommonModal
      isOpen={isOpen}
      onClose={onClose}
      title={initialData?.title ? 'メモを編集' : '新規メモ'}
    >
      <form onSubmit={handleSubmit} className="memory-form">
        {error && <div className="error-message">{error}</div>}
        { formData && <div>
          <div className="form-group">
            <label htmlFor="title">タイトル</label>
            <input
              type="text"
              id="title"
              name="title"
              value={formData.title}
              onChange={handleChange}
              required
            />
          </div>

          <div className="form-group">
            <label htmlFor="notes">メモ内容</label>
            <textarea
              id="notes"
              name="notes"
              value={formData.notes}
              onChange={handleChange}
              rows={5}
              required
            />
          </div>

          <div className="form-group">
            <label htmlFor="tags">タグ（カンマ区切り）</label>
            <input
              type="text"
              id="tags"
              name="tags"
              value={formData.tags}
              onChange={handleChange}
              placeholder="例: 仕事, 重要, 後で"
            />
          </div>

          <div className="form-group">
            <label htmlFor="read_status">ステータス</label>
            <select
              id="read_status"
              name="read_status"
              value={formData.read_status}
              onChange={handleChange}
            >
              <option value="unread">未読</option>
              <option value="reading">読書中</option>
              <option value="completed">完了</option>
            </select>
          </div>

          <div className="form-actions">
            <button
              type="button"
              onClick={onClose}
              className="btn btn-secondary"
            >
              キャンセル
            </button>
            <button
              type="submit"
              className="btn btn-primary"
              disabled={loading}
            >
              {loading ? '保存中...' : '保存'}
            </button>
          </div>
        </div>}
      </form>
    </CommonModal>
  );
};

export default MemoryModal; 