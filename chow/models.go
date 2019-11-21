package chow

// User represents information pertaining to a user
type User struct {
	ID                string   `json:"id"`
	TeamID            string   `json:"team_id"`
	Name              string   `json:"name"`
	Deleted           bool     `json:"deleted"`
	Color             string   `json:"color"`
	RealName          string   `json:"real_name"`
	Tz                string   `json:"tz"`
	TzLabel           string   `json:"tz_label"`
	TzOffset          int      `json:"tz_offset"`
	Profile           Profiles `json:"profile"`
	Inviter           Invite   `json:"inviter"`
	IsAdmin           bool     `json:"is_admin"`
	IsOwner           bool     `json:"is_owner"`
	IsPrimaryOwner    bool     `json:"is_primary_owner"`
	IsRestricted      bool     `json:"is_restricted"`
	IsUltraRestricted bool     `json:"is_ultra_restricted"`
	IsBot             bool     `json:"is_bot"`
	IsAppUser         bool     `json:"is_app_user"`
	Updated           int      `json:"updated"`
	Has2Fa            bool     `json:"has_2fa"`
}

// Profiles structure for more information relating to a user
type Profiles struct {
	Title                 string      `json:"title"`
	Phone                 string      `json:"phone"`
	Skype                 string      `json:"skype"`
	RealName              string      `json:"real_name"`
	RealNameNormalized    string      `json:"real_name_normalized"`
	DisplayName           string      `json:"display_name"`
	DisplayNameNormalized string      `json:"display_name_normalized"`
	Fields                interface{} `json:"fields"`
	StatusText            string      `json:"status_text"`
	StatusEmoji           string      `json:"status_emoji"`
	StatusExpiration      int         `json:"status_expiration"`
	AvatarHash            string      `json:"avatar_hash"`
	Email                 string      `json:"email"`
	Image24               string      `json:"image_24"`
	Image32               string      `json:"image_32"`
	Image48               string      `json:"image_48"`
	Image72               string      `json:"image_72"`
	Image192              string      `json:"image_192"`
	Image512              string      `json:"image_512"`
	StatusTextCanonical   string      `json:"status_text_canonical"`
	Team                  string      `json:"team"`
}

// Invite invitation information required by a user
type Invite struct {
	ID                string   `json:"id"`
	TeamID            string   `json:"team_id"`
	Name              string   `json:"name"`
	Deleted           bool     `json:"deleted"`
	Color             string   `json:"color"`
	RealName          string   `json:"real_name"`
	Tz                string   `json:"tz"`
	TzLabel           string   `json:"tz_label"`
	TzOffset          int      `json:"tz_offset"`
	Profile           Profiles `json:"profile"`
	IsAdmin           bool     `json:"is_admin"`
	IsOwner           bool     `json:"is_owner"`
	IsPrimaryOwner    bool     `json:"is_primary_owner"`
	IsRestricted      bool     `json:"is_restricted"`
	IsUltraRestricted bool     `json:"is_ultra_restricted"`
	IsBot             bool     `json:"is_bot"`
	IsAppUser         bool     `json:"is_app_user"`
	Updated           int      `json:"updated"`
	Has2Fa            bool     `json:"has_2fa"`
}

// UsersData representing information belonging to users in that workspace
type UsersData struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	DateCreate int    `json:"date_create"`
	DateResent int    `json:"date_resent"`
	Bouncing   bool   `json:"bouncing"`
	User       User   `json:"user"`
	Inviter    User   `json:"inviter"`
}
