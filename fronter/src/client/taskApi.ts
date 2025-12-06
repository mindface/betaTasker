import { AddTask, Task } from "../model/task";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export type SuccessResponse<T = any, K extends string = "data"> = {
  success: boolean;
  message?: string;
} & {
  [key in K]?: T;
};

export const fetchTasksClient = async () => {
  const data = await fetchApiJsonCore<undefined,Task[]>({
    endpoint: '/api/task',
    method: 'GET',
    errorMessage: 'error fetchTasksClient タスク一覧取得失敗',
    getKey: 'tasks',
  });

  return data;
};

export const addTaskClient = async (task: AddTask) => {
  const data = await fetchApiJsonCore<AddTask,Task>({
    endpoint: '/api/task',
    method: 'POST',
    body: task,
    errorMessage: 'error addTaskClient タスク追加失敗',
    getKey: 'task',
  });
  return data;
};

export const updateTaskClient = async (task: Task) => {
  const data = await fetchApiJsonCore<Task,Task>({
    endpoint: '/api/task', 
    method: 'PUT',
    body: task,
    errorMessage: 'error updateTaskClient タスク更新失敗',
    getKey: 'task',
  });
  return data;
};

export const deleteTaskClient = async (id: number) => {
  const data = await fetchApiJsonCore<{id:number},Task>({
    endpoint: `/api/task`,
    method: 'DELETE',
    body: { id },
    errorMessage: 'error deleteTaskClient タスク削除失敗',
    getKey: 'task',
  });
  
  return data;
};
