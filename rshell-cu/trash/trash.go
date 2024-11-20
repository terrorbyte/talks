package main

import (
	"bytes"
	_ "embed"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	prompt "github.com/confluentinc/go-prompt"
	"github.com/mattn/go-isatty"
)

//go:embed trash.ascii
var banner string

//go:embed trash.1
var manRoff string

var TrashDir = os.Getenv("TRASHDIR")

func init() {
	_, err := os.Stat("trashfiles")
	if os.IsNotExist(err) {
		check := func(e error) {
			if e != nil {
				panic(e)
			}
		}
		createFile := func(name string, content []byte, mode fs.FileMode) {
			check(os.WriteFile(name, content, mode))
		}
		createDirs := func(name string, mode fs.FileMode) {
			check(os.MkdirAll(name, mode))
		}
		createFile("trash.1", []byte(manRoff), 0o644)

		createDirs("trashfiles/config", 0o755)
		createDirs("trashfiles/nvram", 0o755)
		createDirs("trashfiles/files", 0o755)
		createDirs("trashfiles/files/etc", 0o755)
		createDirs("trashfiles/files/bin", 0o755)
		createDirs("trashfiles/files/lib", 0o755)

		createFile("trashfiles/files/things", []byte{}, 0o644)
		createFile("trashfiles/files/stuff", []byte{}, 0o644)
		for range 10 {
			createFile("trashfiles/files/"+gofakeit.HackerNoun()+"."+gofakeit.FileExtension(), []byte(gofakeit.LoremIpsumSentence(50)), 0o644)
		}
		createFile("trashfiles/files/documents.txt", []byte(gofakeit.AchRouting()+" "+gofakeit.AchAccount()+"\n"), 0o644)
		createFile("trashfiles/nvram/mtd0", bytes.Repeat([]byte{0x1}, 256), 0o640)
		createFile("trashfiles/nvram/mtd1", bytes.Repeat([]byte{0x2}, 8096), 0o640)
		createFile("trashfiles/nvram/mtd2", []byte{0x3}, 0o640)
		createFile("trashfiles/config/admin_password", []byte(gofakeit.BitcoinPrivateKey()), 0o600)
	}
}

var SessionState struct {
	Completion bool
}

var LivePrefixState struct {
	LivePrefix     string
	CurrentSection string
	IsEnable       bool
}

var globalSection = []prompt.Suggest{
	{Text: "nocomplete", Description: "toggle completion"},
	{Text: "getenv", Description: "get environment variable"},
	{Text: "setenv", Description: "set environment variable"},
	{Text: "back", Description: "return to root"},
	{Text: "exit", Description: "exit"},
}

var rootSection = []prompt.Suggest{
	{Text: "help", Description: "print this help"},
	{Text: "files", Description: "filesystem actions"},
	{Text: "net", Description: "networking actions"},
	{Text: "man", Description: "read manual"},
	{Text: "embedded", Description: "embedded device actions"},
	{Text: "enable", Description: "admin mode"},
}

var fileSection = []prompt.Suggest{
	{Text: "ls", Description: "list files"},
	{Text: "cat", Description: "output files"},
	{Text: "find", Description: "find a file by name"},
}

var netSection = []prompt.Suggest{
	{Text: "ip", Description: "IP information"},
	{Text: "ping", Description: "ping IP"},
}

var embeddedSection = []prompt.Suggest{
	{Text: "devices", Description: "print nvram devices"},
	{Text: "nvread", Description: "print hex of nvram target"},
	{Text: "nvwrite", Description: "write hex to nvram target"},
}

func help() {
	switch LivePrefixState.CurrentSection {
	case "files":
		fmt.Printf("global ────────────────────────────\n")
		for _, s := range globalSection {
			fmt.Printf("%s - %s\n", s.Text, s.Description)
		}
		fmt.Printf("\nfiles ─────────────────────────────\n")
		for _, s := range fileSection {
			fmt.Printf("%s - %s\n", s.Text, s.Description)
		}
	case "net":
		fmt.Printf("global ────────────────────────────\n")
		for _, s := range globalSection {
			fmt.Printf("%s - %s\n", s.Text, s.Description)
		}
		fmt.Printf("\nnet ───────────────────────────────\n")
		for _, s := range netSection {
			fmt.Printf("%s - %s\n", s.Text, s.Description)
		}
	case "embedded":
		fmt.Printf("global ────────────────────────────\n")
		for _, s := range globalSection {
			fmt.Printf("%s - %s\n", s.Text, s.Description)
		}
		fmt.Printf("\nembedded ──────────────────────────\n")
		for _, s := range embeddedSection {
			fmt.Printf("%s - %s\n", s.Text, s.Description)
		}
	default:
		fmt.Printf("global ────────────────────────────\n")
		for _, s := range globalSection {
			fmt.Printf("%s - %s\n", s.Text, s.Description)
		}
		fmt.Printf("\ncommands ──────────────────────────\n")
		for _, s := range rootSection {
			fmt.Printf("%s - %s\n", s.Text, s.Description)
		}
		fmt.Printf("\nfiles ─────────────────────────────\n")
		for _, s := range fileSection {
			fmt.Printf("%s - %s\n", s.Text, s.Description)
		}
		fmt.Printf("\nnet ───────────────────────────────\n")
		for _, s := range netSection {
			fmt.Printf("%s - %s\n", s.Text, s.Description)
		}
		fmt.Printf("\nembedded ──────────────────────────\n")
		for _, s := range embeddedSection {
			fmt.Printf("%s - %s\n", s.Text, s.Description)
		}
	}
}

