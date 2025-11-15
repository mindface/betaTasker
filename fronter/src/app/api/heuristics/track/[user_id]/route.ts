import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import { URLs } from '@/constants/url';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";
import { handleBaseRequest, handleError } from "../../../utlts/handleRequest"

const END_POINT_HEURISTICS_TRACK = 'heuristicsTrack';

export async function GET(
  request: NextRequest,
  { params }: { params: Promise<{ user_id: string }> }
) {
  try {
    const { user_id } = await params;

    // クッキーからトークンを取得
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_UNAUTHORIZED])
    }

    // バックエンドAPIにリクエスト
    const backendRes = await fetch(
      `${URLs.heuristicsTrack}/${user_id}`,
      {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        credentials: 'include',
      }
    )

    const data = await backendRes.json();

    if (!backendRes.ok) {
      throw new HttpError(data.status, data.message, data.code)
    }

    return NextResponse.json({ 
        track: data
      }, {
        status: backendRes.status
      });

  } catch (error) {
    return handleError(error,END_POINT_HEURISTICS_TRACK);
  }
}