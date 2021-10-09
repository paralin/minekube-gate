package packet

import "encoding/json"

const biomeRegistryJSON = `
{
  "value": [
    {
      "name": "minecraft:ocean",
      "id": 0,
      "element": {
        "temperature": 0.5,
        "scale": 0.1,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8103167,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": -1,
        "category": "ocean"
      }
    },
    {
      "name": "minecraft:plains",
      "id": 1,
      "element": {
        "temperature": 0.8,
        "scale": 0.05,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7907327,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.4,
        "depth": 0.125,
        "category": "plains"
      }
    },
    {
      "name": "minecraft:desert",
      "id": 2,
      "element": {
        "temperature": 2,
        "scale": 0.05,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 0.125,
        "category": "desert"
      }
    },
    {
      "name": "minecraft:mountains",
      "id": 3,
      "element": {
        "temperature": 0.2,
        "scale": 0.5,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8233727,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.3,
        "depth": 1,
        "category": "extreme_hills"
      }
    },
    {
      "name": "minecraft:forest",
      "id": 4,
      "element": {
        "temperature": 0.7,
        "scale": 0.2,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7972607,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.1,
        "category": "forest"
      }
    },
    {
      "name": "minecraft:taiga",
      "id": 5,
      "element": {
        "temperature": 0.25,
        "scale": 0.2,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8233983,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.2,
        "category": "taiga"
      }
    },
    {
      "name": "minecraft:swamp",
      "id": 6,
      "element": {
        "temperature": 0.8,
        "scale": 0.1,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 2302743,
          "water_color": 6388580,
          "sky_color": 7907327,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "grass_color_modifier": "swamp",
          "foliage_color": 6975545,
          "fog_color": 12638463
        },
        "downfall": 0.9,
        "depth": -0.2,
        "category": "swamp"
      }
    },
    {
      "name": "minecraft:river",
      "id": 7,
      "element": {
        "temperature": 0.5,
        "scale": 0,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8103167,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": -0.5,
        "category": "river"
      }
    },
    {
      "name": "minecraft:nether_wastes",
      "id": 8,
      "element": {
        "temperature": 2,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "music": {
            "sound": "minecraft:music.nether.nether_wastes",
            "replace_current_music": 0,
            "min_delay": 12000,
            "max_delay": 24000
          },
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.nether_wastes.mood",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 3344392,
          "ambient_sound": "minecraft:ambient.nether_wastes.loop",
          "additions_sound": {
            "tick_chance": 0.0111,
            "sound": "minecraft:ambient.nether_wastes.additions"
          }
        },
        "downfall": 0,
        "depth": 0.1,
        "category": "nether"
      }
    },
    {
      "name": "minecraft:the_end",
      "id": 9,
      "element": {
        "temperature": 0.5,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 0,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 10518688
        },
        "downfall": 0.5,
        "depth": 0.1,
        "category": "the_end"
      }
    },
    {
      "name": "minecraft:frozen_ocean",
      "id": 10,
      "element": {
        "temperature_modifier": "frozen",
        "temperature": 0,
        "scale": 0.1,
        "precipitation": "snow",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 3750089,
          "sky_color": 8364543,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": -1,
        "category": "ocean"
      }
    },
    {
      "name": "minecraft:frozen_river",
      "id": 11,
      "element": {
        "temperature": 0,
        "scale": 0,
        "precipitation": "snow",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 3750089,
          "sky_color": 8364543,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": -0.5,
        "category": "river"
      }
    },
    {
      "name": "minecraft:snowy_tundra",
      "id": 12,
      "element": {
        "temperature": 0,
        "scale": 0.05,
        "precipitation": "snow",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8364543,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": 0.125,
        "category": "icy"
      }
    },
    {
      "name": "minecraft:snowy_mountains",
      "id": 13,
      "element": {
        "temperature": 0,
        "scale": 0.3,
        "precipitation": "snow",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8364543,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": 0.45,
        "category": "icy"
      }
    },
    {
      "name": "minecraft:mushroom_fields",
      "id": 14,
      "element": {
        "temperature": 0.9,
        "scale": 0.3,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7842047,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 1,
        "depth": 0.2,
        "category": "mushroom"
      }
    },
    {
      "name": "minecraft:mushroom_field_shore",
      "id": 15,
      "element": {
        "temperature": 0.9,
        "scale": 0.025,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7842047,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 1,
        "depth": 0,
        "category": "mushroom"
      }
    },
    {
      "name": "minecraft:beach",
      "id": 16,
      "element": {
        "temperature": 0.8,
        "scale": 0.025,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7907327,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.4,
        "depth": 0,
        "category": "beach"
      }
    },
    {
      "name": "minecraft:desert_hills",
      "id": 17,
      "element": {
        "temperature": 2,
        "scale": 0.3,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 0.45,
        "category": "desert"
      }
    },
    {
      "name": "minecraft:wooded_hills",
      "id": 18,
      "element": {
        "temperature": 0.7,
        "scale": 0.3,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7972607,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.45,
        "category": "forest"
      }
    },
    {
      "name": "minecraft:taiga_hills",
      "id": 19,
      "element": {
        "temperature": 0.25,
        "scale": 0.3,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8233983,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.45,
        "category": "taiga"
      }
    },
    {
      "name": "minecraft:mountain_edge",
      "id": 20,
      "element": {
        "temperature": 0.2,
        "scale": 0.3,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8233727,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.3,
        "depth": 0.8,
        "category": "extreme_hills"
      }
    },
    {
      "name": "minecraft:jungle",
      "id": 21,
      "element": {
        "temperature": 0.95,
        "scale": 0.2,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7842047,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.9,
        "depth": 0.1,
        "category": "jungle"
      }
    },
    {
      "name": "minecraft:jungle_hills",
      "id": 22,
      "element": {
        "temperature": 0.95,
        "scale": 0.3,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7842047,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.9,
        "depth": 0.45,
        "category": "jungle"
      }
    },
    {
      "name": "minecraft:jungle_edge",
      "id": 23,
      "element": {
        "temperature": 0.95,
        "scale": 0.2,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7842047,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.1,
        "category": "jungle"
      }
    },
    {
      "name": "minecraft:deep_ocean",
      "id": 24,
      "element": {
        "temperature": 0.5,
        "scale": 0.1,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8103167,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": -1.8,
        "category": "ocean"
      }
    },
    {
      "name": "minecraft:stone_shore",
      "id": 25,
      "element": {
        "temperature": 0.2,
        "scale": 0.8,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8233727,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.3,
        "depth": 0.1,
        "category": "none"
      }
    },
    {
      "name": "minecraft:snowy_beach",
      "id": 26,
      "element": {
        "temperature": 0.05,
        "scale": 0.025,
        "precipitation": "snow",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4020182,
          "sky_color": 8364543,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.3,
        "depth": 0,
        "category": "beach"
      }
    },
    {
      "name": "minecraft:birch_forest",
      "id": 27,
      "element": {
        "temperature": 0.6,
        "scale": 0.2,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8037887,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.6,
        "depth": 0.1,
        "category": "forest"
      }
    },
    {
      "name": "minecraft:birch_forest_hills",
      "id": 28,
      "element": {
        "temperature": 0.6,
        "scale": 0.3,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8037887,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.6,
        "depth": 0.45,
        "category": "forest"
      }
    },
    {
      "name": "minecraft:dark_forest",
      "id": 29,
      "element": {
        "temperature": 0.7,
        "scale": 0.2,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7972607,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "grass_color_modifier": "dark_forest",
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.1,
        "category": "forest"
      }
    },
    {
      "name": "minecraft:snowy_taiga",
      "id": 30,
      "element": {
        "temperature": -0.5,
        "scale": 0.2,
        "precipitation": "snow",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4020182,
          "sky_color": 8625919,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.4,
        "depth": 0.2,
        "category": "taiga"
      }
    },
    {
      "name": "minecraft:snowy_taiga_hills",
      "id": 31,
      "element": {
        "temperature": -0.5,
        "scale": 0.3,
        "precipitation": "snow",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4020182,
          "sky_color": 8625919,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.4,
        "depth": 0.45,
        "category": "taiga"
      }
    },
    {
      "name": "minecraft:giant_tree_taiga",
      "id": 32,
      "element": {
        "temperature": 0.3,
        "scale": 0.2,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8168447,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.2,
        "category": "taiga"
      }
    },
    {
      "name": "minecraft:giant_tree_taiga_hills",
      "id": 33,
      "element": {
        "temperature": 0.3,
        "scale": 0.3,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8168447,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.45,
        "category": "taiga"
      }
    },
    {
      "name": "minecraft:wooded_mountains",
      "id": 34,
      "element": {
        "temperature": 0.2,
        "scale": 0.5,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8233727,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.3,
        "depth": 1,
        "category": "extreme_hills"
      }
    },
    {
      "name": "minecraft:savanna",
      "id": 35,
      "element": {
        "temperature": 1.2,
        "scale": 0.05,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7711487,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 0.125,
        "category": "savanna"
      }
    },
    {
      "name": "minecraft:savanna_plateau",
      "id": 36,
      "element": {
        "temperature": 1,
        "scale": 0.025,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7776511,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 1.5,
        "category": "savanna"
      }
    },
    {
      "name": "minecraft:badlands",
      "id": 37,
      "element": {
        "temperature": 2,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "grass_color": 9470285,
          "foliage_color": 10387789,
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 0.1,
        "category": "mesa"
      }
    },
    {
      "name": "minecraft:wooded_badlands_plateau",
      "id": 38,
      "element": {
        "temperature": 2,
        "scale": 0.025,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "grass_color": 9470285,
          "foliage_color": 10387789,
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 1.5,
        "category": "mesa"
      }
    },
    {
      "name": "minecraft:badlands_plateau",
      "id": 39,
      "element": {
        "temperature": 2,
        "scale": 0.025,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "grass_color": 9470285,
          "foliage_color": 10387789,
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 1.5,
        "category": "mesa"
      }
    },
    {
      "name": "minecraft:small_end_islands",
      "id": 40,
      "element": {
        "temperature": 0.5,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 0,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 10518688
        },
        "downfall": 0.5,
        "depth": 0.1,
        "category": "the_end"
      }
    },
    {
      "name": "minecraft:end_midlands",
      "id": 41,
      "element": {
        "temperature": 0.5,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 0,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 10518688
        },
        "downfall": 0.5,
        "depth": 0.1,
        "category": "the_end"
      }
    },
    {
      "name": "minecraft:end_highlands",
      "id": 42,
      "element": {
        "temperature": 0.5,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 0,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 10518688
        },
        "downfall": 0.5,
        "depth": 0.1,
        "category": "the_end"
      }
    },
    {
      "name": "minecraft:end_barrens",
      "id": 43,
      "element": {
        "temperature": 0.5,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 0,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 10518688
        },
        "downfall": 0.5,
        "depth": 0.1,
        "category": "the_end"
      }
    },
    {
      "name": "minecraft:warm_ocean",
      "id": 44,
      "element": {
        "temperature": 0.5,
        "scale": 0.1,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 270131,
          "water_color": 4445678,
          "sky_color": 8103167,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": -1,
        "category": "ocean"
      }
    },
    {
      "name": "minecraft:lukewarm_ocean",
      "id": 45,
      "element": {
        "temperature": 0.5,
        "scale": 0.1,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 267827,
          "water_color": 4566514,
          "sky_color": 8103167,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": -1,
        "category": "ocean"
      }
    },
    {
      "name": "minecraft:cold_ocean",
      "id": 46,
      "element": {
        "temperature": 0.5,
        "scale": 0.1,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4020182,
          "sky_color": 8103167,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": -1,
        "category": "ocean"
      }
    },
    {
      "name": "minecraft:deep_warm_ocean",
      "id": 47,
      "element": {
        "temperature": 0.5,
        "scale": 0.1,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 270131,
          "water_color": 4445678,
          "sky_color": 8103167,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": -1.8,
        "category": "ocean"
      }
    },
    {
      "name": "minecraft:deep_lukewarm_ocean",
      "id": 48,
      "element": {
        "temperature": 0.5,
        "scale": 0.1,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 267827,
          "water_color": 4566514,
          "sky_color": 8103167,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": -1.8,
        "category": "ocean"
      }
    },
    {
      "name": "minecraft:deep_cold_ocean",
      "id": 49,
      "element": {
        "temperature": 0.5,
        "scale": 0.1,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4020182,
          "sky_color": 8103167,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": -1.8,
        "category": "ocean"
      }
    },
    {
      "name": "minecraft:deep_frozen_ocean",
      "id": 50,
      "element": {
        "temperature_modifier": "frozen",
        "temperature": 0.5,
        "scale": 0.1,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 3750089,
          "sky_color": 8103167,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": -1.8,
        "category": "ocean"
      }
    },
    {
      "name": "minecraft:the_void",
      "id": 127,
      "element": {
        "temperature": 0.5,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8103167,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": 0.1,
        "category": "none"
      }
    },
    {
      "name": "minecraft:sunflower_plains",
      "id": 129,
      "element": {
        "temperature": 0.8,
        "scale": 0.05,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7907327,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.4,
        "depth": 0.125,
        "category": "plains"
      }
    },
    {
      "name": "minecraft:desert_lakes",
      "id": 130,
      "element": {
        "temperature": 2,
        "scale": 0.25,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 0.225,
        "category": "desert"
      }
    },
    {
      "name": "minecraft:gravelly_mountains",
      "id": 131,
      "element": {
        "temperature": 0.2,
        "scale": 0.5,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8233727,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.3,
        "depth": 1,
        "category": "extreme_hills"
      }
    },
    {
      "name": "minecraft:flower_forest",
      "id": 132,
      "element": {
        "temperature": 0.7,
        "scale": 0.4,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7972607,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.1,
        "category": "forest"
      }
    },
    {
      "name": "minecraft:taiga_mountains",
      "id": 133,
      "element": {
        "temperature": 0.25,
        "scale": 0.4,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8233983,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.3,
        "category": "taiga"
      }
    },
    {
      "name": "minecraft:swamp_hills",
      "id": 134,
      "element": {
        "temperature": 0.8,
        "scale": 0.3,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 2302743,
          "water_color": 6388580,
          "sky_color": 7907327,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "grass_color_modifier": "swamp",
          "foliage_color": 6975545,
          "fog_color": 12638463
        },
        "downfall": 0.9,
        "depth": -0.1,
        "category": "swamp"
      }
    },
    {
      "name": "minecraft:ice_spikes",
      "id": 140,
      "element": {
        "temperature": 0,
        "scale": 0.45000002,
        "precipitation": "snow",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8364543,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": 0.425,
        "category": "icy"
      }
    },
    {
      "name": "minecraft:modified_jungle",
      "id": 149,
      "element": {
        "temperature": 0.95,
        "scale": 0.4,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7842047,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.9,
        "depth": 0.2,
        "category": "jungle"
      }
    },
    {
      "name": "minecraft:modified_jungle_edge",
      "id": 151,
      "element": {
        "temperature": 0.95,
        "scale": 0.4,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7842047,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.2,
        "category": "jungle"
      }
    },
    {
      "name": "minecraft:tall_birch_forest",
      "id": 155,
      "element": {
        "temperature": 0.6,
        "scale": 0.4,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8037887,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.6,
        "depth": 0.2,
        "category": "forest"
      }
    },
    {
      "name": "minecraft:tall_birch_hills",
      "id": 156,
      "element": {
        "temperature": 0.6,
        "scale": 0.5,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8037887,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.6,
        "depth": 0.55,
        "category": "forest"
      }
    },
    {
      "name": "minecraft:dark_forest_hills",
      "id": 157,
      "element": {
        "temperature": 0.7,
        "scale": 0.4,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7972607,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "grass_color_modifier": "dark_forest",
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.2,
        "category": "forest"
      }
    },
    {
      "name": "minecraft:snowy_taiga_mountains",
      "id": 158,
      "element": {
        "temperature": -0.5,
        "scale": 0.4,
        "precipitation": "snow",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4020182,
          "sky_color": 8625919,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.4,
        "depth": 0.3,
        "category": "taiga"
      }
    },
    {
      "name": "minecraft:giant_spruce_taiga",
      "id": 160,
      "element": {
        "temperature": 0.25,
        "scale": 0.2,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8233983,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.2,
        "category": "taiga"
      }
    },
    {
      "name": "minecraft:giant_spruce_taiga_hills",
      "id": 161,
      "element": {
        "temperature": 0.25,
        "scale": 0.2,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8233983,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.8,
        "depth": 0.2,
        "category": "taiga"
      }
    },
    {
      "name": "minecraft:modified_gravelly_mountains",
      "id": 162,
      "element": {
        "temperature": 0.2,
        "scale": 0.5,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8233727,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.3,
        "depth": 1,
        "category": "extreme_hills"
      }
    },
    {
      "name": "minecraft:shattered_savanna",
      "id": 163,
      "element": {
        "temperature": 1.1,
        "scale": 1.225,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7776767,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 0.3625,
        "category": "savanna"
      }
    },
    {
      "name": "minecraft:shattered_savanna_plateau",
      "id": 164,
      "element": {
        "temperature": 1,
        "scale": 1.2125001,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7776511,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 1.05,
        "category": "savanna"
      }
    },
    {
      "name": "minecraft:eroded_badlands",
      "id": 165,
      "element": {
        "temperature": 2,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "grass_color": 9470285,
          "foliage_color": 10387789,
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 0.1,
        "category": "mesa"
      }
    },
    {
      "name": "minecraft:modified_wooded_badlands_plateau",
      "id": 166,
      "element": {
        "temperature": 2,
        "scale": 0.3,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "grass_color": 9470285,
          "foliage_color": 10387789,
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 0.45,
        "category": "mesa"
      }
    },
    {
      "name": "minecraft:modified_badlands_plateau",
      "id": 167,
      "element": {
        "temperature": 2,
        "scale": 0.3,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "grass_color": 9470285,
          "foliage_color": 10387789,
          "fog_color": 12638463
        },
        "downfall": 0,
        "depth": 0.45,
        "category": "mesa"
      }
    },
    {
      "name": "minecraft:bamboo_jungle",
      "id": 168,
      "element": {
        "temperature": 0.95,
        "scale": 0.2,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7842047,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.9,
        "depth": 0.1,
        "category": "jungle"
      }
    },
    {
      "name": "minecraft:bamboo_jungle_hills",
      "id": 169,
      "element": {
        "temperature": 0.95,
        "scale": 0.3,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7842047,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.9,
        "depth": 0.45,
        "category": "jungle"
      }
    },
    {
      "name": "minecraft:soul_sand_valley",
      "id": 170,
      "element": {
        "temperature": 2,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "particle": {
            "probability": 0.00625,
            "options": {
              "type": "minecraft:ash"
            }
          },
          "music": {
            "sound": "minecraft:music.nether.soul_sand_valley",
            "replace_current_music": 0,
            "min_delay": 12000,
            "max_delay": 24000
          },
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.soul_sand_valley.mood",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 1787717,
          "ambient_sound": "minecraft:ambient.soul_sand_valley.loop",
          "additions_sound": {
            "tick_chance": 0.0111,
            "sound": "minecraft:ambient.soul_sand_valley.additions"
          }
        },
        "downfall": 0,
        "depth": 0.1,
        "category": "nether"
      }
    },
    {
      "name": "minecraft:crimson_forest",
      "id": 171,
      "element": {
        "temperature": 2,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "particle": {
            "probability": 0.025,
            "options": {
              "type": "minecraft:crimson_spore"
            }
          },
          "music": {
            "sound": "minecraft:music.nether.crimson_forest",
            "replace_current_music": 0,
            "min_delay": 12000,
            "max_delay": 24000
          },
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.crimson_forest.mood",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 3343107,
          "ambient_sound": "minecraft:ambient.crimson_forest.loop",
          "additions_sound": {
            "tick_chance": 0.0111,
            "sound": "minecraft:ambient.crimson_forest.additions"
          }
        },
        "downfall": 0,
        "depth": 0.1,
        "category": "nether"
      }
    },
    {
      "name": "minecraft:warped_forest",
      "id": 172,
      "element": {
        "temperature": 2,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7254527,
          "particle": {
            "probability": 0.01428,
            "options": {
              "type": "minecraft:warped_spore"
            }
          },
          "music": {
            "sound": "minecraft:music.nether.warped_forest",
            "replace_current_music": 0,
            "min_delay": 12000,
            "max_delay": 24000
          },
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.warped_forest.mood",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 1705242,
          "ambient_sound": "minecraft:ambient.warped_forest.loop",
          "additions_sound": {
            "tick_chance": 0.0111,
            "sound": "minecraft:ambient.warped_forest.additions"
          }
        },
        "downfall": 0,
        "depth": 0.1,
        "category": "nether"
      }
    },
    {
      "name": "minecraft:basalt_deltas",
      "id": 173,
      "element": {
        "temperature": 2,
        "scale": 0.2,
        "precipitation": "none",
        "effects": {
          "water_fog_color": 4341314,
          "water_color": 4159204,
          "sky_color": 7254527,
          "particle": {
            "probability": 0.118093334,
            "options": {
              "type": "minecraft:white_ash"
            }
          },
          "music": {
            "sound": "minecraft:music.nether.basalt_deltas",
            "replace_current_music": 0,
            "min_delay": 12000,
            "max_delay": 24000
          },
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.basalt_deltas.mood",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 6840176,
          "ambient_sound": "minecraft:ambient.basalt_deltas.loop",
          "additions_sound": {
            "tick_chance": 0.0111,
            "sound": "minecraft:ambient.basalt_deltas.additions"
          }
        },
        "downfall": 0,
        "depth": 0.1,
        "category": "nether"
      }
    },
    {
      "name": "minecraft:dripstone_caves",
      "id": 174,
      "element": {
        "temperature": 0.8,
        "scale": 0.05,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 7907327,
          "mood_sound": {
            "tick_delay": 6000,
            "sound": "minecraft:ambient.cave",
            "offset": 2,
            "block_search_extent": 8
          },
          "fog_color": 12638463
        },
        "downfall": 0.4,
        "depth": 0.125,
        "category": "underground"
      }
    },
    {
      "name": "minecraft:lush_caves",
      "id": 175,
      "element": {
        "temperature": 0.5,
        "scale": 0.2,
        "precipitation": "rain",
        "effects": {
          "water_fog_color": 329011,
          "water_color": 4159204,
          "sky_color": 8103167,
          "fog_color": 12638463
        },
        "downfall": 0.5,
        "depth": 0.1,
        "category": "underground"
      }
    }
  ],
  "type": "minecraft:worldgen/biome"
}
`

// buildBiomeRegistry builds the biome registry from the reference json.
func buildBiomeRegistry() map[string]interface{} {
	var b map[string]interface{}
	_ = json.Unmarshal([]byte(biomeRegistryJSON), &b)
	return b
}
