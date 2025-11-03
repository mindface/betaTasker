import { cookies } from 'next/headers';
import { NextRequest, NextResponse } from 'next/server';
import { URLs } from '@/constants/url';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";

export type Params = { params: Promise<{ id: string }>  };
export async function GET(
  req: NextRequest,
  { params }: Params
) {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_UNAUTHORIZED])
    }

    const { id } = await params;
    if (!id) {
      throw new HttpError(StatusCodes.BadRequest, errorMessages[ErrorCode.PAYLOAD_ID_NOT_FOUND])
    }

    const backendRes = await fetch(`${URLs.memory}/${id}`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    });

    const data = await backendRes.json()
    if (!backendRes.ok) {
      throw new HttpError(data.status, data.message, data.code)
    }

    return NextResponse.json({
        memory: data.memory
      }, {
        status: backendRes.status
      })
  } catch (error) {
    console.error('Memory API エラー:', error)
    if(error instanceof HttpError) {
      return NextResponse.json({
          code: error.code,
          error: `Memory get | ${error.message}`,
        }, {
          status: error.status
        })
    }
  }
}
