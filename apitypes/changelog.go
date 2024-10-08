package apitypes

import (
	"encoding/json"
	"time"

	"github.com/guregu/null/v5"
)

// Represents the Changelog returned by the API via json encoding.
// Implements json un-/marshaling.
type Changelog struct {
	ID          string
	WorkspaceID string
	Subdomain   string
	Domain      null.String
	Title       null.String
	Subtitle    null.String
	Logo        Logo
	Source      Source
	CreatedAt   time.Time
}

func (l Changelog) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID          string    `json:"id"`
		WorkspaceID string    `json:"workspaceId"`
		Subdomain   string    `json:"subdomain"`
		Title       string    `json:"title,omitempty"`
		Domain      string    `json:"domain,omitempty"`
		Subtitle    string    `json:"subtitle,omitempty"`
		Logo        Logo      `json:"logo"`
		Source      Source    `json:"source,omitempty"`
		CreatedAt   time.Time `json:"createdAt"`
	}{
		ID:          l.ID,
		WorkspaceID: l.WorkspaceID,
		Subdomain:   l.Subdomain,
		Domain:      l.Domain.ValueOrZero(),
		Title:       l.Title.ValueOrZero(),
		Subtitle:    l.Subtitle.ValueOrZero(),
		Logo:        l.Logo,
		Source:      l.Source,
		CreatedAt:   l.CreatedAt,
	})
}

func (c *Changelog) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	// We'll store the error (if any) so we can return it if necessary
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	if idRaw, ok := objMap["id"]; ok {
		err = json.Unmarshal(*idRaw, &c.ID)
		if err != nil {
			return err
		}
	}

	if workspaceIdRaw, ok := objMap["workspaceId"]; ok {
		err = json.Unmarshal(*workspaceIdRaw, &c.WorkspaceID)
		if err != nil {
			return err
		}
	}

	if subdomainRaw, ok := objMap["subdomain"]; ok {
		err = json.Unmarshal(*subdomainRaw, &c.Subdomain)
		if err != nil {
			return err
		}
	}

	if domainRaw, ok := objMap["domain"]; ok {
		err = json.Unmarshal(*domainRaw, &c.Domain)
		if err != nil {
			return err
		}
	}

	if titleRaw, ok := objMap["title"]; ok {
		err = json.Unmarshal(*titleRaw, &c.Title)
		if err != nil {
			return err
		}
	}

	if subtitleRaw, ok := objMap["subtitle"]; ok {
		err = json.Unmarshal(*subtitleRaw, &c.Subtitle)
		if err != nil {
			return err
		}
	}

	if logoRaw, ok := objMap["logo"]; ok {
		err = json.Unmarshal(*logoRaw, &c.Logo)
		if err != nil {
			return err
		}
	}

	if createdAtRaw, ok := objMap["createdAt"]; ok {
		err = json.Unmarshal(*createdAtRaw, &c.CreatedAt)
		if err != nil {
			return err
		}
	}

	if sourceRaw, ok := objMap["source"]; ok && sourceRaw != nil {
		c.Source = DecodeSource(*sourceRaw)
	}

	return nil
}

func DecodeSource(in json.RawMessage) Source {
	var sourceMap map[string]json.RawMessage
	err := json.Unmarshal(in, &sourceMap)
	if err != nil {
		return nil
	}

	typeRaw, ok := sourceMap["type"]
	if !ok {
		// No source type specified, so no source is set.
		return nil
	}

	var Type string
	err = json.Unmarshal(typeRaw, &Type)
	if err != nil {
		return nil
	}

	switch Type {
	case string(GitHub):
		var ghSource GHSource
		err = json.Unmarshal(in, &ghSource)
		if err != nil {
			return nil
		}
		return ghSource
	}
	return nil
}

type Logo struct {
	Src    null.String `json:"src"`
	Link   null.String `json:"link"`
	Alt    null.String `json:"alt"`
	Height null.String `json:"height"`
	Width  null.String `json:"width"`
}

type CreateChangelogBody struct {
	Title    null.String `json:"title"`
	Subtitle null.String `json:"subtitle"`
	Logo     Logo        `json:"logo"`
	Domain   null.String `json:"domain"`
}

type UpdateChangelogBody struct {
	CreateChangelogBody
	Subdomain null.String `json:"subdomain"`
}
