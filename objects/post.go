package objects

import (
	"fmt"
	"os/exec"
)
var Scripts string

func post(sql string) (sqlModel []byte) {
	cmd := exec.Command("/bin/php", "-C", Scripts, sql)
	bytes,err := cmd.Output()
	if err != nil {
		fmt.Println("sorry, sqlMode is null, err is", err)
		return nil
	}
	return bytes
}


