package model

import (
	"time"
)

type Spider struct {
	ID           uint      `json:"id"`
	Stype        string    `json:"stype"`
	UserName     string    `json:"user_name"`
	Url          string    `json:"url"`
	OnlineUrl    string    `json:"online_url"`
	DownloadAddr string    `json:"download_addr"`
	Cover        string    `json:"cover"`
	SavePath     string    `json:"save_path"`
	Info         string    `json:"info"`
	AwemeID      string    `json:"aweme_id"`
	HasDown      int       `json:"has_down"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

const (
	DySpider string = "dy"
)

type SpiderUserInfo struct {
	ID               uint      `json:"id"`
	AwemeCount       int       `json:"aweme_count"`
	TotalFavorited   string    `json:"total_favorited"`
	FollowersDetail  string    `json:"followers_detail" gorm:"type:varchar(4000);"`
	Region           string    `json:"region"`
	UID              string    `json:"uid"`
	Nickname         string    `json:"nickname"`
	FollowerCount    int       `json:"follower_count"`
	CustomVerify     string    `json:"custom_verify"`
	ShortID          string    `json:"short_id"`
	FollowingCount   int       `json:"following_count"`
	FavoritingCount  int       `json:"favoriting_count"`
	UniqueID         string    `json:"unique_id"`
	VerificationType int       `json:"verification_type"`
	PlatformSyncInfo string    `json:"platform_sync_info"`
	Geofencing       string    `json:"geofencing"`
	IsGovMediaVip    int       `json:"is_gov_media_vip"`
	AvatarLarger     string    `json:"avatar_larger" gorm:"type:text;"`
	AvatarThumb      string    `json:"avatar_thumb" gorm:"type:text;"`
	PolicyVersion    string    `json:"policy_version"`
	TypeLabel        string    `json:"type_label"`
	OriginalMusician string    `json:"original_musician"`
	Secret           int       `json:"secret"`
	Signature        string    `json:"signature"`
	AvatarMedium     string    `json:"avatar_medium" gorm:"type:text;"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type SpiderUserForm struct {
	StatusCode int `json:"status_code"`
	UserInfo   struct {
		AwemeCount       int    `json:"aweme_count"`
		TotalFavorited   string `json:"total_favorited"`
		FollowersDetail  string `json:"followers_detail"`
		Region           string `json:"region"`
		UID              string `json:"uid"`
		Nickname         string `json:"nickname"`
		FollowerCount    int    `json:"follower_count"`
		CustomVerify     string `json:"custom_verify"`
		ShortID          string `json:"short_id"`
		FollowingCount   int    `json:"following_count"`
		FavoritingCount  int    `json:"favoriting_count"`
		UniqueID         string `json:"unique_id"`
		VerificationType int    `json:"verification_type"`
		PlatformSyncInfo string `json:"platform_sync_info"`
		Geofencing       string `json:"geofencing"`
		IsGovMediaVip    bool   `json:"is_gov_media_vip"`
		AvatarLarger     struct {
			URI     string   `json:"uri"`
			URLList []string `json:"url_list"`
		} `json:"avatar_larger"`
		AvatarThumb struct {
			URI     string   `json:"uri"`
			URLList []string `json:"url_list"`
		} `json:"avatar_thumb"`
		PolicyVersion    interface{} `json:"policy_version"`
		TypeLabel        interface{} `json:"type_label"`
		OriginalMusician struct {
			MusicCount     int `json:"music_count"`
			MusicUsedCount int `json:"music_used_count"`
		} `json:"original_musician"`
		Secret       int    `json:"secret"`
		Signature    string `json:"signature"`
		AvatarMedium struct {
			URI     string   `json:"uri"`
			URLList []string `json:"url_list"`
		} `json:"avatar_medium"`
	} `json:"user_info"`
	Extra struct {
		Now   int64  `json:"now"`
		Logid string `json:"logid"`
	} `json:"extra"`
}

type AwemeList struct {
	ID           uint      `json:"id"`
	ChaList      string    `json:"cha_list"`
	AwemeType    int       `json:"aweme_type"`
	CommentList  string    `json:"comment_list"`
	VideoLabels  string    `json:"video_labels"`
	ImageInfos   string    `json:"image_infos"`
	Geofencing   string    `json:"geofencing"`
	Desc         string    `json:"desc"`
	Video        string    `json:"video"  gorm:"type:text;"`
	Statistics   string    `json:"statistics" gorm:"type:text;"`
	TextExtra    string    `json:"text_extra"  gorm:"type:text;"`
	Promotions   string    `json:"promotions"`
	LongVideo    string    `json:"long_video"`
	LabelTopText string    `json:"label_top_text"`
	Images       string    `json:"images"`
	AwemeID      string    `json:"aweme_id"`
	Author       string    `json:"author"  gorm:"type:text;"`
	VideoText    string    `json:"video_text"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time
}

type AwemeListForm struct {
	Extra struct {
		Now   int64  `json:"now"`
		Logid string `json:"logid"`
	} `json:"extra"`
	StatusCode int `json:"status_code"`
	AwemeList  []struct {
		ChaList     interface{} `json:"cha_list"`
		AwemeType   int         `json:"aweme_type"`
		CommentList interface{} `json:"comment_list"`
		VideoLabels interface{} `json:"video_labels"`
		ImageInfos  interface{} `json:"image_infos"`
		Geofencing  interface{} `json:"geofencing"`
		Desc        string      `json:"desc"`
		Video       struct {
			Cover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"cover"`
			Height   int `json:"height"`
			PlayAddr struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"play_addr"`
			DynamicCover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"dynamic_cover"`
			HasWatermark bool   `json:"has_watermark"`
			Vid          string `json:"vid"`
			DownloadAddr struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"download_addr"`
			Duration    int `json:"duration"`
			IsLongVideo int `json:"is_long_video"`
			Width       int `json:"width"`
			OriginCover struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"origin_cover"`
			Ratio         string `json:"ratio"`
			PlayAddrLowbr struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"play_addr_lowbr"`
			BitRate interface{} `json:"bit_rate"`
		} `json:"video"`
		Statistics struct {
			DiggCount    int    `json:"digg_count"`
			PlayCount    int    `json:"play_count"`
			ShareCount   int    `json:"share_count"`
			ForwardCount int    `json:"forward_count"`
			AwemeID      string `json:"aweme_id"`
			CommentCount int    `json:"comment_count"`
		} `json:"statistics"`
		TextExtra []struct {
			Start       int    `json:"start"`
			End         int    `json:"end"`
			Type        int    `json:"type"`
			HashtagName string `json:"hashtag_name"`
			HashtagID   int64  `json:"hashtag_id"`
			UserID      string `json:"user_id,omitempty"`
		} `json:"text_extra"`
		Promotions   interface{} `json:"promotions"`
		LongVideo    interface{} `json:"long_video"`
		LabelTopText interface{} `json:"label_top_text"`
		Images       interface{} `json:"images"`
		AwemeID      string      `json:"aweme_id"`
		Author       struct {
			CustomVerify     string      `json:"custom_verify"`
			VerificationType int         `json:"verification_type"`
			PlatformSyncInfo interface{} `json:"platform_sync_info"`
			PolicyVersion    interface{} `json:"policy_version"`
			AvatarLarger     struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"avatar_larger"`
			FollowerCount     int         `json:"follower_count"`
			FavoritingCount   int         `json:"favoriting_count"`
			StoryOpen         bool        `json:"story_open"`
			UserCanceled      bool        `json:"user_canceled"`
			UniqueID          string      `json:"unique_id"`
			Secret            int         `json:"secret"`
			Rate              int         `json:"rate"`
			Geofencing        interface{} `json:"geofencing"`
			IsGovMediaVip     bool        `json:"is_gov_media_vip"`
			ShortID           string      `json:"short_id"`
			FollowingCount    int         `json:"following_count"`
			WithCommerceEntry bool        `json:"with_commerce_entry"`
			TotalFavorited    string      `json:"total_favorited"`
			VideoIcon         struct {
				URI     string        `json:"uri"`
				URLList []interface{} `json:"url_list"`
			} `json:"video_icon"`
			WithFusionShopEntry bool   `json:"with_fusion_shop_entry"`
			UID                 string `json:"uid"`
			AvatarMedium        struct {
				URI     string   `json:"uri"`
				URLList []string `json:"url_list"`
			} `json:"avatar_medium"`
			AwemeCount      int           `json:"aweme_count"`
			WithShopEntry   bool          `json:"with_shop_entry"`
			SecUID          string        `json:"sec_uid"`
			HasOrders       bool          `json:"has_orders"`
			Signature       string        `json:"signature"`
			FollowStatus    int           `json:"follow_status"`
			FollowersDetail interface{}   `json:"followers_detail"`
			IsAdFake        bool          `json:"is_ad_fake"`
			Region          string        `json:"region"`
			TypeLabel       []interface{} `json:"type_label"`
			Nickname        string        `json:"nickname"`
			AvatarThumb     struct {
				URLList []string `json:"url_list"`
				URI     string   `json:"uri"`
			} `json:"avatar_thumb"`
			EnterpriseVerifyReason string `json:"enterprise_verify_reason"`
		} `json:"author"`
		VideoText interface{} `json:"video_text"`
	} `json:"aweme_list"`
	MaxCursor int64 `json:"max_cursor"`
	MinCursor int64 `json:"min_cursor"`
	HasMore   bool  `json:"has_more"`
}
