package internal

/*
type Config struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}


func ConfigSetup() {
	var configValues Config

	yamlFile, err := os.ReadFile("/etc/mysqldumpmanager/conf.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &configValues)
	if err != nil {
		cmd.LogError(err.Error())
	}

	u := configValues.User
	p := configValues.Password
	cmd.LogInfo("Fetching data from conf.yaml...")
	cmd.LogInfo("User: " + u)
	cmd.LogInfo("Password: " + p)

	err = os.Setenv("MYSQLDUMPMANAGER_USER", configValues.User)
	if err != nil {
		cmd.LogError(err.Error())
	}
	cmd.LogInfo("Setting up environment variable MYSQLDUMPMANAGER_USER with value: " + os.Getenv("MYSQLDUMPMANAGER_USER"))

	err = os.Setenv("MYSQLDUMPMANAGER_PASSWORD", configValues.Password)
	if err != nil {
		cmd.LogError(err.Error())
	}
	cmd.LogInfo("Setting up environment variable MYSQLDUMPMANAGER_PASSWORD with value: " + os.Getenv("MYSQLDUMPMANAGER_PASSWORD"))

	c := []byte("\n[mysqldump]\nuser: " +
		os.Getenv("MYSQLDUMPMANAGER_USER") +
		"\npassword: " +
		os.Getenv("MYSQLDUMPMANAGER_PASSWORD"))

	f1 := "/etc/my.cnf"
	f2 := "/etc/mysql/my.cnf"
	f3 := "~/.my.cnf"
	f4 := os.Getenv("MYSQL_HOME") + "/.my.cnf"

	fileLocationLoop := []string{f1, f2, f3, f4}

	for _, fileLocation := range fileLocationLoop {
		if cmd.CheckIfFileExists(fileLocation) == false {
			cmd.LogWarn("File " + fileLocation + " does not exist...")
		}
		err = os.WriteFile(fileLocation, c, 0644)
		if err != nil {
			cmd.LogError(err.Error())
		}
	}

}


*/
