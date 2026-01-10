import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import { URLs } from '@/constants/url';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";
import { errorMessages, ErrorCode } from '@/response/errorCodes';

export async function handleBaseRequest(
  method: 'GET' | 'POST' | 'PUT' | 'DELETE',
  endpoint: keyof typeof URLs,
  request?: Request,
  customBody?: object,
  dynamicParams?: Record<string, string>
) {
  try {
    const cookieStore = await cookies();
    const token = cookieStore.get('token')?.value;

    if (!token) {
      throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_UNAUTHORIZED]);
    }

    const headers = {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    };

    let body;
    if (customBody) {
      body = customBody;
    } else if (request) {
      body = await request.json();
    }

    let url;
    if(
      method === 'DELETE' ||
      method === 'PUT' ||
      ( method === 'GET' && body?.id )
    ) {
      url = `${URLs[endpoint]}/${body.id}`;
      console.log('Generated URL with ID:', url);
    } else {
      url = URLs[endpoint];
    }

    if(dynamicParams) {
      const key = Object.keys(dynamicParams)[0];
      switch(key) {
        case 'code':
          url + `/${dynamicParams[key]}`
      }
    }

    const sendData = body ? {
        method: method,
        headers,
        body: JSON.stringify(body)
      } : {
        method,
        headers,
      };
    if(method === 'GET') {
      delete sendData.body;
    }

    const backendRes = await fetch(url, sendData);

    const data = await backendRes.json();
    if (!backendRes.ok) {
      throw new HttpError(data.status, data.message, data.code);
    }

    return { data, status: backendRes.status };
  } catch (error) {
    console.error(`${endpoint} API エラー:`, error);
    throw error; // エラーを再スロー
  }
}


export function handleError(error: unknown, endpoint: string) {
  if (error instanceof HttpError) {
    return NextResponse.json({
      code: error.code,
      error: `${endpoint} API error | ${error.message}`,
    }, {
      status: error.status
    });
  }
  return NextResponse.json({
    code: ErrorCode.SYS_INTERNAL_ERROR,
    error: 'Internal Server Error',
  }, {
    status: StatusCodes.InternalServerError
  });
}
