package config

var base = mergeConfig(
	cityConfig,
	logLevelConfig,
	messageConfig,
	slackConfig,
	timezoneConfig,
	weatherConfig,
)
