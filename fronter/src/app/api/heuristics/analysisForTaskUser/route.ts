import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';

export async function POST(request: Request) {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;
  
    if (!token) {
      return NextResponse.json({ error: '認証トークンが見つかりません' }, { status: 401 })
    }

    const body = await request.json()
    const { userId, taskId } = body;
    
    const backendRes = await fetch(`http://localhost:8080/api/heuristics/analysisForTaskUser`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ userId, taskId })
    })

    if (!backendRes.ok) {
      return NextResponse.json(
        { error: 'バックエンドからの取得に失敗' }, 
        { status: backendRes.status }
      )
    }

    const data = await backendRes.json()

    return NextResponse.json({ analyses: data.analyses || [] }, { status: 200 })
  } catch (error) {
    console.error('Heuristics Analysis For Task User API エラー:', error)
    return NextResponse.json({ error: 'サーバーエラー' }, { status: 500 })
  }
}