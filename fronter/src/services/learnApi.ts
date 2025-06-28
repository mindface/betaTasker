import { LearningData } from '../model/learning';

// APIからlearningDataを取得する関数
export async function fetchLearningData(): Promise<LearningData> {
  const res = await fetch('/api/learning');
  if (!res.ok) {
    throw new Error('learningDataの取得に失敗しました');
  }
  return res.json();
}
