package main

import "fmt"

func main() {

	s1 := map[string]string{"name": "TTHE", "number": "1234"}
	s2 := map[string]string{
		"name": "TTHE", "number": "1234", //如果分行写，最后必须加，
	}
	s3 := map[string]map[string]string{"k1": s1, "k2": s2}

	fmt.Println(s1, s2)
	fmt.Println(s3)

	empty_s1 := map[string]string{} // is empty map, print: map[]
	var empty_s2 map[string]string  // actually is nil, but print result is map[]
	m1 := make(map[string]int)      // is empty map, print: map[]
	fmt.Println(empty_s1)
	fmt.Println(empty_s2)
	fmt.Println(m1)

	for k, v := range s2 {
		fmt.Println(k, v)
	}
	for _, v := range s2 {
		fmt.Println(v)
	}
	for k := range s2 {
		fmt.Println(k)
	}

	name := s2["name"] //TTHE
	nema := s2["nema"] //empty 空串
	//当不存在的时候，能够拿到值，但是是zero value，string的zero value就是空串

	fmt.Println(name)

	fmt.Println(nema)

	name2, ok := s2["name"] //TTHE true
	fmt.Println(name2, ok)

	nema2, ok := s2["nema"] // false

	fmt.Println(nema2, ok)

	if nema3, ok := s2["nema"]; ok {
		fmt.Println("existing nema, value is ", nema3)
	} else {
		fmt.Println("not existing nema")
	}

	name, ok = s2["name"] //TTHE true
	fmt.Println(name, ok)
	delete(s2, "name")    //直接有delete函数支持
	name, ok = s2["name"] // false
	fmt.Println(name, ok)

}
