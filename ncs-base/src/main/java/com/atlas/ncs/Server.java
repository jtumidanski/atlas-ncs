package com.atlas.ncs;

import com.atlas.kafka.consumer.SimpleEventConsumerBuilder;
import com.atlas.ncs.event.consumer.ContinueNpcConversationConsumer;
import com.atlas.ncs.event.consumer.SetReturnTextConsumer;
import com.atlas.ncs.event.consumer.StartNpcConversationConsumer;

public class Server {
   public static void main(String[] args) {
      SimpleEventConsumerBuilder.builder()
            .addConsumer(new StartNpcConversationConsumer())
            .addConsumer(new ContinueNpcConversationConsumer())
            .addConsumer(new SetReturnTextConsumer())
            .initialize();
   }
}
