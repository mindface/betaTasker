import { NextRequest, NextResponse } from 'next/server';
import { cookies } from 'next/headers';
import { URLs } from '@/constants/url';
import { errorMessages, ErrorCode } from '@/response/errorCodes';
import { StatusCodes } from '@/response/statusCodes';
import { HttpError } from "@/response/httpError";
import { handleBaseRequest, handleError } from "../../../utlts/handleRequest"

const END_POINT_HEURISTICS_ANALYZE = 'heuristicsAnalyze'

export async function GET(
  request: NextRequest,
  { params }: { params: Promise<{ id: string }> }
) {
    const { id } = await params;
    try {
      const { data, status } = await handleBaseRequest('GET',END_POINT_HEURISTICS_ANALYZE, undefined, { id });
      return NextResponse.json({ knowledge_patterns: data.knowledge_patterns }, { status });
    } catch (error) {
      return handleError(error,END_POINT_HEURISTICS_ANALYZE);
    }
}