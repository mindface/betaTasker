import { Assessment } from "./assessment";

export interface ResponseMeta {
  total: number;
  total_pages: number;
  page: number;
  per_page: number;
}

export interface ResponseAssessment {
  assessments: Assessment[];
  meta: ResponseMeta;
}