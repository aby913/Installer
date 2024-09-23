package templates

import (
	"text/template"

	"github.com/lithammer/dedent"
)

var TerminusUninstallScriptValues = template.Must(template.New("terminus-uninstall.sh").Parse(
	dedent.Dedent(`#!/bin/bash
set +x
os_type=$(uname -s)
base_dir={{ .BaseDir }}
phase={{ .Phase }}
installer_path="v{{ .Version }}"

all="$1"
args=""
if [ "$all" == "all" ]; then
  args+=" --all"
fi
if [ "${os_type}" == "Darwin" ]; then
  args+=" --minikube"
fi

sudo -E /bin/bash -c "terminus-cli terminus uninstall --base-dir $base_dir --phase $phase $args | tee $base_dir/versions/$installer_path/logs/uninstall.log"

`)))
