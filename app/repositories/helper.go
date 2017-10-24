package repositories

import "fmt"

func ValueExist(s interface{}, f string, v string) (bool, error) {
	count := 0

	err := db.Model(s).Where(fmt.Sprintf("%s = ?", f), v).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}