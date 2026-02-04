import {
  ApplicationError,
  ErrorCode,
  getErrorMessage,
  parseErrorResponse,
} from "./errorCodes";

describe("ApplicationError", () => {
  describe("コンストラクタ", () => {
    it("エラーコードとメッセージで正しくインスタンス化される", () => {
      const error = new ApplicationError(
        ErrorCode.VAL_MISSING_FIELD,
        "カスタムメッセージ",
        "詳細情報",
      );

      expect(error.code).toBe(ErrorCode.VAL_MISSING_FIELD);
      expect(error.message).toBe("カスタムメッセージ");
      expect(error.detail).toBe("詳細情報");
      expect(error.name).toBe("ApplicationError");
      expect(error.timestamp).toBeDefined();
    });

    it("メッセージが指定されない場合、デフォルトメッセージが使用される", () => {
      const error = new ApplicationError(ErrorCode.AUTH_INVALID_CREDENTIALS);

      expect(error.message).toBe("認証情報が無効です");
    });
  });

  describe("toJSON", () => {
    it("正しいJSON形式に変換される", () => {
      const error = new ApplicationError(
        ErrorCode.RES_NOT_FOUND,
        "リソースが見つかりません",
        "ID: 123",
      );
      error.path = "/api/task/123";

      const json = error.toJSON();

      expect(json).toEqual({
        code: ErrorCode.RES_NOT_FOUND,
        message: "リソースが見つかりません",
        detail: "ID: 123",
        timestamp: error.timestamp,
        path: "/api/task/123",
      });
    });
  });
});

describe("getErrorMessage", () => {
  it("既知のエラーコードに対して正しいメッセージを返す", () => {
    const testCases = [
      {
        code: ErrorCode.AUTH_INVALID_CREDENTIALS,
        expected: "認証情報が無効です",
      },
      {
        code: ErrorCode.VAL_MISSING_FIELD,
        expected: "必須項目が入力されていません",
      },
      { code: ErrorCode.RES_NOT_FOUND, expected: "リソースが見つかりません" },
      {
        code: ErrorCode.DB_CONNECTION_FAILED,
        expected: "データベース接続に失敗しました",
      },
      {
        code: ErrorCode.SYS_INTERNAL_ERROR,
        expected: "内部エラーが発生しました",
      },
      {
        code: ErrorCode.NET_CONNECTION_FAILED,
        expected: "ネットワーク接続に失敗しました",
      },
    ];

    testCases.forEach(({ code, expected }) => {
      expect(getErrorMessage(code)).toBe(expected);
    });
  });

  it("未知のエラーコードに対してデフォルトメッセージを返す", () => {
    const unknownCode = "UNKNOWN_CODE" as keyof typeof ErrorCode;
    expect(getErrorMessage(unknownCode)).toBe("エラーが発生しました");
  });
});

describe("parseErrorResponse", () => {
  it("ApplicationErrorインスタンスをそのまま返す", () => {
    const appError = new ApplicationError(ErrorCode.VAL_INVALID_INPUT);
    const result = parseErrorResponse(appError);

    expect(result).toBe(appError);
  });

  it("サーバーレスポンスのエラーコードを解析する", () => {
    const serverError = {
      response: {
        data: {
          code: ErrorCode.AUTH_TOKEN_EXPIRED,
          message: "トークンが期限切れです",
          detail: "再ログインしてください",
        },
      },
    };

    const result = parseErrorResponse(serverError);

    expect(result.code).toBe(ErrorCode.AUTH_TOKEN_EXPIRED);
    expect(result.message).toBe("トークンが期限切れです");
    expect(result.detail).toBe("再ログインしてください");
  });

  it("ネットワークエラーを適切に処理する", () => {
    const connectionError = {
      code: "ECONNREFUSED",
      message: "connect ECONNREFUSED",
    };

    const result = parseErrorResponse(connectionError);

    expect(result.code).toBe(ErrorCode.NET_CONNECTION_FAILED);
    expect(result.message).toBe("ネットワーク接続に失敗しました");
  });

  it("タイムアウトエラーを検出する", () => {
    const timeoutError = {
      code: "ECONNABORTED",
      message: "timeout of 10000ms exceeded",
    };

    const result = parseErrorResponse(timeoutError);

    expect(result.code).toBe(ErrorCode.NET_REQUEST_ABORTED);
    expect(result.message).toBe("リクエストが中断されました");
  });

  it("メッセージからタイムアウトを検出する", () => {
    const error = {
      message: "Request timeout after 5000ms",
    };

    const result = parseErrorResponse(error);

    expect(result.code).toBe(ErrorCode.NET_REQUEST_TIMEOUT);
    expect(result.message).toBe("リクエストがタイムアウトしました");
  });

  it("予期しないエラーを処理する", () => {
    const unknownError = {
      message: "予期しないエラー",
    };

    const result = parseErrorResponse(unknownError);

    expect(result.code).toBe(ErrorCode.SYS_INTERNAL_ERROR);
    expect(result.message).toBe("予期しないエラー");
  });

  it("エラーメッセージがない場合のフォールバック", () => {
    const emptyError = {};

    const result = parseErrorResponse(emptyError);

    expect(result.code).toBe(ErrorCode.SYS_INTERNAL_ERROR);
    expect(result.message).toBe("予期しないエラーが発生しました");
  });
});

describe("エラーコードカテゴリー", () => {
  it("認証系エラーコードが正しく定義されている", () => {
    expect(ErrorCode.AUTH_INVALID_CREDENTIALS).toBe("AUTH_001");
    expect(ErrorCode.AUTH_TOKEN_EXPIRED).toBe("AUTH_002");
    expect(ErrorCode.AUTH_TOKEN_INVALID).toBe("AUTH_003");
    expect(ErrorCode.AUTH_UNAUTHORIZED).toBe("AUTH_004");
    expect(ErrorCode.AUTH_ACCOUNT_DISABLED).toBe("AUTH_005");
  });

  it("バリデーション系エラーコードが正しく定義されている", () => {
    expect(ErrorCode.VAL_INVALID_INPUT).toBe("VAL_001");
    expect(ErrorCode.VAL_MISSING_FIELD).toBe("VAL_002");
    expect(ErrorCode.VAL_INVALID_FORMAT).toBe("VAL_003");
    expect(ErrorCode.VAL_DUPLICATE_ENTRY).toBe("VAL_004");
    expect(ErrorCode.VAL_CONSTRAINT_FAILED).toBe("VAL_005");
  });

  it("リソース系エラーコードが正しく定義されている", () => {
    expect(ErrorCode.RES_NOT_FOUND).toBe("RES_001");
    expect(ErrorCode.RES_ALREADY_EXISTS).toBe("RES_002");
    expect(ErrorCode.RES_ACCESS_DENIED).toBe("RES_003");
    expect(ErrorCode.RES_LOCKED).toBe("RES_004");
  });

  it("ネットワーク系エラーコードが正しく定義されている", () => {
    expect(ErrorCode.NET_CONNECTION_FAILED).toBe("NET_001");
    expect(ErrorCode.NET_REQUEST_TIMEOUT).toBe("NET_002");
    expect(ErrorCode.NET_REQUEST_ABORTED).toBe("NET_003");
  });
});
