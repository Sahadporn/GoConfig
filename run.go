package main

// Importing fmt
import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/viper"

	config "source/config"
	scfg "source/scfg"
)

func index(w http.ResponseWriter, r *http.Request) {
	envVar, exists := os.LookupEnv("ENV_VAR")
	if !exists {
		envVar = "There is no environment variable"
	}
	firstLine := "<h1>Hello World</h1><br><h2>Environment variable: " + envVar + "</h2><br>"

	conf := config.New()
	secondLine := "<h2>Env from struct: " + conf.Env.UIText + " | Number: " + strconv.Itoa(conf.Number) + "</h2><br>"

	viper.BindEnv("env", "ENV_VAR")
	env, ok := viper.Get("env").(string)
	if !ok {
		fmt.Print("Invalid type assertion")
		env = "There is no environment variable"
	}
	thirdLine := "<h2>Env from viper: " + env + "</h2><br>"

	cfg := scfg.New()
	forthLine := "<h2>Env from another method struct: " + cfg.Env + " | Number: " + strconv.Itoa(cfg.Number) + "</h2><br>"

	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	envViper, okv := viper.Get("envvar").(string)
	if !okv {
		fmt.Print("Invalid type assertion")
	}
	fifthLine := "<h2>Env from viper yaml: " + envViper + "</h2><br>"

	fmt.Fprintf(w, firstLine+secondLine+thirdLine+forthLine+fifthLine)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}
