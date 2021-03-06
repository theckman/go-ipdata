// Copyright (c) 2017, 2019 Tim Heckman
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file at the root of this repository.

package ipdata

// IP is a struct that represents the JSON response from the https://ipdata.co
// API.
type IP struct {
	IP           string `json:"ip"`
	ASN          ASN    `json:"asn"`
	Organization string `json:"organisation"`

	City   string `json:"city"`
	Region string `json:"region"`
	RegionCode string `json:"region_code"`
	Postal string `json:"postal"`

	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`

	Flag         string `json:"flag"`
	EmojiFlag    string `json:"emoji_flag"`
	EmojiUnicode string `json:"emoji_unicode"`

	ContinentName string `json:"continent_name"`
	ContinentCode string `json:"continent_code"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	CallingCode string `json:"calling_code"`

	IsEU bool `json:"is_eu"`

	Languages []Language `json:"language,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	TimeZone *TimeZone `json:"time_zone,omitempty"`

	Threat *Threat `json:"threat,omitempty"`
}

func (ip IP) String() string {
	return ip.IP
}

// ASN represents the Autonomous System Number data returned from the API.
type ASN struct {
	ASN    string `json:"asn"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Route  string `json:"route"`
	Type   string `json:"type"`
}

// Language represents the language object within the JSON response from the
// API. This provides information about the language(s) where that IP resides.
type Language struct {
	Name   string `json:"name"`
	Native string `json:"native"`
}

// Currency represents the currency object within the JSON response from the
// API. This provides information about the currency where that IP resides.
type Currency struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Symbol string `json:"symbol"`
	Native string `json:"native"`
	Plural string `json:"plural"`
}

// TimeZone represents the time_zone object within the JSON response from the
// API. This provides information about the timezone where that IP resides.
type TimeZone struct {
	Name         string `json:"name"`
	Abbreviation string `json:"abbr"`
	Offset       string `json:"offset"`
	IsDST        bool   `json:"is_dst"`
	CurrentTime  string `json:"current_time,omitempty"`
}

// Threat represents the threat object within the JSON response from the API.
// This provides information about what type of threat this IP may be.
type Threat struct {
	// IsTOR is true if the IP is associated with a node on the TOR (The Onion
	// Router) network
	IsTOR bool `json:"is_tor"`

	// IsProxy is true if the IP is associated with bring a proxy
	// (HTTP/HTTPS/SSL/SOCKS/CONNECT and transparent proxies)
	IsProxy bool `json:"is_proxy"`

	// IsAnonymous is true if either IsTor or IsProxy are true
	IsAnonymous bool `json:"is_anonymous"`

	// IsKnownAttacker is true if the IP address is a known source of malicious
	// activity (i.e. attacks, malware, botnet activity, etc)
	IsKnownAttacker bool `json:"is_known_attacker"`

	// IsKnownAbuser is true if the IP address is a known source of abuse
	// (i.e. spam, harvesters, registration bots, and other nuisance bots, etc)
	IsKnownAbuser bool `json:"is_known_abuser"`

	// IsThreat is true if either IsKnownAttacker or IsKnownAbuser are true
	IsThreat bool `json:"is_threat"`

	// IsBogon is true if this IP address should be within a bogon filter:
	// https://en.wikipedia.org/wiki/Bogon_filtering
	IsBogon bool `json:"is_bogon"`
}

type bulkIP struct {
	IP           string `json:"ip"`
	ASN          ASN    `json:"asn"`
	Organization string `json:"organisation"`

	City   string `json:"city"`
	Region string `json:"region"`
	Postal string `json:"postal"`

	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`

	Flag         string `json:"flag"`
	EmojiFlag    string `json:"emoji_flag"`
	EmojiUnicode string `json:"emoji_unicode"`

	ContinentName string `json:"continent_name"`
	ContinentCode string `json:"continent_code"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	CallingCode string `json:"calling_code"`

	IsEU bool `json:"is_eu"`

	Languages []Language `json:"language,omitempty"`

	Currency *Currency `json:"currency,omitempty"`

	TimeZone *TimeZone `json:"time_zone,omitempty"`

	Threat *Threat `json:"threat,omitempty"`

	Message string `json:"message"`
}

func bulkToIP(bip bulkIP) *IP {
	ip := IP{
		IP:            bip.IP,
		ASN:           bip.ASN,
		Organization:  bip.Organization,
		City:          bip.City,
		Region:        bip.Region,
		Postal:        bip.Postal,
		CountryName:   bip.CountryName,
		CountryCode:   bip.CountryCode,
		Flag:          bip.Flag,
		EmojiFlag:     bip.EmojiFlag,
		EmojiUnicode:  bip.EmojiUnicode,
		ContinentName: bip.ContinentName,
		ContinentCode: bip.ContinentCode,
		Latitude:      bip.Latitude,
		Longitude:     bip.Longitude,
		CallingCode:   bip.CallingCode,
		IsEU:          bip.IsEU,
		Currency:      bip.Currency,
		TimeZone:      bip.TimeZone,
		Threat:        bip.Threat,
	}

	if len(bip.Languages) > 0 {
		ip.Languages = bip.Languages
	}

	return &ip
}
