package main

import (
	"github.com/golang/mock/gomock"
	"log"
	"testing"
)

func TestGetName(t *testing.T) {
	//新建一个mockController
	ctrl := gomock.NewController(t)
	// 断言 DB.GetName() 方法是否被调用
	defer ctrl.Finish()

	//mock接口
	mock := NewMockOrderDBI(ctrl)
	//模拟传入值与预期的返回值，Times 断言方法将被调用的次数
	mock.EXPECT().GetName(gomock.Eq(1225)).Return("name-1").Times(1)
	mock.EXPECT().GetName(gomock.Eq(12256)).Return("name-2").Times(1)

	//前面定义了传入值与返回值
	//在这里
	if v := mock.GetName(1225); v != "name-1"{
		t.Fatal("expected name-1, but got", v)
	}else{
		log.Println("通过mock取到的name：",v)
	}

	if v := mock.GetName(12256); v != "name-2"{
		t.Fatal("expected name-2, but got", v)
	}else{
		log.Println("通过mock取到的name：",v)
	}
}
