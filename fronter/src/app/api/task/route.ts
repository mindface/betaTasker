import { NextResponse } from "next/server";
import { errorMessages, ErrorCode } from "@/response/errorCodes";
import { StatusCodes } from "@/response/statusCodes";
import { HttpError } from "@/response/httpError";
import { handleBaseRequest, handleError } from "../utlts/handleRequest";

const END_POINT_TASK = "task";

export async function GET() {
  try {
    const { data, status } = await handleBaseRequest("GET", END_POINT_TASK);
    return NextResponse.json(
      {
        tasks: data.tasks,
      },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_TASK);
  }
}

export async function POST(request: Request) {
  try {
    const { data, status } = await handleBaseRequest(
      "POST",
      END_POINT_TASK,
      request,
    );
    return NextResponse.json({ task: data.task }, { status });
  } catch (error) {
    return handleError(error, END_POINT_TASK);
  }
}

export async function PUT(request: Request) {
  try {
    const { data, status } = await handleBaseRequest(
      "PUT",
      END_POINT_TASK,
      request,
    );
    return NextResponse.json({ task: data.task }, { status });
  } catch (error) {
    return handleError(error, END_POINT_TASK);
  }
}

export async function DELETE(request: Request) {
  const body = await request.json();
  const id = body.id;
  if (!id) {
    throw new HttpError(
      StatusCodes.BadRequest,
      errorMessages[ErrorCode.PAYLOAD_ID_NOT_FOUND],
    );
  }
  try {
    const { data, status } = await handleBaseRequest(
      "DELETE",
      END_POINT_TASK,
      request,
      { id },
    );
    return NextResponse.json({ task: data.task }, { status });
  } catch (error) {
    return handleError(error, END_POINT_TASK);
  }
}
