package valorant

type DefaultResponse struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

type MessageData struct {
	Message string `json:"message"`
}

type MMRFetchPlayerResponse struct {
	Version int `json:"Version"`
	Subject string `json:"Subject"`
	NewPlayerExperienceFinished bool `json:"NewPlayerExperienceFinished"`
	QueueSkills map[string]struct {
		TotalGamesNeededForRating int `json:"TotalGamesNeededForRating"`
		TotalGamesNeededForLeaderboard int `json:"TotalGamesNeededForLeaderboard"`
		CurrentSeasonGamesNeededForRating int `json:"CurrentSeasonGamesNeededForRating"`
		SeasonalInfoBySeasonId map[string]struct {
			SeasonId string `json:"SeasonID"`
			NumberOfWins int `json:"NumberOfWins"`
			NumberOfWinsWithPlacements int `json:"NumberOfWinsWithPlacements"`
			NumberOfGames int `json:"NumberOfGames"`
			Rank int `json:"Rank"`
			CapstoneWins int `json:"CapstoneWins"`
			LeaderboardRank int `json:"LeaderboardRank"`
			CompetitiveTier int `json:"CompetitiveTier"`
			RankedRating int `json:"RankedRating"`
			WinsByTier map[string]int `json:"WinsByTier"`
			GamesNeededForRating int `json:"GamesNeededForRating"`
			TotalWinsNeededForRank int `json:"TotalWinsNeededForRank"`
		} `json:"SeasonalInfoBySeasonId"`
	} `json:"QueueSkills"`
	LatestCompetitiveUpdate struct {
		MatchId string `json:"MatchID"`
		MapId string `json:"MapID"`
		SeasonId string `json:"SeasonID"`
		MatchStartTime int `json:"MatchStartTime"`
		TierAfterUpdate int `json:"TierAfterUpdate"`
		TierBeforeUpdate int `json:"TierBeforeUpdate"`
		RankedRatingAfterUpdate int `json:"RankedRatingAfterUpdate"`
		RankedRatingBeforeUpdate int `json:"RankedRatingBeforeUpdate"`
		RankedRatingEarned int `json:"RankedRatingEarned"`
		RankedRatingPerformanceBonus int `json:"RankedRatingPerformanceBonus"`
		CompetitiveMovement string `json:"CompetitiveMovement"`
		AfkPenalty int `json:"AFKPenalty"`
	} `json:"LatestCompetitiveUpdate"`
	IsLeaderboardAnonymized bool `json:"IsLeaderboardAnonymized"`
	IsActRankBadgeHidden bool `json:"IsActRankBadgeHidden"`
}

type FetchCompetitiveUpdatesResponse struct {
	Version int `json:"Version"`
	Subject string `json:"Subject"`
	Matches []struct {
		MatchId string `json:"MatchID"`
		MapId string `json:"MapID"`
		SeasonId string `json:"SeasonID"`
		MatchStartTime int `json:"MatchStartTime"`
		TierAfterUpdate int `json:"TierAfterUpdate"`
		TierBeforeUpdate int `json:"TierBeforeUpdate"`
		RankedRatingAfterUpdate int `json:"RankedRatingAfterUpdate"`
		RankedRatingBeforeUpdate int `json:"RankedRatingBeforeUpdate"`
		RankedRatingEarned int `json:"RankedRatingEarned"`
		RankedRatingPerformanceBonus int `json:"RankedRatingPerformanceBonus"`
		CompetitiveMovement string `json:"CompetitiveMovement"`
		AfkPenalty int `json:"AFKPenalty"`
	} `json:"Matches"`
}