import { NextResponse } from "next/server";
import { handleBaseRequest, handleError } from "../../../utlts/handleRequest";

const END_POINT_ANALYZE_PAGER = "heuristicsAnalyzePager";

export async function GET(request: Request) {
  const url = new URL(request.url);
  const page = url.searchParams.get("page") || "1";
  const limit = url.searchParams.get("limit") || "20"; 
  const task_id = url.searchParams.get("task_id") || "0"; 
  const include = url.searchParams.get("include") || "none"; 

  try {
    const { data, status } = await handleBaseRequest(
      "GET",
      END_POINT_ANALYZE_PAGER,
      undefined,
      undefined,
      { page, limit, task_id, include }
    );
    console.log(data.analyses)
    return NextResponse.json({
        analyses: data.analyses,
        meta: data.meta
      }, {
        status
      });
  } catch (error) {
    return handleError(error, END_POINT_ANALYZE_PAGER);
  }
}
