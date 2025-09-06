import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';

export async function GET() {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      return NextResponse.json({ error: '認証トークンが見つかりません' }, { status: 401 });
    }
    console.log('Heuristics Tracking API トークン:', token);

    const backendRes = await fetch('http://localhost:8080/api/heuristics/tracking', {
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

    return NextResponse.json({ tracking: data.tracking || [] }, { status: 200 });
  } catch (error) {
    console.error('Heuristics Tracking API エラー:', error);
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
    const backendRes = await fetch('http://localhost:8080/api/heuristics/tracking', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body)
    })

    if (!backendRes.ok) {
      return NextResponse.json(
        { error: 'バックエンドからの取得に失敗' },
        { status: backendRes.status }
      )
    }

    const data = await backendRes.json()

    return NextResponse.json(data, { status: 201 })
  } catch (error) {
    console.error('Heuristics Tracking API エラー:', error)
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

    const backendRes = await fetch(`http://localhost:8080/api/heuristics/tracking/${body.id}`, {
      method: 'PUT',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body)
    })

    if (!backendRes.ok) {
      return NextResponse.json(
        { error: '�ï���xn�Xk1W' }, 
        { status: backendRes.status }
      )
    }

    const data = await backendRes.json()
    console.log('Heuristics Tracking API レスポンス:', data)

    return NextResponse.json(data, { status: 201 })
  } catch (error) {
    console.error('Heuristics Tracking API エラー:', error)
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

    const body = await request.json();
    const id = body.id;
    if (!id) {
      return NextResponse.json({ error: 'IDが見つかりません' }, { status: 400 });
    }

    const backendRes = await fetch(`http://localhost:8080/api/heuristics/tracking/${id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });

    if (!backendRes.ok) {
      return NextResponse.json(
        { error: 'バックエンドからの取得に失敗' },
        { status: backendRes.status }
      );
    }

    const data = await backendRes.json();
    return NextResponse.json(data, { status: 200 });
  } catch (error) {
    console.error('Heuristics Tracking API エラー:', error);
    return NextResponse.json({ error: 'サーバーエラー' }, { status: 500 });
  }
}