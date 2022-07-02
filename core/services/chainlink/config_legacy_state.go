package chainlink

import (
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap/zapcore"
)

func (l *legacyGeneralConfig) AppID() uuid.UUID {
	l.appIDOnce.Do(func() {
		l.appID = uuid.NewV4()
	})
	return l.appID
}

func (l *legacyGeneralConfig) SetLogLevel(lvl zapcore.Level) error {
	//TODO implement me
	panic("implement me")
}

func (l *legacyGeneralConfig) SetLogSQL(logSQL bool) {
	//TODO implement me
	panic("implement me")
}
