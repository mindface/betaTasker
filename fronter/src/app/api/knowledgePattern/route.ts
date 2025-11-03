import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';
import { URLs } from '@/constants/url';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from '@/response/httpError';


export async function GET() {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_UNAUTHORIZED])
    }
    console.log('knowledge pattern API トークン:', token);

    const backendRes = await fetch(URLs.knowledgePattern, {
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
        knowledge_patterns: data.knowledge_patterns || []
      }, {
        status: backendRes.status
      })
  } catch (error) {
    console.error('knowledge pattern API エラー:', error);
    if(error instanceof HttpError) {
      return NextResponse.json({
          code: error.code,
          error: `knowledge pattern get | ${error.message}`,
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
      throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_UNAUTHORIZED])
    }

    const body = await request.json()
    const backendRes = await fetch(URLs.knowledgePattern, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body)
    });

    const data = await backendRes.json()
    if (!backendRes.ok) {
      throw new HttpError(data.status, data.message, data.code)
    }

    return NextResponse.json({
        knowledgePatterns: data.knowledge_pattern
      }, {
        status: backendRes.status
      })
  } catch (error) {
    console.error('knowledge pattern API エラー:', error);
    return NextResponse.json({
        error: `knowledge pattern post | ${errorMessages[ErrorCode.DB_QUERY_SAVE_FAILED]}`,
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

    const backendRes = await fetch(`${URLs.knowledgePattern}/${body.id}`, {
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

    console.log('knowledge pattern API レスポンス:', data)
    return NextResponse.json({
        knowledge_pattern: data.knowledge_pattern
      }, { 
        status: backendRes.status
      })

  } catch (error) {
    console.error('KnowledgePattern API エラー:', error)
    return NextResponse.json({
        error: `KnowledgePattern put | ${errorMessages[ErrorCode.DB_QUERY_SAVE_FAILED]}`,
      }, {
        status: StatusCodes.InternalServerError
      })
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
      throw new HttpError(StatusCodes.BadRequest, errorMessages[ErrorCode.PAYLOAD_ID_NOT_FOUND])
    }

    const backendRes = await fetch(`${URLs.knowledgePattern}/${id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    })

    const data = await backendRes.json()
    if (!backendRes.ok) {
      throw new HttpError(data.status, data.message, data.code)
    }

    return NextResponse.json({
        knowledge_pattern: data.knowledge_pattern
      }, {
        status: backendRes.status
      })

  } catch (error) {
    console.error('knowledge pattern API エラー:', error);
    if(error instanceof HttpError) {
      return NextResponse.json({
          code: error.code,
          error: `knowledge pattern delete | ${error.message}`,
        }, {
          status: error.status
        })
    }
  }
}
