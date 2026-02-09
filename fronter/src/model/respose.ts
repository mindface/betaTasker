
export interface ResponseMeta {
  total: number;
  total_pages: number;
  page: number;
  per_page: number;
  limit: number;
}

export type LimitResponse<T, K extends string = "items"> =
  { [P in K]: T[] } & {
    meta: ResponseMeta;
  };
