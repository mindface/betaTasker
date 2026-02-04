import { NextRequest, NextResponse } from "next/server";
import { handleBaseRequest, handleError } from "../../utlts/handleRequest";

const END_POINT_HEURISTICS_ANALYZE = "heuristicsAnalyze";

export async function POST(request: NextRequest) {
  try {
    const { data, status } = await handleBaseRequest(
      "POST",
      END_POINT_HEURISTICS_ANALYZE,
      request,
    );
    return NextResponse.json(
      {
        analyze: data.data.analyze,
      },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_HEURISTICS_ANALYZE);
  }
}
