"use client"
import React from 'react'
import { Task } from "../../model/task";

interface ItemTaskProps {
  task: Task;
  onEdit: (task: Task) => void;
  onDelete: (id: number) => void;
}

const ItemTask: React.FC<ItemTaskProps> = ({ task, onEdit, onDelete }) => {
  return (
    <div className="task-item">
      <div className="task-item__header">
        <h3>{task.title}</h3>
        <div className="task-item__actions">
          <button onClick={() => onEdit(task)} className="btn btn-edit">
            編集
          </button>
          <button onClick={() => onDelete(task.id)} className="btn btn-delete">
            削除
          </button>
        </div>
      </div>
      <div className="task-item__content">
        <p>{task.description}</p>
        {task.status && (
          <span className="task-status">{task.status}</span>
        )}
      </div>
      <div className="task-item__footer">
        <span className="priority">優先度: {task.priority}</span>
        <span className="date">{new Date(task.created_at).toLocaleDateString()}</span>
      </div>
    </div>
  )
}

export default ItemTask;
