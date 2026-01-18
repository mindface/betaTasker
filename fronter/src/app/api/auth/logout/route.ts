// /api/auth/logout.ts
import { NextResponse } from 'next/server'
import { successMessages, SuccessCode } from "@/response/resposeMessage";
import { StatusCodes } from '@/response/statusCodes';
import { URLs } from '@/constants/url';

export async function POST() {
  const response = NextResponse.json({
    message: successMessages[SuccessCode.LOGOUT_SUCCESS],
  },{
    status: StatusCodes.OK,
  })
  const backendRes = await fetch(URLs.logout, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
    })

  const data = await backendRes.json();
  console.log('Login response from backend:', data);


  response.cookies.set('token', '', { maxAge: 0, path: '/' })
  return response
}
