package spotify

type DefaultResponse struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
}

type MessageData struct {
	Message string `json:"message"`
}

type SpotifyRefreshResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
	Scope string `json:"scope"`
}

type SpotifyTopItemsResponse struct {
	Href string `json:"href"`
	Limit int `json:"limit"`
	Next string `json:"next"`
	Offset int `json:"offset"`
	Previous string `json:"previous"`
	Total int `json:"total"`
	Items []SpotifyTopItem `json:"items"`
}

type SpotifyTopItem struct {
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href string `json:"href"`
		Total int `json:"total"`
	} `json:"followers"`
	Genres []string `json:"genres"`
	Href string `json:"href"`
	Id string `json:"id"`
	Images []struct {
		Url string `json:"url"`
		Height int `json:"height"`
		Width int `json:"width"`
	} `json:"images"`
	Name string `json:"name"`
	Popularity int `json:"popularity"`
	Type string `json:"type"`
	Uri string `json:"uri"`
}

type SpotifyAlbum struct {
	AlbumType string `json:"album_type"`
	Artists []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Id string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		Uri string `json:"uri"`
	} `json:"artists"`
	AvailableMarkets []string `json:"available_markets"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href string `json:"href"`
	Id string `json:"id"`
	Images []struct {
		Url string `json:"url"`
		Height int `json:"height"`
		Width int `json:"width"`
	} `json:"images"`
	Name string `json:"name"`
	ReleaseDate string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	TotalTracks int `json:"total_tracks"`
	Type string `json:"type"`
	Uri string `json:"uri"`
}

type SpotifyArtist struct {
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href string `json:"href"`
	Id string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Uri string `json:"uri"`
}

type SpotifyCurrentlyPlayingResponse struct {
	Device struct {
		Id string `json:"id"`
		IsActive bool `json:"is_active"`
		IsPrivateSession bool `json:"is_private_session"`
		IsRestricted bool `json:"is_restricted"`
		Name string `json:"name"`
		Type string `json:"type"`
		VolumePercent int `json:"volume_percent"`
	} `json:"device"`
	RepeatState string `json:"repeat_state"`
	ShuffleState bool `json:"shuffle_state"`
	Timestamp int `json:"timestamp"`
	Context struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Type string `json:"type"`
		Uri string `json:"uri"`
	} `json:"context"`
	ProgressMs int `json:"progress_ms"`
	Item struct {
		Album SpotifyAlbum `json:"album"`
		Artists []SpotifyArtist `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber int `json:"disc_number"`
		DurationMs int `json:"duration_ms"`
		Explicit bool `json:"explicit"`
		ExternalIds struct {
			Isrc string `json:"isrc"`
			Ean string `json:"ean"`
			Upc string `json:"upc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Id string `json:"id"`
		IsPlayable bool `json:"is_playable"`
		LinkedFrom interface{} `json:"linked_from"`
		Restrictions struct {
			Reason string `json:"reason"`
		} `json:"restrictions"`
		Name string `json:"name"`
		Popularity int `json:"popularity"`
		PreviewUrl string `json:"preview_url"`
		TrackNumber int `json:"track_number"`
		Type string `json:"type"`
		Uri string `json:"uri"`
		IsLocal bool `json:"is_local"`
	} `json:"item"`
	CurrentlyPlayingType string `json:"currently_playing_type"`
	Actions struct {
		InterruptingPlayback bool `json:"interrupting_playback"`
		Resuming bool `json:"resuming"`
		Pausing bool `json:"pausing"`
		Seeking bool `json:"seeking"`
		SkippingNext bool `json:"skipping_next"`
		SkippingPrev bool `json:"skipping_prev"`
		TogglingRepeatContext bool `json:"toggling_repeat_context"`
		TogglingShuffle bool `json:"toggling_shuffle"`
		TogglingRepeatTrack bool `json:"toggling_repeat_track"`
		TransferringPlayback bool `json:"transferring_playback"`
	} `json:"actions"`
	IsPlaying bool `json:"is_playing"`
}

type SpotifyRecentlyPlayedResponse struct {
	Href string `json:"href"`
	Limit int `json:"limit"`
	Next string `json:"next"`
	Cursors struct {
		After string `json:"after"`
		Before string `json:"before"`
	} `json:"cursors"`
	Total int `json:"total"`
	Items []SpotifyTopItem `json:"items"`
}

type SpotifyRecentlyPlayedItem struct {
	Track struct {
		Album SpotifyAlbum `json:"album"`
		Artists []SpotifyArtist `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber int `json:"disc_number"`
		DurationMs int `json:"duration_ms"`
		Explicit bool `json:"explicit"`
		ExternalIds struct {
			Isrc string `json:"isrc"`
			Ean string `json:"ean"`
			Upc string `json:"upc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Id string `json:"id"`
		IsPlayable bool `json:"is_playable"`
		LinkedFrom interface{} `json:"linked_from"`
		Restrictions struct {
			Reason string `json:"reason"`
		} `json:"restrictions"`
		Name string `json:"name"`
		Popularity int `json:"popularity"`
		PreviewUrl string `json:"preview_url"`
		TrackNumber int `json:"track_number"`
		Type string `json:"type"`
		Uri string `json:"uri"`
		IsLocal bool `json:"is_local"`
	} `json:"track"`
	PlayedAt string `json:"played_at"`
	Context struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Type string `json:"type"`
		Uri string `json:"uri"`
	} `json:"context"`
}