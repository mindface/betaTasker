export type HttpMethod = "GET" | "POST" | "PUT" | "DELETE";

export interface ResponseError {
  ok: false;
  error: Error;
}
