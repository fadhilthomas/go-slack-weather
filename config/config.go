package config

var base = mergeConfig(
	slackConfig,
	weatherConfig,
	cityConfig,
	timezoneConfig,
)
