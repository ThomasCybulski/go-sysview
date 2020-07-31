package gui

import "os"

// Asset defines an empty function that is needed for astilectron bootstrap
func Asset(name string) ([]byte, error) {
	return []byte{}, nil
}

// AssetDir defines an empty function that is needed for astilectron bootstrap
func AssetDir(name string) ([]string, error) {
	return []string{}, nil
}

// RestoreAssets defines an empty function that is needed for astilectron bootstrap
func RestoreAssets(dir, name string) error {
	return nil
}

// AssetInfo defines an empty function that is needed for astilectron bootstrap
func AssetInfo(name string) (os.FileInfo, error) {
	return nil, nil
}
