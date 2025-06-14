"use client"
import React, { useState, useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '../store'
import { loadMemories, createMemory, updateMemory, removeMemory } from '../features/memory/memorySlice'
import ItemMemory from "./parts/ItemMemory"
import MemoryModal from "./parts/MemoryModal"
import { AddMemory, Memory } from "../model/memory";

export default function SectionMemory() {
  const dispatch = useDispatch()
  const { memories, loading, error } = useSelector((state: RootState) => state.memory)
  const { isAuthenticated } = useSelector((state: RootState) => state.user)
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [editingMemory, setEditingMemory] = useState<AddMemory|Memory|undefined>()

  useEffect(() => {
    dispatch(loadMemories())
  }, [dispatch, isAuthenticated])

  const handleAddMemory = () => {
    // if (!isAuthenticated) {
    //   // TODO: ログインモーダルを表示
    //   return
    // }
    setEditingMemory(undefined)
    setIsModalOpen(true)
  }

  const handleEditMemory = (memory: Memory) => {
    setEditingMemory(memory)
    setIsModalOpen(true)
  }

  const handleSaveMemory = async (memoryData: AddMemory | Memory) => {
    if (editingMemory) {
      // TODO: 編集処理を実装
      console.log(memoryData)
      await dispatch(updateMemory(memoryData as Memory))
    } else {
      await dispatch(createMemory(memoryData as AddMemory))
    }
    setIsModalOpen(false)
  }

  const handleDeleteMemory = async (id: number) => {
    await dispatch(removeMemory(id))
  }

  const tes = (info:string) => {
    const _info = info
    return (test:string) => {
      return test +_info;
    }
  }

  return (
    <div className="section__inner section--memory">
      <div className="section-container">
        <div className="memory-header">
          <h2>メモ</h2>
          <button 
            onClick={() => handleAddMemory()}
            className="btn btn-primary"
          >
            新規メモ
          </button>
        </div>
        {error && (
          <div className="error-message">
            {error}
          </div>
        )}

        {loading ? (
          <div className="loading">読み込み中...</div>
        ) : (
          <div className="memory-list">
            {memories.map((memory,index) => (
              <ItemMemory
                key={`memory-item${index}`}
                memory={memory}
                onEdit={(editMemory) => handleEditMemory(editMemory)}
                onDelete={() => handleDeleteMemory(memory.id)}
              />
            ))}
          </div>
        )}
        <MemoryModal
          initialData={editingMemory}
          isOpen={isModalOpen}
          onClose={() => setIsModalOpen(false)}
          onSave={handleSaveMemory}
        />
      </div>
    </div>
  )
}

