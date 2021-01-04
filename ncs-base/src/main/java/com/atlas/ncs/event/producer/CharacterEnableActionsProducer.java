package com.atlas.ncs.event.producer;

import com.atlas.ncs.EventProducerRegistry;
import com.atlas.csrv.command.EnableActionsCommand;
import com.atlas.csrv.constant.EventConstants;

public final class CharacterEnableActionsProducer {
   public static void enableActions(int characterId) {
      EventProducerRegistry.getInstance().send(EnableActionsCommand.class, EventConstants.TOPIC_ENABLE_ACTIONS, characterId, new EnableActionsCommand(characterId));
   }
}
