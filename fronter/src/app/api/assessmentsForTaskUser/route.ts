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
    console.log("POST /api/assessmentsForTaskUser body:", body)
    const backendRes = await fetch('http://localhost:8080/api/assessmentsForTaskUser', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({task_id: body.taskId, user_id: body.userId}),
    })

  if (!backendRes.ok) {
    const errorRes = await backendRes.json();
    console.error('バックエンドからのエラー:', errorRes);

    return NextResponse.json(
      { error: errorRes.error || 'assessmentsForTaskUserの情報取得に失敗' }, 
      { status: backendRes.status }
    )
  }

    const data = await backendRes.json()

    return NextResponse.json(data, { status: 201 })
  } catch (error) {
    console.error('Task API エラー:', error)
    return NextResponse.json({ error: 'サーバーエラー' }, { status: 500 })
  }
}
