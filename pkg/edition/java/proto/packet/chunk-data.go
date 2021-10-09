package packet

import (
	"io"

	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"go.minekube.com/gate/pkg/edition/java/proto/util"
	"go.minekube.com/gate/pkg/gate/proto"
)

// ChunkData 0x22 sets chunk information for a part of the map
type ChunkData struct {
	// ChunkX is the x coordinate.
	// (block coordinate divided by 16, rounded down)
	ChunkX int
	// ChunkZ is the Z coordinate.
	// (block coordinate divided by 16, rounded down)
	ChunkZ int
	// PrimaryBitMask is a bit set with bits set to 1 for each 16x16x16 chunk
	// section who is included in the data. least significant bit represents the
	// chunk section at the bottom of the chunk column, from the lowest y to 15
	// blocks above.
	PrimaryBitMask []float64
	// Heightmaps is a NBT with containing one long array named MOTION_BLOCKING,
	// which is a heightmap for the highest solid block at each position in the
	// chunk (as a compacted long array with 256 entries at 9 bits per entry
	// totaling 37 longs). The Notchian server also adds a WORLD_SURFACE long
	// array, the purpose of which is unknown, but it's not required for the
	// chunk to be accepted.
	Heightmaps util.NBT
	// Biomes is the Biome IDs, ordered by x then z then y, in 4×4×4 blocks.
	Biomes []int
	// Data is the chunk format data structure
	Data []byte
	// BlockEntities is the list of block entities in the chunk.
	BlockEntities []util.NBT
}

func (r *ChunkData) Encode(c *proto.PacketContext, wr io.Writer) (err error) {
	if err := util.WriteInt(wr, r.ChunkX); err != nil {
		return err
	}
	if err := util.WriteInt(wr, r.ChunkZ); err != nil {
		return err
	}
	if err := util.WriteVarInt(wr, len(r.PrimaryBitMask)); err != nil {
		return err
	}
	for _, bm := range r.PrimaryBitMask {
		if err := util.WriteFloat64(wr, bm); err != nil {
			return err
		}
	}
	nbtEncoder := nbt.NewEncoderWithEncoding(wr, nbt.BigEndian)
	if err := nbtEncoder.Encode(&r.Heightmaps); err != nil {
		return err
	}
	if err := util.WriteVarInt(wr, len(r.Biomes)); err != nil {
		return err
	}
	for _, biome := range r.Biomes {
		if err := util.WriteVarInt(wr, biome); err != nil {
			return err
		}
	}

	err = util.WriteBytes(wr, r.Data)
	if err != nil {
		return err
	}

	if err := util.WriteVarInt(wr, len(r.BlockEntities)); err != nil {
		return err
	}
	for _, blockEntity := range r.BlockEntities {
		if err := nbtEncoder.Encode(&blockEntity); err != nil {
			return err
		}
	}
	return nil
}

func (r *ChunkData) Decode(c *proto.PacketContext, rd io.Reader) (err error) {
	r.ChunkX, err = util.ReadInt(rd)
	if err != nil {
		return err
	}
	r.ChunkZ, err = util.ReadInt(rd)
	if err != nil {
		return err
	}
	bitMaskLen, err := util.ReadVarInt(rd)
	if err != nil {
		return err
	}
	r.PrimaryBitMask = make([]float64, bitMaskLen)
	for i := range r.PrimaryBitMask {
		r.PrimaryBitMask[i], err = util.ReadFloat64(rd)
		if err != nil {
			return err
		}
	}

	nbtDecoder := nbt.NewDecoderWithEncoding(rd, nbt.BigEndian)
	r.Heightmaps = util.NBT{}
	err = nbtDecoder.Decode(&r.Heightmaps)

	biomesCount, err := util.ReadVarInt(rd)
	if err != nil {
		return err
	}
	r.Biomes = make([]int, biomesCount)
	for i := range r.Biomes {
		r.Biomes[i], err = util.ReadVarInt(rd)
		if err != nil {
			return err
		}
	}

	r.Data, err = util.ReadBytes(rd)
	if err != nil {
		return err
	}

	blockEntCount, err := util.ReadVarInt(rd)
	if err != nil {
		return err
	}
	r.BlockEntities = make([]util.NBT, blockEntCount)
	for i := range r.BlockEntities {
		r.BlockEntities[i] = util.NBT{}
		if err := nbtDecoder.Decode(&r.BlockEntities[i]); err != nil {
			return err
		}
	}

	return nil
}

var _ proto.Packet = (*ChunkData)(nil)
