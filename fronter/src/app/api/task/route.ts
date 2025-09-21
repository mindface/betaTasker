import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';
import { URLs } from '@/constants/url';

export async function GET() {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      return NextResponse.json({ error: '認証トークンが見つかりません' }, { status: 401 });
    }
    console.log('Task API トークン:', token);

    const backendRes = await fetch(URLs.task, {
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

    return NextResponse.json({ tasks: data.tasks || [] }, { status: 200 });
  } catch (error) {
    console.error('Task API エラー:', error);
    return NextResponse.json({ error: 'サーバーエラー' }, { status: 500 });
  }
}

export async function POST(request: Request) {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;
  
    if (!token) {
      return NextResponse.json({ error: '認証トークンが見つかりません' }, { status: 401 })
    }

    const body = await request.json()
    const backendRes = await fetch(URLs.task, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body)
    })

    if (!backendRes.ok) {
      return NextResponse.json(
        { error: 'バックエンドへの保存に失敗' }, 
        { status: backendRes.status }
      )
    }

    const data = await backendRes.json()

    return NextResponse.json({ code: data.code, task: data.task, message: data.message }, { status: 201 })
  } catch (error) {
    console.error('Task API エラー:', error)
    return NextResponse.json({ error: 'サーバーエラー' }, { status: 500 })
  }
}

export async function PUT(request: Request) {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      return NextResponse.json({ error: '認証トークンが見つかりません' }, { status: 401 })
    }

    const body = await request.json()

    const backendRes = await fetch(`${URLs.task}/${body.id}`, {
      method: 'PUT',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body)
    })

    if (!backendRes.ok) {
      return NextResponse.json(
        { error: 'バックエンドへの保存に失敗' }, 
        { status: backendRes.status }
      )
    }

    const data = await backendRes.json()
    console.log('Task API レスポンス:', data)

    return NextResponse.json(data, { status: 201 })
  } catch (error) {
    console.error('Task API エラー:', error)
    return NextResponse.json({ error: 'サーバーエラー' }, { status: 500 })
  }
}

export async function DELETE(request: Request) {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      return NextResponse.json({ error: '認証トークンが見つかりません' }, { status: 401 });
    }

    // idはクエリパラメータまたはbodyから取得（ここではbodyから取得する例）
    const body = await request.json();
    const id = body.id;
    if (!id) {
      return NextResponse.json({ error: 'IDが指定されていません' }, { status: 400 });
    }

    const backendRes = await fetch(`${URLs.task}/${id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });

    if (!backendRes.ok) {
      return NextResponse.json(
        { error: 'バックエンドでの削除に失敗' },
        { status: backendRes.status }
      );
    }

    const data = await backendRes.json();
    return NextResponse.json(data, { status: 200 });
  } catch (error) {
    console.error('Task API エラー:', error);
    return NextResponse.json({ error: 'サーバーエラー' }, { status: 500 });
  }
}
