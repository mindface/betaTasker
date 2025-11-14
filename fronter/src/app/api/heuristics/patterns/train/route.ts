import { NextRequest, NextResponse } from 'next/server';
import { handleBaseRequest, handleError } from "../../../utlts/handleRequest"

const END_POINT_HEURISTICS_TRACK = 'heuristicsTrack';

export async function POST(request: NextRequest) {
  try {
    const { data, status } = await handleBaseRequest('POST',END_POINT_HEURISTICS_TRACK);
    return NextResponse.json({
      trains: data.trains }, { status });
  } catch (error) {
    return handleError(error,END_POINT_HEURISTICS_TRACK);
  }
}