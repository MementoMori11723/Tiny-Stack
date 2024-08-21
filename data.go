package main

type datapath map[string]func() result

type result struct {
	data any
	err  error
}

var dataRoutes = datapath{
	"/todos": func() result {
		return result{
      data: []string{"todo1", "todo2", "todo3"}, 
      err: nil,
    }
	},
}

func returnData(path string) any {
	if handler, ok := dataRoutes[path]; ok {
		result := handler()
		if result.err != nil {
			return result.err
		}
		return result.data
	}
	return nil
}
