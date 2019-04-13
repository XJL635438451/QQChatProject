package common

import (
    "bufio"
    "encoding/binary"
    "encoding/json"
    "errors"
    "fmt"
    "net"
    "os"
    "runtime"
    "strings"
    "MyGitHubProject/Wechat/QQChatProject/proto"
)

func GetInput(memName string) (input string, err error) {
    fmt.Printf("%s", memName)
    var inputReader *bufio.Reader
    var platform string
    var sep string

    inputReader = bufio.NewReader(os.Stdin)
    input, err = inputReader.ReadString('\n')
    if err != nil {
        return
    }
    platform = strings.ToLower(runtime.GOOS)
    if platform == "windows" {
        sep = "\r\n"
    } else if platform == "linux" {
        sep = "\n"
    } else {
        err = errors.New(fmt.Sprintf("Unknown platform: %v.", platform))
        return
    }
    input = strings.Trim(input, sep)
    return
}

func SendMsg(conn net.Conn, msg proto.Message) (err error) {
    data, err := json.Marshal(msg)
    if err != nil {
        err = fmt.Errorf("Failed to marshal data[%v].", msg)
        return
    }

    var packLen uint32
    packLen = uint32(len(data))

    var buf [4]byte
    binary.BigEndian.PutUint32(buf[:], packLen)
    //send head
    n, err := conn.Write(buf[:])
    if err != nil || n != 4 {
        err = fmt.Errorf("Failed to sned head data[%v], Error: %v", buf[:], err)
        return
    }
    //send body
    _, err = conn.Write(data[:])
    if err != nil {
        err = fmt.Errorf("Failed to sned body data[%v], Error: %v", data[:], err)
        return
    }
    return
}

func RecvMsg(conn net.Conn) (msg proto.Message, err error) {
    var buf [8192]byte
    //receive head
    n, err := conn.Read(buf[0:4])
    if err != nil || n != 4 {
        err = fmt.Errorf("Failed to receive head data[%v], Error: %v", buf[:], err)
        return
    }
    //receive body
    var packLen uint32
    packLen = binary.BigEndian.Uint32(buf[0:4])
    if packLen > 8192 {
        err = fmt.Errorf("Failed to receive head data[%v], big data, packLen: %v", buf[0:4], packLen)
        return
    }
    n, err = conn.Read(buf[0:packLen])
    if err != nil || n != int(packLen) {
        err = fmt.Errorf("Failed to receive body data[%v], Error: %v", buf[0:packLen], err)
        return
    }

    err = json.Unmarshal(buf[0:packLen], &msg)
    if err != nil {
        err = fmt.Errorf("Failed to unmarshal data[%v].", buf[0:packLen])
        return
    }
    return
}
