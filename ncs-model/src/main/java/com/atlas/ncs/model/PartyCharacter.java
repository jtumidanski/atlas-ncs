package com.atlas.ncs.model;

public record PartyCharacter(int level, int jobId, int mapId) {
   public boolean inMap(int mapId) {
      return mapId == this.mapId;
   }
}
