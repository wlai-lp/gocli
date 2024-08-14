# TL;DR
1. init go project `go mod init gocli`
2. install cobra `go install github.com/spf13/cobra-cli@latest`
3. init cobra project `cobra-cli init`
4. build binary `go build`
5. run `./gocli`

### Notes
- this creates cmd/root.go which contains rootCmd object

## Add command
1. add timezone `cobra-cli add timezone` to create a /cmd/timezone.go
2. run `go run .` to see additional timezone command

### Notes
- this creates cmd/timezone.go AND adds itself to `rootCmd` object

## Add sub command
1. add submodule `cobra-cli add today --parent timezone`
2. it creates today.go in cmd, but the init code use `timezone.AddCommand` which needs to be updated to `timezoneCmd`

# Add Flags
in init of the module use rootCmd.Flags()
```
func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bo-botflow.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&name, "name", "n", "World", "Name to greet")
}
```
## Code pattern
- create a name variable
- reference the &name in Flags()
- use name in your logic
- optional: make required rootCmd.MarkFlagRequired("input")got 

## Flag Usage
`./mycli hello --name Alice`
`./mycli hello -n Alice`

# Package reference
- go mod init creates the top level module example.com/go-project
- all other modules are following that in the folder structure
- create go-project/cmd folder
    - create a root.go folder and has its package set to cmd
- in main.go, you can import
    - example.com/go-project/cmd package which contains all the cmd
    - the use cmd.Public methods

## TODO: update code


# Reference
https://www.digitalocean.com/community/tutorials/how-to-use-the-cobra-package-in-go


mkdir CobraDigitalOcean && cd CobraDigitalOcean
go mod init CobraDigitalOcean

go install github.com/spf13/cobra-cli@latest

cobra-cli init


# Coding Reference

## Map Array Ops
- add to array, string map to array
```
cmMap := make(map[string][]ConversationMessage)
cmMap[el.Group] = append(cmMap[el.Group], el)
```

## return string array
output := [3]string{"abc", "efg", "abc"}
return output[:]

## Init a string array
// Create a slice with the same length as lpBot.Groups
strSlice := make([]string, len(lpBot.Groups))