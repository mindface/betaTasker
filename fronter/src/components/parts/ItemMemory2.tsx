"use client"
import React, { useState } from 'react'
import { Memory } from "../../model/memory";
import { Task } from "../../model/task";
import CommonDialog from "./CommonDialog"

interface ItemMemoryProps {
  memory: Memory;
  tasks: Task[];
  onEdit: (memory: Memory) => void;
  onDelete: (id: number) => void;
}

const ItemMemory: React.FC<ItemMemoryProps> = ({ memory, tasks, onEdit, onDelete }) => {
  const [isTaskModalOpen, setIsTaskModalOpen] = useState(false)

  // const handleOpenMemoryModal = async (memoryId: string) => {
  //   if (memoryCache[memoryId]) {
  //     // キャッシュがあればそれを使う
  //     setSelectedMemory(memoryCache[memoryId]);
  //     setModalOpen(true);
  //   } else {
  //     // なければAPIリクエスト
  //     const res = await fetch(`/api/memory/${memoryId}`);
  //     const data = await res.json();
  //     setMemoryCache(prev => ({ ...prev, [memoryId]: data }));
  //     setSelectedMemory(data);
  //     setModalOpen(true);
  //   }
  // };

  const selectTask = (taskId: number) => {
    const selectedTask = tasks.find(task => task.id === taskId);
    if (!selectedTask) {
      return <div>タスクが見つかりません。</div>;
    }
    return <div>
      <h3>{selectedTask?.title}</h3>
      <p>{selectedTask?.description}</p>
      {(selectedTask.heuristics_insights ?? []).map((heuristic, index) => (
        <div key={index}>
          <h4>{heuristic.data}</h4>
          <p>{heuristic.description}</p>
        </div>
      ))}
      <div>
        {selectedTask?.language_optimizations?.map((opt, index) => (
          <div key={index} className="tag-optimization p-8">
            <p>{opt.abstraction_level}</p>
            <p>{opt.optimized_text}</p>
            <p>{opt.original_text}</p>
          </div>
        ))}
      </div>
    </div>;
  }

  return (
    <div className="card-item">
      <div className="card-item__header">
        <h3>{memory.title}</h3>
        <div className="card-item__actions">
          <button onClick={() => onEdit(memory)} className="btn btn-edit">
            編集
          </button>
          <button onClick={() => onDelete(memory.id)} className="btn btn-delete">
            削除
          </button>
        </div>
      </div>
      <div className="card-item__content card-item__content">
        <p>{memory.notes}</p>
        {memory.tags && (
          <div className="card-item__tags card-item__tags">
            <h3 className="item">{memory.title}</h3>
            <div>
              {memory.tags.split(',').map((tag, index) => (
                <span key={index} className="tag mr-4">
                  {tag.trim()}
                </span>
              ))}
            </div>
          </div>
        )}
      </div>
      <button className="btn" onClick={() => setIsTaskModalOpen(true)}>タスク表示</button>
      <CommonDialog
        isOpen={isTaskModalOpen}
        title="関連タスク選択"
        onClose={() => setIsTaskModalOpen(false)}
        children={
          <div className="dialog-content">
            <label className="title">関連タスク選択: </label>
            <div>
              <div>{selectTask(memory.id)}</div>
            </div>
          </div>
        }
      />
      <div className="card-item__footer card-item__footer">
        <span className="read-status">{memory.read_status}</span>
        <span className="date">{new Date(memory.created_at).toLocaleDateString()}</span>
      </div>
    </div>
  )
}

export default ItemMemory;
