package homebrew

import (
	"io/ioutil"
	"path"
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	tests := []struct {
		File   string
		Result []Package
		Error  bool
	}{
		{
			File:  "broken.json",
			Error: true,
		},
		{
			File: "wercker.json",
			Result: []Package{
				{
					Name:        "wercker-cli",
					FullName:    "wercker/wercker/wercker-cli",
					Description: "wercker command line interface for building and running containers",
					Homepage:    "http://wercker.com",
					Aliases:     []string{},
					Versions: Versions{
						Stable:      "1.0.405",
						Bottle:      false,
						Development: "",
						Head:        "",
					},
					Revision: 0,
					Installed: []InstalledVersion{
						{
							Version:          "1.0.405",
							UsedOptions:      []string{},
							BuiltAsBottle:    false,
							PouredFromBottle: false,
						},
					},
					LinkedKeg:               "1.0.405",
					Pinned:                  false,
					Outdated:                false,
					KegOnly:                 false,
					Dependencies:            []string{},
					RecommendedDependencies: []string{},
					OptionalDependencies:    []string{},
					BuildDependencies:       []string{},
					ConflictsWith:           []string{},
					Requirements:            []Requirement{},
					Options:                 []Option{},
					Bottle:                  map[string]Bottle{},
				},
			},
		},
		{
			File: "git.json",
			Result: []Package{
				{
					Name:        "git",
					FullName:    "git",
					Description: "Distributed revision control system",
					Homepage:    "https://git-scm.com",
					Aliases:     []string{},
					Versions: Versions{
						Stable:      "2.9.0",
						Bottle:      true,
						Development: "",
						Head:        "HEAD",
					},
					Revision: 0,
					Installed: []InstalledVersion{
						{
							Version:          "2.8.1",
							UsedOptions:      []string{},
							BuiltAsBottle:    true,
							PouredFromBottle: true,
						},
						{
							Version:          "2.9.0",
							UsedOptions:      []string{},
							BuiltAsBottle:    true,
							PouredFromBottle: true,
						},
					},
					LinkedKeg: "2.9.0",
					Pinned:    false,
					Outdated:  false,
					KegOnly:   false,
					Dependencies: []string{
						"xz",
						"pcre",
						"gettext",
					},
					RecommendedDependencies: []string{},
					OptionalDependencies: []string{
						"pcre",
						"gettext",
					},
					BuildDependencies: []string{
						"xz",
					},
					ConflictsWith: []string{},
					Requirements:  []Requirement{},
					Options: []Option{
						{
							Option:      "--with-blk-sha1",
							Description: "Compile with the block-optimized SHA1 implementation",
						},
						{
							Option:      "--without-completions",
							Description: "Disable bash/zsh completions from 'contrib' directory",
						},
						{
							Option:      "--with-brewed-openssl",
							Description: "Build with Homebrew OpenSSL instead of the system version",
						},
						{
							Option:      "--with-brewed-curl",
							Description: "Use Homebrew's version of cURL library",
						},
						{
							Option:      "--with-brewed-svn",
							Description: "Use Homebrew's version of SVN",
						},
						{
							Option:      "--with-persistent-https",
							Description: "Build git-remote-persistent-https from 'contrib' directory",
						},
						{
							Option:      "--with-pcre",
							Description: "Build with pcre support",
						},
						{
							Option:      "--with-gettext",
							Description: "Build with gettext support",
						},
					},
					Bottle: map[string]Bottle{
						"stable": Bottle{
							Revision: 0,
							Cellar:   "/usr/local/Cellar",
							Prefix:   "/usr/local",
							RootUrl:  "https://homebrew.bintray.com/bottles",
							Files: map[string]File{
								"el_capitan": {
									URL:    "https://homebrew.bintray.com/bottles/git-2.9.0.el_capitan.bottle.tar.gz",
									Sha256: "6b230fa505e98f2947134271fee52d5fcd16eaff49d9dc0fd9ef933f568b92db",
								},
								"yosemite": {
									URL:    "https://homebrew.bintray.com/bottles/git-2.9.0.yosemite.bottle.tar.gz",
									Sha256: "076146d0cefd7643c02446a279da627cf290702472e7cd9d3756d6a38cf3c13c",
								},
								"mavericks": {
									URL:    "https://homebrew.bintray.com/bottles/git-2.9.0.mavericks.bottle.tar.gz",
									Sha256: "396bb11465e7d13141bfa9c67f41a81628259060761df820193763b9e973fc61",
								},
							},
						},
					},
				},
			},
		},
		{
			File: "mongodb.json",
			Result: []Package{
				{
					Name:        "mongodb",
					FullName:    "mongodb",
					Description: "High-performance, schema-free, document-oriented database",
					Homepage:    "https://www.mongodb.org/",
					Aliases: []string{
						"mongo",
					},
					Versions: Versions{
						Stable:      "3.2.7",
						Bottle:      true,
						Development: "",
						Head:        "",
					},
					Revision: 0,
					Installed: []InstalledVersion{
						{
							Version:          "3.2.4",
							UsedOptions:      []string{},
							BuiltAsBottle:    true,
							PouredFromBottle: true,
						},
					},
					LinkedKeg: "3.2.4",
					Pinned:    false,
					Outdated:  true,
					KegOnly:   false,
					Dependencies: []string{
						"boost",
						"go",
						"scons",
						"openssl",
					},
					RecommendedDependencies: []string{
						"openssl",
					},
					OptionalDependencies: []string{
						"boost",
					},
					BuildDependencies: []string{
						"go",
						"scons",
					},
					ConflictsWith: []string{},
					Requirements: []Requirement{
						{
							Name: "minimummacos",
						},
						{
							Name:           "git",
							DefaultFormula: "git",
						},
					},
					Options: []Option{
						{
							Option:      "--with-boost",
							Description: "Compile using installed boost, not the version shipped with mongodb",
						},
						{
							Option:      "--with-sasl",
							Description: "Compile with SASL support",
						},
						{
							Option:      "--without-openssl",
							Description: "Build without openssl support",
						},
					},
					Bottle: map[string]Bottle{
						"stable": Bottle{
							Revision: 0,
							Cellar:   "/usr/local/Cellar",
							Prefix:   "/usr/local",
							RootUrl:  "https://homebrew.bintray.com/bottles",
							Files: map[string]File{
								"el_capitan": {
									URL:    "https://homebrew.bintray.com/bottles/mongodb-3.2.7.el_capitan.bottle.tar.gz",
									Sha256: "40f0556c93e212a70c6a295b3f6f4a6ac7d243abc0e44dacb9c633dc7725106e",
								},
								"yosemite": {
									URL:    "https://homebrew.bintray.com/bottles/mongodb-3.2.7.yosemite.bottle.tar.gz",
									Sha256: "30c8ac35adc9f0ef0bde8af6c7d58cf02bc14c779c09626f89bafc451c6d5329",
								},
								"mavericks": {
									URL:    "https://homebrew.bintray.com/bottles/mongodb-3.2.7.mavericks.bottle.tar.gz",
									Sha256: "d1fd8b7daf61e0448429fabe10ce554f98151cba30441744aeb1fdbbb0f1c76f",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		data, err := ioutil.ReadFile(path.Join("./fixtures", test.File))
		if err != nil {
			t.Fatal(err)
		}

		actual, err := Unmarshal(data)

		if (err != nil) != test.Error {
			if test.Error {
				t.Fatalf("expected test %s to fail", test.File)
			} else {
				t.Fatalf("expected test %s to pass, got error: '%s'", test.File, err)
			}
		}

		if !reflect.DeepEqual(actual, test.Result) {
			// this is useful for testing, in case things don't match:
			// install from https://github.com/kylelemons/godebug
			// fmt.Println(pretty.Compare(actual, test.Result))

			t.Fatalf("response does not match\n\nexpected:\n%v\n\nactual:\n%v", test.Result, actual)
		}
	}
}
