package awsutils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session

func InitAWSSession(region, profile string) error {
	return initSession(region, profile)
}

func initSession(region, profile string) error {
	var creds *credentials.Credentials
	if profile != "" {
		creds = credentials.NewSharedCredentials("", profile)
	}
	options := session.Options{
		Config: aws.Config{
			Credentials: creds,
			Region: &region,
		},
	}
	var err error
	sess, err = session.NewSessionWithOptions(options)
	return err
}

func GetSession() *session.Session {
	return sess
}