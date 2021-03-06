package main


import (

	"flag"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
)

func main() {
	var use string
	var count int
	fmt.Println("Enter Valid Username")
	fmt.Scanln(&use)
	s :=flag.String("TwitterUserName",use,"Usernames")
	flag.Parse()
	config := oauth1.NewConfig("AccessToken ,  "AccessToken")
	token := oauth1.NewToken("Secret_id", "Secret_id")
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	f, err:= os.Create("test.txt")

	params := &twitter.FollowerListParams{ScreenName: *s , Count:200 }
	followers, resp, err:= client.Followers.List(params)
	fmt.Println(resp , err)
	f.WriteString("UserName:")
	f.WriteString(use)
	f.WriteString("\n")
	f.WriteString("Followers :" + *s)
	for _,users:= range followers.Users{
		f.WriteString("\n" + users.ScreenName)
		count++

	}
	f.WriteString("\n")
	f.WriteString(fmt.Sprintf("No of followers %d\n", count))
	f.WriteString("\n")
	f.WriteString(fmt.Sprintf("No of followers either 200 or more than 200"))
	f.WriteString("\n")
	f.WriteString(fmt.Sprintf("Maximum Limit that can be shown is 200"))
	f.Close()
}
