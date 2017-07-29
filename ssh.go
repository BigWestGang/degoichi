package main

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	ip := "192.168.33.10" //サーバのアドレス
	port := "22"          //ポート番号は文字列で指定
	user := "vagrant"     //ユーザ名
	buf, err := ioutil.ReadFile("./.vagrant/machines/default/virtualbox/private_key.pem")
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

	//空のファイルを作成する
	session.Run("echo -n > empty.txt")
}

func createPemKey() {
	cmdstr := "openssl rsa -in ./.vagrant/machines/default/virtualbox/private_key -outform pem >./.vagrant/machines/default/virtualbox/private_key.pem"
	err := exec.Command("sh", "-c", cmdstr).Run()
	if err != nil {
		log.Println(err)
	}
}
