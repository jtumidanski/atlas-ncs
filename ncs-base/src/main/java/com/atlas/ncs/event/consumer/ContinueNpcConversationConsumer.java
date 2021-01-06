package com.atlas.ncs.event.consumer;

import com.atlas.kafka.consumer.SimpleEventHandler;
import com.atlas.ncs.NPCScriptRegistry;
import com.atlas.ncs.command.ContinueNpcConversationCommand;
import com.atlas.ncs.constant.CommandConstants;
import com.atlas.ncs.processor.TopicDiscoveryProcessor;

public class ContinueNpcConversationConsumer implements SimpleEventHandler<ContinueNpcConversationCommand> {
   @Override
   public void handle(Long key, ContinueNpcConversationCommand command) {
      NPCScriptRegistry.getInstance().action(command.characterId(), command.mode(), command.type(), command.selection());
   }

   @Override
   public Class<ContinueNpcConversationCommand> getEventClass() {
      return ContinueNpcConversationCommand.class;
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
      return TopicDiscoveryProcessor.getTopic(CommandConstants.TOPIC_CONTINUE_NPC_CONVERSATION);
   }
}
