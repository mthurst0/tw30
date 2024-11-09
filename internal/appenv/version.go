package appenv

type VersionInfo struct {
	Version   string
	Commit    string
	BuildTime string
	GoVersion string
}

var Version VersionInfo

func SetVersionInfo(v VersionInfo) {
	Version = v
}

func GetVersionInfo() VersionInfo {
	return Version
}
