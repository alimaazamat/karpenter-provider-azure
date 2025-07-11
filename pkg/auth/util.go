/*
Portions Copyright (c) Microsoft Corporation.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package auth

import (
	"fmt"

	"sigs.k8s.io/karpenter/pkg/operator"
)

func GetUserAgentExtension() string {
	// Note: do not change "karpenter-aks/" prefix, some infra depends on it
	return fmt.Sprintf("karpenter-aks/v%s", operator.Version)
}
