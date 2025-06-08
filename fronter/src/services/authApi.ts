
export const loginApi = async (email: string, password: string) => {
  try {
    const res = await fetch('/api/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password }),
      headers: { 'Content-Type': 'application/json' },
      credentials: 'include',
    });

    const data = await res.json();
    if (!res.ok) throw new Error('ログイン失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};

export const logoutApi = async () => {
  try {
    const res = await fetch('/api/auth/logout', {
      method: 'POST',
      credentials: 'include',
    });
    const data = await res.json();
    if (!res.ok) throw new Error('ログアウト失敗');
    return data;
  } catch (err: any) {
    return { error: err.message };
  }
};
