package memory

type PreSongSelectAddresses struct {
	Status        int64 `sig:"48 83 F8 04 73 1E"`
	SettingsClass int64 `sig:"83 E0 20 85 C0 7E 2F"`
}

type songsFolderD struct {
	SongsFolder string `mem:"[[Settings + 0xB8] + 0x4]"`
}

type PreSongSelectData struct {
	Status uint32 `mem:"[Status - 0x4]"`
}

type staticAddresses struct {
	PreSongSelectAddresses
	Base        int64 `sig:"F8 01 74 04 83 65"`
	MenuMods    int64 `sig:"C8 FF ?? ?? ?? ?? ?? 81 0D ?? ?? ?? ?? 00 08 00 00"`
	PlayTime    int64 `sig:"5E 5F 5D C3 A1 ?? ?? ?? ?? 89 ?? 04"`
	ChatChecker int64 `sig:"0A D7 23 3C 00 00 ?? 01"`
	SkinData    int64 `sig:"75 21 8B 1D"`
	Rulesets    int64 `sig:"7D 15 A1 ?? ?? ?? ?? 85 C0"`
	ChatArea    int64 `sig:"33 47 9D FF 5B 7F FF FF"`
}

type tourneyStaticAddresses struct {
	UserInfo int64 `sig:"52 30 8B C8 E8 ?? ?? ?? ?? 8B C8 8D"`
}

type tourneyD struct {
	IPCState     int32  `mem:"Ruleset + 0x54"`
	LeftStars    int32  `mem:"[Ruleset + 0x1C] + 0x2C"`
	RightStars   int32  `mem:"[Ruleset + 0x20] + 0x2C"`
	BO           int32  `mem:"[Ruleset + 0x20] + 0x30"`
	StarsVisible int8   `mem:"[Ruleset + 0x20] + 0x38"`
	ScoreVisible int8   `mem:"[Ruleset + 0x20] + 0x39"`
	TeamOneName  string `mem:"[[[Ruleset + 0x1C] + 0x20] + 0x144]"`
	TeamTwoName  string `mem:"[[[Ruleset + 0x20] + 0x20] + 0x144]"`
	TeamOneScore int32  `mem:"[Ruleset + 0x1C] + 0x28"`
	TeamTwoScore int32  `mem:"[Ruleset + 0x20] + 0x28"`
	IPCBaseAddr  uint32 `mem:"[[Ruleset + 0x34] + 0x4] + 0x4"`
}

type resultsScreenD struct {
	PlayerName string  `mem:"[[Ruleset + 0x38] + 0x28]"`
	ModsXor1   int32   `mem:"[[Ruleset + 0x38] + 0x1C] + 0xC"`
	ModsXor2   int32   `mem:"[[Ruleset + 0x38] + 0x1C] + 0x8"`
	Mode       int32   `mem:"[Ruleset + 0x38] + 0x64"`
	MaxCombo   int16   `mem:"[Ruleset + 0x38] + 0x68"`
	Score      int32   `mem:"[Ruleset + 0x38] + 0x78"`
	Hit100     int16   `mem:"[Ruleset + 0x38] + 0x88"`
	Hit300     int16   `mem:"[Ruleset + 0x38] + 0x8A"`
	Hit50      int16   `mem:"[Ruleset + 0x38] + 0x8C"`
	HitGeki    int16   `mem:"[Ruleset + 0x38] + 0x8E"`
	HitKatu    int16   `mem:"[Ruleset + 0x38] + 0x90"`
	HitMiss    int16   `mem:"[Ruleset + 0x38] + 0x92"`
	Accuracy   float64 `mem:"[Ruleset + 0x48] + 0xC"`
}

func (staticAddresses) Ruleset() string {
	return "[[Rulesets - 0xB] + 0x4]"
}

func (staticAddresses) Beatmap() string {
	return "[Base - 0xC]"
}

func (PreSongSelectAddresses) Settings() string {
	return "[SettingsClass + 0x8]"
}

func (staticAddresses) PlayContainer() string {
	return "[[[[PlayContainerBase + 0x7] + 0x4] + 0xC4] + 0x4]"
}

func (staticAddresses) Leaderboard() string {
	return "[[[LeaderboardBase+0x1] + 0x4] + 0x74] + 0x24"
}

