package main

import (
	"fmt"
	"github.com/dbferreira/gutils"
	"strings"
)

type user struct {
	name    string
	follows []string
}

type message struct {
	id      int
	content string
	mUser   user
}

// Does not return any value, but uses pointers to update the variables in the main function
func processUsers(s *[]string, u map[string]user) {
	for _, v := range *s {
		splitString := strings.Split(v, " follows ")
		followsString := strings.Split(splitString[1], ",")
		for fi, fv := range followsString {
			followsString[fi] = strings.Trim(fv, " ")
		}
		userName := splitString[0]
		if eUser, exists := u[userName]; exists {
			nFollows := append(followsString, eUser.follows...)
			utils.RemoveDuplicateStrings(&nFollows)
			u[userName] = user{
				name:    userName,
				follows: nFollows,
			}
		} else {
			u[userName] = user{
				name:    userName,
				follows: followsString,
			}
		}
	}
}

func processTweets(s *[]string, t *[]message, u map[string]user) {
	for i, v := range *s {
		splitString := strings.Split(v, "> ")
		sContent := splitString[1]
		sUser := splitString[0]
		sContent = strings.Trim(sContent, " ")
		sUser = strings.Trim(sUser, " ")
		var userStruct user
		if eUser, exists := u[sUser]; exists {
			userStruct = eUser
		} else {
			userStruct = user{
				name: sUser,
			}
			u[sUser] = userStruct
		}
		*t = append(*t, message{
			id:      i,
			content: sContent,
			mUser:   userStruct,
		})
	}
}

func completeUserList(u map[string]user) {
	for _, v := range u {
		users := v.follows
		for _, fv := range users {
			if _, exists := u[fv]; !exists {
				u[fv] = user{
					name: fv,
				}
			}
		}
	}
}

func sendTweets(t *[]message, u map[string]user) {
	for _, v := range u {
		fmt.Println(v.name)
		for _, mv := range *t {
			if mv.mUser.name == v.name {
				fmt.Println("@"+v.name+":", mv.content)
			}
			for _, tv := range v.follows {
				if tv == mv.mUser.name {
					fmt.Println("@"+tv+":", mv.content)
				}
			}
		}
	}
}

func main() {
	Users := make(map[string]user)
	Tweets := make([]message, 0)
	userStrings := utils.Readfile("user.txt")
	tweetStrings := utils.Readfile("tweet.txt")
	processUsers(&userStrings, Users)
	processTweets(&tweetStrings, &Tweets, Users)
	completeUserList(Users)
	sendTweets(&Tweets, Users)
}
