import { NextRequest, NextResponse } from "next/server";
import { handleBaseRequest, handleError } from "../utlts/handleRequest";

const END_POINT_MEMORY_AID = "memoryAid";

export async function GET(request: NextRequest) {
  const { searchParams } = new URL(request.url);
  const code: Record<string, string> = { code: searchParams.get("code") || "" };

  try {
    const { data, status } = await handleBaseRequest(
      "GET",
      END_POINT_MEMORY_AID,
      undefined,
      undefined,
      code,
    );
    return NextResponse.json(
      {
        contexts: data.contexts,
      },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_MEMORY_AID);
  }
}
