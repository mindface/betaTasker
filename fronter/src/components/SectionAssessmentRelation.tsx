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
      <h2>タスク・メモリ・アセスメント連携</h2>
      <div className="relation-container">
        <div className="task-list">
          <h3>タスク一覧</h3>
          <ul>
            {tasks.map(task => (
              <li key={task.id}>
                <button onClick={() => setSelectedTask(task)}>
                  {task.title}
                </button>
              </li>
            ))}
          </ul>
        </div>
        {selectedTask && (
          <div className="memory-list">
            <h3>選択タスクのメモリ一覧</h3>
            <ul>
              {relatedMemories.map(memory => (
                <li key={memory.id}>
                  <span>{memory.title}</span>
                  <button onClick={() => handleAddAssessment(memory)}>
                    アセスメント追加
                  </button>
                </li>
              ))}
            </ul>
          </div>
        )}
        <div className="assessment-list">
          <h3>アセスメント一覧</h3>
          {filteredAssessments.map(assessment => (
            <ItemAssessment key={assessment.id} assessment={assessment} />
          ))}
        </div>
      </div>
      <AssessmentModal
        isOpen={isModalOpen}
        tasks={tasks}
        onClose={() => setIsModalOpen(false)}
        onSave={handleSaveAssessment}
        initialData={selectedMemory ? { task_id: selectedTask?.id || 0, user_id: selectedTask?.user_id || 0, effectiveness_score: 0, effort_score: 0, impact_score: 0, qualitative_feedback: '' } : undefined}
      />
    </div>
  );
}