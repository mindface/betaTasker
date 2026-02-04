import { NextRequest, NextResponse } from "next/server";
import { cookies } from "next/headers";
import { errorMessages, ErrorCode } from "@/response/errorCodes";
import { StatusCodes } from "@/response/statusCodes";
import { HttpError } from "@/response/httpError";
import { handleBaseRequest, handleError } from "../../utlts/handleRequest";

const API_BASE_URL =
  process.env.NEXT_PUBLIC_API_BASE_URL || "http://localhost:8080";

const END_POINT_HEURISTICS_INSIGHTS = "heuristicsInsights";

export async function GET(request: NextRequest) {
  try {
    // クエリパラメータを取得
    const searchParams = request.nextUrl.searchParams;
    const queryString = searchParams.toString();

    // クッキーからトークンを取得
    const cookieStore = await cookies();
    const token = cookieStore.get("token")?.value;

    if (!token) {
      throw new HttpError(
        StatusCodes.Unauthorized,
        errorMessages[ErrorCode.AUTH_UNAUTHORIZED],
      );
    }

    // バックエンドAPIにリクエスト
    const backendRes = await fetch(
      `${API_BASE_URL}/api/heuristics/insights${queryString ? `?${queryString}` : ""}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        credentials: "include",
      },
    );

    const data = await backendRes.json();
    // レスポンスの処理
    if (!backendRes.ok) {
      throw new HttpError(data.status, data.message, data.code);
    }

    return NextResponse.json(
      {
        insights: data.insights,
      },
      {
        status: backendRes.status,
      },
    );
  } catch (error) {
    return handleError(error, END_POINT_HEURISTICS_INSIGHTS);
  }
}
