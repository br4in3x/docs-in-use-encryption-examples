package local

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

var credentials = map[string]string{
	// Mongo Paths + URI
	"MONGODB_URI":     "mongodb://user:password@dedicated-shard-00-00.jo9ls.mongodb.net:27017,dedicated-shard-00-01.jo9ls.mongodb.net:27017,dedicated-shard-00-02.jo9ls.mongodb.net:27017/?ssl=true&replicaSet=atlas-qqgby8-shard-0&authSource=admin&retryWrites=true&w=majority",
	"SHARED_LIB_PATH": "/workspaces/docs-in-use-encryption-examples/cryptshared/lib/mongo_crypt_v1.so",
}

// check if credentials object contains placeholder values
func check_for_placeholders() {
	var error_buffer []string
	placeholder_pattern, _ := regexp.Compile("^<.*>$")
	for key, value := range credentials {
		// check for placeholder text
		if placeholder_pattern.MatchString(string(value)) {
			error_message := fmt.Sprintf("You must fill out the %s field of your credentials object.\n", key)
			error_buffer = append(error_buffer, error_message)
		}
	}
	// raise an error if errors in buffer
	if len(error_buffer) > 0 {
		message := strings.Join(error_buffer[:], "\n")
		log.Fatal(message)
	}
}

// return credentials object and ensure it has been populated
func GetCredentials() map[string]string {
	check_for_placeholders()
	return credentials
}
