import { cookies } from 'next/headers';
import { NextResponse } from 'next/server';
import { URLs } from '@/constants/url';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";
import { handleBaseRequest, handleError } from "../utlts/handleRequest"

const END_POINT_LANGUAGE_OPTIMIZATION = 'languageOptimization';

export async function GET() {
  try {
    const { data, status } = await handleBaseRequest('GET',END_POINT_LANGUAGE_OPTIMIZATION);
    return NextResponse.json({
      language_optimizations: data.language_optimizations }, { status });
  } catch (error) {
    return handleError(error,END_POINT_LANGUAGE_OPTIMIZATION);
  }
}

export async function POST(request: Request) {
  try {
    const { data, status } = await handleBaseRequest('POST',END_POINT_LANGUAGE_OPTIMIZATION,request);
    return NextResponse.json({
      language_optimization: data.language_optimization }, { status });
  } catch (error) {
    return handleError(error,END_POINT_LANGUAGE_OPTIMIZATION);
  }
}

export async function PUT(request: Request) {
  try {
    const { data, status } = await handleBaseRequest('PUT',END_POINT_LANGUAGE_OPTIMIZATION,request);
    return NextResponse.json({
      language_optimization: data.language_optimization }, { status });
  } catch (error) {
    return handleError(error,END_POINT_LANGUAGE_OPTIMIZATION);
  }
}

export async function DELETE(request: Request) {
  try {
    const { data, status } = await handleBaseRequest('PUT',END_POINT_LANGUAGE_OPTIMIZATION,request);
    return NextResponse.json({
      language_optimization: data.language_optimization }, { status });
  } catch (error) {
    return handleError(error,END_POINT_LANGUAGE_OPTIMIZATION);
  }

}
