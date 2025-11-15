import { NextRequest, NextResponse } from 'next/server';
import { handleBaseRequest, handleError } from "../../utlts/handleRequest"
import { memoryUsage } from 'process';

const END_POINT_MEMOERY = 'memory';

export type Params = { params: Promise<{ id: string }>  };
export async function GET(
  reqest: NextRequest,
  { params }: Params
) {
  const { id } = await params;
  try {
    const { data, status } = await handleBaseRequest('GET',END_POINT_MEMOERY,reqest, { id });
    return NextResponse.json({ memory: data.memory }, { status });
  } catch (error) {
    return handleError(error,END_POINT_MEMOERY);
  }
}
