package render

import "github.com/solo-io/skv2/codegen/model"

// renders files used to build the operator
type BuildRenderer struct {
	templateRenderer
}

var defaultBuildInputs = func(build model.Build) inputTemplates {
	return inputTemplates{
		"build/Dockerfile.tmpl": {
			Path: build.Repository + "/Dockerfile",
		},
	}
}

func RenderBuild(build model.Build) ([]OutFile, error) {
	defaultBuildRenderer := BuildRenderer{
		templateRenderer: DefaultTemplateRenderer,
	}
	return defaultBuildRenderer.Render(build)
}

func (r BuildRenderer) Render(build model.Build) ([]OutFile, error) {
	templatesToRender := defaultBuildInputs(build)

	files, err := r.renderCoreTemplates(templatesToRender, build)
	if err != nil {
		return nil, err
	}

	return files, nil
}
