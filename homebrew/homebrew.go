package homebrew

import "encoding/json"

type Versions struct {
	Stable      string `json:"stable"`
	Bottle      bool   `json:"bottle"`
	Development string `json:"devel"`
	Head        string `json:"head"`
}

type InstalledVersion struct {
	Version          string   `json:"version"`
	UsedOptions      []string `json:"used_options"`
	BuiltAsBottle    bool     `json:"built_as_bottle"`
	PouredFromBottle bool     `json:"poured_from_bottle"`
}

type Requirement struct {
	Name           string `json:"name"`
	DefaultFormula string `json:"default_formula"`
	Cask           string `json:"cask"`     // todo(antti): is string ok?
	Download       string `json:"download"` // todo(antti): what shoudl this type be?
}

type Option struct {
	Option      string `json:"option"`
	Description string `json:"description"`
}

type Bottle struct {
	Revision int32           `json:"option"`
	Cellar   string          `json:"cellar"`
	Prefix   string          `json:"prefix"`
	RootUrl  string          `json:"root_url"`
	Files    map[string]File `json:"files"`
}

type File struct {
	URL    string `json:"url"`
	Sha256 string `json:"sha256"`
}

type Package struct {
	Name                    string             `json:"name"`
	FullName                string             `json:"full_name"`
	Description             string             `json:"desc"`
	Homepage                string             `json:"homepage"`
	OldName                 string             `json:"oldname,omitempty"`
	Aliases                 []string           `json:"aliases"`
	Versions                Versions           `json:"versions"`
	Revision                int32              `json:"revision"`
	Installed               []InstalledVersion `json:"installed"`
	LinkedKeg               string             `json:"linked_keg"`
	Pinned                  bool               `json:"pinned"`
	Outdated                bool               `json:"outdated"`
	KegOnly                 bool               `json:"keg_only"` // todo(antti): what is this?
	Dependencies            []string           `json:"dependencies"`
	RecommendedDependencies []string           `json:"recommended_dependencies"`
	OptionalDependencies    []string           `json:"optional_dependencies"`
	BuildDependencies       []string           `json:"build_dependencies"`
	ConflictsWith           []string           `json:"conflicts_with"`
	Caveats                 string             `json:"caveats"`
	Requirements            []Requirement      `json:"requirements"`
	Options                 []Option           `json:"options"`
	Bottle                  map[string]Bottle  `json:"bottle"`
}

func Unmarshal(data []byte) ([]Package, error) {
	var output []Package

	if err := json.Unmarshal(data, &output); err != nil {
		return nil, err
	}

	return output, nil
}
