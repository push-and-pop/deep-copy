package main

const cMonumentTypeNum = 6

type Good struct {
	Path string
}

type Fuck struct {
	Fork byte
}

type Diao struct {
	Num  int64
	Fuck *Fuck
	Good *Good
}

type Cao struct {
	Head int32
}

type Bar struct {
	ArrayPointer [cMonumentTypeNum]*Diao
	ListPointer  []*Diao
	Array        [cMonumentTypeNum]Diao
	Array2       [cMonumentTypeNum]Good
	List         []Diao
	List2        []Good
	Cao          *Cao
}

type Foo struct {
	Age  int64
	Name string
	Bar  *Bar
}

func (pOwn *Foo) DeepCopy() *Foo {
	pFoo := new(Foo)
	if pOwn.Bar != nil {
		pFoo.Bar = new(Bar)
		*pFoo.Bar = *pOwn.Bar
		pOwn.Bar.ArrayPointer = pFoo.Bar.ArrayPointer
		for i := range pOwn.Bar.ArrayPointer {
			if pOwn.Bar.ArrayPointer[i] != nil {
				pFoo.Bar.ArrayPointer[i] = new(Diao)
				*pFoo.Bar.ArrayPointer[i] = *pOwn.Bar.ArrayPointer[i]
				if pOwn.Bar.ArrayPointer[i].Fuck != nil {
					pFoo.Bar.ArrayPointer[i].Fuck = new(Fuck)
					*pFoo.Bar.ArrayPointer[i].Fuck = *pOwn.Bar.ArrayPointer[i].Fuck
				}
				if pOwn.Bar.ArrayPointer[i].Good != nil {
					pFoo.Bar.ArrayPointer[i].Good = new(Good)
					*pFoo.Bar.ArrayPointer[i].Good = *pOwn.Bar.ArrayPointer[i].Good
				}
			}
		}

		pFoo.Bar.ListPointer = make([]*Diao, len(pOwn.Bar.ListPointer))
		for i := range pOwn.Bar.ListPointer {
			pOwn.Bar.ListPointer[i] = pFoo.Bar.ListPointer[i]
			if pOwn.Bar.ListPointer[i] != nil {
				pFoo.Bar.ListPointer[i] = new(Diao)
				*pFoo.Bar.ListPointer[i] = *pOwn.Bar.ListPointer[i]
				if pOwn.Bar.ListPointer[i].Fuck != nil {
					pFoo.Bar.ListPointer[i].Fuck = new(Fuck)
					*pFoo.Bar.ListPointer[i].Fuck = *pOwn.Bar.ListPointer[i].Fuck
				}
				if pOwn.Bar.ListPointer[i].Good != nil {
					pFoo.Bar.ListPointer[i].Good = new(Good)
					*pFoo.Bar.ListPointer[i].Good = *pOwn.Bar.ListPointer[i].Good
				}
			}
		}
		pOwn.Bar.Array = pFoo.Bar.Array
		for i := range pOwn.Bar.Array {
			if pOwn.Bar.Array[i].Fuck != nil {
				pFoo.Bar.Array[i].Fuck = new(Fuck)
				*pFoo.Bar.Array[i].Fuck = *pOwn.Bar.Array[i].Fuck
			}
			if pOwn.Bar.Array[i].Good != nil {
				pFoo.Bar.Array[i].Good = new(Good)
				*pFoo.Bar.Array[i].Good = *pOwn.Bar.Array[i].Good
			}
		}
		pOwn.Bar.Array2 = pFoo.Bar.Array2

		pFoo.Bar.List = make([]Diao, len(pOwn.Bar.List))
		for i := range pOwn.Bar.List {
			pOwn.Bar.List[i] = pFoo.Bar.List[i]
			if pOwn.Bar.List[i].Fuck != nil {
				pFoo.Bar.List[i].Fuck = new(Fuck)
				*pFoo.Bar.List[i].Fuck = *pOwn.Bar.List[i].Fuck
			}
			if pOwn.Bar.List[i].Good != nil {
				pFoo.Bar.List[i].Good = new(Good)
				*pFoo.Bar.List[i].Good = *pOwn.Bar.List[i].Good
			}
		}

		pFoo.Bar.List2 = make([]Good, len(pOwn.Bar.List2))
		copy(pFoo.Bar.List2, pOwn.Bar.List2)
		if pOwn.Bar.Cao != nil {
			pFoo.Bar.Cao = new(Cao)
			*pFoo.Bar.Cao = *pOwn.Bar.Cao
		}
	}
	return pFoo
}
