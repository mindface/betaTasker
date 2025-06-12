"use client"
import React, { useState } from 'react'
import { fetchMemoriesService, addMemoryService, deleteMemoryService } from '../../services/memoryApi'
import { Memory } from "../../model/memory";

interface ItemMemoryProps {
  memory: Memory;
  onEdit: (memory: Memory) => void;
  onDelete: (id: number) => void;
}

const ItemMemory: React.FC<ItemMemoryProps> = ({ memory, onEdit, onDelete }) => {

  return (
    <div className="memory-item">
      <div className="memory-item__header">
        <h3>{memory.title}</h3>
        <div className="memory-item__actions">
          <button onClick={() => onEdit(memory)} className="btn btn-edit">
            編集
          </button>
          <button onClick={() => onDelete(memory.id)} className="btn btn-delete">
            削除
          </button>
        </div>
      </div>
      <div className="memory-item__content">
        <p>{memory.notes}</p>
        {memory.tags && (
          <div className="memory-item__tags">
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
      <div className="memory-item__footer">
        <span className="read-status">{memory.read_status}</span>
        <span className="date">{new Date(memory.created_at).toLocaleDateString()}</span>
      </div>
    </div>
  )
}

export default ItemMemory;
