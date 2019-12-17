package project

import (
	"github.com/giantswarm/versionbundle"
)

func NewVersionBundle() versionbundle.Bundle {
	return versionbundle.Bundle{
		Changelogs: []versionbundle.Changelog{
			{
				Component:   "master-operator",
				Description: "Master",
				Kind:        versionbundle.KindChanged,
			},
		},
		Components: []versionbundle.Component{},
		Name:       "master-operator",
		Version:    BundleVersion(),
	}
}
