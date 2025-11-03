import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import { URLs } from '@/constants/url';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";

export async function GET(
  request: NextRequest,
  { params }: { params: Promise<{ id: string }> }
) {
  try {
    const { id } = await params;

    // クッキーからトークンを取得
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_UNAUTHORIZED])
    }

    // バックエンドAPIにリクエスト
    const response = await fetch(
      `${URLs.heuristicsAnalyze}/${id}`,
      {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        credentials: 'include',
      }
    )

    const data = await response.json();
    // レスポンスの処理
    if (!response.ok) {
      throw new HttpError(data.status, data.message, data.code)
    }

    return NextResponse.json({
        analysis: data.data.analysis,
      }, {
        status: StatusCodes.OK
      })
  } catch (error) {
    console.error('Error fetching analysis:', error)
    if(error instanceof HttpError) {
      return NextResponse.json({
          code: error.code,
          error: `analyze get | ${error.message}`,
        }, {
          status: error.status
        })
    }
  }
}