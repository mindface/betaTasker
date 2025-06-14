import { AddTask, Task } from "../model/task";

export const fetchTasksService = async () => {
  try {
    const res = await fetch('/api/task', {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    console.log(data)
    if (!res.ok) throw new Error('タスク一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const addTaskService = async (task: AddTask) => {
  try {
    const res = await fetch('/api/task', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(task),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('タスク追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateTaskService = async (task: Task) => {
  try {
    const res = await fetch('/api/task', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(task),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('タスク更新失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteTaskService = async (id: string) => {
  try {
    const res = await fetch(`/api/task`, {
      method: 'DELETE',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id }),
    });
    const data = await res.json();
    if (!res.ok) throw new Error('タスク削除失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};
