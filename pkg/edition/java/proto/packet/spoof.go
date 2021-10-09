package packet

import "go.minekube.com/gate/pkg/gate/proto"

// SpoofJoinGame builds a fake join game packet.
func SpoofJoinGame(dimensionName string) *JoinGame {
	if dimensionName == "" {
		dimensionName = "minecraft:overworld"
	}
	dimensionEffects := dimensionName
	coordinateScale := float64(1.0)
	dimensionHeight := int(256)
	dimensionLogicalHeight := int32(dimensionHeight)
	dimensionMinY := int(0)
	dimensionPiglineSafe := false
	dimensionID := int(0)
	dimension := &DimensionData{
		AmbientLight:               0,
		BurningBehaviourIdentifier: "minecraft:infiniburn_overworld",
		Ceiling:                    false,
		CoordinateScale:            &coordinateScale,
		CreateDragonFight:          nil,
		DimensionID:                &dimensionID,
		DoBedsWork:                 false,
		DoRespawnAnchorsWork:       false,
		Effects:                    &dimensionEffects,
		FixedTime:                  nil,
		Height:                     &dimensionHeight,
		LogicalHeight:              dimensionLogicalHeight,
		MinY:                       &dimensionMinY,
		Natural:                    true,
		PiglineSafe:                dimensionPiglineSafe,
		Raids:                      false,
		RegistryIdentifier:         dimensionName,
		Shrunk:                     false,
		Skylight:                   true,
		Ultrawarm:                  false,
	}
	return &JoinGame{
		Difficulty:        0,
		Dimension:         0,
		EntityID:          0,
		Gamemode:          0,
		PreviousGamemode:  255,
		Hardcore:          false,
		MaxPlayers:        0,
		LevelType:         nil,
		ViewDistance:      12,
		ReducedDebugInfo:  false,
		ShowRespawnScreen: true,
		PartialHashedSeed: -2589135725674695292,
		DimensionInfo: &DimensionInfo{
			RegistryIdentifier: dimensionName,
		},
		CurrentDimensionData: dimension,
		DimensionRegistry: &DimensionRegistry{
			Dimensions: []*DimensionData{dimension},
			LevelNames: []string{dimensionName},
		},
		BiomeRegistry: buildBiomeRegistry(),
	}
}

// SpoofPostLoginSequence spoofs the full load-in sequence.
func SpoofPostLoginSequence() []proto.Packet {
	// join game: primary source of info about the server
	joinGame := SpoofJoinGame("")

	// other packets (not sent):
	// - 0x32: player abilities
	// - 0x48: held item change
	// - 0x65: declare recipes
	// - 0x66: tags
	// - 0x1b: entity action
	// - 0x39: unlock recipes

	// position and look
	posAndLook := &PlayerPosAndLook{
		X:               248.69,
		Y:               68,
		Z:               -98,
		Yaw:             155,
		Pitch:           30,
		Flags:           0,
		TeleportID:      0, // possibly needs 1
		DismountVehicle: false,
	}

	return []proto.Packet{
		joinGame,
		posAndLook,
	}
}
