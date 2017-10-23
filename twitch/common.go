package twitch

import "time"

// Channel object
type ChannelS struct {
	Name                         string        `json:"name,omitempty"`
	Status                       string        `json:"status,omitempty"`
	Game                         string        `json:"game,omitempty"`
	Delay                        int           `json:"delay,omitempty"`
	Id                           int           `json:"_id,omitempty"`
	CreatedAt                    string        `json:"created_at,omitempty"`
	UpdatedAt                    string        `json:"updated_at,omitempty"`
	PrimaryTeamName              string        `json:"primary_team_name,omitempty"`
	PrimaryTeamDisplayName       string        `json:"primary_team_display_name,omitempty"`
	Teams                        []TeamS       `json:"teams,omitempty"`
	Title                        string        `json:"title,omitempty"`
	Mature                       bool          `json:"mature,omitempty"`
	AbuseReported                bool          `json:"abuse_reported,omitempty"`
	Banner                       string        `json:"banner,omitempty"`
	VideoBanner                  string        `json:"video_banner,omitempty"`
	Views                        int           `json:"views,omitempty"`
	Followers                    int           `json:"followers,omitempty"`
	Background                   string        `json:"background,omitempty"`
	ProfileBanner                string        `json:"profile_banner,omitempty"`
	ProfileBannerBackgroundColor string        `json:"profile_banner_background_color,omitempty"`
	Links                        ChannelLinksS `json:"_links,omitempty"`
	Logo                         string        `json:"logo,omitempty"`
	Url                          string        `json:"url,omitempty"`
	DisplayName                  string        `json:"display_name,omitempty"`
	Language                     string        `json:"language,omitempty"`
	BroadcasterLanguage          string        `json:"broadcaster_language,omitempty"`

	// authenticated
	StreamKey string `json:"stream_key,omitempty"`
	Login     string `json:"login,omitempty"`
	Email     string `json:"email,omitempty"`
}

// Team object
type TeamS struct {
	Id          int    `json:"_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Background  string `json:"background,omitempty"`
	Banner      string `json:"banner,omitempty"`
	Logo        string `json:"logo,omitempty"`
	Info        string `json:"info,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// Stream oject
type StreamS struct {
	Id                int      `json:"_id,omitempty"`
	Game              string   `json:"game,omitempty"`
	Preview           PreviewS `json:"preview,omitempty"`
	Viewers           int      `json:"viewers,omitempty"`
	Channel           ChannelS `json:"channel,omitempty"`
	VideoHeight       int      `json:"video_height,omitempty"`
	AverageFPS        float64  `json:"average_fps,omitempty"`
	Delay             int      `json:"delay,omitempty"`
	BroadcastPlatform string   `json:"broadcast_platform,omitempty"`
	CommunityId       string   `json:"community_id,omitempty"`
	CommunityIds      []string `json:"community_ids,omitempty"`
	CreatedAt         string   `json:"created_at,omitempty"`
	IsPlaylist        bool     `json:"is_playlist"`
	StreamType        string   `json:"stream_type,omitempty"`
}

// User object
type UserS struct {
	Name        string    `json:"name,omitempty"`
	Logo        string    `json:"logo,omitempty"`
	ID          int       `json:"_id,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Type        string    `json:"type,omitempty"`
	Bio         string    `json:"bio,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// Video object
type VideoS struct {
	Title       string   `json:"title,omitempty"`
	ID          string   `json:"_id,omitempty"`
	Embed       string   `json:"embed,omitempty"`
	Url         string   `json:"url,omitempty"`
	Views       int      `json:"views,omitempty"`
	Preview     string   `json:"preview,omitempty"`
	Length      int      `json:"length,omitempty"`
	Description string   `json:"description,omitempty"`
	BroadcastId int      `json:"broadcast_id"`
	RecordedAt  string   `json:"recorded_at,omitempty"`
	Game        string   `json:"game,omitempty"`
	Channel     ChannelS `json:"channel,omitempty"`
}

type FStreamS struct {
	Stream StreamS `json:"stream,omitempty"`
	Text   string  `json:"text,omitempty"`
	Image  string  `json:"image,omitempty"`
}

type LinksS struct {
	Next string `json:"next,omitempty"`
}

type ChannelLinksS struct {
	Chat         string `json:"chat,omitempty"`
	Commercial   string `json:"commercial,omitempty"`
	Videos       string `json:"videos,omitempty"`
	Teams        string `json:"teams,omitempty"`
	Editors      string `json:"editors,omitempty"`
	Subsciptions string `json:"subscriptions,omitempty"`
	Features     string `json:"features,omitempty"`
	StreamKey    string `json:"stream_key,omitempty"`
	Follows      string `json:"follows,omitempty"`
}

type PreviewS struct {
	Small    string `json:"small,omitempty"`
	Medium   string `json:"medium,omitempty"`
	Large    string `json:"large,omitempty"`
	Template string `json:"template,omitempty"`
}

type ListOptions struct {
	Game       string `url:"game,omitempty"`
	Channel    string `url:"channel,omitempty"`
	Direction  string `url:"direction,omitempty"`
	Period     string `url:"period,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
	Embeddable *bool  `url:"embeddable,omitempty"`
	Hls        *bool  `url:"hls,omitempty"`
	Live       *bool  `url:"live,omitempty"`
}

type M3U8Options struct {
	Token          string `url:"token,omitempty"`
	Sig            string `url:"sig,omitempty"`
	Player         string `url:"player,omitempty"`
	AllowAudioOnly bool   `url:"$allow_audio_only"`
	AllowSource    bool   `url:"allow_source"`
	Type           string `url:"type,omitempty"`
	Random         int    `url:"p,omitempty"`
}
