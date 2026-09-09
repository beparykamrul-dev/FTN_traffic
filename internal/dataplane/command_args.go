package dataplane

import "errors"

var ErrRuntimeCommandInvalid = errors.New("runtime command is invalid")

func ValidateRuntimeArgs(args []string) error {
 if len(args)==0{return ErrRuntimeCommandInvalid}
 for _,a:=range args { if a=="" || a=="-c" {continue}; for _,r:=range a { if r=='\x00' {return ErrRuntimeCommandInvalid} } }
 return nil
}
