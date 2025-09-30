import { Task, AddTask } from "../model/task";
import { Memory, AddMemory } from "../model/memory";
import { Assessment, AddAssessment } from "../model/assessment";
import { taskApiClient } from "./taskApiRefactored";
import { 
  addMemoryService, 
  updateMemoryService, 
  deleteMemoryService,
  fetchMemoriesService 
} from "./memoryApi";
import { 
  addAssessmentService, 
  updateAssessmentService, 
  deleteAssessmentService,
  fetchAssessmentsService 
} from "./assessmentApi";

export interface ApiStrategy<T, CreateT = Omit<T, 'id'>> {
  getAll: () => Promise<T[]>;
  create: (item: CreateT) => Promise<T>;
  update: (item: T) => Promise<T>;
  delete: (id: number) => Promise<{ success: boolean }>;
}

export class TaskApiStrategy implements ApiStrategy<Task, AddTask> {
  async getAll(): Promise<Task[]> {
    return taskApiClient.getTasks();
  }

  async create(task: AddTask): Promise<Task> {
    return taskApiClient.addTask(task);
  }

  async update(task: Task): Promise<Task> {
    return taskApiClient.updateTask(task);
  }

  async delete(id: number): Promise<{ success: boolean }> {
    return taskApiClient.deleteTask(id);
  }
}

export class MemoryApiStrategy implements ApiStrategy<Memory, AddMemory> {
  async getAll(): Promise<Memory[]> {
    const result = await fetchMemoriesService()
    if ('error' in result) {
      throw result.error;
    }
    return result.value;
  }

  async create(memory: AddMemory): Promise<Memory> {
    const result = await addMemoryService(memory);
    if ('error' in result) {
      throw result.error;
    }
    return result.value;
  }

  async update(memory: Memory): Promise<Memory> {
    const result = await updateMemoryService(memory);
    if ('error' in result) {
      throw result.error;
    }
    return result.value;
  }

  async delete(id: number): Promise<{ success: boolean }> {
    const result = await deleteMemoryService(String(id));
    if ('error' in result) {
      throw result.error;
    }
    return { success: true };
  }
}

export class AssessmentApiStrategy implements ApiStrategy<Assessment, AddAssessment> {
  async getAll(): Promise<Assessment[]> {
    const result = await fetchAssessmentsService();
    if ('error' in result) {
      throw result.error;
    }
    return 'value' in result ? result.value : [];
  }

  async create(assessment: AddAssessment): Promise<Assessment> {
    const result = await addAssessmentService(assessment);
    if('error' in result) {
      throw result.error;
    }
    return result;
  }

  async update(assessment: Assessment): Promise<Assessment> {
    const result = await updateAssessmentService(assessment);
    if('error' in result) {
      throw result.error;
    }
    return result;
  }

  async delete(id: number): Promise<{ success: boolean }> {
    const result = await deleteAssessmentService(String(id));
    if('error' in result) {
      throw result.error;
    }
    return { success: true };
  }
}

export const apiStrategies = {
  task: new TaskApiStrategy(),
  memory: new MemoryApiStrategy(),
  assessment: new AssessmentApiStrategy(),
} as const;

export type EntityType = keyof typeof apiStrategies;