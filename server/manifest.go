// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "autotranslate",
  "name": "Autotranslate",
  "description": "Autotranslate plugin for Mattermost 5.22 and later.",
  "homepage_url": "https://github.com/mattermost/mattermost-plugin-autotranslate",
  "support_url": "https://github.com/mattermost/mattermost-plugin-autotranslate/issues",
  "release_notes_url": "https://github.com/mattermost/mattermost-plugin-autotranslate/releases/tag/v0.3.0",
  "version": "0.3.0",
  "min_server_version": "5.22.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "webapp": {
    "bundle_path": "webapp/dist/main.js"
  },
  "settings_schema": {
    "header": "",
    "footer": "",
    "settings": [
      {
        "key": "AWSAccessKeyID",
        "display_name": "AWS Access Key ID:",
        "type": "text",
        "help_text": "The access key ID from AWS.",
        "placeholder": "",
        "default": null
      },
      {
        "key": "AWSSecretAccessKey",
        "display_name": "AWS Secret Access Key:",
        "type": "text",
        "help_text": "The secret access key from AWS.",
        "placeholder": "",
        "default": null
      },
      {
        "key": "AWSRegion",
        "display_name": "AWS Region:",
        "type": "text",
        "help_text": "The region from AWS.",
        "placeholder": "",
        "default": "us-east-1"
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}
