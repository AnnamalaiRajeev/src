package flags

import (
	"flag"
)

var (
	err error;
	String_Flag *string;
	Int_Flag *int;
)

String_Flag, err = func(s1,s2,s3 string) (*string, error){
		s := flag.String(s1,s2,s3)
		return s
}()

Int_Flag, err = func(s1 string ,d1 int ,s3 string) (*int, error){
	s := flag.Int(s1, d1, s3)
	return s
}()

