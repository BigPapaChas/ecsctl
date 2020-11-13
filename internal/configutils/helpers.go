package configutils

import (
	"errors"
	"fmt"
	"os"

	"github.com/ecsctl/ecsctl/internal/printerutils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v3"
)

// Errors
var errCurrentContextNotFound = errors.New("current context not found inside config")
var errContextsNotFound = errors.New("list of contexts not found inside config")
var errSpecifiedContextNotFound = errors.New("could not find current context within defined contexts inside config")

var conf *Config

type Config struct {
	Contexts []*Context
	CurrentContext *string
}

type Context struct {
	Name    *string
	Profile *string
	Region  *string
	Cluster *string
}

type PrintableContext struct {
	Current string  `header:"current"`
	Name    string `header:"name"`
	Profile string `header:"profile"`
	Region  string `header:"region"`
	Cluster string `header:"cluster"`
}

func LoadConfig() {
	conf = &Config{}
	if err := viper.Unmarshal(conf); err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
}

func PrintConfig() error {
	if s, err := yaml.Marshal(conf); err != nil {
		return err
	} else {
		fmt.Printf("%s", string(s))
	}
	return nil
}

func PrintCurrentContext() error {
	if conf.CurrentContext != nil {
		fmt.Println(*conf.CurrentContext)
		return nil
	}
	return errCurrentContextNotFound
}

func UseContext(contextName string) {
	if conf.Contexts != nil {
		for _, c := range conf.Contexts {
			if c.Name != nil {
				if *c.Name == contextName {
					conf.CurrentContext = &contextName
					if err := writeCurrentContext(contextName); err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("Switched to context \"%s\"\n", contextName)
					}
					return
				}
			}
		}
	}
	fmt.Printf("error: no context exists with the name: \"%s\"\n", contextName)
	os.Exit(1)
}

func CreateNewContext() error {
	fmt.Print("Context name: ")
	var name string
	if _, err := fmt.Scanln(&name); err != nil {
		return err
	}
	fmt.Printf("You chose %s", name)
	return nil
}

func PrintContexts() {
	var pContexts []*PrintableContext
	var currentContext string
	if conf.CurrentContext != nil {
		currentContext = *conf.CurrentContext
	}
	if conf.Contexts != nil {
		for _, c := range conf.Contexts {
			pContext := &PrintableContext{}
			if c.Cluster != nil {
				pContext.Cluster = *c.Cluster
			}
			if c.Region != nil {
				pContext.Region = *c.Region
			}
			if c.Profile != nil {
				pContext.Profile = *c.Profile
			}
			if c.Name != nil {
				pContext.Name = *c.Name
				if *c.Name == currentContext {
					pContext.Current = "*"
				}
			}
			pContexts = append(pContexts, pContext)
		}
	}
	printerutils.PrintObjects(pContexts)
}

func writeCurrentContext(contextName string) error {
	viper.Set("currentcontext", contextName)
	return viper.WriteConfig()
}

func GetCluster(cmd *cobra.Command) (string, error) {
	return getFlagStringVal("cluster", cmd)
}

func GetRegion(cmd *cobra.Command) (string, error) {
	return getFlagStringVal("region", cmd)
}

func GetProfile(cmd *cobra.Command) (string, error) {
	return getFlagStringVal("profile", cmd)
}

func getFlagStringVal(name string, cmd *cobra.Command) (string, error) {
	val, err := cmd.Flags().GetString(name)
	if err != nil {
		return "", err
	}

	flag := cmd.Flag(name)

	if flag.Changed {
		// user set flag, return the value found above
		return val, nil
	} else {
		// user didn't set flag, check if value exists within config
		if currentContext, err := conf.getCurrentContext(); err == nil {
			switch name {
			case "profile":
				if currentContext.Profile != nil {
					return *currentContext.Profile, nil
				}
			case "region":
				if currentContext.Region != nil {
					return *currentContext.Region, nil
				}
			case "cluster":
				if currentContext.Cluster != nil {
					return *currentContext.Cluster, nil
				}
			}

		}
		return flag.DefValue, nil
	}
}

func (c *Config) getCurrentContext() (*Context, error) {
	if c.CurrentContext == nil {
		return nil, errCurrentContextNotFound
	}
	if c.Contexts == nil {
		return nil, errContextsNotFound
	}
	for _, context := range c.Contexts {
		if *context.Name == *c.CurrentContext {
			return context, nil
		}
	}
	return nil, errSpecifiedContextNotFound
}
