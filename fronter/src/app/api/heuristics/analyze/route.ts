import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import { URLs } from '@/constants/url';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";

export async function POST(request: NextRequest) {
  try {
    const body = await request.json();
    
    // クッキーからトークンを取得
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_UNAUTHORIZED])
    }

    // バックエンドAPIにリクエスト
    const backendRes = await fetch(
      URLs.heuristicsAnalyze,
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(body),
        credentials: 'include',
      }
    )

    const data = await backendRes.json();
    // レスポンスの処理
    if (!backendRes.ok) {
      throw new HttpError(data.status, data.message, data.code)
    }

    return NextResponse.json({ 
        analyze: data.data.analysis,
      }, {
        status: backendRes.status
      })

  } catch (error) {
    console.error('Error analyze data:', error)
    if(error instanceof HttpError) {
      return NextResponse.json({
          code: error.code,
          error: `analyze post | ${error.message}`,
        }, {
          status: error.status
        })
    }
  }
}