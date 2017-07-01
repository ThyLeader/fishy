package main

import (
	"time"
)

var (
	ScoreGuildKey  = func(guildID string) string { return "exp:guild:" + guildID }
	LocDensityKey  = func(userID string) string { return "user:locationdensity:" + userID }
	LocationKey    = func(userID string) string { return "user:location:" + userID }
	InventoryKey   = func(userID string) string { return "user:inventory:" + userID }
	BlackListKey   = func(userID string) string { return "user:blacklist:" + userID }
	GatherBaitKey  = func(userID string) string { return "user:gatherbait:" + userID }
	UserTrackKey   = func(userID string) string { return "user:" + userID }
	NoInvEEKey     = func(userID string) string { return "ee:" + userID }
	GlobalStatsKey = func(userID string) string { return "statistics:global:" + userID }
	FishInvKey     = func(userID string) string { return "fish:" + userID }
	GuildStatsKey  = func(userID, guildID string) string { return "statistics:" + guildID + ":" + userID }
	RateLimitKey   = func(cmd, userID string) string { return "ratelimit:" + cmd + ":" + userID }
	HourlyCmdTrack = func(cmd string) string { return "tracking:hourly:" + cmd }
	DailyCmdTrack  = func(cmd string) string { return "tracking:daily:" + cmd }
	TotalCmdTrack  = func(cmd string) string { return "tracking:total:" + cmd }
	Morning1       = time.Date(0, 0, 0, 9, 0, 0, 0, time.UTC)
	Morning2       = time.Date(0, 0, 0, 15, 59, 59, 999, time.UTC)
	Night1         = time.Date(0, 0, 0, 16, 0, 0, 0, time.UTC)
	Night2         = time.Date(0, 0, 0, 8, 59, 59, 999, time.UTC)
	CurrentTime    = time.Now().UTC()
)

const (
	FishyTimeout      = 10 * time.Second
	GatherBaitTimeout = 6 * time.Hour
	ScoreGlobalKey    = "exp:global"
)
