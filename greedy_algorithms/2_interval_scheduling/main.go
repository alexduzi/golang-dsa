package main

import (
	"fmt"
	"sort"
)

// Suponha que tenhamos N tarefas para realizar, cada uma com um tempo de começo
// e um tempo de fim. Ao começar uma tarefa, a executamos até o seu fim, e só então
// podemos começar outra. Qual o número máximo de tarefas que é possível realizar?
func main() {
	tasks := []Task{
		Task{0, 6},
		Task{1, 4},
		Task{3, 5},
		Task{3, 8},
		Task{4, 7},
		Task{5, 9},
		Task{6, 10},
		Task{8, 11},
	}
	fmt.Println(mostTasks(8, tasks))
}

func mostTasks(n int, tasks []Task) int {
	// 1. ordenar por tempo de término
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].end < tasks[j].end
	})

	ans, time := 0, 0
	// 2. percorrer as tarefas
	for i := 0; i < n; i++ {
		// 3. verificar se náo tem conflito
		if tasks[i].start >= time {
			// fmt.Printf("%d  %d\n", tasks[i].start, tasks[i].end)
			ans++
			time = tasks[i].end
		}
	}

	return ans
}

type Task struct {
	start int
	end   int
}
