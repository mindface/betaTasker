import { cookies } from "next/headers";
import { NextResponse } from "next/server";
import { URLs } from "@/constants/url";
import { errorMessages, ErrorCode } from "@/response/errorCodes";
import { StatusCodes } from "@/response/statusCodes";
import { HttpError } from "@/response/httpError";

import { handleBaseRequest, handleError } from "../utlts/handleRequest";

const END_POINT_MEMOERY = "memory";

export async function GET() {
  try {
    const { data, status } = await handleBaseRequest("GET", END_POINT_MEMOERY);
    return NextResponse.json({ memories: data.memories }, { status });
  } catch (error) {
    return handleError(error, END_POINT_MEMOERY);
  }
}

export async function POST(request: Request) {
  try {
    const { data, status } = await handleBaseRequest(
      "POST",
      END_POINT_MEMOERY,
      request,
    );
    return NextResponse.json(
      {
        memory: data.memory,
      },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_MEMOERY);
  }
}

export async function PUT(request: Request) {
  try {
    const { data, status } = await handleBaseRequest(
      "PUT",
      END_POINT_MEMOERY,
      request,
    );
    return NextResponse.json(
      {
        memory: data.memory,
      },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_MEMOERY);
  }
}

export async function DELETE(request: Request) {
  try {
    const { data, status } = await handleBaseRequest(
      "DELETE",
      END_POINT_MEMOERY,
      request,
    );
    return NextResponse.json(
      {
        memory: data.memory,
      },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_MEMOERY);
  }
}
