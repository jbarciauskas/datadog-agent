// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

// +build windows

package secrets

// Decrypt encrypted secrets are not available on windows
func Decrypt(data []byte) ([]byte, error) {
	return data, nil
}
