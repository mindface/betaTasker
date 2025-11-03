
export const SuccessCode: Record<string, string> = {
  // ログイン系 (LOGIN_xxx)
  LOGIN_SUCCESS: 'LOGIN_001',
  LOGOUT_SUCCESS: 'LOGOUT_001',

}

export const successMessages: Record<keyof typeof SuccessCode, string> = {
  [SuccessCode.LOGIN_SUCCESS]: 'ログインに成功しました',
  [SuccessCode.LOGOUT_SUCCESS]: 'ログアウトに成功しました',
};
