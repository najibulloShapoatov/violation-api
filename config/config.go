package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"
)

const (
	jsonExt    = ".json"
	configFile = "config.json"
)

// _Database holds database info.
type _Database struct {
	Port   string `json:"port"`
	Host   string `json:"host"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DBName string `json:"dbName"`
} // _Database holds database info.
type _DatabaseSafeCity struct {
	Port   string `json:"port"`
	Host   string `json:"host"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DBName string `json:"dbName"`
}
type _Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
type _PatternPhone struct {
	Code string `json:"code"`
}
type _Kannel struct {
	Link     string `json:"link"`
	Username string `json:"username"`
	Password string `json:"password"`
	Smsc     string `json:"smsc"`
	From     string `json:"from"`
}

//Config config
type Config struct {
	AppName              string             `json:"appName"`
	Database             *_Database         `json:"database"`
	DatabaseSafeCity     *_DatabaseSafeCity `json:"databaseSafecity"`
	Logfile              string             `json:"logfile"`
	SecretKey            string             `json:"secret_key"`
	Server               *_Server           `json:"server"`
	PatternPhone         []string           `json:"patternPhone"`
	DetailsLink          string             `json:"detailsLink"`
	Kannel               *_Kannel           `json:"kannel"`
	Time1                int                `json:"execTime1"`
	Time2                int                `json:"execTime2"`
	TimeMonthly          int                `json:"execTimeMonthly"`
	SusbscribeEndText    string             `json:"susbscribeEndText"`
	SubscribeEndSendTime int                `json:"subscribeEndSendTime"`
	MailingText          string             `json:"mailingText"`
	MailingSendTime      int                `json:"mailingSendTime"`
}

var cfg Config

//readJSON reads config file as JSON into provided interface{}
func readJSON(path string, conf interface{}) error {
	// validate config file
	if len(path) == 0 {
		return errors.New("config: file path cannot be empty")
	}
	if filepath.Ext(path) != jsonExt {
		return errors.New("config: file must be .json")
	}

	// validate config holder
	if conf == nil {
		return errors.New("config: holder is nil")
	}
	if reflect.ValueOf(conf).Kind() != reflect.Ptr {
		return errors.New("config: holder must be a pointer to a type")
	}

	// start read file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(conf)
	if err != nil {
		return err
	}

	return nil
}

//Init func
func Init() {
	err := readJSON(configFile, &cfg)
	if err != nil {
		panic(err)
	}
}

// Get config values
func Get() Config {
	return cfg
}
