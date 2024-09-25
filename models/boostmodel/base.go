package boostmodel

import (
	"github.com/High-Performance-App/API-Shared/models/episodemodel"
	"github.com/High-Performance-App/API-Shared/models/savedmodel"
	"github.com/High-Performance-App/API-Shared/models/storymodel"
)

type Boost struct {
	ID                string                 `json:"id"`
	ScheduledDate     int64                  `json:"scheduledDate"`
	Stories           []storymodel.Story     `json:"stories"`
	Episodes          []episodemodel.Episode `json:"episodes"`
	AfterBoostMessage string                 `json:"after_boost_message"`
	PreBoostMessage   string                 `json:"pre_boost_message"`
}

func (b *Boost) GetSaveType() string {
	return savedmodel.SavedTypeBoost
}
