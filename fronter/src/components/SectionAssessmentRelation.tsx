"use client"
import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '../store';
import { loadTasks } from '../features/task/taskSlice';
import { loadMemories } from '../features/memory/memorySlice';
import { loadAssessments, createAssessment } from '../features/assessment/assessmentSlice';
import { Task } from '../model/task';
import { Memory } from '../model/memory';
import { AddAssessment } from '../model/assessment';
import AssessmentModal from './parts/AssessmentModal';
import ItemAssessment from './parts/ItemAssessment';

export default function SectionAssessmentRelation() {
  const dispatch = useDispatch();
  const { tasks } = useSelector((state: RootState) => state.task);
  const { memories } = useSelector((state: RootState) => state.memory);
  const { assessments } = useSelector((state: RootState) => state.assessment);

  const [selectedTask, setSelectedTask] = useState<Task | null>(null);
  const [selectedMemory, setSelectedMemory] = useState<Memory | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  useEffect(() => {
    dispatch(loadTasks());
    dispatch(loadMemories());
    dispatch(loadAssessments());
  }, [dispatch]);

  // Task選択時に、そのTaskに紐づくMemory一覧を取得
  const relatedMemories = selectedTask
    ? memories.filter(m => m.user_id === selectedTask.user_id)
    : [];

  // Assessment登録
  const handleAddAssessment = (memory: Memory) => {
    setSelectedMemory(memory);
    setIsModalOpen(true);
  };

  const handleSaveAssessment = async (assessmentData: AddAssessment) => {
    if (!selectedTask) {
      console.error('タスクが選択されていません');
      return;
    }
    await dispatch(createAssessment({
      ...assessmentData,
      task_id: selectedTask?.id,
      user_id: selectedTask?.user_id,
    }));
    setIsModalOpen(false);
  };

  // TaskごとにAssessmentを絞り込む
  const filteredAssessments = selectedTask
    ? assessments.filter(a => a.task_id === selectedTask.id)
    : assessments;

  return (
    <div className="section__inner section--assessment-relation">
      <h2 style={{ textAlign: 'center', marginBottom: 24 }}>タスク・メモリ・アセスメント連携</h2>
      <div className="relation-container" style={{ display: 'flex', gap: 32, alignItems: 'flex-start', justifyContent: 'center' }}>
        <div className="task-list" style={{ minWidth: 220, background: '#f8fafc', borderRadius: 8, padding: 16, boxShadow: '0 2px 8px #0001' }}>
          <h3 style={{ borderBottom: '1px solid #e5e7eb', paddingBottom: 8, marginBottom: 12 }}>タスク一覧</h3>
          <ul style={{ listStyle: 'none', padding: 0, margin: 0 }}>
            {tasks.map(task => (
              <li key={task.id} style={{ marginBottom: 8 }}>
                <button
                  onClick={() => setSelectedTask(task)}
                  style={{
                    width: '100%',
                    background: selectedTask?.id === task.id ? '#2563eb' : '#fff',
                    color: selectedTask?.id === task.id ? '#fff' : '#222',
                    border: '1px solid #2563eb',
                    borderRadius: 4,
                    padding: '8px 12px',
                    cursor: 'pointer',
                    fontWeight: selectedTask?.id === task.id ? 'bold' : 'normal',
                    transition: 'all 0.2s',
                  }}
                >
                  {task.title}
                </button>
              </li>
            ))}
          </ul>
        </div>
        {selectedTask && (
          <div className="memory-list" style={{ minWidth: 220, background: '#f1f5f9', borderRadius: 8, padding: 16, boxShadow: '0 2px 8px #0001' }}>
            <h3 style={{ borderBottom: '1px solid #e5e7eb', paddingBottom: 8, marginBottom: 12 }}>選択タスクのメモリ一覧</h3>
            <ul style={{ listStyle: 'none', padding: 0, margin: 0 }}>
              {relatedMemories.map(memory => (
                <li key={memory.id} style={{ marginBottom: 8, display: 'flex', alignItems: 'center', justifyContent: 'space-between' }}>
                  <span style={{ flex: 1 }}>{memory.title}</span>
                  <button
                    onClick={() => handleAddAssessment(memory)}
                    style={{
                      marginLeft: 8,
                      background: '#22c55e',
                      color: '#fff',
                      border: 'none',
                      borderRadius: 4,
                      padding: '6px 12px',
                      cursor: 'pointer',
                      fontWeight: 'bold',
                    }}
                  >
                    アセスメント追加
                  </button>
                </li>
              ))}
            </ul>
          </div>
        )}
        <div className="assessment-list" style={{ flex: 1, background: '#f9fafb', borderRadius: 8, padding: 16, boxShadow: '0 2px 8px #0001', minWidth: 320 }}>
          <h3 style={{ borderBottom: '1px solid #e5e7eb', paddingBottom: 8, marginBottom: 12 }}>アセスメント一覧</h3>
          <div style={{ display: 'flex', flexDirection: 'column', gap: 12 }}>
            {filteredAssessments.map(assessment => (
              <ItemAssessment key={assessment.id} assessment={assessment} />
            ))}
          </div>
        </div>
      </div>
      <AssessmentModal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        onSave={handleSaveAssessment}
        initialData={selectedMemory ? { task_id: selectedTask?.id || 0, user_id: selectedTask?.user_id || 0, effectiveness_score: 0, effort_score: 0, impact_score: 0, qualitative_feedback: '' } : undefined}
      />
    </div>
  );
}