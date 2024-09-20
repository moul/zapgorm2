package zapgorm2_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

func Example() {
	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
	db, _ := gorm.Open(nil, &gorm.Config{Logger: logger})

	// do stuff normally
	_ = db // avoid "unused variable" warn
}

func setupLogsCapture() (*zap.Logger, *observer.ObservedLogs) {
	core, logs := observer.New(zap.WarnLevel)
	return zap.New(core), logs
}

func TestContextFunc(t *testing.T) {
	zaplogger, logs := setupLogsCapture()
	logger := zapgorm2.New(zaplogger)

	type ctxKey string
	key1 := ctxKey("Key")
	key2 := ctxKey("Key2")

	value1 := "Value"
	value2 := "Value2"

	ctx := context.WithValue(context.Background(), key1, value1)
	ctx = context.WithValue(ctx, key2, value2)
	logger.Context = func(ctx context.Context) []zapcore.Field {
		ctxValue, ok := (ctx.Value(key1)).(string)
		require.True(t, ok)
		ctxValue2, ok := (ctx.Value(key2)).(string)
		require.True(t, ok)
		return []zapcore.Field{zap.String(string(key1), ctxValue), zap.String(string(key2), ctxValue2)}
	}

	db, err := gorm.Open(nil, &gorm.Config{Logger: logger})
	require.NoError(t, err)

	db.Logger.Error(ctx, "test")
	require.Equal(t, 1, logs.Len())
	entry := logs.All()[0]
	require.Equal(t, zap.ErrorLevel, entry.Level)
	require.Equal(t, "test", entry.Message)
	require.Equal(t, value1, entry.ContextMap()[string(key1)])
	require.Equal(t, value2, entry.ContextMap()[string(key2)])
}

func TestTrace(t *testing.T) {
	zaplogger, logs := setupLogsCapture()
	emptyCb := func() (s string, i int64) { return }

	t.Run("ignore errors", func(t *testing.T) {
		logger := zapgorm2.Logger{
			ZapLogger:        zaplogger,
			LogLevel:         gormlogger.Warn,
			SlowThreshold:    100 * time.Millisecond,
			SkipCallerLookup: false,
			Context:          nil,
			IgnoreErrors:     errors.Join(gorm.ErrRecordNotFound, gorm.ErrDuplicatedKey),
		}

		logger.Trace(context.Background(), time.Now(), emptyCb, gorm.ErrDuplicatedKey)
		require.Equal(t, 0, logs.Len())
	})

	t.Run("log errors", func(t *testing.T) {
		logger := zapgorm2.New(zaplogger)

		logger.Trace(context.Background(), time.Now(), emptyCb, gorm.ErrDuplicatedKey)
		require.Equal(t, 1, logs.Len())
	})
}
