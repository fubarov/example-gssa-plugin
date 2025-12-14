package example_gssa_plugin

import (
	"encoding/json"
	"fmt"

	"github.com/fubarov/gssa-sdk"
)

type Example struct {
	isConfigured bool
	logger       gssa_sdk.Logger
	settings     *ExampleSettings
}

type ExampleSettings struct {
	AppID string `json:"app_id"`
}

func init() {
	gssa_sdk.RegisterPlugin(&Example{
		isConfigured: false,
		settings:     &ExampleSettings{},
	})
}

func (s *Example) Init(rawConfig json.RawMessage, logger gssa_sdk.Logger) {
	s.logger = logger
	s.isConfigured = true

	if err := json.Unmarshal(rawConfig, s.settings); err != nil {
		s.isConfigured = false
		s.logger.Error(fmt.Sprintf("Failed to parse plugin config: %v", err))
	}

	if s.settings.AppID == "" {
		s.isConfigured = false
		s.logger.Error("Plugin config is invalid: missing app_id")
	}
}

func (s *Example) SearchMoviesByImdbID(target gssa_sdk.TargetMovie) []gssa_sdk.Stream {
	if !s.isConfigured {
		s.logger.Error("Plugin is not configured")
		return []gssa_sdk.Stream{}
	}

	if target.ImdbID == "" {
		s.logger.Error("Invalid target: missing imdb_id")
		return []gssa_sdk.Stream{}
	}

	return []gssa_sdk.Stream{}
}

func (s *Example) SearchSeriesByImdbID(target gssa_sdk.TargetSeries) (streams []gssa_sdk.Stream) {
	if !s.isConfigured {
		s.logger.Error("Plugin is not configured")
		return streams
	}
	if target.ImdbID == "" {
		s.logger.Error("Invalid target: missing imdb_id")
		return streams
	}
	if target.Season == 0 {
		s.logger.Error("Invalid target: missing season")
		return streams
	}
	if target.Episode == 0 {
		s.logger.Error("Invalid target: missing episode")
		return streams
	}
	return streams
}

func (s *Example) GenerateManifest() gssa_sdk.Manifest {
	return gssa_sdk.Manifest{
		ID:            s.settings.AppID,
		Version:       "",
		Name:          "",
		Description:   "",
		Resources:     nil,
		Types:         nil,
		Catalogs:      nil,
		IDPrefixes:    nil,
		BehaviorHints: nil,
	}
}
