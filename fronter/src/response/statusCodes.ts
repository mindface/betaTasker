export const StatusCodes: Record<string, number> = {
  // 2xx Success（成功）
  OK: 200, // 成功
  Created: 201, // リソース作成成功
  Accepted: 202, // リクエスト受理（処理は非同期で完了）
  NoContent: 204, // 成功（レスポンスボディなし）- DELETE成功時など
  ResetContent: 205, // 成功（フォームをリセット）
  PartialContent: 206, // 部分的なコンテンツ（Range requestの応答）

  // 3xx Redirection（リダイレクト）
  MovedPermanently: 301, // 恒久的な移動
  Found: 302, // 一時的な移動
  SeeOther: 303, // 別のURIを参照（POST後のGETリダイレクト）
  NotModified: 304, // 変更なし（キャッシュ有効）
  TemporaryRedirect: 307, // 一時的リダイレクト（メソッド保持）
  PermanentRedirect: 308, // 恒久的リダイレクト（メソッド保持）

  // 4xx Client Error（クライアントエラー）
  BadRequest: 400, // 不正なリクエスト（構文エラー、必須パラメータ欠如）
  Unauthorized: 401, // 認証が必要（トークンなし、期限切れ）
  PaymentRequired: 402, // 支払いが必要（APIクォータ超過など）
  Forbidden: 403, // 権限がない（認証済みだがアクセス禁止）
  NotFound: 404, // リソースが存在しない
  MethodNotAllowed: 405, // HTTPメソッドが許可されていない
  NotAcceptable: 406, // Accept ヘッダーの要件を満たせない
  ProxyAuthRequired: 407, // プロキシ認証が必要
  RequestTimeout: 408, // リクエストタイムアウト
  Conflict: 409, // リソースの競合（デッドロック、重複など）
  Gone: 410, // リソースが恒久的に削除された
  LengthRequired: 411, // Content-Length ヘッダーが必要
  PreconditionFailed: 412, // 前提条件が失敗（If-Match など）
  PayloadTooLarge: 413, // リクエストボディが大きすぎる
  URITooLong: 414, // URIが長すぎる
  UnsupportedMediaType: 415, // サポートされていないメディアタイプ
  RangeNotSatisfiable: 416, // Rangeヘッダーが満たせない
  ExpectationFailed: 417, // Expectヘッダーの要件を満たせない
  ImATeapot: 418, // ティーポット（ジョークRFC）
  MisdirectedRequest: 421, // 誤ったサーバーへのリクエスト
  UnprocessableEntity: 422, // バリデーションエラー（構文は正しいが処理不可）
  Locked: 423, // リソースがロックされている
  FailedDependency: 424, // 依存する処理が失敗
  // TooEarly: 425, // リクエストが早すぎる
  UpgradeRequired: 426, // プロトコルのアップグレードが必要
  PreconditionRequired: 428, // 前提条件が必要
  TooManyRequests: 429, // レート制限超過
  RequestHeaderFieldsTooLarge: 431, // ヘッダーが大きすぎる
  UnavailableForLegalReasons: 451, // 法的理由で利用不可

  // 5xx Server Error（サーバーエラー）
  InternalServerError: 500, // サーバー内部エラー（予期しないエラー）
  NotImplemented: 501, // 未実装の機能
  BadGateway: 502, // 不正なゲートウェイ（上流サーバーからの不正な応答）
  ServiceUnavailable: 503, // サービス利用不可（メンテナンス、過負荷）
  GatewayTimeout: 504, // ゲートウェイタイムアウト（上流サーバーの応答なし）
  HTTPVersionNotSupported: 505, // HTTPバージョン未サポート
  VariantAlsoNegotiates: 506, // サーバー設定エラー
  InsufficientStorage: 507, // ストレージ不足
  LoopDetected: 508, // 無限ループ検出
  NotExtended: 510, // 拡張が必要
  NetworkAuthRequired: 511, // ネットワーク認証が必要
};
