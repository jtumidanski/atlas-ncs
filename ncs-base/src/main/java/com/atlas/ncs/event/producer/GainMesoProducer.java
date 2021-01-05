package com.atlas.ncs.event.producer;

import com.atlas.cos.command.GainMesoCommand;
import com.atlas.cos.constant.CommandConstants;
import com.atlas.ncs.EventProducerRegistry;

public final class GainMesoProducer {
   private GainMesoProducer() {
   }

   public static void command(int characterId, int meso) {
      EventProducerRegistry.getInstance()
            .send(GainMesoCommand.class, CommandConstants.TOPIC_GAIN_MESO, characterId,
                  new GainMesoCommand(characterId, meso));
   }
}
