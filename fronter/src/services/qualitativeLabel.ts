import { AddQualitativeLabel, QualitativeLabel } from "../model/qualitativeLabel";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchQualitativeLabelsService = async () => {
  const data = await fetchApiJsonCore<undefined,QualitativeLabel[]>({
    endpoint: '/api/qualitativeLabel',
    method: 'GET',
    errorMessage: 'error fetchQualitativeLabelsService プロセス最適化一覧取得失敗'
  });
  return data;
};

export const addQualitativeLabelService = async (qualitativeLabel: AddQualitativeLabel) => {
  const data = await fetchApiJsonCore<AddQualitativeLabel,QualitativeLabel>({
    endpoint: '/api/qualitativeLabel',
    method: 'POST',
    body: qualitativeLabel,
    errorMessage: 'error addQualitativeLabelService プロセス最適化追加失敗'
  });
  return data;
};

export const updateQualitativeLabelService = async (qualitativeLabel: QualitativeLabel) => {
  const data = await fetchApiJsonCore<QualitativeLabel,QualitativeLabel>({
    endpoint: '/api/qualitativeLabel',
    method: 'PUT',
    body: qualitativeLabel,
    errorMessage: 'error updateQualitativeLabelService プロセス最適化更新失敗'
  });
  return data;
};

export const deleteQualitativeLabelService = async (id: string) => {
  const data = await fetchApiJsonCore<{id: string},undefined>({
    endpoint: `/api/qualitativeLabel`,
    method: 'DELETE',
    body: { id },
    errorMessage: 'error deleteQualitativeLabelService プロセス最適化削除失敗'
  });
  return data;
};