type menuD struct {
	PreSongSelectData
	//SearchText         string  `mem:"[Ruleset + 0xCC]"`
	//GroupingType       int32   `mem:"Ruleset + 0x104"`
	MenuGameMode       int32   `mem:"[Base - 0x33]"`
	Plays              int32   `mem:"[Base - 0x33] + 0xC"`
	Artist             string  `mem:"[[Beatmap] + 0x18]"`
	ArtistOriginal     string  `mem:"[[Beatmap] + 0x1C]"`
	Title              string  `mem:"[[Beatmap] + 0x24]"`
	TitleOriginal      string  `mem:"[[Beatmap] + 0x28]"`
	AR                 float32 `mem:"[Beatmap] + 0x2C"`
	CS                 float32 `mem:"[Beatmap] + 0x30"`
	HP                 float32 `mem:"[Beatmap] + 0x34"`
	OD                 float32 `mem:"[Beatmap] + 0x38"`
	StarRatingStruct   uint32  `mem:"[Beatmap] + 0x8C"`
	AudioFilename      string  `mem:"[[Beatmap] + 0x64]"`
	BackgroundFilename string  `mem:"[[Beatmap] + 0x68]"`
	Folder             string  `mem:"[[Beatmap] + 0x74]"`
	Creator            string  `mem:"[[Beatmap] + 0x78]"`
	Name               string  `mem:"[[Beatmap] + 0x7C]"`
	Path               string  `mem:"[[Beatmap] + 0x90]"`
	Difficulty         string  `mem:"[[Beatmap] + 0xAC]"`
	MapID              int32   `mem:"[Beatmap] + 0xC8"`
	SetID              int32   `mem:"[Beatmap] + 0xCC"`
	BeatmapMode        int32   `mem:"[Beatmap] + 0x118"`
	RankedStatus       int32   `mem:"[Beatmap] + 0x128"` // unknown, unsubmitted, pending/wip/graveyard, unused, ranked, approved, qualified
	MD5                string  `mem:"[[Beatmap] + 0x6C]"`
	ObjectCount        int32   `mem:"[Beatmap] + 0xF4"`
	//Tags               string  `mem:"[[Beatmap] + 0x20]"`
	//Length       int32 `mem:"[Beatmap] + 0x12C"`
	//AudioLeadIn          int32   `mem:"[Beatmap] + 0xC0"`
	//DrainTime            int32   `mem:"[Beatmap] + 0xE8"`
	//DrainTime2           int32   `mem:"[Beatmap] + 0xEC"`
	//ScoreMenu            int32   `mem:"[Beatmap] + 0xFC"` // Local, global, mod, friend, country
	//PreviewTime  int32 `mem:"[Beatmap] + 0x118"`
}

type mainMenuD struct {
	AudioVelocityBase uint32 `mem:"[Ruleset + 0x44] + 0x10"`
}

type allTimesD struct {
	PlayTime      int32  `mem:"[PlayTime + 0x5]"`
	MenuMods      uint32 `mem:"[MenuMods + 0x9]"`
	ChatStatus    int8   `mem:"ChatChecker - 0x20"`
	SkinFolder    string `mem:"[[[SkinData + 4] + 0] + 68]"`
	ShowInterface int8   `mem:"[Settings + 0x4] + 0xC"`
}
type gameplayD struct {
	Retries             int32   `mem:"[Base - 0x33] + 0x8"`
	PlayerName          string  `mem:"[[[Ruleset + 0x60] + 0x38] + 0x28]"`
	ModsXor1            int32   `mem:"[[[Ruleset + 0x60] + 0x38] + 0x1C] + 0xC"`
	ModsXor2            int32   `mem:"[[[Ruleset + 0x60] + 0x38] + 0x1C] + 0x8"`
	HitErrors           []int32 `mem:"[[[Ruleset + 0x60] + 0x38] + 0x38]"`
	Mode                int32   `mem:"[[Ruleset + 0x60] + 0x38] + 0x64"`
	MaxCombo            int16   `mem:"[[Ruleset + 0x60] + 0x38] + 0x68"`
	Score               int32   `mem:"[[Ruleset + 0x60] + 0x38] + 0x78"`
	ScoreV2             int32   `mem:"Ruleset + 0xF8"`
	Hit100              int16   `mem:"[[Ruleset + 0x60] + 0x38] + 0x88"`
	Hit300              int16   `mem:"[[Ruleset + 0x60] + 0x38] + 0x8A"`
	Hit50               int16   `mem:"[[Ruleset + 0x60] + 0x38] + 0x8C"`
	HitGeki             int16   `mem:"[[Ruleset + 0x60] + 0x38] + 0x8E"`
	HitKatu             int16   `mem:"[[Ruleset + 0x60] + 0x38] + 0x90"`
	HitMiss             int16   `mem:"[[Ruleset + 0x60] + 0x38] + 0x92"`
	Combo               int16   `mem:"[[Ruleset + 0x60] + 0x38] + 0x94"`
	PlayerHPSmooth      float64 `mem:"[[Ruleset + 0x60] + 0x40] + 0x14"`
	PlayerHP            float64 `mem:"[[Ruleset + 0x60] + 0x40] + 0x1C"`
	Accuracy            float64 `mem:"[[Ruleset + 0x60] + 0x48] + 0xC"`
	LeaderBoard         uint32  `mem:"[Ruleset + 0x74] + 0x24"`
	KeyOverlayArrayAddr uint32  `mem:"[[Ruleset + 0xA8] + 0x10] + 0x4"` //has to be at the end due to mem not liking dead pointers, TODO: Fix this mem-side
}
