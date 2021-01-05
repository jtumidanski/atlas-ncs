package com.atlas.ncs.model;

public record Character(int id, String name, int level, int mapId, int x, int y, int hp, byte gender, int jobId, int face, int hair,
                        int remainingSp, int meso) {
}
