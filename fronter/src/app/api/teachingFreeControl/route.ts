import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';
import { URLs } from '@/constants/url';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";
import { handleBaseRequest, handleError } from "../utlts/handleRequest"

const END_POINT_TEACHING_FREE_CONTROL = 'teachingFreeControl';

export async function GET() {
  try {
    const { data, status } = await handleBaseRequest('GET',END_POINT_TEACHING_FREE_CONTROL);
    return NextResponse.json({ 
      teaching_free_controls: data.teaching_free_controls }, { status });
  } catch (error) {
    return handleError(error,END_POINT_TEACHING_FREE_CONTROL);
  }
}

export async function POST(request: Request) {
  try {
    const { data, status } = await handleBaseRequest('POST',END_POINT_TEACHING_FREE_CONTROL,request);
    return NextResponse.json({ 
      teaching_free_control: data.teaching_free_control }, { status });
  } catch (error) {
    return handleError(error,END_POINT_TEACHING_FREE_CONTROL);
  }
}

export async function PUT(request: Request) {
  const body = await request.json();
  const id = body.id;
  if (!id) {
    return NextResponse.json({
      error: errorMessages[ErrorCode.PAYLOAD_ID_NOT_FOUND],
      status: StatusCodes.BadRequest
    })
  }
  try {
    const { data, status } = await handleBaseRequest('PUT',END_POINT_TEACHING_FREE_CONTROL,request,{ id });
    return NextResponse.json({ 
      teaching_free_control: data.teaching_free_control }, { status });
  } catch (error) {
    return handleError(error,END_POINT_TEACHING_FREE_CONTROL);
  }
}

export async function DELETE(request: Request) {
  const body = await request.json();
  const id = body.id;
  if (!id) {
    return NextResponse.json({
      error: errorMessages[ErrorCode.PAYLOAD_ID_NOT_FOUND],
      status: StatusCodes.BadRequest
    })
  }
  try {
    const { data, status } = await handleBaseRequest('DELETE',END_POINT_TEACHING_FREE_CONTROL,request,{ id });
    return NextResponse.json({ 
      teaching_free_control: data.teaching_free_control }, { status });
  } catch (error) {
    return handleError(error,END_POINT_TEACHING_FREE_CONTROL);
  }
}
