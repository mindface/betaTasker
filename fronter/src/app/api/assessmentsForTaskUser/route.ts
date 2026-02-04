import { NextResponse } from "next/server";
import { handleBaseRequest, handleError } from "../utlts/handleRequest";

const END_POINT_ASSESSMENT_FOR_TASK_USER = "assessmentsForTaskUser";

export async function POST(request: Request) {
  const body = await request.json();
  const customBody = { task_id: body.taskId, user_id: body.userId };
  try {
    const { data, status } = await handleBaseRequest(
      "POST",
      END_POINT_ASSESSMENT_FOR_TASK_USER,
      request,
      customBody,
    );
    return NextResponse.json(
      {
        assessments: data.assessments,
      },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_ASSESSMENT_FOR_TASK_USER);
  }
}
