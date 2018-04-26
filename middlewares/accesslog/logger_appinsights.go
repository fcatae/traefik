package accesslog

import (
	"strings"
	"time"

	"github.com/jjcollinge/logrus-appinsights"
	"github.com/sirupsen/logrus"
)

func createAppInsightHook(logger *logrus.Logger, clientName string, clientKey string) {
	hook, err := logrus_appinsights.New(clientName, logrus_appinsights.Config{
		InstrumentationKey: clientKey,
		MaxBatchSize:       2,               // optional
		MaxBatchInterval:   time.Second * 1, // optional
	})
	if err != nil || hook == nil {
		panic(err)
	}

	hook.SetLevels(logrus.AllLevels)

	// ignore fields
	hook.AddIgnore("private")
	logger.AddHook(hook)

	// logger.Add()
}

func parseAppInsightMoniker(filename string) (string, string, bool) {
	const appInsightPrefix = "appinsights://"

	if isAppInsights := strings.HasPrefix(filename, appInsightPrefix); isAppInsights {
		parameters := strings.TrimPrefix(filename, appInsightPrefix)
		components := strings.Split(parameters, ":")

		if len(components) != 2 {
			panic("inccorect syntax for application insights")
		}

		clientName, clientKey := components[0], components[1]

		return clientName, clientKey, true
	}

	return "", "", false
}
