import { NextResponse } from "next/server";
import { errorMessages, ErrorCode } from "@/response/errorCodes";
import { StatusCodes } from "@/response/statusCodes";
import { HttpError } from "@/response/httpError";
import { handleBaseRequest, handleError } from "../utlts/handleRequest";

const END_POINT_QUALITATIVE_LABEL = "qualitativeLabel";

export async function GET() {
  try {
    const { data, status } = await handleBaseRequest(
      "GET",
      END_POINT_QUALITATIVE_LABEL,
    );
    return NextResponse.json(
      {
        qualitative_labels: data.qualitative_labels,
      },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_QUALITATIVE_LABEL);
  }
}

export async function POST(request: Request) {
  try {
    const { data, status } = await handleBaseRequest(
      "POST",
      END_POINT_QUALITATIVE_LABEL,
      request,
    );
    return NextResponse.json(
      {
        qualitative_label: data.qualitative_label,
      },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_QUALITATIVE_LABEL);
  }
}

export async function PUT(request: Request) {
  try {
    const { data, status } = await handleBaseRequest(
      "PUT",
      END_POINT_QUALITATIVE_LABEL,
      request,
    );
    return NextResponse.json(
      {
        qualitative_label: data.qualitative_label,
      },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_QUALITATIVE_LABEL);
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
      END_POINT_QUALITATIVE_LABEL,
      request,
      { id },
    );
    return NextResponse.json(
      {
        qualitative_label: data.qualitative_label,
      },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_QUALITATIVE_LABEL);
  }
}
