package memo

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	MemoDir string `json:"memo_dir"`
}

func DefaultConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".memo", "config.json"), nil
}

func SaveConfig(path string, cfg Config) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	body, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, body, 0o644)
}

func LoadOrDefaultConfig(path string) (Config, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			cwd, e := os.Getwd()
			if e != nil {
				return Config{}, e
			}
			return Config{MemoDir: cwd}, nil
		}
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(body, &cfg); err != nil {
		return Config{}, err
	}
	if cfg.MemoDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return Config{}, err
		}
		cfg.MemoDir = cwd
	}
	return cfg, nil
}
