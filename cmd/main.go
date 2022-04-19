package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func helloWorld(uri, username, password string) (string, error) {
	fmt.Println("hello")
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	fmt.Println("hello2")
	if err != nil {
		return "", err
	}
	defer driver.Close()

	fmt.Println("hello3")
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	fmt.Println("hello4")

	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]interface{}{"message": "hello, world"})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	fmt.Println("hello5")
	if err != nil {
		return "", err
	}

	return greeting.(string), nil
}

func main(){
	trunsctionResult, err := helloWorld("neo4j://neo4j:7687","neo4j","password")
	fmt.Println(trunsctionResult,err)
}