func ls() {
	root := TrashDir
	if TrashDir == "" {
		root = "trashfiles/files"
	}
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
		if path != "." {
			info, err := d.Info()
			if err != nil {
				fmt.Printf("%s\n", err.Error())

				return err
			}
			// gross. this is messy to get underlying interface
			s := info.Sys()
			if s == nil {
				return nil
			}
			sys := s.(*syscall.Stat_t)
			luser, err := user.LookupId(strconv.Itoa(int(sys.Uid)))
			if err != nil {
				luser.Username = strconv.Itoa(int(sys.Uid))
			}
			group, err := user.LookupGroupId(strconv.Itoa(int(sys.Gid)))
			if err != nil {
				group.Name = strconv.Itoa(int(sys.Gid))
			}
			fmt.Printf("%s\t%s\t%s\t%d\t%s\t%s\n", info.Mode().String(),
				luser.Username, group.Name,
				info.Size(), info.ModTime().Format(time.RFC822), path)

		}
		return err
	})
}

func find(args string) {
	root := TrashDir
	if TrashDir == "" {
		root = "trashfiles/files"
	}
	a := []string{root, "-iname"}
	a = append(a, strings.Split(args, " ")...)
	cmd := exec.Command("find", a...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func cat(args string) {
	root := TrashDir
	if TrashDir == "" {
		root = "trashfiles/files"
	}
	if strings.Contains(args, "..") {
		fmt.Printf("Path traversal? This incident will be reported.\n")

		os.Exit(1)
	}
	data, err := os.ReadFile(root + "/" + args)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	os.Stdout.Write(data)
	fmt.Printf("\n")
}

func ip() {
	cmd := exec.Command("ip", "addr")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func ping(args string) {
	a := []string{"-c", "ping " + args}
	cmd := exec.Command("sh", a...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func shell() {
	cmd := exec.Command("sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func man() {
	cmd := exec.Command("man", "./trash.1")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func devices() {
	root := TrashDir
	if TrashDir == "" {
		root = "trashfiles/nvram"
	}
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
		if path != "." {
			fmt.Printf("%s\n", path)
		}
		return err
	})
}

func nvread(args string) {
	a := strings.Split(args, " ")
	if len(a) != 1 {
		fmt.Println("Incorrect number of arguments, expected: nvread <slot>")

		return
	}
	root := TrashDir
	if TrashDir == "" {
		root = "trashfiles/nvram"
	}
	data, err := os.ReadFile(root + "/" + a[0])
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Println(hex.Dump(data))
}

func nvwrite(args string) {
	a := strings.Split(args, " ")
	if len(a) != 3 {
		fmt.Println("Incorrect number of arguments, expected: nvwrite <slot> <offset> <hex-string>")

		return
	}
	data, err := hex.DecodeString(a[2])
	if err != nil {
		fmt.Println("<hex-string> could not be decoded.")

		return
	}
	offset, err := strconv.Atoi(a[1])
	if err != nil {
		fmt.Println("<offset> offset could not be converted.")

		return
	}
	root := TrashDir
	if TrashDir == "" {
		root = "trashfiles/nvram"
	}
	file, err := os.OpenFile(root+"/"+a[0], os.O_RDWR, 0o644)
	if err != nil {
		fmt.Printf("Error file: %s\n", err.Error())

		return
	}
	defer file.Close()
	offsetWriter := io.NewOffsetWriter(file, int64(offset))
	_, err = offsetWriter.Write(data)
	if err != nil {
		fmt.Printf("Error write: %s\n", err.Error())

		return
	}
}

func fileExecutor(in string) {
	var base string
	var args string
	if strings.Contains(in, " ") {
		base = strings.SplitN(in, " ", 2)[0]
		args = strings.SplitN(in, " ", 2)[1]
	} else {
		base = in
	}
	switch base {
	case "":
	case "files":
		fileExecutor(args)
	case "ls":
		ls()
	case "cat":
		cat(args)
	case "find":
		find(args)
	default:
		fmt.Printf("files command %s not found\n", base)
	}
}

func devExecutor(in string) {
	var base string
	// var args string
	if strings.Contains(in, " ") {
		base = strings.SplitN(in, " ", 2)[0]
		// args = strings.SplitN(in, " ", 2)[1]
	} else {
		base = in
	}
	switch base {
	case "shell":
		shell()
	case "crash":
		panic("oh $#@!")
	case "", "dev":
	default:
		fmt.Printf("dev command %s not found\n", base)
	}
}

func netExecutor(in string) {
	var base string
	var args string
	if strings.Contains(in, " ") {
		base = strings.SplitN(in, " ", 2)[0]
		args = strings.SplitN(in, " ", 2)[1]
	} else {
		base = in
	}
	switch base {
	case "":
	case "net":
		netExecutor(args)
	case "ip":
		ip()
	case "ping":
		ping(args)
	default:
		fmt.Printf("net command %s not found\n", in)
	}
}

func embeddedExecutor(in string) {
	var base string
	var args string
	if strings.Contains(in, " ") {
		base = strings.SplitN(in, " ", 2)[0]
		args = strings.SplitN(in, " ", 2)[1]
	} else {
		base = in
	}
	switch base {
	case "":
	case "embedded":
		embeddedExecutor(args)
	case "nvread":
		nvread(args)
	case "devices":
		devices()
	case "nvwrite":
		nvwrite(args)

	default:
		fmt.Printf("util command %s not found\n", in)
	}
}

func executor(in string) {
	var base string
	var args string
	if strings.Contains(in, " ") {
		base = strings.SplitN(in, " ", 2)[0]
		args = strings.SplitN(in, " ", 2)[1]
	} else {
		base = in
	}
	if in == "" {
		return
	}

	if base == "help" {
		help()
		return
	}
	if base == "exit" {
		os.Exit(0)
		return
	}
	if base == "back" {
		LivePrefixState.IsEnable = false
		LivePrefixState.CurrentSection = ""
		LivePrefixState.LivePrefix = in
		return
	}
	if base == "nocomplete" {
		SessionState.Completion = !SessionState.Completion
		return
	}
	if base == "enable" {
		fmt.Printf("Password: ")
		var pass string
		_, err := fmt.Scanf("%s", &pass)
		if err != nil {
			log.Fatal(err)
		}
		root := TrashDir
		if TrashDir == "" {
			root = "trashfiles/config/admin_password"
		}
		data, err := os.ReadFile(root)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
		fmt.Printf("\n")
		if strings.TrimSpace(string(data)) == strings.TrimSpace(pass) {
			shell()

			return
		}

		return
	}
	if base == "setenv" {
		if strings.Contains(args, " ") {
			key := strings.Split(args, " ")[0]
			value := strings.Split(args, " ")[1]
			os.Setenv(key, value)
		}

		return
	}
	if base == "getenv" {
		fmt.Println(os.Getenv(args))
		return
	}
	if base == "dev" {
		devExecutor(args)
		return
	}
	if base == "man" {
		man()
		return
	}
	if LivePrefixState.CurrentSection == "" {
		switch base {
		case "files":
			LivePrefixState.LivePrefix = base + "$ "
			LivePrefixState.CurrentSection = "files"
			LivePrefixState.IsEnable = true
		case "net":
			LivePrefixState.LivePrefix = base + "$ "
			LivePrefixState.CurrentSection = "net"
			LivePrefixState.IsEnable = true
		case "embedded":
			LivePrefixState.LivePrefix = base + "$ "
			LivePrefixState.CurrentSection = "embedded"
			LivePrefixState.IsEnable = true
		default:
			LivePrefixState.LivePrefix = "$ "
			LivePrefixState.CurrentSection = ""
			LivePrefixState.IsEnable = false
		}
	}

	switch LivePrefixState.CurrentSection {
	case "files":
		fileExecutor(in)
	case "net":
		netExecutor(in)
	case "embedded":
		embeddedExecutor(in)
	default:
		fmt.Printf("command %s not found\n", base)
	}
}

func completer(in prompt.Document) []prompt.Suggest {
	if !SessionState.Completion {
		return []prompt.Suggest{}
	}
	switch LivePrefixState.CurrentSection {
	case "files":
		return prompt.FilterHasPrefix(append(fileSection, globalSection...), in.GetWordBeforeCursor(), true)
	case "net":
		return prompt.FilterHasPrefix(append(netSection, globalSection...), in.GetWordBeforeCursor(), true)
	case "embedded":
		return prompt.FilterHasPrefix(append(embeddedSection, globalSection...), in.GetWordBeforeCursor(), true)
	// case "enable":
	//	return []prompt.Suggest{}
	default:
		return prompt.FilterHasPrefix(append(rootSection, globalSection...), in.GetWordBeforeCursor(), true)
	}
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func main() {
	fmt.Println(banner)
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Err: %s", err.Error())

		return
	}
	SessionState.Completion = true
	fmt.Printf("Welcome to trash %s. Type `help` for command information\n", currentUser.Username)
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println("trash requires a full tty")

		return
	}

	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("$ "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionDescriptionBGColor(prompt.Black),
		prompt.OptionDescriptionTextColor(prompt.White),
		prompt.OptionSelectedDescriptionBGColor(prompt.White),
		prompt.OptionSelectedDescriptionTextColor(prompt.Black),
		prompt.OptionSuggestionBGColor(prompt.Black),
		prompt.OptionSuggestionTextColor(prompt.DarkGreen),
		prompt.OptionSelectedSuggestionBGColor(prompt.DarkGreen),
		prompt.OptionSelectedSuggestionTextColor(prompt.Black),
		prompt.OptionTitle("trash"),
	)
	p.Run()
}
