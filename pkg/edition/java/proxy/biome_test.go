package proxy

import "testing"

func TestBuildBiomeRegistry(t *testing.T) {
	v := buildBiomeRegistry()
	if v["type"].(string) != "minecraft:worldgen/biome" {
		t.Fail()
	}
}
