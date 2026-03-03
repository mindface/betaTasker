"use client";
import React, { useState } from "react";
import { Task } from "../../model/task";
import { useDispatch, useSelector } from "react-redux";
import { getMemory } from "../../features/memory/memorySlice";

import MemoryModal from "./MemoryModal";
import ListDialog from "./ListDialog";
import HeuristicsAnalysisDialog from "./HeuristicsAnalysisDIalog";
import KnowledgePatternsDIalog from "./KnowledgePatternsDIalog";

import { Memory } from "../../model/memory";
import { RootState } from "../../store";

interface ItemTaskProps {
  task: Task;
  onEdit: (task: Task) => void;
  onDelete: (id: number) => void;
  onSetTaskId?: (id: number) => void;
}

type OptimizationsType = NonNullable<Task["language_optimizations"]>[number];

const ItemTask = ({ task, onEdit, onDelete, onSetTaskId }: ItemTaskProps) => {
  const dispath = useDispatch();
  const { memoryItem, memoryLoading, memoryError } = useSelector(
    (state: RootState) => state.memory,
  );

  const [isModalOpen, setIsModalOpen] = useState(false);

  const getMemoryAction = (memoryId: number) => {
    dispath(getMemory(memoryId));
    setIsModalOpen(true);
  };

  return (
    <>
      <div className="card-item">
        <div className="card-item__header">
          <h3 className="p-b-2">{task.title}</h3>
          <div className="card-item__actions">
            <button onClick={() => onEdit(task)} className="btn btn-edit">
              編集
            </button>
            <button
              onClick={() => onDelete(task.id)}
              className="btn btn-delete"
            >
              削除
            </button>
            <button
              onClick={() => {
                if (onSetTaskId) {
                  onSetTaskId(task.id);
                }
              }}
              className="btn"
            >
              アセスメントの確認
            </button>
          </div>
        </div>
        <div className="card-item__content">
          <p className="pb-1">
            {task.memory_id && (
              <button
                onClick={() => getMemoryAction(task.memory_id as number)}
                className="btn btn-edit"
              >
                記録を確認する
              </button>
            )}
          </p>
          <div className="task-for-memory-view">
            〈 記録ID: {task.memory_id}〉
          </div>
          <p className="p-b-1">{task.title}</p>
          <p className="p-b-4">{task.description}</p>
          {task.status && <span className="card-status">{task.status}</span>}
        </div>
        <div className="card-item__footer">
          <span className="priority">優先度: {task.priority}</span>
          <span className="date">
            {new Date(task.created_at).toLocaleDateString()}
          </span>
        </div>
        {task.language_optimizations && (
          <ListDialog<OptimizationsType>
            viewData={task.language_optimizations}
            title="Language Optimizations"
            btnText="Language Optimizationsを確認"
            renderItem={(item, index) => (
              <div className="language_optimizations-item p-8 bg-gray m-b-8">
                <p className="p-b-8">
                  <span className="abstraction-level p-8">
                    abstraction_level:
                  </span>
                  {item.abstraction_level}
                </p>
                <p className="p-b-4">original_text: {item.original_text}</p>
                <p className="p-b-4">optimized_text: {item.optimized_text}</p>
                {/* <p>domain: {item.domain}</p>
              <p>abstraction_level: {item.abstraction_level}</p>
              <p>precision: {item.precision}</p>
              <p>clarity: {item.clarity}</p>
              <p>completeness: {item.completeness}</p>
              <p>context: {item.context}</p>
              <p>transformation: {item.transformation}</p> */}
              </div>
            )}
          />
        )}
        <HeuristicsAnalysisDialog task={task} />
        {task.knowledge_patterns && <KnowledgePatternsDIalog knowledgePatterns={task.knowledge_patterns} />}
      </div>
      {memoryItem && (
        <MemoryModal
          initialData={memoryItem as Memory}
          isOpen={isModalOpen}
          isViewType={true}
          onClose={() => setIsModalOpen(false)}
          onSave={() => setIsModalOpen(false)}
        />
      )}
    </>
  );
};

export default ItemTask;
