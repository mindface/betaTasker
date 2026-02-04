import {
  AddTeachingFreeControl,
  TeachingFreeControl,
} from "../model/teachingFreeControl";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchTeachingFreeControlClient = async () => {
  const data = await fetchApiJsonCore<undefined, TeachingFreeControl[]>({
    endpoint: "/api/teachingFreeControl",
    method: "GET",
    errorMessage:
      "error fetchTeachingFreeControlClient テックコントロールサービス一覧取得失敗",
    getKey: "teaching_free_controls",
  });
  return data;
};

export const addTeachingFreeControlClient = async (
  teachingFreeControl: AddTeachingFreeControl,
) => {
  const data = await fetchApiJsonCore<
    AddTeachingFreeControl,
    TeachingFreeControl
  >({
    endpoint: "/api/teachingFreeControl",
    method: "POST",
    body: teachingFreeControl,
    errorMessage:
      "error addTeachingFreeControlClient テックコントロールサービス追加失敗",
    getKey: "teaching_free_control",
  });
  return data;
};

export const updateTeachingFreeControlClient = async (
  teachingFreeControl: TeachingFreeControl,
) => {
  const data = await fetchApiJsonCore<
    AddTeachingFreeControl,
    TeachingFreeControl
  >({
    endpoint: "/api/teachingFreeControl",
    method: "PUT",
    body: teachingFreeControl,
    errorMessage:
      "error updateTeachingFreeControlClient テックコントロールサービス更新失敗",
    getKey: "teaching_free_control",
  });
  return data;
};

export const deleteTeachingFreeControlClient = async (id: string) => {
  const data = await fetchApiJsonCore<{ id: string }, TeachingFreeControl>({
    endpoint: `/api/teachingFreeControl`,
    method: "DELETE",
    body: { id },
    errorMessage:
      "error deleteTeachingFreeControlClient テックコントロールサービス削除失敗",
    getKey: "teaching_free_control",
  });
  return data;
};
