package main

import (
	_ "github.com/spf13/viper/remote"
)

func main() {
	//! INSTALL VIPER
	// go get github.com/spf13/viper

	//! PRECEDENCE
	// Set
	// Flag
	// Env
	// Config
	// key/value store
	// default

	//! 1. SET && GET && REGISTERING AND USING ALIASES
	// viper := viper.New()
	// viper.Set("key", "Value")
	// viper.RegisterAlias("keyAlias", "key")
	// fmt.Println("GET key: ", viper.Get("key"))
	// fmt.Println("GET keyAlias: ", viper.Get("keyAlias"))

	// Get(key string) : interface{}
	// GetBool(key string) : bool
	// GetFloat64(key string) : float64
	// GetInt(key string) : int
	// GetIntSlice(key string) : []int
	// GetString(key string) : string
	// GetStringMap(key string) : map[string]interface{}
	// GetStringMapString(key string) : map[string]string
	// GetStringSlice(key string) : []string
	// GetTime(key string) : time.Time
	// GetDuration(key string) : time.Duration
	// IsSet(key string) : bool
	// AllSettings() : map[string]interface{}

	// viper := viper.New()
	// viper.Set("GetBool", true)
	// viper.Set("GetBool", "string")
	// viper.Set("GetFloat64", 123.123)
	// viper.Set("GetInt", 123)
	// viper.Set("GetIntSlice", []int{1, 2, 3})
	// viper.Set("GetString", []string{"mot", "hai", "ba"})
	// viper.Set("GetStringMap", map[string]interface{}{"mot": "mot", "hai": "hai", "ba": "ba"})
	// viper.Set("GetStringMapString", map[string]string{"mot": "mot", "hai": "hai", "ba": "ba"})
	// viper.Set("GetStringSlice", []string{"mot", "hai", "ba"})
	// viper.Set("GetTime", time.Now())
	// viper.Set("GetDuration", time.Hour)

	// fmt.Println("GetBool: ", viper.GetBool("GetBool"))
	// fmt.Println("GetFloat64: ", viper.GetFloat64("GetFloat64"))
	// fmt.Println("GetInt: ", viper.GetInt("GetInt"))
	// fmt.Println("GetIntSlice: ", viper.GetIntSlice("GetIntSlice"))
	// fmt.Println("GetString: ", viper.GetString("GetString"))
	// fmt.Println("GetStringMap: ", viper.GetStringMap("GetStringMap"))
	// fmt.Println("GetStringMapString: ", viper.GetStringMapString("GetStringMapString"))
	// fmt.Println("GetStringSlice: ", viper.GetStringSlice("GetStringSlice"))
	// fmt.Println("GetTime: ", viper.GetTime("GetTime"))
	// fmt.Println("GetDuration: ", viper.GetDuration("GetDuration"))
	// fmt.Println("GetBool: ", viper.IsSet("GetBool"))
	// fmt.Println("AllSettings: ", viper.AllSettings())

	//! 2. USE FLAG

	//! 2.1 FLAG
	// var host string
	// var port int

	// pflag.StringVar(&host, "host", "localhost_flag", "des")
	// pflag.IntVar(&port, "port", 5432, "port")
	// pflag.Parse()
	// viper.BindPFlags(pflag.CommandLine)

	// fmt.Println("host: ", viper.GetString("host"))
	// fmt.Println("port: ", viper.GetInt("port"))

	// var frac Fraction
	// flag.Var(&frac, "frac", "A fraction value")
	// flag.Parse()
	// viper.BindFlagValue("frac", frac)
	// viper.SetEnvKeyReplacer(strings.NewReplacer("+", "_"))

	//! 2.2 FLAG INTERFACE

	// var fraction Fraction
	// flag.Var(&fraction, "fraction", "TEST")
	// flag.Parse()
	// viper.BindFlagValue("frac", fraction)
	// fmt.Println(viper.AllSettings())

	//! 3. .ENV
	// AutomaticEnv()
	// BindEnv(string...) : error
	// SetEnvPrefix(string)
	// SetEnvKeyReplacer(string...) *strings.Replacer
	// AllowEmptyEnv(bool)

	// go get github.com/joho/godotenv

	// godotenv.Load(".env")

	// viper := viper.New()

	// viper.SetEnvPrefix("MYAPP")
	// viper.AutomaticEnv()
	// viper.BindEnv("HOST")
	// // viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// fmt.Println(viper.Get("HOST"))
	// fmt.Println(viper.Get("MYAPP_PORT"))

	//! 4. CONFIG && WATCH CONFIG && RE-READING && WRITING CONFIG FILE && READ CONFIG FROM IO.READER
	// viper.SetConfigFile("./config/config.json")
	// viper.SetConfigName("config")
	// viper.SetConfigType("json")

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }

	// fmt.Println(viper.GetString("host"))
	// fmt.Println(viper.GetInt("port"))

	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Config file changed:", e.Name)
	// })
	// viper.WatchConfig()
	// select {}

	// viper := viper.New()
	// viper.SetConfigName("config")
	// viper.SetConfigType("json")
	// viper.AddConfigPath("config/")

	// viper.Set("host", "localhost_set")
	// viper.Set("port", "5432")

	// err := viper.WriteConfig() //
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }

	// err := viper.SafeWriteConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }

	// err := viper.WriteConfigAs("config/config.yaml")
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }

	// err := viper.SafeWriteConfigAs("config/config.yaml")
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }

	// configData := `
	//       database:
	//           host: localhost
	//           port: 5432
	//           username: postgres
	//           password: secret
	//   `
	// reader := strings.NewReader(configData)

	// if err := viper.ReadConfig(reader); err != nil {
	// 	panic(fmt.Errorf("error reading config: %w", err))
	// }

	// fmt.Println("Database host:", viper.GetString("database.host"))
	// fmt.Println("Database port:", viper.GetInt("database.port"))
	// fmt.Println("Database username:", viper.GetString("database.username"))
	// fmt.Println("Database password:", viper.GetString("database.password"))

	//! 5. KEY/STORE
	// viper.AddRemoteProvider("consul", "localhost:8500", "config")
	// viper.SetConfigType("json")
	// err := viper.ReadRemoteConfig()

	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %w", err))
	// }

	// fmt.Println("host: ", viper.GetString("host"))
	// fmt.Println("port:", viper.GetInt("port"))

	// go func() {
	// 	for {
	// 		time.Sleep(time.Second * 2)

	// 		err := viper.WatchRemoteConfig()
	// 		if err != nil {
	// 			fmt.Println("unable to read")
	// 			continue
	// 		}

	// 		fmt.Println("host: ", viper.GetString("host"))
	// 		fmt.Println("port:", viper.GetInt("port"))
	// 	}
	// }()
	// select {}

	//! 6. SETDEFAULT
	// viper.SetDefault("host", "localhsot_default")
	// viper.SetDefault("port", 5432)

	// fmt.Println("host: ", viper.GetString("host"))
	// fmt.Println("port:", viper.GetInt("port"))

	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("config")
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Config file changed:", e.Name)
	// })
	// viper.WatchConfig()
	// select {}

}

//! FARATION
// type Fraction struct {
// 	numerator   int
// 	denominator int
// }

// func (f *Fraction) String() string {
// 	return fmt.Sprintf("%d/%d", f.numerator, f.denominator)
// }

// func (f *Fraction) Set(s string) error {
// 	parts := strings.Split(s, "/")
// 	if len(parts) != 2 {
// 		return fmt.Errorf("invalid fraction format: %s", s)
// 	}
// 	numerator, err := strconv.Atoi(parts[0])
// 	if err != nil {
// 		return fmt.Errorf("invalid numerator: %s", parts[0])
// 	}
// 	denominator, err := strconv.Atoi(parts[1])
// 	if err != nil {
// 		return fmt.Errorf("invalid denominator: %s", parts[1])
// 	}
// 	if denominator == 0 {
// 		return fmt.Errorf("denominator cannot be zero")
// 	}
// 	f.numerator = numerator
// 	f.denominator = denominator
// 	return nil
// }

// func (f Fraction) HasChanged() bool    { return false }
// func (f Fraction) Name() string        { return "fraction" }
// func (f Fraction) ValueString() string { return f.String() }
// func (f Fraction) ValueType() string   { return f.String() }

//! END FARATION
