import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';
import { URLs } from '@/constants/url';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";

export async function GET() {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_UNAUTHORIZED])
    }
    console.log('language optimization API トークン:', token);

    const backendRes = await fetch(URLs.languageOptimization, {
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
        language_optimizations: data
      }, {
        status: backendRes.status
      });
  } catch (error) {
    console.error('language optimization API エラー:', error);
    if(error instanceof HttpError) {
      return NextResponse.json({
          code: error.code,
          error: `language optimization put | ${error.message}`,
        }, {
          status: error.status
        })
    }
  }
}

export async function POST(request: Request) {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;
  
    if (!token) {
      return NextResponse.json({
        error: errorMessages[ErrorCode.AUTH_UNAUTHORIZED],
      }, {
        status: StatusCodes.Unauthorized
      })
    }

    const body = await request.json()
    const backendRes = await fetch(URLs.languageOptimization, {
      method: 'POST',
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

    return NextResponse.json({
        language_optimization: data
      }, {
        status: backendRes.status
      })
  } catch (error) {
    console.error('Language optimization API エラー:', error)
    return NextResponse.json({
        error: `language optimization API エラー | ${errorMessages[ErrorCode.DB_QUERY_SAVE_FAILED]}`,
      }, {
        status: StatusCodes.InternalServerError
      })
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

    const backendRes = await fetch(`${URLs.languageOptimization}/${body.id}`, {
      method: 'PUT',
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

    console.log('language optimization API レスポンス:', data)

    return NextResponse.json({
        language_optimization: data.language_optimization
      }, {
        status: backendRes.status
      })
  } catch (error) {
    console.error('language optimization API エラー:', error)
    if(error instanceof HttpError) {
      return NextResponse.json({
          code: error.code,
          error: `language optimization  put | ${error.message}`,
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
      return NextResponse.json({
        error: errorMessages[ErrorCode.AUTH_UNAUTHORIZED],
      }, {
        status: StatusCodes.Unauthorized
      })
    }

    // idはクエリパラメータまたはbodyから取得（ここではbodyから取得する例）
    const body = await request.json();
    const id = body.id;
    if (!id) {
      throw new HttpError(StatusCodes.BadRequest, errorMessages[ErrorCode.PAYLOAD_ID_NOT_FOUND])
    }

    const backendRes = await fetch(`${URLs.languageOptimization}/${id}`, {
      method: 'DELETE',
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
        language_optimization: data.language_optimization,
      }, {
        status: backendRes.status
      })
  } catch (error) {
    console.error('language optimization API エラー:', error);
    if(error instanceof HttpError) {
      return NextResponse.json({
          code: error.code,
          error: `language optimization get | ${error.message}`,
        }, {
          status: error.status
        })
    }
  }
}
