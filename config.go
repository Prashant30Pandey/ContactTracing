package config

type configuration struct {
	Environment string
	Mongo       MongoConfiguration
}

type MongoConfiguration struct {
	Server     string
	Database   string
	Collection string
}

func GetConfig() Configuration {
	conf := Configuration{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.SetConfigPath("./config")

	err := viper.ReadConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unamrshal(&conf)
	if err != nil {
		panic(err)
	}

	return conf
}
