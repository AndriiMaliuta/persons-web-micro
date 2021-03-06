package models

import "time"

type Cat struct {
	CatId        int
	Id           string
	Name         string
	Color        string
	Registry     string
	Origin       string
	Breed        string
	CountryCodes string `json:"country_codes"`
	WikipediaUrl string `json:"wikipedia_url"`
	Age          int
	Indoor       int
	Hairless     int
	Adaptability int
	Intelligence int
	PersonId     int64
}

type City struct {
	Id         int
	Name       string
	Population int32
}

type Person struct {
	PersonId    int64  `json:"person_id"`
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	Status      string `json:"status"`
	CountryCode string `json:"countryCode"`
	Age         int
	Cars        int
	Engaged     bool
	HasChildren bool      `json:"has_children"`
	CreatedAt   time.Time `json:"created_at"`
	CountryId   int       `json:"country_id"`
}

type Country struct {
	Id            int    `json:"id" json:"id,omitempty"`
	CountryId     string `json:"country_id,omitempty"`
	Area          int64  `json:"area" json:"area,omitempty"`
	Population    int64  `json:"population,omitempty"`
	Status        string `json:"status,omitempty"`
	UnMember      string `json:"un_member,omitempty"`
	Cca2          string `json:"cca2,omitempty"`
	Code          string `json:"code,omitempty"`
	Name          string `json:"name,omitempty"`
	ContinentName string `json:"continent_name,omitempty"`
	Region        string `json:"region,omitempty"`
	Capital       string `json:"capital,omitempty"`
	Languages     string `json:"languages,omitempty"`
	Borders       string `json:"borders,omitempty"`
	Continents    string `json:"continents,omitempty"`
	Independent   bool   `json:"Independent,omitempty"`
}

type CountryJson struct {
	Name struct {
		Common     string `json:"common"`
		Official   string `json:"official"`
		NativeName struct {
			Ukr struct {
				Official string `json:"official"`
				Common   string `json:"common"`
			} `json:"ukr"`
		} `json:"nativeName"`
	} `json:"name"`
	Cca2        string `json:"cca2"`
	Independent bool   `json:"independent"`
	Status      string `json:"status"`
	UnMember    bool   `json:"unMember"`
	Idd         struct {
		Root     string   `json:"root"`
		Suffixes []string `json:"suffixes"`
	} `json:"idd"`
	Capital   []string `json:"capital"`
	Region    string   `json:"region"`
	Subregion string   `json:"subregion"`
	Languages struct {
		Ukr string `json:"ukr"`
	} `json:"languages"`
	Landlocked bool     `json:"landlocked"`
	Borders    []string `json:"borders"`
	Area       int      `json:"area"`
	Flag       string   `json:"flag"`
	Maps       struct {
		GoogleMaps     string `json:"googleMaps"`
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
	Population int    `json:"population"`
	Fifa       string `json:"fifa"`
	Car        struct {
		Signs []string `json:"signs"`
		Side  string   `json:"side"`
	} `json:"car"`
	Timezones  []string `json:"timezones"`
	Continents []string `json:"continents"`
	Flags      struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags"`
	CoatOfArms struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"coatOfArms"`
	StartOfWeek string `json:"startOfWeek"`
	CapitalInfo struct {
		Latlng []float64 `json:"latlng"`
	} `json:"capitalInfo"`
	PostalCode struct {
		Format string `json:"format"`
		Regex  string `json:"regex"`
	} `json:"postalCode"`
}

type Comment struct {
	Id        int64     `json:"id,omitempty"`
	Body      string    `json:"body,omitempty"`
	Category  string    `json:"category,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	AuthorId  int64     `json:"author_id" json:"authorId,omitempty"`
}
