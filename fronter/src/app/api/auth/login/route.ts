import { NextRequest, NextResponse } from 'next/server';
import { URLs } from '@/constants/url';
import { StatusCodes } from '@/response/statusCodes';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { SuccessCode, successMessages } from "@/response/resposeMessage";
import { HttpError } from "@/response/httpError";
import { handleError } from "../../utlts/handleRequest"

export async function POST(req: NextRequest) {
  const { email, password } = await req.json();
  console.log('Login request received:', { email, password });
  if (!email || !password) {
    throw new HttpError(StatusCodes.BadRequest, errorMessages[ErrorCode.PAYLOAD_EMAIL_AND_PASSWORD_NOT_FOUND])
  }

  try {
    const backendRes = await fetch(URLs.login, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      })

    const data = await backendRes.json();
    console.log('Login response from backend:', data);

    if (!backendRes.ok) {
      throw new HttpError(data.status, data.message, data.code)
    }

    if (backendRes.ok && data.token) {
      const response = NextResponse.json({
          message: successMessages[SuccessCode.LOGIN_SUCCESS],
          token: data.token,
          user: data.user
        },{
          status: StatusCodes.OK,
        })

      response.cookies.set({
        name: 'token',
        value: data.token,
        httpOnly: true,
        secure: process.env.NODE_ENV === 'production', // 開発中は false でも可
        sameSite: 'strict',
        path: '/',
        maxAge: 60 * 60 * 24,
      })

      return response;
    }
    throw new HttpError(StatusCodes.Unauthorized, errorMessages[ErrorCode.AUTH_INVALID_CREDENTIALS])
  } catch (error) {
    console.error('Login error:', error);
    return handleError(error,'Login');
  }
}
