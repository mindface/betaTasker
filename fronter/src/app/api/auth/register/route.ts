import { NextResponse } from 'next/server';
import { URLs } from '@/constants/url';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";

export async function POST(request: Request) {
  try {
    const body = await request.json();
    const backendRes = await fetch(URLs.register, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    })

    const data = await backendRes.json()
    if (!backendRes.ok) {
      throw new HttpError(data.status, data.message, data.code)
    }
    return NextResponse.json({
      user: data.user,
      status: backendRes.status
    })
  } catch (error) {
    if (error instanceof HttpError) {
      return NextResponse.json(
        {
          code: error.code,
          error: error.message
        },
        { status: error.status }
      )
    }
  }
}
