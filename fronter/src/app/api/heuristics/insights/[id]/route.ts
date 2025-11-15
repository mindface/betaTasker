import { NextRequest, NextResponse } from 'next/server';
import { handleBaseRequest, handleError } from "../../../utlts/handleRequest"

const END_POINT_HEURISTICS_INSIGHTS = 'heuristicsInsights'

export async function GET(
  request: NextRequest,
  { params }: { params: Promise<{ id: string }> }
) {
    const { id } = await params;
    try {
      const { data, status } = await handleBaseRequest('GET',END_POINT_HEURISTICS_INSIGHTS, undefined, { id });
      return NextResponse.json({ knowledge_patterns: data.knowledge_patterns }, { status });
    } catch (error) {
      return handleError(error,END_POINT_HEURISTICS_INSIGHTS);
    }
}