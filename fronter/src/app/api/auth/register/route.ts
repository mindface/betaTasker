import { NextResponse } from 'next/server';
import { URLs } from '@/constants/url';

export async function POST(request: Request) {
  try {
    const body = await request.json();
    const backendRes = await fetch(URLs.register, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    });

    const data = await backendRes.json()
    if (!backendRes.ok) {
      return NextResponse.json({ error: data.error || '登録に失敗しました' }, { status: backendRes.status });
    }
    return NextResponse.json(data, { status: 200 });
  } catch (error) {
    return NextResponse.json({ error: 'サーバーエラー' }, { status: 500 });
  }
}
