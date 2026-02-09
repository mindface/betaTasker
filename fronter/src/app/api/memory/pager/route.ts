import { NextResponse } from "next/server";
import { handleBaseRequest, handleError } from "../../utlts/handleRequest";

const END_POINT_MEMORY_PAGER = "memoryPager";

export async function GET(request: Request) {
  const url = new URL(request.url);
  const page = url.searchParams.get("page") || "1";
  const limit = url.searchParams.get("limit") || "20"; 
 
  try {
    const { data, status } = await handleBaseRequest(
      "GET",
      END_POINT_MEMORY_PAGER,
      undefined,
      undefined,
      { page, limit }
    );
    return NextResponse.json({
        memories: data.memories,
        meta: data.meta
      }, {
        status
      });
  } catch (error) {
    return handleError(error, END_POINT_MEMORY_PAGER);
  }
}
