// filepath: src/helper/authApi.ts
export async function loginApi(email: string, password: string): Promise<{ token?: string; error?: string }> {
  try {
    const res = await fetch('/api/auth/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password })
    });
    const data = await res.json();
    if (res.ok && data.token) {
      return { token: data.token };
    } else {
      return { error: data.error || 'ログイン失敗' };
    }
  } catch (err: any) {
    return { error: err.message || '通信エラー' };
  }
}
