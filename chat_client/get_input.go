package main

import (
    "errors"
    "fmt"
    "strconv"
    "time"
    "MyGitHubProject/Wechat/QQChatProject/proto"
    "MyGitHubProject/Wechat/QQChatProject/common"
)

//userId and passwd is global variable
func initUserInput() (err error) {
    fmt.Println("Welcom to chat client...")
    fmt.Println("Starting to login or register...")
    fmt.Println("Please input your id and password.")

    var i int
    //user input userId
    for i = 0; i < 3; i++ {
        input, err := common.GetInput("UserId: ")
        if err != nil {
            fmt.Println(err)
            continue
        }
        //convert string to int
        userId, err = strconv.Atoi(input)
        if err != nil {
            fmt.Println("UserId must be number. Error: ", err)
            continue
        }
        break
    }
    if i >= 3 {
        err = fmt.Errorf("Input userId error is greater than three times.")
        return
    }

    //user input passwd
    for i = 0; i < 3; i++ {
        var err error
        passwd, err = common.GetInput("Password: ")
        if err != nil {
            fmt.Println(err)
            continue
        }
        break
    }
    if i >= 3 {
        err = fmt.Errorf("Input password error is greater than three times.")
        return
    }

    fmt.Printf("userId: %d passwd: %s\n", userId, passwd)
    return
}

func userRegisterMessage(regCmd *proto.RegisterData) (err error) {
    fmt.Println("You already input your userid and paaaword, next input other message...")
    var i int
    //user input Nick
    for i = 0; i < 3; i++ {
        input, err := common.GetInput("Nick: ")
        if err != nil {
            fmt.Println(err)
            continue
        }
        regCmd.User.Nick = input
        break
    }
    if i >= 3 {
        err = errors.New("Input nick error is greater than three times.")
        return
    }

    //user input Sex
    for i = 0; i < 3; i++ {
        input, err := common.GetInput("Sex(male/female): ")
        if err != nil {
            fmt.Println(err)
            continue
        }
        sex := input
        if sex == common.SexMale || sex == common.SexFemale {
            regCmd.User.Sex = sex
        } else {
            fmt.Printf("Input error.")
            continue
        }
        break
    }
    if i >= 3 {
        err = errors.New("Input sex error is greater than three times.")
        return
    }

    //user input Header, the user image display will be considered later.
    for i = 0; i < 3; i++ {
        input, err := common.GetInput("Header: ")
        if err != nil {
            fmt.Println(err)
            continue
        }
        regCmd.User.Header = input
        break
    }
    if i >= 3 {
        err = errors.New("Input header error is greater than three times.")
        return
    }

    regCmd.User.LastLogin = time.Now()
    regCmd.User.Status = common.UserOnline
    return
}
