package com.atlas.ncs.event.consumer;

import com.atlas.kafka.consumer.SimpleEventHandler;
import com.atlas.ncs.NPCScriptRegistry;
import com.atlas.ncs.command.SetReturnTextCommand;
import com.atlas.ncs.constant.CommandConstants;
import com.atlas.ncs.processor.TopicDiscoveryProcessor;

public class SetReturnTextConsumer implements SimpleEventHandler<SetReturnTextCommand> {
   @Override
   public void handle(Long key, SetReturnTextCommand command) {
      NPCScriptRegistry.getInstance()
            .getCM(command.characterId())
            .ifPresent(cm -> cm.setText(command.text()));
   }

   @Override
   public Class<SetReturnTextCommand> getEventClass() {
      return SetReturnTextCommand.class;
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
      return TopicDiscoveryProcessor.getTopic(CommandConstants.TOPIC_SET_RETURN_TEXT);
   }
}
