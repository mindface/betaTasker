import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';

export async function GET() {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      return NextResponse.json({ error: '認証トークンが見つかりません' }, { status: 401 });
    }
    console.log('Memory API トークン:', token);

    const backendRes = await fetch('http://localhost:8080/api/memory', {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });

    if (!backendRes.ok) {
      return NextResponse.json({ error: 'バックエンドからの取得に失敗' }, { status: backendRes.status });
    }

    const data = await backendRes.json();
    console.log('Memory API レスポンス:', data);

    return NextResponse.json({ memories: data.memories || [] }, { status: 200 });
  } catch (error) {
    console.error('Memory API エラー:', error);
    return NextResponse.json({ error: 'サーバーエラー' }, { status: 500 });
  }
}
