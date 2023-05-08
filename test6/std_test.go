package test6

import (
	"errors"
	"testing"
)

func TestIsNotNegative(log *testing.T){
	err:=errors.New("is negative")
	if IsNotNegative(-1) {
		log.Log("ok")
	}else{
		log.Error(err)
	}
	if IsNotNegative(1) {
		log.Log("ok")
	}else{
		log.Error(err)
	}
		
}