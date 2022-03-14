/*
 * cfutil.go
 *
 * Copyright 2022 Immanuel Jeyaraj
 *
 * Author: Immanuel Jeyaraj <irj@sefier.com>
 *
 * Created date: 3 June 2019
 */

package sfutil

import (
	"crypto/sha256"
	"os"
	"os/user"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// Path_separator provides the path separator for various OS
func Path_separator() string {
	if runtime.GOOS == "windows" {
		return ("\\")
	} else {
		return ("/")
	}
}

// reports whether the named file or directory exists.
func File_is_exists(f string) bool {
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// Return the current user name
func CurrentUser() string {
	user, err := user.Current()
	if err != nil {
		return ""
	}

	if runtime.GOOS == "windows" {
		s := strings.Split(user.Username, "\\")
		return s[1]
	}
	return user.Username
}

func GetSplitValues(completestring string) []string {
	var tmpVal, arrStr []string

	tmpVal = strings.Split(completestring, ",")

	for _, each := range tmpVal {
		arrStr = append(arrStr, strings.Trim(each, " "))
	}

	return arrStr
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func GenerateAPIKey(emailid string) [32]byte {
	return sha256.Sum256([]byte(emailid + time.Now().String()))
}

// func InsertAPIKey(emailid string) {
// 	apikey := GenerateAPIKey(emailid)
// 	DBPool.Query("INSERT INTO `tbl_apilist` (email_address, api_key, status) VALUES ('" + emailid + "', '" + fmt.Sprintf("%x", apikey) +
// 		"', 1);")
// }

// func ValidateAPI(APIKey string) bool {
// 	var checkval int64

// 	if APIKey == AdminAPIKey() {
// 		return true
// 	}

// 	rows, err := DBPool.Query("SELECT count(1) FROM `tbl_apilist` WHERE `api_key` = '" + APIKey + "' AND status = 1;")
// 	if err != nil {
// 		//###		tracelog.Error(fmt.Errorf(fmt.Sprint(err)), "API Validate", "SQL-units")
// 		return false
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		rows.Scan(&checkval)
// 	}

// 	if checkval == 1 {
// 		return true
// 	}

// 	return false

// }
