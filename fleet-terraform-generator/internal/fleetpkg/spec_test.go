package fleetpkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnnotateFileMetadata(t *testing.T) {
	integ := Integration{
		Manifest: Manifest{
			Vars: []Var{
				{
					Name: "proxy_url",
				},
			},
			PolicyTemplates: []PolicyTemplate{
				{
					Vars: []Var{
						{
							Name: "proxy_url",
						},
					},
					Inputs: []Input{
						{
							Vars: []Var{
								{
									Name: "proxy_url",
								},
							},
						},
					},
				},
			},
		},
		DataStreams: map[string]*DataStream{
			"log": {
				Manifest: DataStreamManifest{
					Streams: []Stream{
						{
							Vars: []Var{
								{
									Name: "interval",
								},
							},
						},
					},
				},
			},
		},
	}

	const file = "manifest.yml"
	annotateFileMetadata(file, &integ)
	assert.Equal(t, file, integ.Manifest.Vars[0].Meta.File)
	assert.Equal(t, file, integ.Manifest.PolicyTemplates[0].Vars[0].Meta.File)
	assert.Equal(t, file, integ.Manifest.PolicyTemplates[0].Inputs[0].Vars[0].Meta.File)
	assert.Equal(t, file, integ.DataStreams["log"].Manifest.Streams[0].Vars[0].Meta.File)
}
