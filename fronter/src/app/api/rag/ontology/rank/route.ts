import { NextResponse } from "next/server";
import { handleBaseRequest, handleError } from "@/app/api/utlts/handleRequest";

const END_POINT_RAG_ONTOLOGY_RANK = "ragOntologyRank";

export async function POST(request: Request) {
  try {
    const { data, status } = await handleBaseRequest(
      "POST",
      END_POINT_RAG_ONTOLOGY_RANK,
      request,
    );
    return NextResponse.json(data, { status });
  } catch (error) {
    return handleError(error, END_POINT_RAG_ONTOLOGY_RANK);
  }
}
