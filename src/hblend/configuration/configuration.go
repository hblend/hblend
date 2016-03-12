package configuration

type Configuration struct {
	Init     bool
	ShowHelp bool
	Dir      ConfigurationDirs
}

type ConfigurationDirs struct {
	Components string
	Www        string
}
