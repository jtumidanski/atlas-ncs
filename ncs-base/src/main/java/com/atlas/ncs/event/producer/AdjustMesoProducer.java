package com.atlas.ncs.event.producer;

import com.atlas.cos.command.AdjustMesoCommand;
import com.atlas.cos.constant.CommandConstants;
import com.atlas.ncs.EventProducerRegistry;

public final class AdjustMesoProducer {
   private AdjustMesoProducer() {
   }

   public static void command(int characterId, int meso) {
      EventProducerRegistry.getInstance().send(CommandConstants.TOPIC_ADJUST_MESO, characterId,
            new AdjustMesoCommand(characterId, meso, true));
   }
}
