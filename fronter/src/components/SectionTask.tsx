"use client"
import React, { useState, useEffect } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { RootState } from '../store'
import { loadTasks, createTask, updateTask, removeTask } from '../features/task/taskSlice'
import ItemTask from "./parts/ItemTask"
import TaskModal from "./parts/TaskModal"
import AssessmentListModal from "./parts/AssessmentListModal"
import { AddTask, Task } from "../model/task";
import { loadMemories } from '../features/memory/memorySlice';

export default function SectionTask() {
  const dispatch = useDispatch()
  const { tasks, taskLoading, taskError } = useSelector((state: RootState) => state.task)
  const { isAuthenticated } = useSelector((state: RootState) => state.user)
  const { memories } = useSelector((state: RootState) => state.memory)
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [editingTask, setEditingTask] = useState<AddTask|Task|undefined>()
  const [TaskId,setTaskId] = useState<number>(-1);

  useEffect(() => {
    dispatch(loadTasks())
    dispatch(loadMemories())
  }, [dispatch, isAuthenticated])

  const handleAddTask = () => {
    setEditingTask(undefined)
    setIsModalOpen(true)
  }

  const handleEditTask = (task: Task) => {
    setEditingTask(task)
    setIsModalOpen(true)
  }

  const handleSaveTask = async (taskData: AddTask | Task) => {
    if (editingTask) {
      await dispatch(updateTask(taskData as Task))
    } else {
      await dispatch(createTask(taskData as AddTask))
    }
    setIsModalOpen(false)
  }

  const handleDeleteTask = async (id: number) => {
    await dispatch(removeTask(id))
  }

  return (
    <div className="section__inner section--task">
      <div className="section-container">
        <div className="task-header">
          <h2>タスク</h2>
          <button
            onClick={() => handleAddTask()}
            className="btn btn-primary"
          >
            新規タスク
          </button>
        </div>
        {taskError && (
          <div className="error-message">
            {taskError}
          </div>
        )}
        {taskLoading ? (
          <div className="loading">読み込み中...</div>
        ) : (
          <div className="task-list">
            {tasks.map((task: Task, index: number) => (
              <ItemTask
                key={`task-item${index}`}
                task={task}
                onEdit={(editTask: Task) => handleEditTask(editTask)}
                onDelete={() => handleDeleteTask(task.id)}
                onSetTaskId={(id: number) => setTaskId(id)}
              />
            ))}
          </div>
        )}
        {TaskId}
        <AssessmentListModal
          taskId={TaskId}
          isOpen={TaskId !== -1}
          onClose={() => setTaskId(-1)}
          onSave={() => {}}   
        />
        <TaskModal
          initialData={editingTask}
          isOpen={isModalOpen}
          onClose={() => setIsModalOpen(false)}
          onSave={handleSaveTask}
          memories={memories}
        />

      </div>
    </div>
  )
}
