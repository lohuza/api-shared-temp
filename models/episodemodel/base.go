package episodemodel

import (
	"github.com/lohuza/api-shared-temp/models/savedmodel"
	"github.com/lohuza/api-shared-temp/models/storymodel"
)

type Episode struct {
	ID              string             `json:"id"`
	Type            string             `json:"type"`
	Title           string             `json:"title"`
	Number          string             `json:"number"`
	Duration        float32            `json:"duration"`
	Description     string             `json:"description"`
	Category        EpisodeCategory    `json:"category"`
	AudioLink       string             `json:"audio_link"`
	DownloadLink    string             `json:"download_link"`
	VideoPlaybackId string             `json:"video_playback_id"`
	VideoAssetID    string             `json:"-"`
	EpisodeGuest    string             `json:"episode_guest"`
	Thumbnail       string             `json:"thumbnail"`
	PublicationDate uint64             `json:"publication_date"`
	ExternalLinks   []ExternalLink     `json:"external_links"`
	RewindStories   []Rewind           `json:"rewind_stories"`
	Stories         []storymodel.Story `json:"stories"`
	Tags            []string           `json:"tags"`
	SubCategories   []string           `json:"sub_categories"`
}

type ExternalLink struct {
	Type  string `json:"type"`
	Url   string `json:"url"`
	Title string `json:"title"`
}

type EpisodeCategory struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Icon           string `json:"icon"`
	ColorMain      string `json:"color_main"`
	ColorSecondary string `json:"color_secondary"`
}

type Rewind struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Title       string  `json:"title"`
	Duration    float32 `json:"duration"`
	Thumbnail   string  `json:"thumbnail"`
	ContentLink string  `json:"content_link"`
}

func (e *Episode) GetSaveType() string {
	return savedmodel.SavedTypeEpisode
}
