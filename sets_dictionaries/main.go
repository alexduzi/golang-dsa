package main

func main() {

}

// Set representa um conjunto de elementos únicos usando um mapa internamente
type Set struct {
	elements map[any]struct{}
}

// NewSet cria e retorna um novo conjunto vazio
func NewSet() Set {
	return Set{
		elements: make(map[any]struct{}),
	}
}

// Add insere um elemento no conjunto (ignorado se já existir)
func (set *Set) Add(elem any) {
	set.elements[elem] = struct{}{}
}

// Remove exclui um elemento do conjunto
func (set *Set) Remove(elem any) {
	delete(set.elements, elem)
}

// IsEmpty retorna true se o conjunto não contiver nenhum elemento
func (set *Set) IsEmpty(elem any) bool {
	return set.Size() == 0
}

// Size retorna a quantidade de elementos no conjunto
func (set *Set) Size() int {
	return len(set.elements)
}

// List retorna todos os elementos do conjunto em um slice
func (set *Set) List() (list []any) {
	for k := range set.elements {
		list = append(list, k)
	}
	return
}

// Has verifica se um elemento pertence ao conjunto
func (set *Set) Has(elem any) bool {
	_, ok := set.elements[elem]
	return ok
}

// Copy retorna uma cópia independente do conjunto
func (set *Set) Copy() Set {
	newSet := NewSet()
	for k := range set.elements {
		newSet.Add(k)
	}
	return newSet
}

// Union retorna um novo conjunto com todos os elementos do conjunto atual e dos conjuntos passados
func (set *Set) Union(sets ...Set) Set {
	s := set.Copy()
	for _, inSet := range sets {
		for _, elem := range inSet.elements {
			s.Add(elem)
		}
	}
	return s
}

// Intersection retorna um novo conjunto com apenas os elementos presentes em todos os conjuntos
func (set *Set) Intersection(sets ...Set) Set {
	inter := set.Copy()
	for k := range inter.elements {
		for _, set := range sets {
			if !set.Has(k) {
				inter.Remove(k)
				break
			}
		}
	}
	return inter
}

// Difference retorna os elementos do conjunto atual que não existem em t
func (set Set) Difference(t Set) Set {
	for k := range t.elements {
		if set.Has(k) {
			delete(set.elements, k)
		}
	}
	return set
}

// IsSubset retorna true se todos os elementos do conjunto atual estiverem contidos em t
func (set Set) IsSubset(t Set) bool {
	for k := range set.elements {
		if !t.Has(k) {
			return false
		}
	}
	return true
}
