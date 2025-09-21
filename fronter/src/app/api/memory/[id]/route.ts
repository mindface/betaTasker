import { cookies } from 'next/headers';
import { NextRequest, NextResponse } from 'next/server';
import { URLs } from '@/constants/url';

export type Params = { params: Promise<{ id: string }>  };
export async function GET(
  req: NextRequest,
  { params }: Params
) {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      return NextResponse.json({ error: '認証トークンが見つかりません' }, { status: 401 });
    }

    const { id } = await params;
    if (!id) {
      return NextResponse.json({ error: 'IDが指定されていません' }, { status: 400 });
    }

    const backendRes = await fetch(`${URLs.memory}/${id}`, {
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

    return NextResponse.json(data, { status: 200 });
  } catch (error) {
    console.error('Memory API エラー:', error);
    return NextResponse.json({ error: 'サーバーエラー' }, { status: 500 });
  }
}
