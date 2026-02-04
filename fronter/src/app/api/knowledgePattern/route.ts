import { NextResponse } from "next/server";
import { handleBaseRequest, handleError } from "../utlts/handleRequest";

const END_POINT_KNOWLEDGE_PATTERN = "knowledgePattern";

export async function GET() {
  try {
    const { data, status } = await handleBaseRequest(
      "GET",
      END_POINT_KNOWLEDGE_PATTERN,
    );
    return NextResponse.json(
      { knowledge_patterns: data.knowledge_patterns },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_KNOWLEDGE_PATTERN);
  }
}

export async function POST(request: Request) {
  try {
    const { data, status } = await handleBaseRequest(
      "POST",
      END_POINT_KNOWLEDGE_PATTERN,
      request,
    );
    return NextResponse.json(
      { knowledge_pattern: data.knowledge_pattern },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_KNOWLEDGE_PATTERN);
  }
}

export async function PUT(request: Request) {
  try {
    const { data, status } = await handleBaseRequest(
      "PUT",
      END_POINT_KNOWLEDGE_PATTERN,
      request,
    );
    return NextResponse.json(
      { knowledge_pattern: data.knowledge_pattern },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_KNOWLEDGE_PATTERN);
  }
}

export async function DELETE(request: Request) {
  try {
    const { data, status } = await handleBaseRequest(
      "DELETE",
      END_POINT_KNOWLEDGE_PATTERN,
      request,
    );
    return NextResponse.json(
      { knowledge_pattern: data.knowledge_pattern },
      { status },
    );
  } catch (error) {
    return handleError(error, END_POINT_KNOWLEDGE_PATTERN);
  }
}
