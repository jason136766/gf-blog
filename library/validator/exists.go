package validator

import (
	"context"
	"errors"
	"strings"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/util/gvalid"
)

func init() {
	err := gvalid.RegisterRule("exists", existsChecker)
	if err != nil {
		panic(err)
	}

	err = gvalid.RegisterRule("not-exists", existsChecker)
	if err != nil {
		panic(err)
	}
}

// existsChecker 校验字段值是否存在
func existsChecker(ctx context.Context, rule string, value interface{}, message string, data interface{}) error {
	if value == nil {
		return nil
	}
	val := value.(string)
	var spliced []string
	isExists := true

	if strings.HasPrefix(rule, "exists:") {
		spliced = strings.Split(strings.TrimPrefix(rule, "exists:"), ",")
	} else if strings.HasPrefix(rule, "not-exists:") {
		spliced = strings.Split(strings.TrimPrefix(rule, "not-exists:"), ",")
		isExists = false
	}

	count, err := g.Model(spliced[0]).Count(spliced[1]+"= ?", val)
	if err != nil {
		return err
	}

	if (isExists && count == 0) || (!isExists && count != 0) {
		msg := strings.Replace(message, ":value", val, -1)
		return errors.New(msg)
	}
	return nil
}
