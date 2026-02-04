import { NextResponse } from "next/server";
import { handleBaseRequest, handleError } from "../../utlts/handleRequest";

const END_POINT_REGISTER = "register";

export async function POST(request: Request) {
  try {
    const { data, status } = await handleBaseRequest(
      "POST",
      END_POINT_REGISTER,
      request,
    );
    return NextResponse.json(
      { knowledge_patterns: data.knowledge_patterns },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_REGISTER);
  }
}
