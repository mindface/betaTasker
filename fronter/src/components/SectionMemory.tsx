"use client";
import React, { useState, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../store";
import {
  loadMemories,
  createMemory,
  updateMemory,
  removeMemory,
  getMemoriesLimit,
} from "../features/memory/memorySlice";
import ItemMemory from "./parts/ItemMemory";
import MemoryModal from "./parts/MemoryModal";
import { AddMemory, Memory } from "../model/memory";
import MemoryAidList from "./MemoryAidList";
import PageNation from "./parts/PageNation";

export default function SectionMemory() {
  const dispatch = useDispatch();
  const { memories, memoryLoading, memoryError, memoriesPage, memoriesLimit, memoriesTotal, memoriesTotalPages   } = useSelector(
    (state: RootState) => state.memory,
  );
  const { isAuthenticated } = useSelector((state: RootState) => state.user);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingMemory, setEditingMemory] = useState<
    AddMemory | Memory | undefined
  >();
  const [aidCode, setAidCode] = useState("MA-Q-02"); // デフォルト値は任意

  useEffect(() => {
    // dispatch(loadMemories());
    dispatch(getMemoriesLimit({ page: 1, limit: 20 }));
  }, [dispatch, isAuthenticated]);

  const handleAddMemory = () => {
    // if (!isAuthenticated) {
    //   // TODO: ログインモーダルを表示
    //   return
    // }
    setEditingMemory(undefined);
    setIsModalOpen(true);
  };

  const handleEditMemory = (memory: Memory) => {
    setEditingMemory(memory);
    setIsModalOpen(true);
  };

  const handleSaveMemory = async (memoryData: AddMemory | Memory) => {
    if (editingMemory) {
      try {
        await dispatch(updateMemory(memoryData as Memory));
        await dispatch(loadMemories());
      } catch (error) {
        console.error("メモの更新に失敗しました:", error);
      }
    } else {
      try {
        await dispatch(createMemory(memoryData as AddMemory));
        await dispatch(loadMemories());
      } catch (error) {
        console.error("メモの更新に失敗しました:", error);
      }
    }
    setIsModalOpen(false);
  };

  const handleDeleteMemory = async (id: number) => {
    await dispatch(removeMemory(id));
  };

  const handlePageChange = (newPage: number) => {
    dispatch(getMemoriesLimit({ page: newPage, limit: memoriesLimit }));
  }

  return (
    <div className="section__inner section--memory p-8">
      <div className="section-container">
        <div className="memory-header p-b-8">
          <h2>メモ</h2>
          <button onClick={() => handleAddMemory()} className="btn btn-primary">
            新規メモ
          </button>
          <div className="p-b-10">
            {/* aidCode入力欄とMemoryAidList表示を追加 */}
            <div>
              <label>MemoryAidコード: </label>
              <input
                value={aidCode}
                onChange={(e) => setAidCode(e.target.value)}
                style={{ marginRight: 8 }}
              />
            </div>
          </div>
        </div>
        <MemoryAidList code={aidCode} />
        {/* 既存のメモリ表示 */}
        {memoryLoading && <div className="error-message">{memoryLoading}</div>}
        {memoryError ? (
          <div className="loading">読み込み中...</div>
        ) : (
          <div className="memory-list card-list">
            {memories.map((memory, index) => (
              <ItemMemory
                key={`memory-item${index}`}
                memory={memory}
                onEdit={(editMemory) => handleEditMemory(editMemory)}
                onDelete={() => handleDeleteMemory(memory.id)}
              />
            ))}
          </div>
        )}
        <div className="memory-pagination p-t-8">
          <PageNation
            page={memoriesPage}
            limit={memoriesLimit}
            totalPages={memoriesTotalPages}
            onChange={(newPage: number) => {
              handlePageChange(newPage);
            }}
          />
        </div>
        <MemoryModal
          isViewType={false}
          initialData={editingMemory}
          isOpen={isModalOpen}
          onClose={() => setIsModalOpen(false)}
          onSave={handleSaveMemory}
        />
      </div>
    </div>
  );
}
