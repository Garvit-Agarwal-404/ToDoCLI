package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.load(&todos)
	CmdFlags := NewCmdFlags()
	CmdFlags.Execute(&todos)
	storage.Save(todos)
}
