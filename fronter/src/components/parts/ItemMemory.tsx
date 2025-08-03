"use client"
import React, { useState } from 'react'
import { Memory } from "../../model/memory";


interface ItemMemoryProps {
  memory: Memory;
  onEdit: (memory: Memory) => void;
  onDelete: (id: number) => void;
}

const ItemMemory: React.FC<ItemMemoryProps> = ({ memory, onEdit, onDelete }) => {

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
            <p className="item">{memory.title}</p>
            <p>
              {memory.tags.split(',').map((tag, index) => (
                <span key={index} className="tag">
                  {tag.trim()}
                </span>
              ))}
            </p>
          </div>
        )}
      </div>
      <div className="card-item__footer card-item__footer">
        <span className="read-status">{memory.read_status}</span>
        <span className="date">{new Date(memory.created_at).toLocaleDateString()}</span>
      </div>
    </div>
  )
}

export default ItemMemory;
