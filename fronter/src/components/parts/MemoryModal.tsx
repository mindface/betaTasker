import React, { useState, useEffect } from 'react';
import CommonDialog from "./CommonDialog";
import { useSelector } from 'react-redux';
import { RootState } from '../../store';
import { AddMemory, Memory } from "../../model/memory";

interface MemoryModalProps {
  isOpen: boolean;
  isViewType: boolean;
  onClose: () => void;
  onSave: (memoryData: (AddMemory|Memory)) => void;
  initialData?: (AddMemory|Memory);
}

const initiaSetlData = {
    title: '',
    notes: '',
    tags: '',
    read_status: 'unread',
    factor: '',
    process: '',
    evaluation_axis: '',
    information_amount: '',
  }

const MemoryModal = ({
  isOpen,
  isViewType,
  onClose,
  onSave,
  initialData,
}: MemoryModalProps) => {
  const [formData, setFormData] = useState<(AddMemory|Memory|undefined)>(initiaSetlData as AddMemory);
  const { memoryLoading, memoryError } = useSelector((state: RootState) => state.memory);

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
    }else {
      setFormData(initiaSetlData as AddMemory);
    }
  }, [initialData]);

  return (
    <CommonDialog
      isOpen={isOpen}
      onClose={onClose}
      title={initialData?.title ? 'メモを編集' : '新規メモ'}
    >
      <form onSubmit={handleSubmit} className="memory-form card-form">
        {memoryError && <div className="error-message">{memoryError}</div>}
        <div>
          <div className="form-group">
            <label htmlFor="title">タイトル</label>
            <input
              type="text"
              id="title"
              name="title"
              value={formData?.title || ''}
              onChange={handleChange}
              disabled={isViewType}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="notes">メモ内容</label>
            <textarea
              id="notes"
              name="notes"
              value={formData?.notes || ''}
              onChange={handleChange}
              rows={5}
              disabled={isViewType}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="tags">タグ（カンマ区切り）</label>
            <input
              type="text"
              id="tags"
              name="tags"
              value={formData?.tags || ''}
              onChange={handleChange}
              placeholder="例: 仕事, 重要, 後で"
              disabled={isViewType}
            />
          </div>
          <div className="form-group">
            <label htmlFor="read_status">ステータス</label>
            <select
              id="read_status"
              name="read_status"
              value={formData?.read_status || ''}
              onChange={handleChange}
              disabled={isViewType}
            >
              <option value="unread">未読</option>
              <option value="reading">読書中</option>
              <option value="completed">完了</option>
            </select>
          </div>
          <div className="form-group">
            <label htmlFor="factor">因子</label>
            <input
              type="text"
              id="factor"
              name="factor"
              value={formData?.factor || ''}
              onChange={handleChange}
              placeholder="例: 環境, 動機, 習慣 など"
              disabled={isViewType}
            />
          </div>
          <div className="form-group">
            <label htmlFor="process">プロセス仮説</label>
            <input
              type="text"
              id="process"
              name="process"
              value={formData?.process || ''}
              onChange={handleChange}
              placeholder="例: どのように学んだか・使ったか"
              disabled={isViewType}
            />
          </div>
          <div className="form-group">
            <label htmlFor="evaluation_axis">評価軸</label>
            <input
              type="text"
              id="evaluation_axis"
              name="evaluation_axis"
              value={formData?.evaluation_axis || ''}
              onChange={handleChange}
              placeholder="例: 理解度, 応用度, 継続性 など"
              disabled={isViewType}
            />
          </div>
          <div className="form-group">
            <label htmlFor="information_amount">情報量</label>
            <input
              type="text"
              id="information_amount"
              name="information_amount"
              value={formData?.information_amount || ''}
              onChange={handleChange}
              placeholder="例: 参考文献数やメモの分量など"
              disabled={isViewType}
            />
          </div>

          { !isViewType &&
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
                disabled={memoryLoading}
              >
                {memoryLoading ? '保存中...' : '保存'}
              </button>
            </div>
          }
        </div>
      </form>
    </CommonDialog>
  );
};

export default MemoryModal;