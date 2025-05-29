package main

import "fmt"

type Enty struct {
	C1 string
	C2 string
}

type IRepo1 interface {
	Save(d Enty) Enty
	List() string
}
type IRepoProps = IRepo1

type IEntyRepo struct {
	Save IRepo1.Save
	List IRepo1.List
}

// func Save (d Enty) Enty {
// 	return d
// }

// func List () string {
// 	return "Lista aqui"
// }

// func SaveEntyVarOBJRepo() IEntyRepo {
// 	res := IEntyRepo{
// 		Save: Save,
// 		List: List,
// 	}
// 	return res
// }

var memoryEntyRepository = func () IEntyRepo {
	res := IEntyRepo{
		Save: func(d Enty) Enty {
		},
		List: func() string {
			return "lista aqui"
		},
	}
	return res

}

var SaveEntyDTO1 = Enty{
	C1: "Reinaldo",
	C2: "Zachars",
}

func main() {

	create := memoryEntyRepository.Save(SaveEntyDTO1)

	res := create
	fmt.Println(res)
}
