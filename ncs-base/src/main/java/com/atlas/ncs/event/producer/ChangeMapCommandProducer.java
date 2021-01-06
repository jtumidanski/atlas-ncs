package com.atlas.ncs.event.producer;

import com.atlas.cos.command.ChangeMapCommand;
import com.atlas.cos.constant.CommandConstants;
import com.atlas.ncs.EventProducerRegistry;

public final class ChangeMapCommandProducer {
   private ChangeMapCommandProducer() {
   }

   public static void changeMap(int worldId, int channelId, int characterId, int mapId, int portalId) {
      EventProducerRegistry.getInstance().send(CommandConstants.TOPIC_CHANGE_MAP_COMMAND, characterId,
            new ChangeMapCommand(worldId, channelId, characterId, mapId, portalId));
   }
}
