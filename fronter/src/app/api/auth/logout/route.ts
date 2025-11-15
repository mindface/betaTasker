// /api/auth/logout.ts
import { NextResponse } from 'next/server'
import { successMessages, SuccessCode } from "@/response/resposeMessage";
import { StatusCodes } from '@/response/statusCodes';

export async function POST() {
  const response = NextResponse.json({
    message: successMessages[SuccessCode.LOGOUT_SUCCESS],
  },{
    status: StatusCodes.OK,
  })
  response.cookies.set('token', '', { maxAge: 0, path: '/' })
  return response
}
