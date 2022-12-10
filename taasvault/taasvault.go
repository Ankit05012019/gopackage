package taasvault

import (
	"fmt"
	"log"

	vault "github.com/hashicorp/vault/api"
)

//const password string = "<PASSWORD>"

func ReadData(VaultAddr, VaultToken, SecretPath string) (map[string]interface{}, error) {

	config := vault.DefaultConfig()
	config.Address = VaultAddr

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}

	client.SetToken(VaultToken)

	secret, err := client.Logical().Read(SecretPath)
	if err != nil {
		log.Fatalf(
			"Unable to read the super secret password from the vault: %v",
			err,
		)
		return nil, err
	}

	dat, ok := secret.Data["data"]
	if !ok {
		fmt.Errorf("secret: not found")
	}

	// convert the interface into map  and print the value of the password
	data, ok := dat.(map[string]interface{})
	//log.Printf("Passord is %s", data["Password"])

	return data, nil
}
