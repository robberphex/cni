// Copyright 2015 CNI authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package invoke

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FindInPath returns the full path of the plugin by searching in the provided path
func FindInPath(plugin string, paths []string) (string, error) {
	if plugin == "" {
		return "", fmt.Errorf("no plugin name provided")
	}

	if strings.ContainsRune(plugin, os.PathSeparator) {
		return "", fmt.Errorf("invalid plugin name: %s", plugin)
	}

	if len(paths) == 0 {
		return "", fmt.Errorf("no paths provided")
	}

	for _, path := range paths {
		for _, fe := range ExecutableFileExtensions {
			fullpath := filepath.Join(path, plugin) + fe
			if fi, err := os.Stat(fullpath); err == nil && fi.Mode().IsRegular() {
				return fullpath, nil
			}
		}
	}

	return "", fmt.Errorf("failed to find plugin %q in path %s", plugin, paths)
}

func ListInPath(paths []string) ([]string, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("no paths provided")
	}

	var plugins []string
	for _, path := range paths {
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			pluginName := entry.Name()
			executable := false
			for _, fe := range ExecutableFileExtensions {
				if strings.HasSuffix(pluginName, fe) {
					fullpath := filepath.Join(path, pluginName)
					fi, err := os.Stat(fullpath)
					if err != nil {
						return nil, err
					}
					if fi.Mode().IsRegular() {
						executable = true
						break
					}
				}
			}
			if executable {
				plugins = append(plugins, pluginName)
			}
		}
	}
	return plugins, nil
}
