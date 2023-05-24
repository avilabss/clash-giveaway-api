package clashofclans

const (
	RoleNotMember = "NOT_MEMBER"
	RoleMember    = "MEMBER"
	RoleLeader    = "LEADER"
	RoleAdmin     = "ADMIN"
	RoleCoLeader  = "COLEADER"
)

const (
	WarFrequencyUnknown             = "UNKNOWN"
	WarFrequencyAlways              = "ALWAYS"
	WarFrequencyMoreThanOncePerWeek = "MORE_THAN_ONCE_PER_WEEK"
	WarFrequencyOncePerWeek         = "ONCE_PER_WEEK"
	WarFrequencyLessThanOncePerWeek = "LESS_THAN_ONCE_PER_WEEK"
	WarFrequencyNever               = "NEVER"
	WarFrequencyAny                 = "ANY"
)

const (
	ClanTypeOpen       = "OPEN"
	ClanTypeInviteOnly = "INVITE_ONLY"
	ClanTypeClosed     = "CLOSED"
)

const (
	TimeFormat = "20060102T080000.000Z"
)

type Clan struct {
	WarLeague                   WarLeague     `json:"warLeague"`
	CapitalLeague               CapitalLeague `json:"capitalLeague"`
	MemberList                  []ClanMember  `json:"memberList"`
	Tag                         string        `json:"tag"`
	RequiredVersusTrophies      int           `json:"requiredVersusTrophies"`
	IsWarLogPublic              bool          `json:"isWarLogPublic"`
	WarFrequency                string        `json:"warFrequency"`
	ClanLevel                   int           `json:"clanLevel"`
	WarWinStreak                int           `json:"warWinStreak"`
	WarWins                     int           `json:"warWins"`
	WarTies                     int           `json:"warTies"`
	WarLosses                   int           `json:"warLosses"`
	ClanPoints                  int           `json:"clanPoints"`
	RequiredTownhallLevel       int           `json:"requiredTownhallLevel"`
	ChatLanguage                Language      `json:"chatLanguage"`
	IsFamilyFriendly            bool          `json:"isFamilyFriendly"`
	ClanBuilderBasePoints       int           `json:"clanBuilderBasePoints"`
	ClanVersusPoints            int           `json:"clanVersusPoints"`
	ClanCapitalPoints           int           `json:"clanCapitalPoints"`
	RequiredTrophies            int           `json:"requiredTrophies"`
	RequiredBuilderBaseTrophies int           `json:"requiredBuilderBaseTrophies"`
	Labels                      []Label       `json:"labels"`
	Name                        string        `json:"name"`
	Location                    Location      `json:"location"`
	Type                        string        `json:"type"`
	Members                     int           `json:"members"`
	Description                 string        `json:"description"`
	ClanCapital                 ClanCapital   `json:"clanCapital"`
	BadgeUrls                   BadgeUrls     `json:"badgeUrls"`
}

type WarLeague struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CapitalLeague struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ClanMember struct {
	League              League            `json:"league"`
	BuilderBaseLeague   BuilderBaseLeague `json:"builderBaseLeague"`
	VersusTrophies      int               `json:"versusTrophies"`
	Tag                 string            `json:"tag"`
	Name                string            `json:"name"`
	Role                string            `json:"role"`
	ExpLevel            int               `json:"expLevel"`
	ClanRank            int               `json:"clanRank"`
	PreviousClanRank    int               `json:"previousClanRank"`
	Donations           int               `json:"donations"`
	DonationsReceived   int               `json:"donationsReceived"`
	Trophies            int               `json:"trophies"`
	BuilderBaseTrophies int               `json:"builderBaseTrophies"`
	PlayerHouse         PlayerHouse       `json:"playerHouse"`
}

type League struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	IconUrls IconUrls `json:"iconUrls"`
}

type IconUrls struct {
	Small  string `json:"small"`
	Tiny   string `json:"tiny"`
	Medium string `json:"medium"`
}

type BuilderBaseLeague struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PlayerHouse struct {
	Elements []PlayerHouseElement `json:"elements"`
}

type PlayerHouseElement struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

type Language struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	LanguageCode string `json:"languageCode"`
}

type Label struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	IconUrls struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
	} `json:"iconUrls"`
}

type Location struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	LocalizedName string `json:"localizedName"`
	IsCountry     bool   `json:"isCountry"`
	CountryCode   string `json:"countryCode"`
}

type ClanCapital struct {
	CapitalHallLevel int        `json:"capitalHallLevel"`
	Districts        []District `json:"districts"`
}

type District struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	DistrictHallLevel int    `json:"districtHallLevel"`
}

type BadgeUrls struct {
	Small  string `json:"small"`
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

type GoldPass struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}
