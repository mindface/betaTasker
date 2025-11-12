import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';
import { URLs } from '@/constants/url';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";

import { handleBaseRequest, handleError } from "../utlts/handleRequest"

const END_POINT_MEMOERY = 'memory';

export async function GET() {
  try {
    const { data, status } = await handleBaseRequest('GET',END_POINT_MEMOERY);
    return NextResponse.json({ memories: data.memories }, { status });
  } catch (error) {
    return handleError(error,END_POINT_MEMOERY);
  }
}

export async function POST(request: Request) {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;
  
    if (!token) {
      throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_UNAUTHORIZED])
    }

    const body = await request.json()
    const backendRes = await fetch(URLs.memory, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body)
    })

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
    console.error('memory API エラー:', error)
    if(error instanceof HttpError) {
      return NextResponse.json({
          code: error.code,
          error: `memory post | ${error.message}`,
        }, {
          status: error.status
        })
    }
  }
}

export async function PUT(request: Request) {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_UNAUTHORIZED])
    }

    const body = await request.json()

    const backendRes = await fetch(`${URLs.memory}/${body.id}`, {
      method: 'PUT',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body)
    })

    if (!backendRes.ok) {
      return NextResponse.json({
          error: errorMessages[ErrorCode.DB_QUERY_FAILED],
        }, {
          status: backendRes.status
        })
    }

    const data = await backendRes.json()
    console.log('Memory API レスポンス:', data)

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
          error: `memory put | ${error.message}`,
        }, {
          status: error.status
        })
    }
  }
}

export async function DELETE(request: Request) {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_UNAUTHORIZED])
    }

    // idはクエリパラメータまたはbodyから取得（ここではbodyから取得する例）
    const body = await request.json();
    const id = body.id;
    if (!id) {
      return NextResponse.json({
        error: errorMessages[ErrorCode.PAYLOAD_ID_NOT_FOUND],
        status: StatusCodes.BadRequest
      });
    }

    const backendRes = await fetch(`${URLs.memory}/${id}`, {
      method: 'DELETE',
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
        memory: data.memory
      }, {
        status: backendRes.status
      });
  } catch (error) {
    console.error('Memory API エラー:', error);
    if(error instanceof HttpError) {
      return NextResponse.json({
          code: error.code,
          error: `memory get | ${error.message}`,
        }, {
          status: error.status
        })
    }
  }
}
