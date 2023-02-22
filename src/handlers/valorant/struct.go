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

type FetchMatchDetailsResponse struct {
	MatchInfo struct {
		MatchId string `json:"matchId"`
		MapId string `json:"mapId"`
		GamePodId string `json:"gamePodId"`
		GameLoopZone string `json:"gameLoopZone"`
		GameServerAddress string `json:"gameServerAddress"`
		GameVersion string `json:"gameVersion"`
		GameLengthMillis int `json:"gameLengthMillis"`
		GameStartMillis int `json:"gameStartMillis"`
		ProvisioningFlowId string `json:"provisioningFlowID"`
		IsCompleted bool `json:"isCompleted"`
		CustomGameName string `json:"customGameName"`
		ForcePostProcessing bool `json:"forcePostProcessing"`
		QueueId string `json:"queueID"`
		GameMode string `json:"gameMode"`
		IsRanked bool `json:"isRanked"`
		IsMatchSampled bool `json:"isMatchSampled"`
		SeasonId string `json:"seasonId"`
		CompletionState string `json:"completionState"`
		PlatformType string `json:"platformType"`
		PartyRrPenalties interface{} `json:"partyRRPenalties"`
		ShouldMatchDisablePenalties bool `json:"shouldMatchDisablePenalties"`
	} `json:"matchInfo"`
	Players []struct {
		Subject string `json:"subject"`
		GameName string `json:"gameName"`
		TagLine string `json:"tagLine"`
		PlatformInfo struct {
			PlatformType string `json:"platformType"`
			PlatformOs string `json:"platformOS"`
			PlatformOsVersion string `json:"platformOSVersion"`
			PlatformChipset string `json:"platformChipset"`
		} `json:"platformInfo"`
		TeamId string `json:"teamId"`
		PartyId string `json:"partyId"`
		CharacterId string `json:"characterId"`
		Stats struct {
			Score int `json:"score"`
			RoundsPlayed int `json:"roundsPlayed"`
			Kills int `json:"kills"`
			Deaths int `json:"deaths"`
			Assists int `json:"assists"`
			PlayTimeMillis int `json:"playtimeMillis"`
			AbilityCasts struct {
				GrenadeCasts int `json:"grenadeCasts"`
				Ability1Casts int `json:"ability1Casts"`
				Ability2Casts int `json:"ability2Casts"`
				UltimateCasts int `json:"ultimateCasts"`
			} `json:"abilityCasts"`
		} `json:"stats"`
		RoundDamage []struct {
			Round int `json:"round"`
			Receiver string `json:"receiver"`
			Damage int `json:"damage"`
		} `json:"roundDamage"`
		CompetitiveTier int `json:"competitiveTier"`
		IsObserver bool `json:"isObserver"`
		PlayerCard string `json:"playerCard"`
		PlayerTitle string `json:"playerTitle"`
		PreferredLevelBorder string `json:"preferredLevelBorder"`
		AccountLevel int `json:"accountLevel"`
		SessionPlaytimeMinutes int `json:"sessionPlaytimeMinutes"`
		XpModifications struct {
			Value int `json:"value"`
			Id string `json:"ID"`
		} `json:"xpModifications"`
		BehaviorFactors struct {
			AfkRounds int `json:"afkRounds"`
			Collisions float64 `json:"collisions"`
			DamageParticipationOutgoing int `json:"damageParticipationOutgoing"`
			FriendlyFireIncoming int `json:"friendlyFireIncoming"`
			FriendlyFireOutgoing int `json:"friendlyFireOutgoing"`
			MouseMovement int `json:"mouseMovement"`
			StayedInSpawnRounds int `json:"stayedInSpawnRounds"`
		} `json:"behaviorFactors"`
		NewPlayerExperienceDetails struct {
			BasicMovement struct {
				IdleTimeMillis int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"basicMovement"`
			BasicGunSkill struct {
				IdleTimeMillis int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"basicGunSkill"`
			AdaptiveBots struct {
				AdaptiveBotAverageDurationMillisAllAttempts int `json:"adaptiveBotAverageDurationMillisAllAttempts"`
				AdaptiveBotAverageDurationMillisFirstAttempt int `json:"adaptiveBotAverageDurationMillisFirstAttempt"`
				KillDetailsFirstAttempt interface{} `json:"killDetailsFirstAttempt"`
				IdleTimeMillis int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"adaptiveBots"`
			Ability struct {
				IdleTimeMillis int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"ability"`
			BombPlant struct {
				IdleTimeMillis int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"bombPlant"`
			DefendBombSite struct {
				Success bool `json:"success"`
				IdleTimeMillis int `json:"idleTimeMillis"`
				ObjectiveCompleteTimeMillis int `json:"objectiveCompleteTimeMillis"`
			} `json:"defendBombSite"`
			SettingsStatus struct {
				IsMouseSensitivityDefault bool `json:"isMouseSensitivityDefault"`
				IsCrosshairDefault bool `json:"isCrosshairDefault"`
			} `json:"settingsStatus"`
		} `json:"newPlayerExperienceDetails"`
	} `json:"players"`
	Bots interface{} `json:"bots"`
	Coaches []struct {
		Subject string `json:"subject"`
		TeamId string `json:"teamId"`
	} `json:"coaches"`
	Teams []struct {
		TeamId string `json:"teamId"`
		Won bool `json:"won"`
		RoundsPlayed int `json:"roundsPlayed"`
		RoundsWon int `json:"roundsWon"`
		NumPoints int `json:"numPoints"`
	} `json:"teams"`
	RoundResults []struct {
		RoundNum int `json:"roundNum"`
		RoundResult string `json:"roundResult"`
		RoundCeremony string `json:"roundCeremony"`
		WinningTeam string `json:"winningTeam"`
		BombPlanter string `json:"bombPlanter"`
		BombDefuser string `json:"bombDefuser"`
		PlantRoundTime int `json:"plantRoundTime"`
		PlantPlayerLocations []struct {
			Subject string `json:"subject"`
			ViewRadians float64 `json:"viewRadians"`
			Location struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"location"`
		} `json:"plantPlayerLocations"`
		PlantLocation struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"plantLocation"`
		PlantSite string `json:"plantSite"`
		DefuseRoundTime int `json:"defuseRoundTime"`
		DefusePlayerLocations []struct {
			Subject string `json:"subject"`
			ViewRadians float64 `json:"viewRadians"`
			Location struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"location"`
		} `json:"defusePlayerLocations"`
		DefuseLocation struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"defuseLocation"`
		PlayerStats []struct {
			Subject string `json:"subject"`
			Kills struct {
				GameTime int `json:"gameTime"`
				RoundTime int `json:"roundTime"`
				Killer string `json:"killer"`
				Victim string `json:"victim"`
				VictimLocation struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"victimLocation"`
				Assistants []string `json:"assistants"`
				PlayerLocations []struct {
					Subject string `json:"subject"`
					ViewRadians float64 `json:"viewRadians"`
					Location struct {
						X int `json:"x"`
						Y int `json:"y"`
					} `json:"location"`
				} `json:"playerLocations"`
				FinishingDamage struct {
					DamageType string `json:"damageType"`
					DamageItem string `json:"damageItem"`
					IsSecondaryFireMode bool `json:"isSecondaryFireMode"`
				} `json:"finishingDamage"`
			} `json:"kills"`
			Damage []struct {
				Receiver string `json:"receiver"`
				Damage int `json:"damage"`
				Legshots int `json:"legshots"`
				Bodyshots int `json:"bodyshots"`
				Headshots int `json:"headshots"`
			} `json:"damage"`
			Score int `json:"score"`
			Economy struct {
				LoadoutValue int `json:"loadoutValue"`
				Weapon string `json:"weapon"`
				Armor string `json:"armor"`
				Remaining int `json:"remaining"`
				Spent int `json:"spent"`
			} `json:"economy"`
			Ability struct {
				GrenadeEffects interface{} `json:"grenadeEffects"`
				Ability1Effects interface{} `json:"ability1Effects"`
				Ability2Effects interface{} `json:"ability2Effects"`
				UltimateEffects interface{} `json:"ultimateEffects"`
			} `json:"ability"`
			WasAfk bool `json:"wasAfk"`
			WasPenalized bool `json:"wasPenalized"`
			StayedInSpawn bool `json:"stayedInSpawn"`
		} `json:"playerStats"`
		RoundResultCode string `json:"roundResultCode"`
		PlayerEconomies []struct {
			Subject string `json:"subject"`
			LoadoutValue int `json:"loadoutValue"`
			Weapon string `json:"weapon"`
			Armor string `json:"armor"`
			Remaining int `json:"remaining"`
			Spent int `json:"spent"`
		} `json:"playerEconomies"`
		PlayerScores []struct {
			Subject string `json:"subject"`
			Score int `json:"score"`
		} `json:"playerScores"`
	} `json:"roundResults"`
	Kills []struct {
		GameTime int `json:"gameTime"`
		RoundTime int `json:"roundTime"`
		Killer string `json:"killer"`
		Victim string `json:"victim"`
		VictimLocation struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"victimLocation"`
		Assistants []string `json:"assistants"`
		PlayerLocations []struct {
			Subject string `json:"subject"`
			ViewRadians float64 `json:"viewRadians"`
			Location struct {
				X int `json:"x"`
				Y int `json:"y"`
			} `json:"location"`
		} `json:"playerLocations"`
		FinishingDamage struct {
			DamageType string `json:"damageType"`
			DamageItem string `json:"damageItem"`
			IsSecondaryFireMode bool `json:"isSecondaryFireMode"`
		} `json:"finishingDamage"`
	} `json:"kills"`
}
