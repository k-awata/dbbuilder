package main

import (
	"io"

	"go.yaml.in/yaml/v2"
)

type Project struct {
	LogFile           string             `yaml:"log_file"`
	Users             []User             `yaml:"users"`
	AuthUsers         []AuthUser         `yaml:"auth_users"`
	Teams             []Team             `yaml:"teams"`
	MasterDBs         []MasterDB         `yaml:"master_dbs"`
	Foreigns          []Foreign          `yaml:"foreigns"`
	ExtractDBs        []ExtractDB        `yaml:"extract_dbs"`
	WorkingExtractDBs []WorkingExtractDB `yaml:"working_extract_dbs"`
	DynamicDBSets     []DynamicDBSet     `yaml:"dynamic_dbsets"`
	DBSets            []StaticDBSet      `yaml:"dbsets"`
	MDBs              []MDB              `yaml:"mdbs"`
	TTFonts           []TTFont           `yaml:"ttfonts"`
	PostCommands      []string           `yaml:"post_commands"`
}

type User struct {
	Name        string `yaml:"name"`
	Password    string `yaml:"password"`
	Description string `yaml:"description"`
	Free        bool   `yaml:"free,omitempty"`
}

type AuthUser struct {
	Name        string   `yaml:"name"`
	DefaultUser string   `yaml:"default_user"`
	Users       []string `yaml:"users,omitempty"`
}

type Team struct {
	Name        string   `yaml:"name"`
	Members     []string `yaml:"members"`
	Description string   `yaml:"description,omitempty"`
}

type MasterDB struct {
	Name          string  `yaml:"name"`
	Type          string  `yaml:"type"`
	RefOnly       bool    `yaml:"ref_only"`
	DBNumber      uint32  `yaml:"db_number"`
	AreaNumber    uint16  `yaml:"area_number,omitempty"`
	ExplicitClaim bool    `yaml:"explicit_claim,omitempty"`
	Protected     bool    `yaml:"protected,omitempty"`
	ExtractNumber uint16  `yaml:"extract_number,omitempty"`
	FileNumber    uint16  `yaml:"file_number,omitempty"`
	Description   string  `yaml:"description,omitempty"`
	CreateElement Element `yaml:"create_element"`
}

type Element struct {
	Type string `yaml:"type"`
	Name string `yaml:"name"`
}

type Foreign struct {
	Project    string   `yaml:"project"`
	Username   string   `yaml:"username,omitempty"`
	Password   string   `yaml:"password,omitempty"`
	CopyDBs    []CopyDB `yaml:"copy_dbs,omitempty"`
	IncludeDBs []string `yaml:"include_dbs,omitempty"`
}

type CopyDB struct {
	From       string `yaml:"from"`
	To         string `yaml:"to,omitempty"`
	AreaNumber uint16 `yaml:"area_number,omitempty"`
	FileNumber uint16 `yaml:"file_number,omitempty"`
}

type ExtractDB struct {
	Owner         string `yaml:"owner"`
	Name          string `yaml:"name"`
	Variant       bool   `yaml:"variant,omitempty"`
	Session       uint32 `yaml:"session,omitempty"`
	AreaNumber    uint16 `yaml:"area_number,omitempty"`
	ExplicitClaim bool   `yaml:"explicit_claim,omitempty"`
	ExtractNumber uint16 `yaml:"extract_number,omitempty"`
	Description   string `yaml:"description,omitempty"`
}

type WorkingExtractDB struct {
	Owner         string   `yaml:"owner"`
	Users         []string `yaml:"users"`
	Variant       bool     `yaml:"variant,omitempty"`
	AreaNumber    uint16   `yaml:"area_number,omitempty"`
	ExplicitClaim bool     `yaml:"explicit_claim,omitempty"`
	Description   string   `yaml:"description,omitempty"`
}

type DynamicDBSet struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description,omitempty"`
	Criteria    string `yaml:"criteria"`
	OrderBy     string `yaml:"order_by"`
}

type StaticDBSet struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description,omitempty"`
	DBs         []string `yaml:"dbs"`
}

type MDB struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	CurrentDBs  []string `yaml:"current_dbs"`
	DeferredDBs []string `yaml:"deferred_dbs,omitempty"`
}

type TTFont struct {
	Index       uint8  `yaml:"index"`
	FaceName    string `yaml:"face_name"`
	Description string `yaml:"description,omitempty"`
}

func LoadProject(i io.Reader) (*Project, error) {
	in, err := io.ReadAll(i)
	if err != nil {
		return nil, err
	}
	var out Project
	if err := yaml.Unmarshal(in, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
