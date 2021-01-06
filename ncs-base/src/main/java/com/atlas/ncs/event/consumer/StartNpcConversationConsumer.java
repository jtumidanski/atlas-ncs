package com.atlas.ncs.event.consumer;

import com.atlas.kafka.consumer.SimpleEventHandler;
import com.atlas.ncs.NPCScriptRegistry;
import com.atlas.ncs.command.StartNpcConversationCommand;
import com.atlas.ncs.constant.CommandConstants;
import com.atlas.ncs.processor.TopicDiscoveryProcessor;

public class StartNpcConversationConsumer implements SimpleEventHandler<StartNpcConversationCommand> {
   @Override
   public void handle(Long key, StartNpcConversationCommand command) {
      NPCScriptRegistry.getInstance().start(command.worldId(), command.channelId(), command.mapId(), command.characterId(),
            command.npcId(), command.npcObjectId());
   }

   @Override
   public Class<StartNpcConversationCommand> getEventClass() {
      return StartNpcConversationCommand.class;
   }

   @Override
   public String getConsumerId() {
      return "NPC Conversation Service";
   }

   @Override
   public String getBootstrapServers() {
      return System.getenv("BOOTSTRAP_SERVERS");
   }

   @Override
   public String getTopic() {
      return TopicDiscoveryProcessor.getTopic(CommandConstants.TOPIC_START_NPC_CONVERSATION);
   }
}
