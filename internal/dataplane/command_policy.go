package dataplane

import "errors"

var ErrCommandNotAllowed = errors.New("local runtime command is not allowed")
var ErrRuntimeCommandInvalid = errors.New("runtime command is invalid")

var allowedRuntimeCommands = map[string]bool{"ip":true,"tc":true,"nft":true,"vtysh":true,"birdc":true,"gobgp":true}

func ValidateRuntimeCommand(name string) error {if !allowedRuntimeCommands[name]{return ErrCommandNotAllowed};return nil}
func ValidateRuntimeArgs(args []string) error {for _,a:=range args{if a==""{return ErrRuntimeCommandInvalid};for _,r:=range a{if r=='\x00'{return ErrRuntimeCommandInvalid}}};return nil}
