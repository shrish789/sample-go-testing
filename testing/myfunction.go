package testing

func MyFunction(d MyDependency) (string, error) {
    return d.GetMessage()
}
