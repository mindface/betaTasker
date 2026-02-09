import { NextResponse } from "next/server";
import { handleBaseRequest, handleError } from "../../utlts/handleRequest";

const END_POINT_TASK_PAGER = "taskPager";

export async function GET(request: Request) {
  const url = new URL(request.url);
  const page = url.searchParams.get("page") || "1";
  const limit = url.searchParams.get("limit") || "20"; 
 
  try {
    const { data, status } = await handleBaseRequest(
      "GET",
      END_POINT_TASK_PAGER,
      undefined,
      undefined,
      { page, limit }
    );
    return NextResponse.json({
        tasks: data.tasks,
        meta: data.meta
      }, {
        status
      });
  } catch (error) {
    return handleError(error, END_POINT_TASK_PAGER);
  }
}
