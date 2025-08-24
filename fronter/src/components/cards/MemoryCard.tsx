"use client"
import React from 'react';
import GenericItemCard from '../parts/GenericItemCard';
import { Memory } from '../../model/memory';
import { useItemOperations } from '../../hooks/useItemOperations';

interface MemoryCardProps {
  memory: Memory;
  onRefresh?: () => void;
}

export default function MemoryCard({ memory, onRefresh }: MemoryCardProps) {
  const { deleteItem, updateItem } = useItemOperations('memory', {
    onDeleteSuccess: onRefresh,
    onUpdateSuccess: onRefresh
  });

  const handleUpdate = async (item: Memory) => {
    await updateItem(item);
  };

  const renderContent = (memory: Memory) => (
    <div className="memory-details">
      <div className="detail-item">
        <span className="label">タグ:</span>
        <span className="tags">{memory.tags}</span>
      </div>
      <div className="detail-item">
        <span className="label">ステータス:</span>
        <span className={`read-status status-${memory.read_status}`}>
          {memory.read_status === 'unread' && '未読'}
          {memory.read_status === 'read' && '既読'}
          {memory.read_status === 'archived' && 'アーカイブ'}
        </span>
      </div>
      {memory.source_type && (
        <div className="detail-item">
          <span className="label">ソースタイプ:</span>
          <span className="source-type">{memory.source_type}</span>
        </div>
      )}
    </div>
  );

  const renderEditForm = (memory: Memory, onChange: (memory: Memory) => void) => (
    <form className="memory-edit-form">
      <div className="form-group">
        <label htmlFor="title">タイトル</label>
        <input
          type="text"
          id="title"
          value={memory.title}
          onChange={(e) => onChange({ ...memory, title: e.target.value })}
        />
      </div>
      <div className="form-group">
        <label htmlFor="notes">メモ</label>
        <textarea
          id="notes"
          value={memory.notes}
          onChange={(e) => onChange({ ...memory, notes: e.target.value })}
          rows={5}
        />
      </div>
      <div className="form-group">
        <label htmlFor="tags">タグ</label>
        <input
          type="text"
          id="tags"
          value={memory.tags}
          onChange={(e) => onChange({ ...memory, tags: e.target.value })}
          placeholder="カンマ区切りでタグを入力"
        />
      </div>
      <div className="form-group">
        <label htmlFor="source_type">ソースタイプ</label>
        <input
          type="text"
          id="source_type"
          value={memory.source_type || ''}
          onChange={(e) => onChange({ ...memory, source_type: e.target.value })}
        />
      </div>
      <div className="form-group">
        <label htmlFor="read_status">ステータス</label>
        <select
          id="read_status"
          value={memory.read_status}
          onChange={(e) => onChange({ ...memory, read_status: e.target.value as Memory['read_status'] })}
        >
          <option value="unread">未読</option>
          <option value="read">既読</option>
          <option value="archived">アーカイブ</option>
        </select>
      </div>
    </form>
  );

  return (
    <GenericItemCard
      item={memory}
      itemType="memory"
      onDelete={deleteItem}
      onUpdate={handleUpdate}
      renderContent={renderContent}
      renderEditForm={renderEditForm}
      getTitle={(memory) => memory.title}
      getDescription={(memory) => memory.notes}
      getId={(memory) => memory.id}
    />
  );
}