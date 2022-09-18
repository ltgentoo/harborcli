/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "context"
    "fmt"
    "github.com/mittwald/goharbor-client/v5/apiv2"
    modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
    "github.com/spf13/cobra"

)

// registryCmd represents the registry command
var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
	    createRegistry()
		fmt.Println("registry called")
	},
}

func init() {
	createCmd.AddCommand(registryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var (
    url2 = "192.168.0.2:1121"
    username2 = "admin"
    password2 = "Harbor12345!"
)

func NewClient(url,username,password string)(*apiv2.RESTClient,error){
    client,err := apiv2.NewRESTClientForHost(url,username,password,nil)
    if err != nil {
        panic(err)
    }
    return client,err
}

func createRegistry() {
    var (
        name string
        url string
        desc string
        username string
        password string
    )
    fmt.Printf("请输入registry名称：")
    fmt.Scan(&name)
    fmt.Printf("请输入registry地址：")
    fmt.Scan(&url)
    fmt.Printf("请输入registry描述：")
    fmt.Scan(&desc)
    fmt.Printf("请输入用户名：")
    fmt.Scan(&username)
    fmt.Printf("请输入密码：")
    fmt.Scan(&password)

    //client,err := apiv2.NewRESTClientForHost(url,username,password,nil)
    client,error := NewClient(url2,username2,password2)
    if error != nil {
        fmt.Println(error)
    }


    registry := &modelv2.Registry{
        Credential:   &modelv2.RegistryCredential{
            AccessKey: username,
            AccessSecret: password,
            Type: "test",
        },
        Description:  desc,
        Insecure:     true,
        Name:         name,
        URL:          url,
    }
    var ctx context.Context
    err := client.NewRegistry(ctx,registry)
    if err != nil {
        panic(err)
    }
}
