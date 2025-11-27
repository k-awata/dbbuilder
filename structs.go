package main

import (
	"encoding/json"
	"io"
)

type Project struct {
	LogFile        string         `json:"logFile"`
	Users          []User         `json:"users"`
	AuthUsers      []AuthUser     `json:"authUsers"`
	Teams          []Team         `json:"teams"`
	MasterDBs      []MasterDB     `json:"masterDbs"`
	CopyDBs        []CopyDB       `json:"copyDbs"`
	IncludeDBs     []ForeignDB    `json:"includeDbs"`
	ExtractDBs     []ExtractDB    `json:"extractDbs"`
	DynamicDBSets  []DynamicDBSet `json:"dynamicDbSets"`
	DBSets         []StaticDBSet  `json:"dbSets"`
	MDBs           []MDB          `json:"mdbs"`
	TTFonts        []TTFont       `json:"ttFonts"`
	AdditionalCmds []string       `json:"additionalCmds"`
}

type User struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	Description string `json:"description"`
	Free        bool   `json:"free,omitempty"`
}

type AuthUser struct {
	Name        string   `json:"name"`
	DefaultUser string   `json:"defaultUser"`
	Users       []string `json:"users,omitempty"`
}

type Team struct {
	Name        string   `json:"name"`
	Members     []string `json:"members"`
	Description string   `json:"description,omitempty"`
}

type MasterDB struct {
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	ReferenceOnly bool    `json:"referenceOnly"`
	DBNumber      uint32  `json:"dbNumber"`
	AreaNumber    uint16  `json:"areaNumber,omitempty"`
	ExplicitClaim bool    `json:"explicitClaim,omitempty"`
	Protected     bool    `json:"protected,omitempty"`
	ExtractNumber uint16  `json:"extractNumber,omitempty"`
	FileNumber    uint16  `json:"fileNumber,omitempty"`
	Description   string  `json:"description,omitempty"`
	CreateElement Element `json:"createElement"`
}

type Element struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type CopyDB struct {
	From       ForeignDB `json:"from"`
	To         string    `json:"to"`
	AreaNumber uint16    `json:"areaNumber,omitempty"`
	FileNumber uint16    `json:"fileNumber,omitempty"`
}

type ForeignDB struct {
	Name     string `json:"name"`
	Project  string `json:"project"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type ExtractDB struct {
	Name          string `json:"name"`
	Owner         string `json:"owner"`
	Variant       bool   `json:"variant,omitempty"`
	Session       uint32 `json:"session,omitempty"`
	AreaNumber    uint16 `json:"areaNumber,omitempty"`
	ExplicitClaim bool   `json:"explicitClaim,omitempty"`
	ExtractNumber uint16 `json:"extractNumber,omitempty"`
	Description   string `json:"description,omitempty"`
}

type DynamicDBSet struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Criteria    string `json:"criteria"`
	OrderBy     string `json:"orderBy"`
}

type StaticDBSet struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	DBs         []string `json:"dbs"`
}

type MDB struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CurrentDBs  []string `json:"currentDBs"`
	DeferredDBs []string `json:"deferredDBs,omitempty"`
}

type TTFont struct {
	Index       uint8  `json:"index"`
	FaceName    string `json:"faceName"`
	Description string `json:"description,omitempty"`
}

func LoadProject(i io.Reader) (*Project, error) {
	j, err := io.ReadAll(i)
	if err != nil {
		return nil, err
	}
	var b Project
	if err := json.Unmarshal(j, &b); err != nil {
		return nil, err
	}
	return &b, nil
}
