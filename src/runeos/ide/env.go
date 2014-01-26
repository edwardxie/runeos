package ide

func SetEnv() {}

func GetEnv(n string) interface{} { return globalEnv[n] }

func GetEnvMap() map[string]interface{} { return globalEnv }
