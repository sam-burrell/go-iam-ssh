package main

import (
  "os"
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/iam"
)

func exitError(msg string, args ...interface{}) {
  fmt.Fprintf(os.Stderr, msg+"\n", args...)
  os.Exit(1)
}

func main() {
  svc := iam.New(session.New())
  input := &iam.ListUsersInput{}
  result, err := svc.ListUsers(input)
  if err != nil {
    exitError("Unable to list users, %v", err)
  }
  for _, value := range result.Users {
    input := &iam.ListSSHPublicKeysInput{
      UserName: value.UserName,
    }
    result, err := svc.ListSSHPublicKeys(input)
    if err != nil {
      exitError("Unable to list SSH keys, %v", err)
    }
    for _, value := range result.SSHPublicKeys {
      input := &iam.GetSSHPublicKeyInput{
        Encoding: aws.String("SSH"),
        SSHPublicKeyId: value.SSHPublicKeyId,
        UserName: value.UserName,
      }
      result, err := svc.GetSSHPublicKey(input)
      if err != nil {
        exitError("Unable to get SSH keys, %v", err)
      }
      fmt.Println(aws.StringValue(result.SSHPublicKey.SSHPublicKeyBody))
    }
  }
}
