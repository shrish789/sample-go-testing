package services


func (w *Workflow) Example() string {
	example := w.Db.Example()

	if (example) {
		return "Hello World! Example function"
	}
	return "Yo!"
}
