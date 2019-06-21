package yamltmpl

func applyOverrides(args map[string]interface{}, overrides ...map[string]interface{}) {
	for _, override := range overrides {
		for k, v := range override {
			args[k] = v
		}
	}
}
