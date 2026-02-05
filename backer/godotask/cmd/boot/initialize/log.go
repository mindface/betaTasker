package initialize

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog() {
	// lumberjack がログファイル管理を担当
	logFile := &lumberjack.Logger{
		Filename:   "logs/app.log", // 出力先
		MaxSize:    100,            // MB 単位
		MaxAge:     14,             // 日数（14日で削除）
		MaxBackups: 10,             // 世代数
		Compress:   true,           // gzip 圧縮
	}

	// 標準出力 + ファイル両方に出す場合
	multi := zerolog.MultiLevelWriter(
		os.Stdout,
		logFile,
	)

	log.Logger = zerolog.New(multi).
		With().
		Timestamp().
		Logger()

	// ログレベル（本番想定）
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// タイムゾーン（任意）
	zerolog.TimeFieldFormat = time.RFC3339
}
