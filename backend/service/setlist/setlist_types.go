package setlist

type SetlistArtist struct {
	Type         string `json:"type"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Page         int    `json:"page"`
	Total        int    `json:"total"`
	Artist       []struct {
		Mbid           string `json:"mbid"`
		Name           string `json:"name"`
		SortName       string `json:"sortName"`
		Disambiguation string `json:"disambiguation"`
		URL            string `json:"url"`
	} `json:"artist"`
}

type Artist_MBID_Setlists struct {
	Type         string `json:"type"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Page         int    `json:"page"`
	Total        int    `json:"total"`
	Setlist      []struct {
		ID          string `json:"id"`
		VersionID   string `json:"versionId"`
		EventDate   string `json:"eventDate"`
		LastUpdated string `json:"lastUpdated"`
		Artist      struct {
			Mbid           string `json:"mbid"`
			Name           string `json:"name"`
			SortName       string `json:"sortName"`
			Disambiguation string `json:"disambiguation"`
			URL            string `json:"url"`
		} `json:"artist"`
		Venue struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			City struct {
				ID        string `json:"id"`
				Name      string `json:"name"`
				State     string `json:"state"`
				StateCode string `json:"stateCode"`
				Coords    struct {
					Lat  float64 `json:"lat"`
					Long float64 `json:"long"`
				} `json:"coords"`
				Country struct {
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"country"`
			} `json:"city"`
			URL string `json:"url"`
		} `json:"venue"`
		Sets struct {
			Set []struct {
				Song []struct {
					Name string `json:"name"`
					With struct {
						Mbid           string `json:"mbid"`
						Name           string `json:"name"`
						SortName       string `json:"sortName"`
						Disambiguation string `json:"disambiguation"`
						URL            string `json:"url"`
					} `json:"with"`
				} `json:"song"`
			} `json:"set"`
		} `json:"sets"`
		URL  string `json:"url"`
		Tour struct {
			Name string `json:"name"`
		} `json:"tour,omitempty"`
		Info string `json:"info,omitempty"`
	} `json:"setlist"`
}
