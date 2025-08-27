package config

import (
    "go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger() {
    var err error
    // 生产模式可以用 zap.NewProduction()
    Log, err = zap.NewDevelopment()
    if err != nil {
        panic(err)
    }
}
