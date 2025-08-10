import { AddTask, Task } from "../model/task";

const API_BASE_URL = '/api/task';

class TaskApiClient {
  private async request<T>(
    url: string,
    options?: RequestInit
  ): Promise<T> {
    const response = await fetch(url, {
      ...options,
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json',
        ...options?.headers,
      },
    });

    if (!response.ok) {
      const errorMessage = `API request failed: ${response.status} ${response.statusText}`;
      throw new Error(errorMessage);
    }

    const data = await response.json();
    return data;
  }

  async getTasks(): Promise<Task[]> {
    return this.request<Task[]>(API_BASE_URL, {
      method: 'GET',
    });
  }

  async getTask(id: number): Promise<Task> {
    return this.request<Task>(`${API_BASE_URL}/${id}`, {
      method: 'GET',
    });
  }

  async addTask(task: AddTask): Promise<Task> {
    return this.request<Task>(API_BASE_URL, {
      method: 'POST',
      body: JSON.stringify(task),
    });
  }

  async updateTask(task: Task): Promise<Task> {
    return this.request<Task>(API_BASE_URL, {
      method: 'PUT',
      body: JSON.stringify(task),
    });
  }

  async deleteTask(id: number): Promise<{ success: boolean }> {
    return this.request<{ success: boolean }>(API_BASE_URL, {
      method: 'DELETE',
      body: JSON.stringify({ id }),
    });
  }

  async updateTaskStatus(
    id: number,
    status: 'todo' | 'in_progress' | 'completed'
  ): Promise<Task> {
    return this.request<Task>(`${API_BASE_URL}/${id}/status`, {
      method: 'PATCH',
      body: JSON.stringify({ status }),
    });
  }

  async updateTaskPriority(
    id: number,
    priority: number
  ): Promise<Task> {
    return this.request<Task>(`${API_BASE_URL}/${id}/priority`, {
      method: 'PATCH',
      body: JSON.stringify({ priority }),
    });
  }
}

export const taskApiClient = new TaskApiClient();