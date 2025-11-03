import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import { URLs } from '@/constants/url';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";

export async function GET(request: NextRequest) {
  try {
    // クエリパラメータを取得
    const searchParams = request.nextUrl.searchParams;
    const queryString = searchParams.toString();

    // クッキーからトークンを取得
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      return NextResponse.json({
        error: errorMessages[ErrorCode.AUTH_UNAUTHORIZED],
      }, {
        status: StatusCodes.Unauthorized
      })
    }
    console.log('Patterns API トークン:', queryString);
    console.log('searchParams API トークン:', searchParams);

    // バックエンドAPIにリクエスト
    const backendRes = await fetch(
      `${URLs.heuristicsPatterns}${queryString ? `?${queryString}` : ''}`,
      {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        credentials: 'include',
      }
    )

    const data = await backendRes.json()
    if (!backendRes.ok) {
      throw new HttpError(data.status, data.message, data.code)
    }

    return NextResponse.json({
        patterns: data.patterns || [],
      },{
        status: backendRes.status
      });
    
  } catch (error) {
    console.error('Error fetching patterns:', error);
    if(error instanceof HttpError) {
      return NextResponse.json({
          code: error.code,
          error: `patterns get | ${error.message}`,
        }, {
          status: error.status
        })
    }
  }
}