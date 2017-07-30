package main

import (
  "fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	ip := "192.168.33.10" //サーバのアドレス
	port := "22"          //ポート番号は文字列で指定
	user := "vagrant"     //ユーザ名
  keyFile := "./.vagrant/machines/default/virtualbox/private_key.pem"
	buf, err := readPemKey(keyFile)
  if err != nil {
    panic(err)
  }
	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		panic(err)
	}
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", ip+":"+port, config)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

  session, err := conn.NewSession()
  if err != nil {
    log.Println(err)
  }
  defer session.Close()
  fmt.Println("command:sudo apt-get update -y")
	mes, err := session.CombinedOutput("sudo apt-get update -y")
  s := string(mes[:])
  fmt.Println(s)

  session2, err := conn.NewSession()
  if err != nil {
    log.Println(err)
  }
  defer session2.Close()

  fmt.Println("command: sudo -E apt-get upgrade -y")
  mes2, err := session2.CombinedOutput("export DEBIAN_FRONTEND=noninteractive;sudo -E apt-get upgrade -y")
  s2 := string(mes2[:])
  fmt.Println(s2)
}

func readPemKey(keyname string) ([]byte, error) {
  buf, err := ioutil.ReadFile("./.vagrant/machines/default/virtualbox/private_key.pem")
  if err != nil {
    createPemKey()
    buf, err := ioutil.ReadFile("./.vagrant/machines/default/virtualbox/private_key.pem")
    return buf, err
  }
  return buf, err
}

func createPemKey() {
  fmt.Println("Now Creating pem key...")
	cmdstr := "openssl rsa -in ./.vagrant/machines/default/virtualbox/private_key -outform pem >./.vagrant/machines/default/virtualbox/private_key.pem"
	err := exec.Command("sh", "-c", cmdstr).Run()
	if err != nil {
		log.Println(err)
	}
  fmt.Println("complete!")
}
