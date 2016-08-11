package settings

import (
    "encoding/json"
    "os"
    "log"
    "fmt"
    "github.com/leonharvey/go-web-boilerplate/lib/utilities"
)

var Config struct {
    Host string `json:"host"`
    AbsolutePath string `json:"absolutePath"`
    Favicon string `json:"favicon"`
    Static string `json:"static"`
    DevelopmentMode bool `json:"developmentMode"`
    ListenPort int `json:"listenPort"`
    LoggerEnabled bool `json:"loggerEnabled"`
    Secret string `json:"secret"`
    ForceSSL bool `json:"forceSSL"`
    DatabaseType string `json:"databaseType"`
}

func init() {
    fmt.Println("Loading configuration")
    
    cwd, err := os.Getwd()
    
    if err != nil {
        log.Fatal(err)
    }
    
    configFile, err := os.Open(cwd + "/config.json")
    
    if err != nil {
        log.Fatal(err.Error())
    }

    jsonParser := json.NewDecoder(configFile)
    
    if err = jsonParser.Decode(&Config); err != nil {
        log.Fatal(err.Error())
    }
    Config.AbsolutePath = cwd
    Config.Secret = utilities.RandomString(64)
    
    return
}
