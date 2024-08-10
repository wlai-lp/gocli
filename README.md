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


## TODO: update code


# Reference
https://www.digitalocean.com/community/tutorials/how-to-use-the-cobra-package-in-go


mkdir CobraDigitalOcean && cd CobraDigitalOcean
go mod init CobraDigitalOcean

go install github.com/spf13/cobra-cli@latest

cobra-cli init