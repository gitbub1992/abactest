package main

import "fmt"

const modelText = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj,act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
`

//r.sub.Type == r.obj.GetType() &&  r.sub.Group == r.obj.Group  &&  r.sub.Method == r.act.Method

func DefineFunction(group, userType string) bool {
	if group == "gptn.mediator1" && userType == "user" {
		return true
	} else if group == "gptn.mediator1" && userType == "client" {
		return true
	} else {
		return false
	}
}

func DefineFunction2(method float64) bool {
	if method == 1 || method == 3 {
		return true
	} else {
		return false
	}
}

func DefineFunctionWrapper(args ...interface{}) (interface{}, error) {
	group := args[0].(string)
	userType := args[1].(string)
	return bool(DefineFunction(group,userType)),nil
}

func DefineFunctionWrapper2(args ...interface{}) (interface{}, error) {
	method := args[0].(float64)
	return bool(DefineFunction2(method)),nil
}


func main() {
	TestTeacherEnterSchoolGate()
	fmt.Println("--------")
	user1 := &User{
		Address:  "P13VBemDosoqQvQX6XaF84LsZMaRF7smxaF",
		UserName: "User1",
		CertId:   12345678,
		Type:     "user",
		Group:    "A",
		Credit:   50,
	}

	user2 := &User{
		Address:  "P1R8oJsCypC2BgLRuxpnXW9S9gK9swYXQf",
		UserName: "User2",
		CertId:   87654321,
		Type:     "user",
		Group:    "B",
		Credit:   90,
	}
	//  用户 palletone  属于 group1 , group2
	users := []*User{user1, user2}
	TestTokenPermission(users)
	fmt.Println("--------")
	TestContractPermission(users)
	fmt.Println("--------")
	TestContractInstallPermission(users)
}
