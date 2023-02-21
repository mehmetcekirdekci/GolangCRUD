package services

import (
	"github.com/mehmetcekirdekci/GolangCRUD/app/product/types"
	"strings"
)

func ResultMessageBuilder(isOperationSuccesful bool, err error) string {
	var resultMessageBuilder strings.Builder
	if isOperationSuccesful {
		resultMessageBuilder.WriteString(types.SuccesfulOperation)
	} else {
		resultMessageBuilder.WriteString(types.OperationFailed)
		if err != nil {
			resultMessageBuilder.WriteString(" ")
			resultMessageBuilder.WriteString(err.Error())
		}
	}
	return resultMessageBuilder.String()
}
