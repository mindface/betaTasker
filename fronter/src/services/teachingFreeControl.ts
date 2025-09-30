import { AddTeachingFreeControl, TeachingFreeControl } from "../model/teachingFreeControl";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchTeachingFreeControlService = async () => {
  const data = await fetchApiJsonCore<undefined,TeachingFreeControl[]>({
    endpoint: '/api/teachingFreeControl',
    method: 'GET',
    errorMessage: 'error fetchTeachingFreeControlService テックコントロールサービス一覧取得失敗',
  });
  return data;
};

export const addTeachingFreeControlService = async (teachingFreeControl: AddTeachingFreeControl) => {
  const data = await fetchApiJsonCore<AddTeachingFreeControl,TeachingFreeControl>({
    endpoint: '/api/teachingFreeControl',
    method: 'POST',
    body: teachingFreeControl,
    errorMessage: 'error addTeachingFreeControlService テックコントロールサービス追加失敗',
  });
  return data;
};

export const updateTeachingFreeControlService = async (teachingFreeControl: TeachingFreeControl) => {
  const data = await fetchApiJsonCore<AddTeachingFreeControl,TeachingFreeControl>({
    endpoint: '/api/teachingFreeControl',
    method: 'PUT',
    body: teachingFreeControl,
    errorMessage: 'error updateTeachingFreeControlService テックコントロールサービス更新失敗',
  });
  return data;
};

export const deleteTeachingFreeControlService = async (id: string) => {
  const data = await fetchApiJsonCore<{id:string},TeachingFreeControl>({
    endpoint: `/api/teachingFreeControl`,
    method: 'DELETE',
    body: { id },
    errorMessage: 'error deleteTeachingFreeControlService テックコントロールサービス削除失敗',
  });
  return data;
};
