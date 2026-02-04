import { ResponseError } from "@/model/fetchTypes";
import { User, LoginUserInfo } from "../model/user";
import { fetchApiJsonCore } from "@/utils/fetchApi";

export const loginApi = async (
  email: string,
  password: string,
): Promise<ResponseError | { token: string; user: User }> => {
  const data = await fetchApiJsonCore<
    { email: string; password: string },
    { token: string; user: User }
  >({
    endpoint: "/api/auth/login",
    method: "POST",
    body: { email, password },
    errorMessage: "error loginApi ログイン失敗",
  });
  if ("error" in data) {
    return data;
  }
  return data.value;
};

export const logoutApi = async (): Promise<ResponseError | undefined> => {
  const data = await fetchApiJsonCore<undefined, undefined>({
    endpoint: "/api/auth/logout",
    method: "POST",
    errorMessage: "error logoutApi ログアウトに失敗",
  });
  if ("error" in data) {
    return data;
  }
  return undefined;
};

export const regApi = async (
  user: LoginUserInfo,
): Promise<ResponseError | { token: string; user: User }> => {
  const data = await fetchApiJsonCore<
    LoginUserInfo,
    { token: string; user: User }
  >({
    endpoint: "/api/auth/register",
    method: "POST",
    body: user,
    errorMessage: "error regApi 登録に失敗",
  });
  if ("error" in data) {
    return data;
  }
  return data.value;
};
