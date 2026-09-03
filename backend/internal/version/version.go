package version

var (
	Version = "1.1.29"
	Repo    = "stripathi02123-tech/nexora-panel--2"
)

func Current() string {
	if Version == "" {
		return "dev"
	}
	return Version
}
