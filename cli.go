package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	"gopkg.in/ini.v1"
)

type Cli struct {
	file    *ini.File
	command string //保存命令
	key     string //保存key
	pass    string //新的password
	length  int    //密钥长度
}

var (
	c = Cli{}
)

//判断指令
func (cli *Cli) execCommand() error {
	cli.command = strings.ToLower(cli.command)
	var err error
	switch cli.command {
	case commandGen:
		err = cli.genPass()
	case commandGet:
		err = cli.getPass()
	case commandDel:
		err = cli.delPass()
	case commandSet:
		err = cli.setPass()
	default:
		err = errors.New("invalid command")
	}
	return err
}

func (cli *Cli) genPass() error {
	if cli.key == "" {
		return errors.New("invalid key")
	}
	section := cli.file.Section("pass")
	key, _ := section.GetKey(cli.key)
	if key != nil {
		fmt.Fprintf(os.Stdout, "key %s aready exist, overwrite it, y/n?", cli.key)
		ok, err := cli.getConfirm()
		if err != nil {
			return err
		}
		if !ok {
			return nil
		}
	}
	fmt.Fprintf(os.Stdout, fmt.Sprintf("input password lentgh(6-16):"))
	for {
		var input string
		fmt.Scan(&input)
		if n, err := strconv.ParseInt(input, 10, 64); err != nil {
			fmt.Fprintf(os.Stdout, "please input number between(6-16):")
			continue
		} else {
			cli.length = int(n)
		}
		if cli.length < 6 || cli.length > 16 {
			fmt.Fprintf(os.Stdout, "please input number between(6-16):")
			continue
		}
		break
	}
	cli.pass = genPass(cli.length)

	switch {
	case key != nil:
		key.SetValue(cli.pass)
	default:
		_, err := section.NewKey(cli.key, cli.pass)
		if err != nil {
			return err
		}
	}
	fmt.Fprintf(os.Stdout, fmt.Sprintf("%s: %s\n", cli.key, cli.pass))
	err := cli.file.SaveTo(fileName)
	return err
}

func (cli *Cli) getPass() error {
	section := cli.file.Section("pass")

	if cli.key != "" {
		key, _ := section.GetKey(cli.key)
		if key == nil {
			return nil
		}
		fmt.Fprintf(os.Stdout, fmt.Sprintf("%s: %s\n", cli.key, key.String()))
		return nil
	}

	keys := section.Keys()
	var pass []string

	for i := range keys {
		pass = append(pass, fmt.Sprintf("%s: %s\n", keys[i].Name(), keys[i].Value()))
	}
	for i := range pass {
		fmt.Fprintf(os.Stdout, "%s", pass[i])
	}
	return nil
}

func (cli *Cli) delPass() error {
	section := cli.file.Section("pass")
	if cli.key == "" {
		return errors.New("invalid key")
	}
	key, _ := section.GetKey(cli.key)
	if key == nil {
		return nil
	}
	fmt.Fprintf(os.Stdout, "confirm DEL the password of %s y/n?", cli.key)
	ok, err := cli.getConfirm()
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	section.DeleteKey(cli.key)
	err = cli.file.SaveTo(fileName)
	return err
}

func (cli *Cli) setPass() error {
	var err error
	if cli.key == "" {
		err = errors.New("invalid key")
		return err
	}
	if len(cli.pass) > 0 && len(cli.pass) < 6 {
		err = errors.New("pass length must more than 6")
		return err
	}
	if len(cli.pass) == 0 {
		fmt.Fprintf(os.Stdout, fmt.Sprintf("input password lentgh(6-16):"))
		for {
			var input string
			fmt.Scan(&input)
			if n, err := strconv.ParseInt(input, 10, 64); err != nil {
				fmt.Fprintf(os.Stdout, "please input number between(6-16):")
				continue
			} else {
				cli.length = int(n)
			}
			if cli.length < 6 || cli.length > 16 {
				fmt.Fprintf(os.Stdout, "please input number between(6-16):")
				continue
			}
			break
		}
		cli.pass = genPass(cli.length)
	}
	section := cli.file.Section("pass")
	key, _ := section.GetKey(cli.key)
	switch {
	case key == nil:
		key, err = section.NewKey(cli.key, cli.pass)
		if err != nil {
			return err
		}
	default:
		fmt.Fprintf(os.Stdout, "confirm SET the password of key %s y/n?", cli.key)
		ok, err := cli.getConfirm()
		if err != nil {
			return err
		}
		if !ok {
			return nil
		}
		key.SetValue(cli.pass)
	}
	err = cli.file.SaveTo(fileName)
	return err
}

func (cli *Cli) loadIni() error {
	var err error
	cli.file, err = ini.Load(fileName)
	if err != nil {
		return err
	}
	return nil
}

func (cli *Cli) getConfirm() (bool, error) {
	var flag string
	for {
		_, err := fmt.Scan(&flag)
		if err != nil {
			return false, err
		}
		flag = strings.ToLower(flag)
		switch flag {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		}
	}
}

func genPass(l int) string {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	var charCnt = 2
	if l > 12 {
		charCnt = 4
	}
	var bytes []byte
	for i := 0; i < charCnt; i++ {
		b := chars[rnd.Intn(20)]
		bytes = append(bytes, b)
	}
	for i := 0; i < l-charCnt; i++ {
		b := code[rnd.Intn(62)]
		bytes = append(bytes, b)
	}
	for i := l - 1; i > 0; i-- {
		randIndex := rnd.Intn(i)
		bytes[i], bytes[randIndex] = bytes[randIndex], bytes[i]
	}
	return string(bytes)
}

func getHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func initFile() error {
	home := getHomeDir()
	filePath := path.Join(home, "secret.ini")

	file, err := os.Open(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		file, err = os.Create(filePath)
		if err != nil {
			return err
		}
	}
	file.Close()
	fileName = filePath
	err = c.loadIni()
	return err
}

func showHelp() {
	fmt.Fprintf(os.Stderr, helpText)
}
