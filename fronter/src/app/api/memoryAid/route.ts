import { cookies } from 'next/headers';
import { NextRequest, NextResponse } from 'next/server';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";

export async function GET(req: NextRequest) {
  const cookieStore = await cookies();
  const token = cookieStore.get('token')?.value;
  const { searchParams } = new URL(req.url);
  const code = searchParams.get('code') || '';
  if (!token) {
    throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_UNAUTHORIZED])
  }
  try {
    const backendRes = await fetch(`http://localhost:8080/api/memory/aid/${code}`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      })
    const data = await backendRes.json();
    if (!backendRes.ok) {
      throw new HttpError(data.status, data.message, data.code)
    }
    return NextResponse.json({
        contexts: data.contexts,
      }, {
        status: backendRes.status
      })
  } catch (error) {
    console.error('memory aid API エラー:', error);
    if(error instanceof HttpError) {
      return NextResponse.json({
          code: error.code,
          error: `memory aid get | ${error.message}`,
        }, {
          status: error.status
        })
    }
  }
}
