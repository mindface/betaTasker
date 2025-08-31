import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:8080';

export async function GET(request: NextRequest) {
  try {
    // クエリパラメータを取得
    const searchParams = request.nextUrl.searchParams;
    const queryString = searchParams.toString();

    // クッキーからトークンを取得
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      return NextResponse.json({ error: '認証トークンが見つかりません' }, { status: 401 });
    }
    console.log('Patterns API トークン:', queryString);
    console.log('searchParams API トークン:', searchParams);
    
    // バックエンドAPIにリクエスト
    const response = await fetch(
      `${API_BASE_URL}/api/heuristics/patterns${queryString ? `?${queryString}` : ''}`,
      {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        credentials: 'include',
      }
    );

    // レスポンスの処理
    if (!response.ok) {
      const error = await response.json();
      return NextResponse.json(
        { 
          success: false,
          error: error.message || 'Failed to fetch patterns',
          code: error.code 
        },
        { status: response.status }
      );
    }

    const data = await response.json();
    console.log(data)
    return NextResponse.json(data);
    
  } catch (error) {
    console.error('Error fetching patterns:', error);
    return NextResponse.json(
      { 
        success: false,
        error: 'Internal server error',
        message: error instanceof Error ? error.message : 'Unknown error'
      },
      { status: 500 }
    );
  }
}