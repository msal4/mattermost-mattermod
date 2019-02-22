// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package server

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/mattermost/mattermost-server/mlog"
)

type LabelResponse struct {
	Label   string
	Message string
}

type Repository struct {
	Owner                      string
	Name                       string
	BuildStatusContext         string
	JenkinsServer              string
	InstanceSetupScript        string
	InstanceSetupUpgradeScript string
	JobName                    string
}

type JenkinsCredentials struct {
	URL      string
	Username string
	ApiToken string
}

type PRServerConfig struct {
	ListenAddress     string
	GithubAccessToken string
	Username          string

	TickRate               int
	SpinmintExpirationHour int

	DriverName string
	DataSource string

	Repositories []*Repository

	SetupSpinmintTag                   string
	SetupSpinmintMessage               string
	SetupSpinmintDoneMessage           string
	SetupSpinmintFailedMessage         string
	DestroyedSpinmintMessage           string
	DestroyedExpirationSpinmintMessage string
	SpinmintsUseHttps                  bool

	SetupSpinmintUpgradeTag         string
	SetupSpinmintUpgradeMessage     string
	SetupSpinmintUpgradeDoneMessage string

	BuildMobileAppTag           string
	BuildMobileAppInitMessage   string
	BuildMobileAppDoneMessage   string
	BuildMobileAppFailedMessage string

	StartLoadtestTag     string
	StartLoadtestMessage string

	SignedCLAURL          string
	NeedsToSignCLAMessage string

	PrLabels    []LabelResponse
	IssueLabels []LabelResponse

	JenkinsCredentials map[string]*JenkinsCredentials

	AWSCredentials struct {
		Id     string
		Secret string
		Token  string
	}

	AWSRegion        string
	AWSImageId       string
	AWSKeyName       string
	AWSInstanceType  string
	AWSHostedZoneId  string
	AWSSecurityGroup string
	AWSDnsSuffix     string

	MattermostWebhookURL    string
	MattermostWebhookFooter string

	LoggerConfiguration struct {
		EnableConsole bool
		ConsoleJson   bool
		ConsoleLevel  string
		EnableFile    bool
		FileJson      bool
		FileLevel     string
		FileLocation  string
	}
}

var Config *PRServerConfig = &PRServerConfig{}

func FindConfigFile(fileName string) string {
	if _, err := os.Stat("/tmp/" + fileName); err == nil {
		fileName, _ = filepath.Abs("/tmp/" + fileName)
	} else if _, err := os.Stat("./config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("./config/" + fileName)
	} else if _, err := os.Stat("../config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("../config/" + fileName)
	} else if _, err := os.Stat(fileName); err == nil {
		fileName, _ = filepath.Abs(fileName)
	}

	return fileName
}

func LoadConfig(fileName string) {
	fileName = FindConfigFile(fileName)
	mlog.Info(fmt.Sprintf("Loading %v", fileName))

	file, err := os.Open(fileName)
	if err != nil {
		mlog.Critical(fmt.Sprintf("Error opening config file=%v, err=%v", fileName, err.Error()))
		panic(err.Error())
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(Config)
	if err != nil {
		mlog.Critical(fmt.Sprintf("Error decoding config file=%v, err=", fileName, err.Error()))
		panic(err.Error())
	}
}

func (config *PRServerConfig) GetRepository(owner, name string) (*Repository, bool) {
	for _, repo := range config.Repositories {
		if repo.Owner == owner && repo.Name == name {
			return repo, true
		}
	}

	return nil, false
}

func (config *PRServerConfig) GetAwsConfig() *aws.Config {
	var creds *credentials.Credentials = nil
	if config.AWSCredentials.Id != "" {
		creds = credentials.NewStaticCredentials(
			config.AWSCredentials.Id,
			config.AWSCredentials.Secret,
			config.AWSCredentials.Token,
		)
	}

	return &aws.Config{
		Credentials: creds,
		Region:      &config.AWSRegion,
	}
}
