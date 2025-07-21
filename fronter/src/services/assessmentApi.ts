import { AddAssessment, Assessment } from "../model/assessment";

export const fetchAssessmentsService = async () => {
  try {
    const res = await fetch('/api/assessment', {
      method: 'GET',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('アセスメント一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const getAssessmentsForTaskUserService = async (userId: number,taskId:number) => {
  try {
    const res = await fetch('/api/assessmentsForTaskUser', {
      method: 'POST',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ userId, taskId }),
    });
    const data = await res.json();
    if (!res.ok) throw new Error('アセスメント一覧取得失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const addAssessmentService = async (assessment: AddAssessment) => {
  try {
    const res = await fetch('/api/assessment', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(assessment),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('アセスメント追加失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const updateAssessmentService = async (assessment: Assessment) => {
  try {
    const res = await fetch('/api/assessment', {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(assessment),
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('アセスメント更新失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const deleteAssessmentService = async (id: string) => {
  try {
    const res = await fetch(`/api/assessment`, {
      method: 'DELETE',
      credentials: 'include',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ id }),
    });
    const data = await res.json();
    if (!res.ok) throw new Error('アセスメント削除失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};
