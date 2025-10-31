import { ResponseError, HttpMethod } from "@/model/fetchTypes";

interface FetchOptions<TBody> {
  endpoint: string;
  method?: HttpMethod;
  body?: TBody;
  errorMessage: string;
  getKey?: string;
}

type Result<T> = { ok: true; value: T } | ResponseError;

// TB=TBody | TR=TResponse 
export const fetchApiJsonCore = async <TB,TR>({
  endpoint,
  method = 'GET',
  body,
  errorMessage,
  getKey,
}: FetchOptions<TB>): Promise<Result<TR>> => {
  try {
    const res = await fetch(endpoint, {
      method,
      credentials: 'include',
      headers: body ? { 'Content-Type': 'application/json' } : undefined,
      body: body ? JSON.stringify(body) : undefined,
    });

    const data = await res.json();
    console.log(123, res);
    if (!res.ok) {
      return { ok: false, error: new Error(errorMessage) };
    }
    return { ok: true, value: getKey ? data[getKey] as TR : data as TR };
  } catch (err: unknown) {
    return { ok: false, error: err instanceof Error ? err : new Error(String(err)) };
  }
};