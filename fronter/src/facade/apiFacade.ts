import { Task, AddTask } from "../model/task";
import { Memory, AddMemory } from "../model/memory";
import { Assessment, AddAssessment } from "../model/assessment";
import {
  fetchTasksClient,
  addTaskClient,
  updateTaskClient,
  deleteTaskClient,
} from "../client/taskApi";
import {
  addMemoryClient,
  updateMemoryClient,
  deleteMemoryClient,
  fetchMemoriesClient,
} from "../client/memoryApi";
import {
  addAssessmentClient,
  updateAssessmentClient,
  deleteAssessmentClient,
  fetchAssessmentsClient,
} from "../client/assessmentApi";
import { LimitResponse } from "@/model/respose";

export interface ApiFacade<T, CreateT = Omit<T, "id">> {
  getAll: () => Promise<LimitResponse<T>>;
  create: (item: CreateT) => Promise<T>;
  update: (item: T) => Promise<T>;
  delete: (id: number) => Promise<{ success: boolean }>;
}

// Code that tested Claude and verified the coding methodology.
export class TaskApiFacade implements ApiFacade<Task, AddTask> {
  async getAll(): Promise<LimitResponse<Task>> {
    const response = await fetchTasksClient();
    if ("error" in response) {
      throw response;
    }
    return { items: response.tasks, meta: response.meta };
  }

  async create(task: AddTask): Promise<Task> {
    const response = await addTaskClient(task);
    if ("error" in response) {
      throw response.error;
    }
    return response.value;
  }

  async update(task: Task): Promise<Task> {
    const response = await updateTaskClient(task);
    if ("error" in response) {
      throw response.error;
    }
    return response.value;
  }

  async delete(id: number): Promise<{ success: boolean }> {
    const response = await deleteTaskClient(id);
    if ("error" in response) {
      throw response.error;
    }
    return { success: true };
  }
}

export class MemoryApiFacade implements ApiFacade<Memory, AddMemory> {
  async getAll(): Promise<LimitResponse<Memory>> {
    const result = await fetchMemoriesClient();
    if ("error" in result) {
      throw result.error;
    }
    return { items: result.memories, meta: result.meta };
  }

  async create(memory: AddMemory): Promise<Memory> {
    const result = await addMemoryClient(memory);
    if ("error" in result) {
      throw result.error;
    }
    return result.value;
  }

  async update(memory: Memory): Promise<Memory> {
    const result = await updateMemoryClient(memory);
    if ("error" in result) {
      throw result.error;
    }
    return result.value;
  }

  async delete(id: number): Promise<{ success: boolean }> {
    const result = await deleteMemoryClient(String(id));
    if ("error" in result) {
      throw result.error;
    }
    return { success: true };
  }
}

export class AssessmentApiFacade
  implements ApiFacade<Assessment, AddAssessment>
{
  async getAll(): Promise<LimitResponse<Assessment>> {
    const result = await fetchAssessmentsClient();
    if ("error" in result) {
      throw result;
    }
    return { items: result.assessments, meta: result.meta };
  }

  async create(assessment: AddAssessment): Promise<Assessment> {
    const result = await addAssessmentClient(assessment);
    if ("error" in result) {
      throw result.error;
    }
    return result;
  }

  async update(assessment: Assessment): Promise<Assessment> {
    const result = await updateAssessmentClient(assessment);
    if ("error" in result) {
      throw result.error;
    }
    return result;
  }

  async delete(id: number): Promise<{ success: boolean }> {
    const result = await deleteAssessmentClient(String(id));
    if ("error" in result) {
      throw result.error;
    }
    return { success: true };
  }
}

export const apiFacade = {
  task: new TaskApiFacade(),
  memory: new MemoryApiFacade(),
  assessment: new AssessmentApiFacade(),
} as const;

export type EntityType = keyof typeof apiFacade;
