package service

import (
	"image"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/iineva/ipa-server/pkg/uuid"
)

type AppInfoType int
type AppInfo struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Version     string      `json:"version"`
	Identifier  string      `json:"identifier"`
	Build       string      `json:"build"`
	Channel     string      `json:"channel"`
	Date        time.Time   `json:"date"`
	Size        int64       `json:"size"`
	NoneIcon    bool        `json:"noneIcon"`
	Type        AppInfoType `json:"type"`
	Env         string      `json:"env"`
	ProjectID   int         `json:"project_id"`
	PlatformID  int         `json:"platform_id"`
	Description string      `json:"description"`
	Region      string      `json:"region"`
	// Metadata   plist.Plist `json:"metadata"` // metadata from Info.plist
}

const (
	AppInfoTypeIpa     = AppInfoType(0)
	AppInfoTypeApk     = AppInfoType(1)
	AppInfoTypeAab     = AppInfoType(2)
	AppInfoTypeUnknown = AppInfoType(-1)
)

func (t AppInfoType) StorageName() string {
	switch t {
	case AppInfoTypeIpa:
		return "ipa.ipa"
	case AppInfoTypeApk:
		return "apk.apk"
	case AppInfoTypeAab:
		return "aab.aab"
	default:
		return "unknown"
	}
}

func FileType(n string) AppInfoType {
	ext := strings.ToLower(path.Ext(n))
	switch ext {
	case ".ipa":
		return AppInfoTypeIpa
	case ".apk":
		return AppInfoTypeApk
	case ".aab":
		return AppInfoTypeAab
	default:
		return AppInfoTypeUnknown
	}
}

type AppList []*AppInfo

func (a AppList) Len() int           { return len(a) }
func (a AppList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AppList) Less(i, j int) bool { return a[i].Date.After(a[j].Date) }

type Package interface {
	Name() string
	Version() string
	Identifier() string
	Build() string
	Channel() string
	Icon() image.Image
	Size() int64
}

type PackageExt interface {
	Env() string
	PlatformID() int
	ProjectID() int
	Description() string
	Region() string
}

type PackageExter struct {
	env         string
	platformID  int
	projectID   int
	description string
	region      string
}

func ParsePackageExt(filename, description string) *PackageExter {
	//Plinko_v1.0.58_cv108_2305121850_GOOGLE_6_810_release-cn.apk
	re := regexp.MustCompile(`.*_(\d+)_(\d+)_(\w+)(?:-(\w+))?\.apk`)
	matches := re.FindStringSubmatch(filename)
	if len(matches) >= 4 {
		platformID, _ := strconv.ParseInt(matches[1], 10, 0)
		projectID, _ := strconv.ParseInt(matches[2], 10, 0)
		env := matches[3]
		region := ""
		if len(matches) >= 5 {
			region = matches[4]
		}
		return &PackageExter{
			env:         env,
			platformID:  int(platformID),
			projectID:   int(projectID),
			description: description,
			region:      region,
		}
	}
	return &PackageExter{}
}

func (p *PackageExter) Env() string {
	return p.env
}

func (p *PackageExter) PlatformID() int {
	return p.platformID
}

func (p *PackageExter) ProjectID() int {
	return p.projectID
}

func (p *PackageExter) Description() string {
	return p.description
}

func (p *PackageExter) Region() string {
	return p.region
}

func NewAppInfo(i Package, t AppInfoType, pext PackageExt) *AppInfo {
	return &AppInfo{
		ID:          uuid.NewString(),
		Name:        i.Name(),
		Version:     i.Version(),
		Identifier:  i.Identifier(),
		Build:       i.Build(),
		Channel:     i.Channel(),
		Date:        time.Now(),
		Size:        i.Size(),
		Type:        t,
		NoneIcon:    i.Icon() == nil,
		Env:         pext.Env(),
		ProjectID:   pext.ProjectID(),
		PlatformID:  pext.PlatformID(),
		Description: pext.Description(),
		Region:      pext.Region(),
	}
}

func (a *AppInfo) IconStorageName() string {
	if a.NoneIcon {
		return ""
	}
	return filepath.Join(a.Identifier, a.ID, "icon.png")
}

func (a *AppInfo) PackageStorageName() string {
	return filepath.Join(a.Identifier, a.ID, a.Type.StorageName())
}
