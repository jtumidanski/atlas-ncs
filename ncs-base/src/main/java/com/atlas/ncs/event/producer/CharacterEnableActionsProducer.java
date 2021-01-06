package com.atlas.ncs.event.producer;

import com.atlas.csrv.command.EnableActionsCommand;
import com.atlas.csrv.constant.CommandConstants;
import com.atlas.ncs.EventProducerRegistry;

public final class CharacterEnableActionsProducer {
   public static void enableActions(int characterId) {
      EventProducerRegistry.getInstance().send(CommandConstants.TOPIC_ENABLE_ACTIONS, characterId,
            new EnableActionsCommand(characterId));
   }
}
