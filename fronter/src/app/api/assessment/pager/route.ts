import { NextResponse } from "next/server";
import { handleBaseRequest, handleError } from "../../utlts/handleRequest";

const END_POINT_ASSESSMENT_PAGER = "assessmentPager";

export async function GET(request: Request) {
  const url = new URL(request.url);
  const page = url.searchParams.get("page") || "1";
  const limit = url.searchParams.get("limit") || "20"; 
 
  try {
    const { data, status } = await handleBaseRequest(
      "GET",
      END_POINT_ASSESSMENT_PAGER,
      undefined,
      undefined,
      { page, limit }
    );
    return NextResponse.json({
        assessments: data.assessments,
        meta: data.meta
      }, {
        status
      });
  } catch (error) {
    return handleError(error, END_POINT_ASSESSMENT_PAGER);
  }
}
