import { AddQualitativeLabel, QualitativeLabel } from "../model/qualitativeLabel";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const fetchQualitativeLabelsClient = async () => {
  const data = await fetchApiJsonCore<undefined,QualitativeLabel[]>({
    endpoint: '/api/qualitativeLabel',
    method: 'GET',
    errorMessage: 'error fetchQualitativeLabelsClient プロセス最適化一覧取得失敗',
    getKey: 'qualitative_labels',    
  });
  return data;
};

export const addQualitativeLabelClient = async (qualitativeLabel: AddQualitativeLabel) => {
  const data = await fetchApiJsonCore<AddQualitativeLabel,QualitativeLabel>({
    endpoint: '/api/qualitativeLabel',
    method: 'POST',
    body: qualitativeLabel,
    errorMessage: 'error addQualitativeLabelClient プロセス最適化追加失敗',
    getKey: 'qualitative_label',
  });
  return data;
};

export const updateQualitativeLabelClient = async (qualitativeLabel: QualitativeLabel) => {
  const data = await fetchApiJsonCore<QualitativeLabel,QualitativeLabel>({
    endpoint: '/api/qualitativeLabel',
    method: 'PUT',
    body: qualitativeLabel,
    errorMessage: 'error updateQualitativeLabelClient プロセス最適化更新失敗',
    getKey: 'qualitative_label',
  });
  return data;
};

export const deleteQualitativeLabelClient = async (id: string) => {
  const data = await fetchApiJsonCore<{id: string},undefined>({
    endpoint: `/api/qualitativeLabel`,
    method: 'DELETE',
    body: { id },
    errorMessage: 'error deleteQualitativeLabelClient プロセス最適化削除失敗',
    getKey: 'qualitative_label',
  });
  return data;
};
