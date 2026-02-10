"use client";
import { useState, useEffect, Suspense } from "react";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../store";
import { createTask, updateTask, removeTask } from "../features/task/taskSlice";
import ItemTask from "./parts/ItemTask";
import TaskModal from "./parts/TaskModal";
import PageNation from "./parts/PageNation";
import AssessmentListModal from "./parts/AssessmentListModal";
import { AddTask, Task } from "../model/task";
import { loadMemories } from "../features/memory/memorySlice";
import { getTasksLimit } from "../features/task/taskSlice";

export default function SectionTask() {
  const dispatch = useDispatch();
  const { isAuthenticated } = useSelector((state: RootState) => state.user);
  const { memories } = useSelector((state: RootState) => state.memory);
  const { tasks, taskError, tasksPage, tasksLimit, tasksTotal, tasksTotalPages } = useSelector((state: RootState) => state.task);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingTask, setEditingTask] = useState<AddTask | Task | undefined>();
  const [TaskId, setTaskId] = useState<number>(-1);

  useEffect(() => {
    dispatch(loadMemories());
    dispatch(getTasksLimit({ page: 1, limit: 20 }));
  }, [dispatch, isAuthenticated]);

  const handleAddTask = () => {
    setEditingTask(undefined);
    setIsModalOpen(true);
  };

  const handleEditTask = (task: Task) => {
    setEditingTask(task);
    setIsModalOpen(true);
  };

  // const handleSaveTask = async (taskData: AddTask | Task) => {
  //   if (editingTask) {
  //     await dispatch(updateTask(taskData as Task));
  //   } else {
  //     await dispatch(createTask(taskData as AddTask));
  //   }
  //   setIsModalOpen(false);
  // };

  const handleDeleteTask = async (id: number) => {
    await dispatch(removeTask(id));
  };

  const handlePageChange = (newPage: number) => {
    dispatch(getTasksLimit({ page: newPage, limit: tasksLimit }));
  };

  if(taskError !== null) {
    throw new Error(taskError.name + ": " + taskError.message);
  }

  return (
    <div className="section__inner section--task">
      <div className="section-container">
        <div className="task-header p-b-8">
          <h2>タスク一覧</h2>
          <button onClick={() => handleAddTask()} className="btn btn-primary">
            新規タスク
          </button>
        </div>
        <div className="task-list card-list">
          <Suspense fallback={<p>Waiting...</p>}>
            {(tasks ?? []).map((task: Task, index: number) => (
              <ItemTask
                key={`task-item${index}`}
                task={task}
                onEdit={(editTask: Task) => handleEditTask(editTask)}
                onDelete={() => handleDeleteTask(task.id)}
                // onSetTaskId={(id: number) => setTaskId(id)}
              />
            ))}
          </Suspense>
        </div>
        <div className="task-pagination p-t-8">
          <PageNation
            page={tasksPage}
            limit={tasksLimit}
            totalPages={tasksTotalPages}
            onChange={(newPage: number) => {
              handlePageChange(newPage);
            }}
          />
        </div>
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
          memories={memories}
        />
      </div>
    </div>
  );
}
