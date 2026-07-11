package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	vaultToken := "root"
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "8080"
		log.Println("PORT environment variable not set, defaulting to", port)
	}
	vaultUrl := os.Getenv("VAULT_ADDR")
	if vaultUrl == "" {
		vaultUrl = "http://vault:8200"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received Request - Port forwarding is working.")
		// If the JWT path is setup then get the new token from Vault using the k8s auth
		jwtPath := os.Getenv("JWT_PATH")
		if jwtPath != "" {
			jwtFile, err := os.ReadFile(jwtPath)
			if err != nil {
				fmt.Println("Error reading JWT file at", jwtPath, ": ", err)
				return
			}
			jwt := string(jwtFile)
			fmt.Println("Read JTW:", jwt)
			authPath := "auth/kubernetes/login"
			// create the payload for Vault authentication
			pl := VaultJWTPayload{Role: "webapp", JWT: jwt}
			jwtPayload, err := json.Marshal(pl)
			if err != nil {
				fmt.Println("Error encoding Vault request JSON:", err)
				return
			}
			fmt.Printf("jwtPayload: %s\n", string(jwtPayload))
			// Send a request to Vault to retrieve a token
			vaultLoginResponse := &VaultLoginResponse{}
			err = SendRequest(vaultUrl+"/v1/"+authPath, "", "POST", jwtPayload, vaultLoginResponse)
			if err != nil {
				fmt.Println("Error getting response from Vault k8s login:", err)
				return
			}
			fmt.Printf("vaultLoginResponse: %v\n", vaultLoginResponse)
			vaultToken = vaultLoginResponse.Auth.ClientToken
			fmt.Printf("Retrieved token: %v\n", vaultToken)
		}
		secretsPath := "secret/data/webapp/config"
		// Send a request to Vault using the token to retrieve the secret
		vaultSecretResponse := &VaultSecretResponse{}
		err := SendRequest(vaultUrl+"/v1/"+secretsPath, vaultToken, "GET", nil, &vaultSecretResponse)
		if err != nil {
			fmt.Println("Error getting secret from Vault:", err)
			return
		} else {
			fmt.Println("Successfully retrieved secret from Vault.")
		}
		fmt.Printf("Vault Secret Response: %+v\n", vaultSecretResponse)
		secretData := vaultSecretResponse.Data.Data
		fmt.Printf("Secret Data: %v", secretData)
	})
	log.Println("Listening on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %x", err)
	}
}

func SendRequest(url string, token string, requestType string, payload []byte, target any) error {
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("X-Vault-Token", token)
	}
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to Vault:", err)
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}
