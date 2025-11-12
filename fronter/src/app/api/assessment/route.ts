import { NextResponse } from 'next/server';
import { handleBaseRequest, handleError } from "../utlts/handleRequest"

const END_POINT_ASSESSMENT = 'assessment';

export async function GET() {
  try {
    const { data, status } = await handleBaseRequest('GET',END_POINT_ASSESSMENT);
    return NextResponse.json({ assessments: data.assessments }, { status });
  } catch (error) {
    return handleError(error,END_POINT_ASSESSMENT);
  }
}

export async function POST(request: Request) {
  try {
    const { data, status } = await handleBaseRequest('POST',END_POINT_ASSESSMENT,request);
    return NextResponse.json({ assessment: data.assessment }, { status });
  } catch (error) {
    return handleError(error,END_POINT_ASSESSMENT);
  }
}

export async function PUT(request: Request) {
  try {
    const { data, status } = await handleBaseRequest('PUT',END_POINT_ASSESSMENT,request);
    return NextResponse.json({ assessment: data.assessment }, { status });
  } catch (error) {
    return handleError(error,END_POINT_ASSESSMENT);
  }
}

export async function DELETE(request: Request) {
  try {
    const { data, status } = await handleBaseRequest('DELETE',END_POINT_ASSESSMENT,request);
    return NextResponse.json({ assessment: data.assessment }, { status });
  } catch (error) {
    return handleError(error,END_POINT_ASSESSMENT);
  }
} 
