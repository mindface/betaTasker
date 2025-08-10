"use client"
import React from 'react';
import GenericItemCard from '../parts/GenericItemCard';
import { Task } from '../../model/task';
import { useItemOperations } from '../../hooks/useItemOperations';

interface TaskCardProps {
  task: Task;
  onRefresh?: () => void;
}

export default function TaskCard({ task, onRefresh }: TaskCardProps) {
  const { deleteItem, updateItem } = useItemOperations('task', {
    onDeleteSuccess: onRefresh,
    onUpdateSuccess: onRefresh
  });

  const handleUpdate = async (item: Task) => {
    await updateItem(item);
  };

  const renderContent = (task: Task) => (
    <div className="task-details">
      <div className="detail-item">
        <span className="label">ステータス:</span>
        <span className={`status status-${task.status}`}>
          {task.status === 'todo' && '未着手'}
          {task.status === 'in_progress' && '進行中'}
          {task.status === 'completed' && '完了'}
        </span>
      </div>
      <div className="detail-item">
        <span className="label">優先度:</span>
        <span className="priority">{task.priority}</span>
      </div>
      {task.date && (
        <div className="detail-item">
          <span className="label">期限:</span>
          <span className="due-date">{new Date(task.date).toLocaleDateString('ja-JP')}</span>
        </div>
      )}
    </div>
  );

  const renderEditForm = (task: Task, onChange: (task: Task) => void) => (
    <form className="task-edit-form">
      <div className="form-group">
        <label htmlFor="title">タイトル</label>
        <input
          type="text"
          id="title"
          value={task.title}
          onChange={(e) => onChange({ ...task, title: e.target.value })}
        />
      </div>
      <div className="form-group">
        <label htmlFor="description">説明</label>
        <textarea
          id="description"
          value={task.description}
          onChange={(e) => onChange({ ...task, description: e.target.value })}
          rows={3}
        />
      </div>
      <div className="form-group">
        <label htmlFor="status">ステータス</label>
        <select
          id="status"
          value={task.status}
          onChange={(e) => onChange({ ...task, status: e.target.value as Task['status'] })}
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
          value={task.priority}
          onChange={(e) => onChange({ ...task, priority: Number(e.target.value) })}
          min={1}
          max={5}
        />
      </div>
    </form>
  );

  return (
    <GenericItemCard
      item={task}
      itemType="task"
      onDelete={deleteItem}
      onUpdate={handleUpdate}
      renderContent={renderContent}
      renderEditForm={renderEditForm}
      getTitle={(task) => task.title}
      getDescription={(task) => task.description}
      getId={(task) => task.id}
    />
  );
}